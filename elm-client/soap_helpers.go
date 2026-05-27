package elmsoap

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/packer-plugin-sdk/packer"
)

// ErrWorkTicketNotInActiveFilter is returned by GetTaskStateActiveFilter when
// the work ticket is not in the active (pending operations) list, indicating
// it has moved to the completed list.
var ErrWorkTicketNotInActiveFilter = errors.New("work ticket not found in active filter")

// SoapHelper provides methods for querying and monitoring ELM operations.
// It holds session state (SOAP client, credentials, URL) established during login,
// eliminating repeated parameter passing across helper functions.
type SoapHelper struct {
	Client             ApiSoap
	Cookie             string
	Token              string
	URL                string
	InsecureSkipVerify bool
}

// HeaderCaptureTransport is a custom HTTP transport that captures headers and cookies.
type HeaderCaptureTransport struct {
	Rt            http.RoundTripper
	Cookie        string
	Unidesk_token string
}

func (t *HeaderCaptureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("🔹 Starting RoundTrip...")
	if t.Cookie != "" {
		req.Header.Set("Cookie", t.Cookie)
	}
	if t.Unidesk_token != "" {
		req.Header.Set("Unidesk_token", t.Unidesk_token)
	}
	log.Printf("🔹 SOAP Request Headers:")
	for key, values := range req.Header {
		if strings.Contains(strings.ToLower(key), "password") {
			log.Printf("Ignore the header: %s, due to sensitive information.", key)
			continue
		}
		for _, value := range values {
			lkey := strings.ToLower(key)
			if strings.Contains(lkey, "cookie") || strings.Contains(lkey, "token") {
				if len(value) > 4 {
					log.Printf("%s: ...%s", key, value[len(value)-4:])
				} else {
					log.Printf("%s: ...", key)
				}
			} else {
				log.Printf("%s: %s", key, value)
			}
		}
	}
	// Log Request Body
	if req.Body != nil {
		var reqBuf bytes.Buffer
		body := io.TeeReader(req.Body, &reqBuf)
		req.Body = io.NopCloser(&reqBuf)

		data, err := io.ReadAll(body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			return nil, err
		}
		if !strings.Contains(strings.ToLower(string(data)), "password") && !strings.Contains(strings.ToLower(string(data)), "token") {
			log.Printf("🔹 SOAP Request body: %s", string(data))
		}
	}
	resp, err := t.Rt.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if cookies := resp.Header["Set-Cookie"]; len(cookies) > 0 {
		t.Cookie = strings.Split(cookies[0], ";")[0]
	}
	log.Printf("🔸 SOAP Response Headers:")
	for key, values := range resp.Header {
		if strings.Contains(strings.ToLower(key), "password") {
			log.Printf("Ignore the header: %s, due to sensitive information.", key)
			continue
		}
		for _, value := range values {
			lkey := strings.ToLower(key)
			if strings.Contains(lkey, "cookie") || strings.Contains(lkey, "token") {
				if len(value) > 4 {
					log.Printf("%s: ...%s", key, value[len(value)-4:])
				} else {
					log.Printf("%s: ...", key)
				}
			} else {
				log.Printf("%s: %s", key, value)
			}
		}
	}
	// Log Response Body
	if resp.Body != nil {
		var resBuf bytes.Buffer
		body := io.TeeReader(resp.Body, &resBuf)
		resp.Body = io.NopCloser(&resBuf)

		data, _ := io.ReadAll(body)
		respStr := string(data)

		// Mask token if present
		re := regexp.MustCompile(`<Token>([^<]+)</Token>`)
		respStr = re.ReplaceAllStringFunc(respStr, func(tokenTag string) string {
			matches := re.FindStringSubmatch(tokenTag)
			if len(matches) == 2 {
				token := matches[1]
				if len(token) > 4 {
					return "<Token>xxx" + token[len(token)-4:] + "</Token>"
				}
				return "<Token>xxx</Token>"
			}
			return tokenTag
		})

		if strings.Contains(strings.ToLower(respStr), "password") {
			log.Printf("Ignore the SOAP Response body due to sensitive info.")
		} else {
			log.Printf("🔸 SOAP Response body: %s", respStr)
		}
	}

	log.Printf("🔹 End RoundTrip...")
	return resp, nil
}

type ApplayeringOperationType string

const (
	REVISION_OS_LAYER       = "REVISION_OS_LAYER"
	REVISION_PLATFORM_LAYER = "REVISION_PLATFORM_LAYER"
	REVISION_APP_LAYER      = "REVISION_APP_LAYER"
	CREATE_PLATFORM_LAYER   = "CREATE_PLATFORM_LAYER"
	CREATE_APP_LAYER        = "CREATE_APP_LAYER"

	//only for debugging
	CONNECT_REVISION_OS_VM_ONLY       = "CONNECT_REVISION_OS_VM_ONLY"
	CONNECT_REVISION_PLATFORM_VM_ONLY = "CONNECT_REVISION_PLATFORM_VM_ONLY"
	CONNECT_REVISION_APP_VM_ONLY      = "CONNECT_REVISION_APP_VM_ONLY"
	CONNECT_CREATE_PLATFORM_VM_ONLY   = "CONNECT_CREATE_PLATFORM_VM_ONLY"
	CONNECT_CREATE_APP_VM_ONLY        = "CONNECT_CREATE_APP_VM_ONLY"

	APPLAYERING_OPERATION_TYPE_UNDEFINED = "APPLAYERING_OPERATION_TYPE_UNDEFINED"
)

func IsValidApplayeringOperationType(optionalType string) bool {
	switch ApplayeringOperationType(optionalType) {
	case REVISION_OS_LAYER,
		REVISION_PLATFORM_LAYER,
		REVISION_APP_LAYER,
		CREATE_PLATFORM_LAYER,
		CREATE_APP_LAYER,
		CONNECT_REVISION_OS_VM_ONLY,
		CONNECT_REVISION_PLATFORM_VM_ONLY,
		CONNECT_REVISION_APP_VM_ONLY,
		CONNECT_CREATE_PLATFORM_VM_ONLY,
		CONNECT_CREATE_APP_VM_ONLY,
		APPLAYERING_OPERATION_TYPE_UNDEFINED:
		return true
	default:
		return false
	}
}

func GetAllSupportedApplayeringOperationTypes() []string {
	return []string{
		REVISION_OS_LAYER,
		REVISION_PLATFORM_LAYER,
		REVISION_APP_LAYER,
		CREATE_PLATFORM_LAYER,
		CREATE_APP_LAYER,
	}
}

func IsValidDiskFormat(diskFormat string) bool {
	switch DiskFormat(diskFormat) {
	case DiskFormatVhdDiskFormat,
		DiskFormatVmdkDiskFormat,
		DiskFormatVmdkSparseDiskFormat,
		DiskFormatVhdxDiskFormat,
		DiskFormatQCow2DiskFormat:
		return true
	}
	return false
}
func GetAllSupportedDiskFormats() []string {
	return []string{
		string(DiskFormatVhdDiskFormat),
		string(DiskFormatVmdkDiskFormat),
		string(DiskFormatVmdkSparseDiskFormat),
		string(DiskFormatVhdxDiskFormat),
		string(DiskFormatQCow2DiskFormat),
	}
}

// CheckWebResultError safely extracts an ELM error from a *WebResultBase,
// handling nil embedded pointers that Go's XML decoder does not allocate
// when the corresponding XML elements are absent.
func CheckWebResultError(w *WebResultBase) error {
	if w == nil || w.ResultBase == nil || w.ResultBase.Error == nil || w.ResultBase.Error.Message == "" {
		return nil
	}
	return fmt.Errorf("%s", w.ResultBase.Error.Message)
}

func (s *SoapHelper) GetPlatformConnectorConfigId(platformConnectorConfigName string) (string, error) {
	queryPlatformConnectorConfigSummaryRequest := &QueryPlatformConnectorConfigSummary{
		Query: &PlatformConnectorConfigSummaryQuery{},
	}
	queryPlatformConnectorConfigSummaryResponse, err := s.Client.QueryPlatformConnectorConfigSummary(queryPlatformConnectorConfigSummaryRequest)
	if err != nil {
		return "", fmt.Errorf("GetPlatformConnectorConfigId::error calling QueryPlatformConnectorConfigSummary: %v", err)
	}
	if r := queryPlatformConnectorConfigSummaryResponse.QueryPlatformConnectorConfigSummaryResult; r != nil {
		if err := CheckWebResultError(r.WebResultBase); err != nil {
			return "", fmt.Errorf("QueryPlatformConnectorConfigSummary: %w", err)
		}
	}
	for _, platformConnectorConfig := range queryPlatformConnectorConfigSummaryResponse.QueryPlatformConnectorConfigSummaryResult.Configurations.PlatformConnectorConfigSummary {
		if platformConnectorConfig.Name == platformConnectorConfigName {
			return platformConnectorConfig.Id, nil
		}
	}
	return "", fmt.Errorf("platform connector config not found for name: %s", platformConnectorConfigName)
}

func (s *SoapHelper) GetOsLayerId(osLayerName string) (int64, error) {
	osLayerQueryRequest := &QueryOsLayers{
		Query: &OsLayersQuery{
			Filter: osLayerName,
		},
	}
	queryOsLayersResponse, err := s.Client.QueryOsLayers(osLayerQueryRequest)
	if err != nil {
		return 0, fmt.Errorf("error calling QueryOsLayers: %v", err)
	}
	if queryOsLayersResponse.QueryOsLayersResult == nil {
		return 0, fmt.Errorf("os layer not found for name: %s", osLayerName)
	}
	if err := CheckWebResultError(queryOsLayersResponse.QueryOsLayersResult.WebResultBase); err != nil {
		return 0, fmt.Errorf("QueryOsLayers: %w", err)
	}
	if queryOsLayersResponse.QueryOsLayersResult.OsLayers == nil {
		return 0, fmt.Errorf("os layer not found for name: %s", osLayerName)
	}
	for _, osLayer := range queryOsLayersResponse.QueryOsLayersResult.OsLayers.LayerEntitySummary {
		if osLayer.Name == osLayerName {
			return osLayer.Id, nil
		}
	}
	return 0, fmt.Errorf("os layer not found for name: %s", osLayerName)
}

func (s *SoapHelper) GetOsLayerRevisionId(osLayerName string, osLayerVersionName string) (int64, error) {
	osLayerQueryRequest := &QueryOsLayers{
		Query: &OsLayersQuery{
			Filter: osLayerName,
		},
	}
	queryOsLayersResponse, err := s.Client.QueryOsLayers(osLayerQueryRequest)
	if err != nil {
		return 0, fmt.Errorf("error calling QueryOsLayers: %v", err)
	}
	if queryOsLayersResponse.QueryOsLayersResult == nil {
		return 0, fmt.Errorf("os layer revision not found for name: %s and version: %s", osLayerName, osLayerVersionName)
	}
	if err := CheckWebResultError(queryOsLayersResponse.QueryOsLayersResult.WebResultBase); err != nil {
		return 0, fmt.Errorf("QueryOsLayers: %w", err)
	}
	if queryOsLayersResponse.QueryOsLayersResult.OsLayers == nil {
		return 0, fmt.Errorf("os layer revision not found for name: %s and version: %s", osLayerName, osLayerVersionName)
	}
	for _, osLayer := range queryOsLayersResponse.QueryOsLayersResult.OsLayers.LayerEntitySummary {
		if osLayer.Name == osLayerName {
			queryOsLayerDetailsResponse, err := s.Client.QueryOsLayerDetails(&QueryOsLayerDetails{
				Query: &LayerDetailsQuery{
					Id: osLayer.Id,
				}})
			if err != nil {
				return 0, fmt.Errorf("error calling QueryOsLayerDetails: %v", err)
			}
			r := queryOsLayerDetailsResponse.QueryOsLayerDetailsResult
			if r == nil || r.LayerDetailsResultOfOsLayerRevisionDetail == nil {
				continue
			}
			if err := CheckWebResultError(r.LayerDetailsResultOfOsLayerRevisionDetail.WebResultBase); err != nil {
				return 0, fmt.Errorf("QueryOsLayerDetails: %w", err)
			}
			if r.Revisions == nil {
				continue
			}
			for _, revision := range r.Revisions.OsLayerRevisionDetail {
				if revision.DisplayedVersion == osLayerVersionName {
					return revision.Id, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("os layer revision not found for name: %s and version: %s", osLayerName, osLayerVersionName)
}

func (s *SoapHelper) GetPlatformLayerId(platformLayerName string) (int64, error) {
	platformLayerQueryRequest := &QueryPlatformLayers{
		Query: &PlatformLayersQuery{
			Filter: platformLayerName,
		},
	}
	queryPlatformLayersResponse, err := s.Client.QueryPlatformLayers(platformLayerQueryRequest)
	if err != nil {
		return 0, fmt.Errorf("error calling QueryPlatformLayers: %v", err)
	}
	if queryPlatformLayersResponse.QueryPlatformLayersResult == nil {
		return 0, fmt.Errorf("platform layer not found for name: %s", platformLayerName)
	}
	if err := CheckWebResultError(queryPlatformLayersResponse.QueryPlatformLayersResult.WebResultBase); err != nil {
		return 0, fmt.Errorf("QueryPlatformLayers: %w", err)
	}
	if queryPlatformLayersResponse.QueryPlatformLayersResult.PlatformLayers == nil {
		return 0, fmt.Errorf("platform layer not found for name: %s", platformLayerName)
	}
	for _, platformLayer := range queryPlatformLayersResponse.QueryPlatformLayersResult.PlatformLayers.LayerEntitySummary {
		if platformLayer.Name == platformLayerName {
			return platformLayer.Id, nil
		}
	}
	return 0, fmt.Errorf("platform layer not found for name: %s", platformLayerName)
}

func (s *SoapHelper) GetPlatformLayerRevisionId(platformLayerName string, platformLayerVersionName string) (int64, error) {
	platformLayerQueryRequest := &QueryPlatformLayers{
		Query: &PlatformLayersQuery{
			Filter: platformLayerName,
		},
	}
	queryPlatformLayersResponse, err := s.Client.QueryPlatformLayers(platformLayerQueryRequest)
	if err != nil {
		return 0, fmt.Errorf("error calling QueryPlatformLayers: %v", err)
	}
	if queryPlatformLayersResponse.QueryPlatformLayersResult == nil {
		return 0, fmt.Errorf("platform layer revision not found for name: %s and version: %s", platformLayerName, platformLayerVersionName)
	}
	if err := CheckWebResultError(queryPlatformLayersResponse.QueryPlatformLayersResult.WebResultBase); err != nil {
		return 0, fmt.Errorf("QueryPlatformLayers: %w", err)
	}
	if queryPlatformLayersResponse.QueryPlatformLayersResult.PlatformLayers == nil {
		return 0, fmt.Errorf("platform layer revision not found for name: %s and version: %s", platformLayerName, platformLayerVersionName)
	}
	for _, platformLayer := range queryPlatformLayersResponse.QueryPlatformLayersResult.PlatformLayers.LayerEntitySummary {
		if platformLayer.Name == platformLayerName {
			queryOsLayerDetailsResponse, err := s.Client.QueryPlatformLayerDetails(&QueryPlatformLayerDetails{
				Query: &LayerDetailsQuery{
					Id: platformLayer.Id,
				}})
			if err != nil {
				return 0, fmt.Errorf("error calling QueryPlatformLayerDetails: %v", err)
			}
			r := queryOsLayerDetailsResponse.QueryPlatformLayerDetailsResult
			if r == nil || r.LayerDetailsResultOfPlatformLayerRevisionDetail == nil {
				continue
			}
			if err := CheckWebResultError(r.LayerDetailsResultOfPlatformLayerRevisionDetail.WebResultBase); err != nil {
				return 0, fmt.Errorf("QueryPlatformLayerDetails: %w", err)
			}
			if r.Revisions == nil {
				continue
			}
			for _, revision := range r.Revisions.PlatformLayerRevisionDetail {
				if revision.DisplayedVersion == platformLayerVersionName {
					return revision.Id, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("platform layer revision not found for name: %s and version: %s", platformLayerName, platformLayerVersionName)
}

func (s *SoapHelper) GetAppLayerRevisionId(appLayerName string, appLayerVersionName string) (int64, error) {
	appLayerQueryRequest := &QueryApplicationLayers{
		Query: &AppLayersQuery{
			Filter: appLayerName,
		},
	}
	queryAppLayersResponse, err := s.Client.QueryApplicationLayers(appLayerQueryRequest)
	if err != nil {
		return 0, fmt.Errorf("error calling QueryApplicationLayers: %v", err)
	}
	if queryAppLayersResponse.QueryApplicationLayersResult == nil {
		return 0, fmt.Errorf("app layer revision not found for name: %s and version: %s", appLayerName, appLayerVersionName)
	}
	if err := CheckWebResultError(queryAppLayersResponse.QueryApplicationLayersResult.WebResultBase); err != nil {
		return 0, fmt.Errorf("QueryApplicationLayers: %w", err)
	}
	if queryAppLayersResponse.QueryApplicationLayersResult.AppLayers == nil {
		return 0, fmt.Errorf("app layer revision not found for name: %s and version: %s", appLayerName, appLayerVersionName)
	}
	for _, appLayer := range queryAppLayersResponse.QueryApplicationLayersResult.AppLayers.LayerEntitySummary {
		if appLayer.Name == appLayerName {
			queryOsLayerDetailsResponse, err := s.Client.QueryApplicationLayerDetails(&QueryApplicationLayerDetails{
				Query: &LayerDetailsQuery{
					Id: appLayer.Id,
				}})
			if err != nil {
				return 0, fmt.Errorf("error calling QueryApplicationLayerDetails: %v", err)
			}
			r := queryOsLayerDetailsResponse.QueryApplicationLayerDetailsResult
			if r == nil || r.LayerDetailsResultOfAppLayerRevisionDetail == nil {
				continue
			}
			if err := CheckWebResultError(r.LayerDetailsResultOfAppLayerRevisionDetail.WebResultBase); err != nil {
				return 0, fmt.Errorf("QueryApplicationLayerDetails: %w", err)
			}
			if r.Revisions == nil {
				continue
			}
			for _, revision := range r.Revisions.AppLayerRevisionDetail {
				if revision.DisplayedVersion == appLayerVersionName {
					return revision.Id, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("app layer revision not found for name: %s and version: %s", appLayerName, appLayerVersionName)
}

func (s *SoapHelper) GetAppLayerId(appLayerName string) (int64, error) {
	appLayerQueryRequest := &QueryApplicationLayers{
		Query: &AppLayersQuery{
			Filter: appLayerName,
		},
	}
	queryAppLayersResponse, err := s.Client.QueryApplicationLayers(appLayerQueryRequest)
	if err != nil {
		return 0, fmt.Errorf("error calling QueryApplicationLayers: %v", err)
	}
	if queryAppLayersResponse.QueryApplicationLayersResult == nil {
		return 0, fmt.Errorf("app layer not found for name: %s", appLayerName)
	}
	if err := CheckWebResultError(queryAppLayersResponse.QueryApplicationLayersResult.WebResultBase); err != nil {
		return 0, fmt.Errorf("QueryApplicationLayers: %w", err)
	}
	if queryAppLayersResponse.QueryApplicationLayersResult.AppLayers == nil {
		return 0, fmt.Errorf("app layer not found for name: %s", appLayerName)
	}
	for _, appLayer := range queryAppLayersResponse.QueryApplicationLayersResult.AppLayers.LayerEntitySummary {
		if appLayer.Name == appLayerName {
			return appLayer.Id, nil
		}
	}
	return 0, fmt.Errorf("app layer not found for name: %s", appLayerName)
}

func (s *SoapHelper) GetDefaultFileShareId() (int64, error) {
	log.Printf("[TRACE] GetDefaultFileShareId: calling QueryRemoteFileShares")
	queryResult, err := s.Client.QueryRemoteFileShares(&QueryRemoteFileShares{Query: &RemoteFileSharesQuery{}})
	if err != nil {
		log.Printf("[TRACE] GetDefaultFileShareId: QueryRemoteFileShares error: %v", err)
		return 0, fmt.Errorf("error calling QueryRemoteFileShares: %v", err)
	}
	result := queryResult.QueryRemoteFileSharesResult
	if result == nil {
		log.Printf("[TRACE] GetDefaultFileShareId: QueryRemoteFileSharesResult is nil")
		return 0, fmt.Errorf("no file shares found; please configure at least one file share in ELM")
	}
	if err := CheckWebResultError(result.WebResultBase); err != nil {
		return 0, fmt.Errorf("QueryRemoteFileShares: %w", err)
	}
	if result.RemoteShares == nil {
		log.Printf("[TRACE] GetDefaultFileShareId: RemoteShares is nil")
		return 0, fmt.Errorf("no file shares found; please configure at least one file share in ELM")
	}
	shares := result.RemoteShares.RemoteFileShareSummary
	log.Printf("[TRACE] GetDefaultFileShareId: found %d share(s)", len(shares))
	for i, s := range shares {
		log.Printf("[TRACE] GetDefaultFileShareId: share[%d] Id=%d Path=%s", i, s.Id, s.SharePath)
	}
	if len(shares) == 0 {
		return 0, fmt.Errorf("no file shares found; please configure at least one file share in ELM")
	}
	// RemoteFileShareSummary has no IsDefault field; return the first share found.
	log.Printf("[TRACE] GetDefaultFileShareId: using share Id=%d Path=%s", shares[0].Id, shares[0].SharePath)
	return shares[0].Id, nil
}

// GetLatestAppLayerRevision returns the most recent (highest Revision number)
// revision detail for the named app layer. Used for optional base_version_name
// and size defaulting when base version is auto-detected.
func (s *SoapHelper) GetLatestAppLayerRevision(appLayerName string) (*AppLayerRevisionDetail, error) {
	queryResponse, err := s.Client.QueryApplicationLayers(&QueryApplicationLayers{
		Query: &AppLayersQuery{Filter: appLayerName},
	})
	if err != nil {
		return nil, fmt.Errorf("error calling QueryApplicationLayers: %v", err)
	}
	if queryResponse.QueryApplicationLayersResult == nil {
		return nil, fmt.Errorf("app layer not found for name: %s", appLayerName)
	}
	if err := CheckWebResultError(queryResponse.QueryApplicationLayersResult.WebResultBase); err != nil {
		return nil, fmt.Errorf("QueryApplicationLayers: %w", err)
	}
	if queryResponse.QueryApplicationLayersResult.AppLayers == nil {
		return nil, fmt.Errorf("app layer not found for name: %s", appLayerName)
	}
	for _, appLayer := range queryResponse.QueryApplicationLayersResult.AppLayers.LayerEntitySummary {
		if appLayer.Name == appLayerName {
			detailsResponse, err := s.Client.QueryApplicationLayerDetails(&QueryApplicationLayerDetails{
				Query: &LayerDetailsQuery{Id: appLayer.Id},
			})
			if err != nil {
				return nil, fmt.Errorf("error calling QueryApplicationLayerDetails: %v", err)
			}
			r := detailsResponse.QueryApplicationLayerDetailsResult
			if r == nil || r.LayerDetailsResultOfAppLayerRevisionDetail == nil {
				continue
			}
			if err := CheckWebResultError(r.LayerDetailsResultOfAppLayerRevisionDetail.WebResultBase); err != nil {
				return nil, fmt.Errorf("QueryApplicationLayerDetails: %w", err)
			}
			revisions := r.Revisions
			if revisions == nil || len(revisions.AppLayerRevisionDetail) == 0 {
				return nil, fmt.Errorf("no revisions found for app layer: %s", appLayerName)
			}
			var latest *AppLayerRevisionDetail
			for _, rev := range revisions.AppLayerRevisionDetail {
				if latest == nil || rev.Revision > latest.Revision {
					latest = rev
				}
			}
			return latest, nil
		}
	}
	return nil, fmt.Errorf("app layer not found for name: %s", appLayerName)
}

// GetPlatformLayerRevisionDetailByName returns the revision detail for the named
// platform layer at the specified version. Used to read base version
// platform type IDs as defaults when revision fields are omitted.
func (s *SoapHelper) GetPlatformLayerRevisionDetailByName(platformLayerName, versionName string) (*PlatformLayerRevisionDetail, error) {
	queryResponse, err := s.Client.QueryPlatformLayers(&QueryPlatformLayers{
		Query: &PlatformLayersQuery{Filter: platformLayerName},
	})
	if err != nil {
		return nil, fmt.Errorf("error calling QueryPlatformLayers: %v", err)
	}
	if queryResponse.QueryPlatformLayersResult == nil {
		return nil, fmt.Errorf("platform layer revision not found for name: %s and version: %s", platformLayerName, versionName)
	}
	if err := CheckWebResultError(queryResponse.QueryPlatformLayersResult.WebResultBase); err != nil {
		return nil, fmt.Errorf("QueryPlatformLayers: %w", err)
	}
	if queryResponse.QueryPlatformLayersResult.PlatformLayers == nil {
		return nil, fmt.Errorf("platform layer revision not found for name: %s and version: %s", platformLayerName, versionName)
	}
	for _, layer := range queryResponse.QueryPlatformLayersResult.PlatformLayers.LayerEntitySummary {
		if layer.Name == platformLayerName {
			detailsResponse, err := s.Client.QueryPlatformLayerDetails(&QueryPlatformLayerDetails{
				Query: &LayerDetailsQuery{Id: layer.Id},
			})
			if err != nil {
				return nil, fmt.Errorf("error calling QueryPlatformLayerDetails: %v", err)
			}
			r := detailsResponse.QueryPlatformLayerDetailsResult
			if r == nil || r.LayerDetailsResultOfPlatformLayerRevisionDetail == nil {
				continue
			}
			if err := CheckWebResultError(r.LayerDetailsResultOfPlatformLayerRevisionDetail.WebResultBase); err != nil {
				return nil, fmt.Errorf("QueryPlatformLayerDetails: %w", err)
			}
			if r.Revisions == nil {
				continue
			}
			for _, rev := range r.Revisions.PlatformLayerRevisionDetail {
				if rev.DisplayedVersion == versionName {
					return rev, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("platform layer revision not found for name: %s and version: %s", platformLayerName, versionName)
}

// GetOsLayerRevisionSizeMiB returns the SizeMegs field of the named OS layer
// revision. Used to default version_size_gb when not specified.
func (s *SoapHelper) GetOsLayerRevisionSizeMiB(layerName, versionName string) (int32, error) {
	queryResponse, err := s.Client.QueryOsLayers(&QueryOsLayers{
		Query: &OsLayersQuery{Filter: layerName},
	})
	if err != nil {
		return 0, fmt.Errorf("error calling QueryOsLayers: %v", err)
	}
	if queryResponse.QueryOsLayersResult == nil {
		return 0, fmt.Errorf("os layer revision not found for name: %s and version: %s", layerName, versionName)
	}
	if err := CheckWebResultError(queryResponse.QueryOsLayersResult.WebResultBase); err != nil {
		return 0, fmt.Errorf("QueryOsLayers: %w", err)
	}
	if queryResponse.QueryOsLayersResult.OsLayers == nil {
		return 0, fmt.Errorf("os layer revision not found for name: %s and version: %s", layerName, versionName)
	}
	for _, layer := range queryResponse.QueryOsLayersResult.OsLayers.LayerEntitySummary {
		if layer.Name == layerName {
			detailsResponse, err := s.Client.QueryOsLayerDetails(&QueryOsLayerDetails{
				Query: &LayerDetailsQuery{Id: layer.Id},
			})
			if err != nil {
				return 0, fmt.Errorf("error calling QueryOsLayerDetails: %v", err)
			}
			r := detailsResponse.QueryOsLayerDetailsResult
			if r == nil || r.LayerDetailsResultOfOsLayerRevisionDetail == nil {
				continue
			}
			if err := CheckWebResultError(r.LayerDetailsResultOfOsLayerRevisionDetail.WebResultBase); err != nil {
				return 0, fmt.Errorf("QueryOsLayerDetails: %w", err)
			}
			if r.Revisions == nil {
				continue
			}
			for _, rev := range r.Revisions.OsLayerRevisionDetail {
				if rev.DisplayedVersion == versionName {
					return rev.SizeMegs, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("os layer revision not found for name: %s and version: %s", layerName, versionName)
}

// GetPlatformLayerRevisionSizeMiB returns the SizeMegs field of the named platform
// layer revision. Used to default version_size_gb when not specified.
func (s *SoapHelper) GetPlatformLayerRevisionSizeMiB(layerName, versionName string) (int32, error) {
	queryResponse, err := s.Client.QueryPlatformLayers(&QueryPlatformLayers{
		Query: &PlatformLayersQuery{Filter: layerName},
	})
	if err != nil {
		return 0, fmt.Errorf("error calling QueryPlatformLayers: %v", err)
	}
	if queryResponse.QueryPlatformLayersResult == nil {
		return 0, fmt.Errorf("platform layer revision not found for name: %s and version: %s", layerName, versionName)
	}
	if err := CheckWebResultError(queryResponse.QueryPlatformLayersResult.WebResultBase); err != nil {
		return 0, fmt.Errorf("QueryPlatformLayers: %w", err)
	}
	if queryResponse.QueryPlatformLayersResult.PlatformLayers == nil {
		return 0, fmt.Errorf("platform layer revision not found for name: %s and version: %s", layerName, versionName)
	}
	for _, layer := range queryResponse.QueryPlatformLayersResult.PlatformLayers.LayerEntitySummary {
		if layer.Name == layerName {
			detailsResponse, err := s.Client.QueryPlatformLayerDetails(&QueryPlatformLayerDetails{
				Query: &LayerDetailsQuery{Id: layer.Id},
			})
			if err != nil {
				return 0, fmt.Errorf("error calling QueryPlatformLayerDetails: %v", err)
			}
			r := detailsResponse.QueryPlatformLayerDetailsResult
			if r == nil || r.LayerDetailsResultOfPlatformLayerRevisionDetail == nil {
				continue
			}
			if err := CheckWebResultError(r.LayerDetailsResultOfPlatformLayerRevisionDetail.WebResultBase); err != nil {
				return 0, fmt.Errorf("QueryPlatformLayerDetails: %w", err)
			}
			if r.Revisions == nil {
				continue
			}
			for _, rev := range r.Revisions.PlatformLayerRevisionDetail {
				if rev.DisplayedVersion == versionName {
					return rev.SizeMegs, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("platform layer revision not found for name: %s and version: %s", layerName, versionName)
}

// GetAppLayerRevisionSizeMiB returns the SizeMegs field of the named app layer
// revision. Used to default version_size_gb when not specified.
func (s *SoapHelper) GetAppLayerRevisionSizeMiB(layerName, versionName string) (int32, error) {
	queryResponse, err := s.Client.QueryApplicationLayers(&QueryApplicationLayers{
		Query: &AppLayersQuery{Filter: layerName},
	})
	if err != nil {
		return 0, fmt.Errorf("error calling QueryApplicationLayers: %v", err)
	}
	if queryResponse.QueryApplicationLayersResult == nil {
		return 0, fmt.Errorf("app layer revision not found for name: %s and version: %s", layerName, versionName)
	}
	if err := CheckWebResultError(queryResponse.QueryApplicationLayersResult.WebResultBase); err != nil {
		return 0, fmt.Errorf("QueryApplicationLayers: %w", err)
	}
	if queryResponse.QueryApplicationLayersResult.AppLayers == nil {
		return 0, fmt.Errorf("app layer revision not found for name: %s and version: %s", layerName, versionName)
	}
	for _, appLayer := range queryResponse.QueryApplicationLayersResult.AppLayers.LayerEntitySummary {
		if appLayer.Name == layerName {
			detailsResponse, err := s.Client.QueryApplicationLayerDetails(&QueryApplicationLayerDetails{
				Query: &LayerDetailsQuery{Id: appLayer.Id},
			})
			if err != nil {
				return 0, fmt.Errorf("error calling QueryApplicationLayerDetails: %v", err)
			}
			r := detailsResponse.QueryApplicationLayerDetailsResult
			if r == nil || r.LayerDetailsResultOfAppLayerRevisionDetail == nil {
				continue
			}
			if err := CheckWebResultError(r.LayerDetailsResultOfAppLayerRevisionDetail.WebResultBase); err != nil {
				return 0, fmt.Errorf("QueryApplicationLayerDetails: %w", err)
			}
			if r.Revisions == nil {
				continue
			}
			for _, rev := range r.Revisions.AppLayerRevisionDetail {
				if rev.DisplayedVersion == versionName {
					return rev.SizeMegs, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("app layer revision not found for name: %s and version: %s", layerName, versionName)
}

func (s *SoapHelper) GetWorkTicketIdByOperationTypeAndLayerName(operationType ApplayeringOperationType, layerName string) (int64, error) {
	queryWorkTicketsRequest := &QueryWorkTickets{
		Query: &WorkTicketsQuery{
			Filter: &WorkTicketsQueryFilter{},
		},
	}
	queryWorkTicketsResponse, err := s.Client.QueryWorkTickets(queryWorkTicketsRequest)
	if err != nil {
		return 0, fmt.Errorf("error calling QueryWorkTickets: %v", err)
	}
	if queryWorkTicketsResponse.QueryWorkTicketsResult == nil {
		return 0, fmt.Errorf("no work ticket found for operation type: %s and layer name: %s", operationType, layerName)
	}
	if err := CheckWebResultError(queryWorkTicketsResponse.QueryWorkTicketsResult.WebResultBase); err != nil {
		return 0, fmt.Errorf("QueryWorkTickets: %w", err)
	}
	if queryWorkTicketsResponse.QueryWorkTicketsResult.WorkTickets == nil {
		return 0, fmt.Errorf("no work ticket found for operation type: %s and layer name: %s", operationType, layerName)
	}
	for _, workTicket := range queryWorkTicketsResponse.QueryWorkTicketsResult.WorkTickets.WorkTicketResult {
		for _, workItem := range workTicket.WorkItems.WorkItemResult {
			if *workItem.State == WorkItemStateActionRequired {
				return workItem.Id, nil
			}
		}
	}
	return 0, fmt.Errorf("no work ticket found for operation type: %s and layer name: %s", operationType, layerName)
}

func (s *SoapHelper) GetWorkTicketId(ui packer.Ui, operationType ApplayeringOperationType, layerName string) (int64, error) {
	queryWorkTicketsAsPendingOp := &QueryWorkTicketsAsPendingOp{
		Query: &WorkTicketsQuery{
			Filter: &WorkTicketsQueryFilter{
				XsiType: "WorkTicketsQueryByActiveFilter",
			},
		},
	}
	queryWorkTicketsAsPendingOpResponse, err := s.Client.QueryWorkTicketsAsPendingOp(queryWorkTicketsAsPendingOp)
	if err != nil {
		return 0, err
	}
	if r := queryWorkTicketsAsPendingOpResponse.QueryWorkTicketsAsPendingOpResult; r != nil {
		if err := CheckWebResultError(r.WebResultBase); err != nil {
			return 0, fmt.Errorf("QueryWorkTicketsAsPendingOp: %w", err)
		}
	}
	operationResult := queryWorkTicketsAsPendingOpResponse.QueryWorkTicketsAsPendingOpResult.OperationResult
	if operationResult == nil {
		return 0, fmt.Errorf("no operation result found for operation type: %s and layer name: %s", operationType, layerName)
	}
	operationToWorkTicket := map[ApplayeringOperationType]string{
		REVISION_OS_LAYER:       "CreateOsLayerRevisionWorkTicketDescription",
		REVISION_PLATFORM_LAYER: "CreatePlatformLayerRevisionWorkTicketDescription",
		REVISION_APP_LAYER:      "CreateAppLayerRevisionWorkTicketDescription",
		CREATE_PLATFORM_LAYER:   "CreatePlatformLayerWorkTicketDescription",
		CREATE_APP_LAYER:        "CreateAppLayerWorkTicketDescription",
	}
	for _, workTicket := range operationResult.WorkTickets.WorkTicketResult {
		if workTicket.TitleResourceId == operationToWorkTicket[operationType] || *workTicket.ResourceArgs.Astring[0] == layerName {
			return workTicket.Id, nil
		}
	}
	return 0, fmt.Errorf("no operation result found for operation type: %s and layer name: %s", operationType, layerName)
}

func (s *SoapHelper) GetIpByWorkTicketId(workTicketId int64) (string, error) {
	log.Printf("GetIpByWorkTicketId:: Try to get IP for work ticket id: %d", workTicketId)
	task, err := s.GetTaskByWorkTicketId(workTicketId)
	if err != nil {
		return "", err
	}
	ip := ""

	re := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
	for _, workitem := range task.WorkItems.WorkItemResult {
		if *workitem.ItemType == "WorkItem" {
			continue
		}
		if *workitem.State != WorkItemStateActionRequired {
			log.Printf("GetIpByWorkTicketId:: Work Ticket Id %d :: Work Item ID %d 's state is %s, not in ActionRequired state", workTicketId, workitem.Id, *workitem.State)
			continue
		}
		ips := re.FindAllString(workitem.Status, -1)
		if len(ips) > 0 {
			ip = ips[0]
			return ip, nil
		}
	}
	return "", nil
}

func (s *SoapHelper) GetTaskStateActiveFilter(workTicketId int64) (string, error) {
	log.Printf("GetTaskStateActiveFilter start...")
	task, err := s.GetTaskByWorkTicketId(workTicketId)
	if err != nil {
		return "", err
	}
	state := string(*task.State)
	log.Printf("GetTaskStateActiveFilter end. state: %s", state)
	return state, nil
}

func (s *SoapHelper) queryWorkTicketByFilter(workTicketId int64, filterType string, notFoundErr error) (WorkTicketResult, error) {
	soapRequest := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:s1="http://microsoft.com/wsdl/types/" xmlns:tns="http://www.unidesk.com/" xmlns:tm="http://microsoft.com/wsdl/mime/textMatching/">
	<soap:Body xmlns:xsd="http://www.w3.org/2001/XMLSchema">
		<QueryWorkTicketsAsPendingOp xmlns="http://www.unidesk.com/">
			<query>
				<Filter xsi:type="%s"/>
			</query>
		</QueryWorkTicketsAsPendingOp>
	</soap:Body>
</soap:Envelope>`, filterType)
	soapAction := `http://www.unidesk.com/QueryWorkTicketsAsPendingOp`

	req, _ := http.NewRequest("POST", s.URL, bytes.NewBufferString(soapRequest))
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", soapAction)
	req.Header.Set("Cookie", s.Cookie)
	req.Header.Set("Unidesk_token", s.Token)

	client := &http.Client{
		Transport: &HeaderCaptureTransport{
			Rt: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: s.InsecureSkipVerify,
				},
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return WorkTicketResult{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("queryWorkTicketByFilter(%s):: read body failed. %v", filterType, err)
		return WorkTicketResult{}, err
	}
	var result SoapEnvelopeQueryWorkTicketsAsPendingOpResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return WorkTicketResult{}, err
	}
	for _, workTicket := range result.Body.QueryWorkTicketsAsPendingOpResponse.QueryWorkTicketsAsPendingOpResult.OperationResult.WorkTickets.WorkTicketResult {
		if workTicket.Id == workTicketId {
			return *workTicket, nil
		}
	}
	return WorkTicketResult{}, notFoundErr
}

func (s *SoapHelper) GetTaskByWorkTicketId(workTicketId int64) (WorkTicketResult, error) {
	log.Printf("GetTaskByWorkTicketId start...")
	return s.queryWorkTicketByFilter(workTicketId, "WorkTicketsQueryByActiveFilter", ErrWorkTicketNotInActiveFilter)
}

func (s *SoapHelper) GetTaskCompletedFilter(workTicketId int64) (WorkTicketResult, error) {
	log.Printf("GetTaskCompletedFilter start...")
	result, err := s.queryWorkTicketByFilter(workTicketId, "WorkTicketsQueryByCompletedFilter", fmt.Errorf("work ticket %d not found", workTicketId))
	if err == nil {
		log.Printf("GetTaskCompletedFilter end. workTicketResult:%v", result)
	}
	return result, err
}

func BuildServerURL(input string) (string, error) {
	const defaultScheme = "https"
	const defaultPath = "/Unidesk.Web/API.asmx"

	// If scheme is missing, prepend one
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
		input = defaultScheme + "://" + input
	}

	// Parse the input as URL
	parsedURL, err := url.Parse(input)
	if err != nil {
		return "", fmt.Errorf("invalid server input: %w", err)
	}

	// Set path if it's not already there
	if parsedURL.Path == "" || parsedURL.Path == "/" {
		parsedURL.Path = defaultPath
	}

	return parsedURL.String(), nil
}

func Login2(username, password, unideskurl string, insecureSkipVerify bool) (string, string, error) {
	log.Printf("Login start...")
	soapRequest := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:s1="http://microsoft.com/wsdl/types/" xmlns:tns="http://www.unidesk.com/" xmlns:tm="http://microsoft.com/wsdl/mime/textMatching/">
	<soap:Body xmlns:xsd="http://www.w3.org/2001/XMLSchema">
		<Login xmlns="http://www.unidesk.com/">
			<command>
				<UserName>%s</UserName>
				<Password>%s</Password>
				<Culture>en-US</Culture>
				<RememberMe>false</RememberMe>
			</command>
		</Login>
	</soap:Body>
</soap:Envelope>`, username, password)
	req, _ := http.NewRequest("POST", unideskurl, bytes.NewBufferString(soapRequest))

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "http://www.unidesk.com/Login")

	client := &http.Client{
		Transport: &HeaderCaptureTransport{
			Rt: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecureSkipVerify,
				},
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	setCookies, ok := resp.Header["Set-Cookie"]
	if !ok {
		log.Printf("set-cookie cannot be found in response header")
		err = fmt.Errorf("set-cookie cannot be found in response header")
		return "", "", err
	}
	cookie := ""
	if cookies := setCookies; len(cookies) > 0 {
		cookie = strings.Split(cookies[0], ";")[0]
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read login response body:%v", err)
		return "", "", err
	}

	var result SoapEnvelopeLoginResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return "", "", err
	}
	token := result.Body.LoginResponse.LoginResult.Token
	return cookie, token, nil
}

// CancelWorkTicket requests cancellation of a running ELM work ticket.
// It is a best-effort call; errors are logged but not fatal.
func (s *SoapHelper) CancelWorkTicket(workTicketId int64) error {
	ids := &ArrayOfLong{Long: []int64{workTicketId}}
	req := &CancelWorkTickets{
		Command: &CancelWorkTicketsCommand{
			WorkTicketIds: ids,
		},
	}
	resp, err := s.Client.CancelWorkTickets(req)
	if err != nil {
		return fmt.Errorf("CancelWorkTicket(%d): %w", workTicketId, err)
	}
	if resp.CancelWorkTicketsResult != nil {
		if err := CheckWebResultError(resp.CancelWorkTicketsResult.WebResultBase); err != nil {
			return fmt.Errorf("CancelWorkTicket(%d): %w", workTicketId, err)
		}
	}
	return nil
}
