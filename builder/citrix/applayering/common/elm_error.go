package common

import elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"

func FormatELMError(e *elmsoap.ApplicationError) string {
	if e != nil && e.Message != "" {
		return e.Message
	}
	return "no error details available"
}

// GetCreateLayerResultError safely extracts the ApplicationError from a CreateLayerResult,
// handling nil embedded pointers (WebResultBase/ResultBase) that Go's XML decoder
// does not allocate when the corresponding XML elements are absent.
func GetCreateLayerResultError(r *elmsoap.CreateLayerResult) *elmsoap.ApplicationError {
	if r == nil || r.WebResultBase == nil || r.WebResultBase.ResultBase == nil {
		return nil
	}
	return r.WebResultBase.ResultBase.Error
}

// GetRevisionResultError safely extracts the ApplicationError from a CreateRevisionResult.
func GetRevisionResultError(r *elmsoap.CreateRevisionResult) *elmsoap.ApplicationError {
	if r == nil || r.WebResultBase == nil || r.WebResultBase.ResultBase == nil {
		return nil
	}
	return r.WebResultBase.ResultBase.Error
}
