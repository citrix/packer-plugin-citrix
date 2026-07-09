// Copyright (c) Citrix, Inc.

package elmsoap

// baseApiSoap satisfies the very large ApiSoap interface by panicking
// on every method. Per-test mocks should embed *baseApiSoap and override only
// the SOAP methods they exercise.
//
// This file is auto-generated test scaffolding — DO NOT edit method stubs by
// hand. If the ApiSoap interface changes, regenerate from elm-client/elm_client.go:
//
//   awk '/^type ApiSoap interface/{flag=1;next} flag && /^}/{exit} flag && /^\t[A-Z]/ {...}' \
//       elm-client/elm_client.go
//
// Keeping it in *_test.go ensures it never ships in the production binary.

import (
	"context"

	"github.com/hooklift/gowsdl/soap"
)

// Compile-time assertion that *baseApiSoap satisfies ApiSoap.
var _ ApiSoap = (*baseApiSoap)(nil)

// baseApiSoap is the empty receiver that holds all panic stubs below.
type baseApiSoap struct{}

// Suppress "imported and not used" if the generated stub set ever shrinks.
// context.Context and *soap.Client both appear in ApiSoap method signatures.
var (
	_ = context.TODO
	_ *soap.Client
)

func (b *baseApiSoap) Login(request *Login) (*LoginResponse, error) {
	panic("baseApiSoap.Login not implemented in test — embed and override this method")
}

func (b *baseApiSoap) LoginContext(ctx context.Context, request *Login) (*LoginResponse, error) {
	panic("baseApiSoap.LoginContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) Logout(request *Logout) (*LogoutResponse, error) {
	panic("baseApiSoap.Logout not implemented in test — embed and override this method")
}

func (b *baseApiSoap) LogoutContext(ctx context.Context, request *Logout) (*LogoutResponse, error) {
	panic("baseApiSoap.LogoutContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) IsAuthenticated(request *IsAuthenticated) (*IsAuthenticatedResponse, error) {
	panic("baseApiSoap.IsAuthenticated not implemented in test — embed and override this method")
}

func (b *baseApiSoap) IsAuthenticatedContext(ctx context.Context, request *IsAuthenticated) (*IsAuthenticatedResponse, error) {
	panic("baseApiSoap.IsAuthenticatedContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) Log(request *Log) (*LogResponse, error) {
	panic("baseApiSoap.Log not implemented in test — embed and override this method")
}

func (b *baseApiSoap) LogContext(ctx context.Context, request *Log) (*LogResponse, error) {
	panic("baseApiSoap.LogContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementAppliance(request *QueryManagementAppliance) (*QueryManagementApplianceResponse, error) {
	panic("baseApiSoap.QueryManagementAppliance not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementApplianceContext(ctx context.Context, request *QueryManagementAppliance) (*QueryManagementApplianceResponse, error) {
	panic("baseApiSoap.QueryManagementApplianceContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementApplianceDetails(request *QueryManagementApplianceDetails) (*QueryManagementApplianceDetailsResponse, error) {
	panic("baseApiSoap.QueryManagementApplianceDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementApplianceDetailsContext(ctx context.Context, request *QueryManagementApplianceDetails) (*QueryManagementApplianceDetailsResponse, error) {
	panic("baseApiSoap.QueryManagementApplianceDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementApplianceDetailsWithParam(request *QueryManagementApplianceDetailsWithParam) (*QueryManagementApplianceDetailsWithParamResponse, error) {
	panic("baseApiSoap.QueryManagementApplianceDetailsWithParam not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementApplianceDetailsWithParamContext(ctx context.Context, request *QueryManagementApplianceDetailsWithParam) (*QueryManagementApplianceDetailsWithParamResponse, error) {
	panic("baseApiSoap.QueryManagementApplianceDetailsWithParamContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryActivity(request *QueryActivity) (*QueryActivityResponse, error) {
	panic("baseApiSoap.QueryActivity not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryActivityContext(ctx context.Context, request *QueryActivity) (*QueryActivityResponse, error) {
	panic("baseApiSoap.QueryActivityContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) BrowseContainer(request *BrowseContainer) (*BrowseContainerResponse, error) {
	panic("baseApiSoap.BrowseContainer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) BrowseContainerContext(ctx context.Context, request *BrowseContainer) (*BrowseContainerResponse, error) {
	panic("baseApiSoap.BrowseContainerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RegisterImportableVM(request *RegisterImportableVM) (*RegisterImportableVMResponse, error) {
	panic("baseApiSoap.RegisterImportableVM not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RegisterImportableVMContext(ctx context.Context, request *RegisterImportableVM) (*RegisterImportableVMResponse, error) {
	panic("baseApiSoap.RegisterImportableVMContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteGoldImage(request *DeleteGoldImage) (*DeleteGoldImageResponse, error) {
	panic("baseApiSoap.DeleteGoldImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteGoldImageContext(ctx context.Context, request *DeleteGoldImage) (*DeleteGoldImageResponse, error) {
	panic("baseApiSoap.DeleteGoldImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUsers(request *QueryUsers) (*QueryUsersResponse, error) {
	panic("baseApiSoap.QueryUsers not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUsersContext(ctx context.Context, request *QueryUsers) (*QueryUsersResponse, error) {
	panic("baseApiSoap.QueryUsersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryGroups(request *QueryGroups) (*QueryGroupsResponse, error) {
	panic("baseApiSoap.QueryGroups not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryGroupsContext(ctx context.Context, request *QueryGroups) (*QueryGroupsResponse, error) {
	panic("baseApiSoap.QueryGroupsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUpgradeAvailability(request *QueryUpgradeAvailability) (*QueryUpgradeAvailabilityResponse, error) {
	panic("baseApiSoap.QueryUpgradeAvailability not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUpgradeAvailabilityContext(ctx context.Context, request *QueryUpgradeAvailability) (*QueryUpgradeAvailabilityResponse, error) {
	panic("baseApiSoap.QueryUpgradeAvailabilityContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ClearUpgradeStatus(request *ClearUpgradeStatus) (*ClearUpgradeStatusResponse, error) {
	panic("baseApiSoap.ClearUpgradeStatus not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ClearUpgradeStatusContext(ctx context.Context, request *ClearUpgradeStatus) (*ClearUpgradeStatusResponse, error) {
	panic("baseApiSoap.ClearUpgradeStatusContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemAppAssignments(request *QueryDirectoryItemAppAssignments) (*QueryDirectoryItemAppAssignmentsResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemAppAssignments not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemAppAssignmentsContext(ctx context.Context, request *QueryDirectoryItemAppAssignments) (*QueryDirectoryItemAppAssignmentsResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemAppAssignmentsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemDetails(request *QueryDirectoryItemDetails) (*QueryDirectoryItemDetailsResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemDetailsContext(ctx context.Context, request *QueryDirectoryItemDetails) (*QueryDirectoryItemDetailsResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SearchDirectoryItem(request *SearchDirectoryItem) (*SearchDirectoryItemResponse, error) {
	panic("baseApiSoap.SearchDirectoryItem not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SearchDirectoryItemContext(ctx context.Context, request *SearchDirectoryItem) (*SearchDirectoryItemResponse, error) {
	panic("baseApiSoap.SearchDirectoryItemContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SearchDirectoryItemPendingOp(request *SearchDirectoryItemPendingOp) (*SearchDirectoryItemPendingOpResponse, error) {
	panic("baseApiSoap.SearchDirectoryItemPendingOp not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SearchDirectoryItemPendingOpContext(ctx context.Context, request *SearchDirectoryItemPendingOp) (*SearchDirectoryItemPendingOpResponse, error) {
	panic("baseApiSoap.SearchDirectoryItemPendingOpContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryWorkTicketsAsPendingOp(request *QueryWorkTicketsAsPendingOp) (*QueryWorkTicketsAsPendingOpResponse, error) {
	panic("baseApiSoap.QueryWorkTicketsAsPendingOp not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryWorkTicketsAsPendingOpContext(ctx context.Context, request *QueryWorkTicketsAsPendingOp) (*QueryWorkTicketsAsPendingOpResponse, error) {
	panic("baseApiSoap.QueryWorkTicketsAsPendingOpContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateDirectoryItem(request *CreateDirectoryItem) (*CreateDirectoryItemResponse, error) {
	panic("baseApiSoap.CreateDirectoryItem not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateDirectoryItemContext(ctx context.Context, request *CreateDirectoryItem) (*CreateDirectoryItemResponse, error) {
	panic("baseApiSoap.CreateDirectoryItemContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryAuditLog(request *QueryAuditLog) (*QueryAuditLogResponse, error) {
	panic("baseApiSoap.QueryAuditLog not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryAuditLogContext(ctx context.Context, request *QueryAuditLog) (*QueryAuditLogResponse, error) {
	panic("baseApiSoap.QueryAuditLogContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryAuditLogDetail(request *QueryAuditLogDetail) (*QueryAuditLogDetailResponse, error) {
	panic("baseApiSoap.QueryAuditLogDetail not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryAuditLogDetailContext(ctx context.Context, request *QueryAuditLogDetail) (*QueryAuditLogDetailResponse, error) {
	panic("baseApiSoap.QueryAuditLogDetailContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditDirectoryItem(request *EditDirectoryItem) (*EditDirectoryItemResponse, error) {
	panic("baseApiSoap.EditDirectoryItem not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditDirectoryItemContext(ctx context.Context, request *EditDirectoryItem) (*EditDirectoryItemResponse, error) {
	panic("baseApiSoap.EditDirectoryItemContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteDirectoryItem(request *DeleteDirectoryItem) (*DeleteDirectoryItemResponse, error) {
	panic("baseApiSoap.DeleteDirectoryItem not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteDirectoryItemContext(ctx context.Context, request *DeleteDirectoryItem) (*DeleteDirectoryItemResponse, error) {
	panic("baseApiSoap.DeleteDirectoryItemContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteDirectoryJunction(request *DeleteDirectoryJunction) (*DeleteDirectoryJunctionResponse, error) {
	panic("baseApiSoap.DeleteDirectoryJunction not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteDirectoryJunctionContext(ctx context.Context, request *DeleteDirectoryJunction) (*DeleteDirectoryJunctionResponse, error) {
	panic("baseApiSoap.DeleteDirectoryJunctionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryInfrastructure(request *QueryInfrastructure) (*QueryInfrastructureResponse, error) {
	panic("baseApiSoap.QueryInfrastructure not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryInfrastructureContext(ctx context.Context, request *QueryInfrastructure) (*QueryInfrastructureResponse, error) {
	panic("baseApiSoap.QueryInfrastructureContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateInfrastructure(request *CreateInfrastructure) (*CreateInfrastructureResponse, error) {
	panic("baseApiSoap.CreateInfrastructure not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateInfrastructureContext(ctx context.Context, request *CreateInfrastructure) (*CreateInfrastructureResponse, error) {
	panic("baseApiSoap.CreateInfrastructureContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateInfrastructure(request *UpdateInfrastructure) (*UpdateInfrastructureResponse, error) {
	panic("baseApiSoap.UpdateInfrastructure not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateInfrastructureContext(ctx context.Context, request *UpdateInfrastructure) (*UpdateInfrastructureResponse, error) {
	panic("baseApiSoap.UpdateInfrastructureContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestInfrastructure(request *TestInfrastructure) (*TestInfrastructureResponse, error) {
	panic("baseApiSoap.TestInfrastructure not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestInfrastructureContext(ctx context.Context, request *TestInfrastructure) (*TestInfrastructureResponse, error) {
	panic("baseApiSoap.TestInfrastructureContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDataCenters(request *QueryDataCenters) (*QueryDataCentersResponse, error) {
	panic("baseApiSoap.QueryDataCenters not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDataCentersContext(ctx context.Context, request *QueryDataCenters) (*QueryDataCentersResponse, error) {
	panic("baseApiSoap.QueryDataCentersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryComputeResources(request *QueryComputeResources) (*QueryComputeResourcesResponse, error) {
	panic("baseApiSoap.QueryComputeResources not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryComputeResourcesContext(ctx context.Context, request *QueryComputeResources) (*QueryComputeResourcesResponse, error) {
	panic("baseApiSoap.QueryComputeResourcesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryEnvironment(request *QueryEnvironment) (*QueryEnvironmentResponse, error) {
	panic("baseApiSoap.QueryEnvironment not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryEnvironmentContext(ctx context.Context, request *QueryEnvironment) (*QueryEnvironmentResponse, error) {
	panic("baseApiSoap.QueryEnvironmentContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryWorkTickets(request *QueryWorkTickets) (*QueryWorkTicketsResponse, error) {
	panic("baseApiSoap.QueryWorkTickets not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryWorkTicketsContext(ctx context.Context, request *QueryWorkTickets) (*QueryWorkTicketsResponse, error) {
	panic("baseApiSoap.QueryWorkTicketsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableVMs(request *QueryImportableVMs) (*QueryImportableVMsResponse, error) {
	panic("baseApiSoap.QueryImportableVMs not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableVMsContext(ctx context.Context, request *QueryImportableVMs) (*QueryImportableVMsResponse, error) {
	panic("baseApiSoap.QueryImportableVMsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerIcons(request *QueryLayerIcons) (*QueryLayerIconsResponse, error) {
	panic("baseApiSoap.QueryLayerIcons not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerIconsContext(ctx context.Context, request *QueryLayerIcons) (*QueryLayerIconsResponse, error) {
	panic("baseApiSoap.QueryLayerIconsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelWorkTickets(request *CancelWorkTickets) (*CancelWorkTicketsResponse, error) {
	panic("baseApiSoap.CancelWorkTickets not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelWorkTicketsContext(ctx context.Context, request *CancelWorkTickets) (*CancelWorkTicketsResponse, error) {
	panic("baseApiSoap.CancelWorkTicketsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelWorkItems(request *CancelWorkItems) (*CancelWorkItemsResponse, error) {
	panic("baseApiSoap.CancelWorkItems not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelWorkItemsContext(ctx context.Context, request *CancelWorkItems) (*CancelWorkItemsResponse, error) {
	panic("baseApiSoap.CancelWorkItemsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPendingOperation(request *QueryPendingOperation) (*QueryPendingOperationResponse, error) {
	panic("baseApiSoap.QueryPendingOperation not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPendingOperationContext(ctx context.Context, request *QueryPendingOperation) (*QueryPendingOperationResponse, error) {
	panic("baseApiSoap.QueryPendingOperationContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelPendingOperation(request *CancelPendingOperation) (*CancelPendingOperationResponse, error) {
	panic("baseApiSoap.CancelPendingOperation not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelPendingOperationContext(ctx context.Context, request *CancelPendingOperation) (*CancelPendingOperationResponse, error) {
	panic("baseApiSoap.CancelPendingOperationContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportOs(request *ImportOs) (*ImportOsResponse, error) {
	panic("baseApiSoap.ImportOs not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportOsContext(ctx context.Context, request *ImportOs) (*ImportOsResponse, error) {
	panic("baseApiSoap.ImportOsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryOsLayers(request *QueryOsLayers) (*QueryOsLayersResponse, error) {
	panic("baseApiSoap.QueryOsLayers not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryOsLayersContext(ctx context.Context, request *QueryOsLayers) (*QueryOsLayersResponse, error) {
	panic("baseApiSoap.QueryOsLayersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryOsLayerDetails(request *QueryOsLayerDetails) (*QueryOsLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryOsLayerDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryOsLayerDetailsContext(ctx context.Context, request *QueryOsLayerDetails) (*QueryOsLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryOsLayerDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerRevisions(request *QueryLayerRevisions) (*QueryLayerRevisionsResponse, error) {
	panic("baseApiSoap.QueryLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerRevisionsContext(ctx context.Context, request *QueryLayerRevisions) (*QueryLayerRevisionsResponse, error) {
	panic("baseApiSoap.QueryLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteOsLayerRevisions(request *DeleteOsLayerRevisions) (*DeleteOsLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeleteOsLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteOsLayerRevisionsContext(ctx context.Context, request *DeleteOsLayerRevisions) (*DeleteOsLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeleteOsLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteAppLayerRevisions(request *DeleteAppLayerRevisions) (*DeleteAppLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeleteAppLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteAppLayerRevisionsContext(ctx context.Context, request *DeleteAppLayerRevisions) (*DeleteAppLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeleteAppLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeletePlatformLayerRevisions(request *DeletePlatformLayerRevisions) (*DeletePlatformLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeletePlatformLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeletePlatformLayerRevisionsContext(ctx context.Context, request *DeletePlatformLayerRevisions) (*DeletePlatformLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeletePlatformLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateApplicationLayer(request *CreateApplicationLayer) (*CreateApplicationLayerResponse, error) {
	panic("baseApiSoap.CreateApplicationLayer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateApplicationLayerContext(ctx context.Context, request *CreateApplicationLayer) (*CreateApplicationLayerResponse, error) {
	panic("baseApiSoap.CreateApplicationLayerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CloneLayer(request *CloneLayer) (*CloneLayerResponse, error) {
	panic("baseApiSoap.CloneLayer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CloneLayerContext(ctx context.Context, request *CloneLayer) (*CloneLayerResponse, error) {
	panic("baseApiSoap.CloneLayerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreatePlatformLayer(request *CreatePlatformLayer) (*CreatePlatformLayerResponse, error) {
	panic("baseApiSoap.CreatePlatformLayer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreatePlatformLayerContext(ctx context.Context, request *CreatePlatformLayer) (*CreatePlatformLayerResponse, error) {
	panic("baseApiSoap.CreatePlatformLayerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) FinalizeLayerRevision(request *FinalizeLayerRevision) (*FinalizeLayerRevisionResponse, error) {
	panic("baseApiSoap.FinalizeLayerRevision not implemented in test — embed and override this method")
}

func (b *baseApiSoap) FinalizeLayerRevisionContext(ctx context.Context, request *FinalizeLayerRevision) (*FinalizeLayerRevisionResponse, error) {
	panic("baseApiSoap.FinalizeLayerRevisionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerInstallDisk(request *QueryLayerInstallDisk) (*QueryLayerInstallDiskResponse, error) {
	panic("baseApiSoap.QueryLayerInstallDisk not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerInstallDiskContext(ctx context.Context, request *QueryLayerInstallDisk) (*QueryLayerInstallDiskResponse, error) {
	panic("baseApiSoap.QueryLayerInstallDiskContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryApplicationLayers(request *QueryApplicationLayers) (*QueryApplicationLayersResponse, error) {
	panic("baseApiSoap.QueryApplicationLayers not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryApplicationLayersContext(ctx context.Context, request *QueryApplicationLayers) (*QueryApplicationLayersResponse, error) {
	panic("baseApiSoap.QueryApplicationLayersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryApplicationLayerDetails(request *QueryApplicationLayerDetails) (*QueryApplicationLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryApplicationLayerDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryApplicationLayerDetailsContext(ctx context.Context, request *QueryApplicationLayerDetails) (*QueryApplicationLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryApplicationLayerDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformLayers(request *QueryPlatformLayers) (*QueryPlatformLayersResponse, error) {
	panic("baseApiSoap.QueryPlatformLayers not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformLayersContext(ctx context.Context, request *QueryPlatformLayers) (*QueryPlatformLayersResponse, error) {
	panic("baseApiSoap.QueryPlatformLayersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformLayerDetails(request *QueryPlatformLayerDetails) (*QueryPlatformLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryPlatformLayerDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformLayerDetailsContext(ctx context.Context, request *QueryPlatformLayerDetails) (*QueryPlatformLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryPlatformLayerDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateIcon(request *CreateIcon) (*CreateIconResponse, error) {
	panic("baseApiSoap.CreateIcon not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateIconContext(ctx context.Context, request *CreateIcon) (*CreateIconResponse, error) {
	panic("baseApiSoap.CreateIconContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetItemsAssociatedWithIcon(request *GetItemsAssociatedWithIcon) (*GetItemsAssociatedWithIconResponse, error) {
	panic("baseApiSoap.GetItemsAssociatedWithIcon not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetItemsAssociatedWithIconContext(ctx context.Context, request *GetItemsAssociatedWithIcon) (*GetItemsAssociatedWithIconResponse, error) {
	panic("baseApiSoap.GetItemsAssociatedWithIconContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteIcon(request *DeleteIcon) (*DeleteIconResponse, error) {
	panic("baseApiSoap.DeleteIcon not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteIconContext(ctx context.Context, request *DeleteIcon) (*DeleteIconResponse, error) {
	panic("baseApiSoap.DeleteIconContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateAppLayerRevision(request *CreateAppLayerRevision) (*CreateAppLayerRevisionResponse, error) {
	panic("baseApiSoap.CreateAppLayerRevision not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateAppLayerRevisionContext(ctx context.Context, request *CreateAppLayerRevision) (*CreateAppLayerRevisionResponse, error) {
	panic("baseApiSoap.CreateAppLayerRevisionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreatePlatformLayerRevision(request *CreatePlatformLayerRevision) (*CreatePlatformLayerRevisionResponse, error) {
	panic("baseApiSoap.CreatePlatformLayerRevision not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreatePlatformLayerRevisionContext(ctx context.Context, request *CreatePlatformLayerRevision) (*CreatePlatformLayerRevisionResponse, error) {
	panic("baseApiSoap.CreatePlatformLayerRevisionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateOsLayerRevision(request *CreateOsLayerRevision) (*CreateOsLayerRevisionResponse, error) {
	panic("baseApiSoap.CreateOsLayerRevision not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateOsLayerRevisionContext(ctx context.Context, request *CreateOsLayerRevision) (*CreateOsLayerRevisionResponse, error) {
	panic("baseApiSoap.CreateOsLayerRevisionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpgradeElm(request *UpgradeElm) (*UpgradeElmResponse, error) {
	panic("baseApiSoap.UpgradeElm not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpgradeElmContext(ctx context.Context, request *UpgradeElm) (*UpgradeElmResponse, error) {
	panic("baseApiSoap.UpgradeElmContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUpgradeStatus(request *QueryUpgradeStatus) (*QueryUpgradeStatusResponse, error) {
	panic("baseApiSoap.QueryUpgradeStatus not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUpgradeStatusContext(ctx context.Context, request *QueryUpgradeStatus) (*QueryUpgradeStatusResponse, error) {
	panic("baseApiSoap.QueryUpgradeStatusContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetWelcomeWindowState(request *GetWelcomeWindowState) (*GetWelcomeWindowStateResponse, error) {
	panic("baseApiSoap.GetWelcomeWindowState not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetWelcomeWindowStateContext(ctx context.Context, request *GetWelcomeWindowState) (*GetWelcomeWindowStateResponse, error) {
	panic("baseApiSoap.GetWelcomeWindowStateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GatherDiagnostics(request *GatherDiagnostics) (*GatherDiagnosticsResponse, error) {
	panic("baseApiSoap.GatherDiagnostics not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GatherDiagnosticsContext(ctx context.Context, request *GatherDiagnostics) (*GatherDiagnosticsResponse, error) {
	panic("baseApiSoap.GatherDiagnosticsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportLdapItem(request *ImportLdapItem) (*ImportLdapItemResponse, error) {
	panic("baseApiSoap.ImportLdapItem not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportLdapItemContext(ctx context.Context, request *ImportLdapItem) (*ImportLdapItemResponse, error) {
	panic("baseApiSoap.ImportLdapItemContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetNotifications(request *GetNotifications) (*GetNotificationsResponse, error) {
	panic("baseApiSoap.GetNotifications not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetNotificationsContext(ctx context.Context, request *GetNotifications) (*GetNotificationsResponse, error) {
	panic("baseApiSoap.GetNotificationsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditLayer(request *EditLayer) (*EditLayerResponse, error) {
	panic("baseApiSoap.EditLayer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditLayerContext(ctx context.Context, request *EditLayer) (*EditLayerResponse, error) {
	panic("baseApiSoap.EditLayerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PreviewLicense(request *PreviewLicense) (*PreviewLicenseResponse, error) {
	panic("baseApiSoap.PreviewLicense not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PreviewLicenseContext(ctx context.Context, request *PreviewLicense) (*PreviewLicenseResponse, error) {
	panic("baseApiSoap.PreviewLicenseContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetLicense(request *GetLicense) (*GetLicenseResponse, error) {
	panic("baseApiSoap.GetLicense not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetLicenseContext(ctx context.Context, request *GetLicense) (*GetLicenseResponse, error) {
	panic("baseApiSoap.GetLicenseContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetLicense(request *SetLicense) (*SetLicenseResponse, error) {
	panic("baseApiSoap.SetLicense not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetLicenseContext(ctx context.Context, request *SetLicense) (*SetLicenseResponse, error) {
	panic("baseApiSoap.SetLicenseContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CurrentHttpCertificate(request *CurrentHttpCertificate) (*CurrentHttpCertificateResponse, error) {
	panic("baseApiSoap.CurrentHttpCertificate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CurrentHttpCertificateContext(ctx context.Context, request *CurrentHttpCertificate) (*CurrentHttpCertificateResponse, error) {
	panic("baseApiSoap.CurrentHttpCertificateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) StoreHttpCertificate(request *StoreHttpCertificate) (*StoreHttpCertificateResponse, error) {
	panic("baseApiSoap.StoreHttpCertificate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) StoreHttpCertificateContext(ctx context.Context, request *StoreHttpCertificate) (*StoreHttpCertificateResponse, error) {
	panic("baseApiSoap.StoreHttpCertificateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PreviewHttpCertificate(request *PreviewHttpCertificate) (*PreviewHttpCertificateResponse, error) {
	panic("baseApiSoap.PreviewHttpCertificate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PreviewHttpCertificateContext(ctx context.Context, request *PreviewHttpCertificate) (*PreviewHttpCertificateResponse, error) {
	panic("baseApiSoap.PreviewHttpCertificateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GenerateHttpCertificate(request *GenerateHttpCertificate) (*GenerateHttpCertificateResponse, error) {
	panic("baseApiSoap.GenerateHttpCertificate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GenerateHttpCertificateContext(ctx context.Context, request *GenerateHttpCertificate) (*GenerateHttpCertificateResponse, error) {
	panic("baseApiSoap.GenerateHttpCertificateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetStorageRecords(request *GetStorageRecords) (*GetStorageRecordsResponse, error) {
	panic("baseApiSoap.GetStorageRecords not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetStorageRecordsContext(ctx context.Context, request *GetStorageRecords) (*GetStorageRecordsResponse, error) {
	panic("baseApiSoap.GetStorageRecordsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageSummary(request *QueryImageSummary) (*QueryImageSummaryResponse, error) {
	panic("baseApiSoap.QueryImageSummary not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageSummaryContext(ctx context.Context, request *QueryImageSummary) (*QueryImageSummaryResponse, error) {
	panic("baseApiSoap.QueryImageSummaryContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageDetail(request *QueryImageDetail) (*QueryImageDetailResponse, error) {
	panic("baseApiSoap.QueryImageDetail not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageDetailContext(ctx context.Context, request *QueryImageDetail) (*QueryImageDetailResponse, error) {
	panic("baseApiSoap.QueryImageDetailContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageByLayerRevision(request *QueryImageByLayerRevision) (*QueryImageByLayerRevisionResponse, error) {
	panic("baseApiSoap.QueryImageByLayerRevision not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageByLayerRevisionContext(ctx context.Context, request *QueryImageByLayerRevision) (*QueryImageByLayerRevisionResponse, error) {
	panic("baseApiSoap.QueryImageByLayerRevisionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRemoteFileShares(request *QueryRemoteFileShares) (*QueryRemoteFileSharesResponse, error) {
	panic("baseApiSoap.QueryRemoteFileShares not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRemoteFileSharesContext(ctx context.Context, request *QueryRemoteFileShares) (*QueryRemoteFileSharesResponse, error) {
	panic("baseApiSoap.QueryRemoteFileSharesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteRemoteFileShares(request *DeleteRemoteFileShares) (*DeleteRemoteFileSharesResponse, error) {
	panic("baseApiSoap.DeleteRemoteFileShares not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteRemoteFileSharesContext(ctx context.Context, request *DeleteRemoteFileShares) (*DeleteRemoteFileSharesResponse, error) {
	panic("baseApiSoap.DeleteRemoteFileSharesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectors(request *QueryPlatformConnectors) (*QueryPlatformConnectorsResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectors not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorsContext(ctx context.Context, request *QueryPlatformConnectors) (*QueryPlatformConnectorsResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfig(request *QueryPlatformConnectorConfig) (*QueryPlatformConnectorConfigResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfig not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfigContext(ctx context.Context, request *QueryPlatformConnectorConfig) (*QueryPlatformConnectorConfigResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfigContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfigSummary(request *QueryPlatformConnectorConfigSummary) (*QueryPlatformConnectorConfigSummaryResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfigSummary not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfigSummaryContext(ctx context.Context, request *QueryPlatformConnectorConfigSummary) (*QueryPlatformConnectorConfigSummaryResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfigSummaryContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfigDetails(request *QueryPlatformConnectorConfigDetails) (*QueryPlatformConnectorConfigDetailsResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfigDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfigDetailsContext(ctx context.Context, request *QueryPlatformConnectorConfigDetails) (*QueryPlatformConnectorConfigDetailsResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfigDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectory(request *QueryDirectory) (*QueryDirectoryResponse, error) {
	panic("baseApiSoap.QueryDirectory not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryContext(ctx context.Context, request *QueryDirectory) (*QueryDirectoryResponse, error) {
	panic("baseApiSoap.QueryDirectoryContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetRemoteFileShare(request *SetRemoteFileShare) (*SetRemoteFileShareResponse, error) {
	panic("baseApiSoap.SetRemoteFileShare not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetRemoteFileShareContext(ctx context.Context, request *SetRemoteFileShare) (*SetRemoteFileShareResponse, error) {
	panic("baseApiSoap.SetRemoteFileShareContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestRemoteFileShare(request *TestRemoteFileShare) (*TestRemoteFileShareResponse, error) {
	panic("baseApiSoap.TestRemoteFileShare not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestRemoteFileShareContext(ctx context.Context, request *TestRemoteFileShare) (*TestRemoteFileShareResponse, error) {
	panic("baseApiSoap.TestRemoteFileShareContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetAvailableDrives(request *GetAvailableDrives) (*GetAvailableDrivesResponse, error) {
	panic("baseApiSoap.GetAvailableDrives not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetAvailableDrivesContext(ctx context.Context, request *GetAvailableDrives) (*GetAvailableDrivesResponse, error) {
	panic("baseApiSoap.GetAvailableDrivesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryCreate(request *DirectoryCreate) (*DirectoryCreateResponse, error) {
	panic("baseApiSoap.DirectoryCreate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryCreateContext(ctx context.Context, request *DirectoryCreate) (*DirectoryCreateResponse, error) {
	panic("baseApiSoap.DirectoryCreateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryUpdate(request *DirectoryUpdate) (*DirectoryUpdateResponse, error) {
	panic("baseApiSoap.DirectoryUpdate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryUpdateContext(ctx context.Context, request *DirectoryUpdate) (*DirectoryUpdateResponse, error) {
	panic("baseApiSoap.DirectoryUpdateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryDelete(request *DirectoryDelete) (*DirectoryDeleteResponse, error) {
	panic("baseApiSoap.DirectoryDelete not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryDeleteContext(ctx context.Context, request *DirectoryDelete) (*DirectoryDeleteResponse, error) {
	panic("baseApiSoap.DirectoryDeleteContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRecipes(request *QueryRecipes) (*QueryRecipesResponse, error) {
	panic("baseApiSoap.QueryRecipes not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRecipesContext(ctx context.Context, request *QueryRecipes) (*QueryRecipesResponse, error) {
	panic("baseApiSoap.QueryRecipesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypes(request *QueryPlatformTypes) (*QueryPlatformTypesResponse, error) {
	panic("baseApiSoap.QueryPlatformTypes not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypesContext(ctx context.Context, request *QueryPlatformTypes) (*QueryPlatformTypesResponse, error) {
	panic("baseApiSoap.QueryPlatformTypesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditPlatformTypesAssociations(request *EditPlatformTypesAssociations) (*EditPlatformTypesAssociationsResponse, error) {
	panic("baseApiSoap.EditPlatformTypesAssociations not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditPlatformTypesAssociationsContext(ctx context.Context, request *EditPlatformTypesAssociations) (*EditPlatformTypesAssociationsResponse, error) {
	panic("baseApiSoap.EditPlatformTypesAssociationsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateAppLayerAssignment(request *UpdateAppLayerAssignment) (*UpdateAppLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdateAppLayerAssignment not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateAppLayerAssignmentContext(ctx context.Context, request *UpdateAppLayerAssignment) (*UpdateAppLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdateAppLayerAssignmentContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RemoveAppLayerAssignment(request *RemoveAppLayerAssignment) (*RemoveAppLayerAssignmentResponse, error) {
	panic("baseApiSoap.RemoveAppLayerAssignment not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RemoveAppLayerAssignmentContext(ctx context.Context, request *RemoveAppLayerAssignment) (*RemoveAppLayerAssignmentResponse, error) {
	panic("baseApiSoap.RemoveAppLayerAssignmentContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateOsLayerAssignment(request *UpdateOsLayerAssignment) (*UpdateOsLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdateOsLayerAssignment not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateOsLayerAssignmentContext(ctx context.Context, request *UpdateOsLayerAssignment) (*UpdateOsLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdateOsLayerAssignmentContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdatePlatformLayerAssignment(request *UpdatePlatformLayerAssignment) (*UpdatePlatformLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdatePlatformLayerAssignment not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdatePlatformLayerAssignmentContext(ctx context.Context, request *UpdatePlatformLayerAssignment) (*UpdatePlatformLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdatePlatformLayerAssignmentContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypeHelpUrl(request *QueryPlatformTypeHelpUrl) (*QueryPlatformTypeHelpUrlResponse, error) {
	panic("baseApiSoap.QueryPlatformTypeHelpUrl not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypeHelpUrlContext(ctx context.Context, request *QueryPlatformTypeHelpUrl) (*QueryPlatformTypeHelpUrlResponse, error) {
	panic("baseApiSoap.QueryPlatformTypeHelpUrlContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypeLayerHelpUrl(request *QueryPlatformTypeLayerHelpUrl) (*QueryPlatformTypeLayerHelpUrlResponse, error) {
	panic("baseApiSoap.QueryPlatformTypeLayerHelpUrl not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypeLayerHelpUrlContext(ctx context.Context, request *QueryPlatformTypeLayerHelpUrl) (*QueryPlatformTypeLayerHelpUrlResponse, error) {
	panic("baseApiSoap.QueryPlatformTypeLayerHelpUrlContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFreeDisks(request *QueryFreeDisks) (*QueryFreeDisksResponse, error) {
	panic("baseApiSoap.QueryFreeDisks not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFreeDisksContext(ctx context.Context, request *QueryFreeDisks) (*QueryFreeDisksResponse, error) {
	panic("baseApiSoap.QueryFreeDisksContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExpandStorage(request *ExpandStorage) (*ExpandStorageResponse, error) {
	panic("baseApiSoap.ExpandStorage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExpandStorageContext(ctx context.Context, request *ExpandStorage) (*ExpandStorageResponse, error) {
	panic("baseApiSoap.ExpandStorageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateLayerFolder(request *UpdateLayerFolder) (*UpdateLayerFolderResponse, error) {
	panic("baseApiSoap.UpdateLayerFolder not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateLayerFolderContext(ctx context.Context, request *UpdateLayerFolder) (*UpdateLayerFolderResponse, error) {
	panic("baseApiSoap.UpdateLayerFolderContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) AnalyzeLayerRevisions(request *AnalyzeLayerRevisions) (*AnalyzeLayerRevisionsResponse, error) {
	panic("baseApiSoap.AnalyzeLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) AnalyzeLayerRevisionsContext(ctx context.Context, request *AnalyzeLayerRevisions) (*AnalyzeLayerRevisionsResponse, error) {
	panic("baseApiSoap.AnalyzeLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDiskUsageEstimate(request *QueryDiskUsageEstimate) (*QueryDiskUsageEstimateResponse, error) {
	panic("baseApiSoap.QueryDiskUsageEstimate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDiskUsageEstimateContext(ctx context.Context, request *QueryDiskUsageEstimate) (*QueryDiskUsageEstimateResponse, error) {
	panic("baseApiSoap.QueryDiskUsageEstimateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUserRoles(request *QueryUserRoles) (*QueryUserRolesResponse, error) {
	panic("baseApiSoap.QueryUserRoles not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUserRolesContext(ctx context.Context, request *QueryUserRoles) (*QueryUserRolesResponse, error) {
	panic("baseApiSoap.QueryUserRolesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemGroupMembership(request *QueryDirectoryItemGroupMembership) (*QueryDirectoryItemGroupMembershipResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemGroupMembership not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemGroupMembershipContext(ctx context.Context, request *QueryDirectoryItemGroupMembership) (*QueryDirectoryItemGroupMembershipResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemGroupMembershipContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetUserSetting(request *SetUserSetting) (*SetUserSettingResponse, error) {
	panic("baseApiSoap.SetUserSetting not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetUserSettingContext(ctx context.Context, request *SetUserSetting) (*SetUserSettingResponse, error) {
	panic("baseApiSoap.SetUserSettingContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateServerStorageSlot(request *CreateServerStorageSlot) (*CreateServerStorageSlotResponse, error) {
	panic("baseApiSoap.CreateServerStorageSlot not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateServerStorageSlotContext(ctx context.Context, request *CreateServerStorageSlot) (*CreateServerStorageSlotResponse, error) {
	panic("baseApiSoap.CreateServerStorageSlotContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ReadServerStorageSlot(request *ReadServerStorageSlot) (*ReadServerStorageSlotResponse, error) {
	panic("baseApiSoap.ReadServerStorageSlot not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ReadServerStorageSlotContext(ctx context.Context, request *ReadServerStorageSlot) (*ReadServerStorageSlotResponse, error) {
	panic("baseApiSoap.ReadServerStorageSlotContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateServerStorageSlot(request *UpdateServerStorageSlot) (*UpdateServerStorageSlotResponse, error) {
	panic("baseApiSoap.UpdateServerStorageSlot not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateServerStorageSlotContext(ctx context.Context, request *UpdateServerStorageSlot) (*UpdateServerStorageSlotResponse, error) {
	panic("baseApiSoap.UpdateServerStorageSlotContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteServerStorageSlot(request *DeleteServerStorageSlot) (*DeleteServerStorageSlotResponse, error) {
	panic("baseApiSoap.DeleteServerStorageSlot not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteServerStorageSlotContext(ctx context.Context, request *DeleteServerStorageSlot) (*DeleteServerStorageSlotResponse, error) {
	panic("baseApiSoap.DeleteServerStorageSlotContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RunCacheCommand(request *RunCacheCommand) (*RunCacheCommandResponse, error) {
	panic("baseApiSoap.RunCacheCommand not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RunCacheCommandContext(ctx context.Context, request *RunCacheCommand) (*RunCacheCommandResponse, error) {
	panic("baseApiSoap.RunCacheCommandContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DefineFileShare(request *DefineFileShare) (*DefineFileShareResponse, error) {
	panic("baseApiSoap.DefineFileShare not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DefineFileShareContext(ctx context.Context, request *DefineFileShare) (*DefineFileShareResponse, error) {
	panic("baseApiSoap.DefineFileShareContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditFileShare(request *EditFileShare) (*EditFileShareResponse, error) {
	panic("baseApiSoap.EditFileShare not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditFileShareContext(ctx context.Context, request *EditFileShare) (*EditFileShareResponse, error) {
	panic("baseApiSoap.EditFileShareContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFileShareSummary(request *QueryFileShareSummary) (*QueryFileShareSummaryResponse, error) {
	panic("baseApiSoap.QueryFileShareSummary not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFileShareSummaryContext(ctx context.Context, request *QueryFileShareSummary) (*QueryFileShareSummaryResponse, error) {
	panic("baseApiSoap.QueryFileShareSummaryContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFileShareDetails(request *QueryFileShareDetails) (*QueryFileShareDetailsResponse, error) {
	panic("baseApiSoap.QueryFileShareDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFileShareDetailsContext(ctx context.Context, request *QueryFileShareDetails) (*QueryFileShareDetailsResponse, error) {
	panic("baseApiSoap.QueryFileShareDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteFileShares(request *DeleteFileShares) (*DeleteFileSharesResponse, error) {
	panic("baseApiSoap.DeleteFileShares not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteFileSharesContext(ctx context.Context, request *DeleteFileShares) (*DeleteFileSharesResponse, error) {
	panic("baseApiSoap.DeleteFileSharesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ReorderFileShares(request *ReorderFileShares) (*ReorderFileSharesResponse, error) {
	panic("baseApiSoap.ReorderFileShares not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ReorderFileSharesContext(ctx context.Context, request *ReorderFileShares) (*ReorderFileSharesResponse, error) {
	panic("baseApiSoap.ReorderFileSharesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdatePasswords(request *UpdatePasswords) (*UpdatePasswordsResponse, error) {
	panic("baseApiSoap.UpdatePasswords not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdatePasswordsContext(ctx context.Context, request *UpdatePasswords) (*UpdatePasswordsResponse, error) {
	panic("baseApiSoap.UpdatePasswordsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DefaultFileShareMessages(request *DefaultFileShareMessages) (*DefaultFileShareMessagesResponse, error) {
	panic("baseApiSoap.DefaultFileShareMessages not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DefaultFileShareMessagesContext(ctx context.Context, request *DefaultFileShareMessages) (*DefaultFileShareMessagesResponse, error) {
	panic("baseApiSoap.DefaultFileShareMessagesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryExportableRevisions(request *QueryExportableRevisions) (*QueryExportableRevisionsResponse, error) {
	panic("baseApiSoap.QueryExportableRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryExportableRevisionsContext(ctx context.Context, request *QueryExportableRevisions) (*QueryExportableRevisionsResponse, error) {
	panic("baseApiSoap.QueryExportableRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRemoteFileShareAvailableSpace(request *QueryRemoteFileShareAvailableSpace) (*QueryRemoteFileShareAvailableSpaceResponse, error) {
	panic("baseApiSoap.QueryRemoteFileShareAvailableSpace not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRemoteFileShareAvailableSpaceContext(ctx context.Context, request *QueryRemoteFileShareAvailableSpace) (*QueryRemoteFileShareAvailableSpaceResponse, error) {
	panic("baseApiSoap.QueryRemoteFileShareAvailableSpaceContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExportLayerRevisions(request *ExportLayerRevisions) (*ExportLayerRevisionsResponse, error) {
	panic("baseApiSoap.ExportLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExportLayerRevisionsContext(ctx context.Context, request *ExportLayerRevisions) (*ExportLayerRevisionsResponse, error) {
	panic("baseApiSoap.ExportLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableRevisions(request *QueryImportableRevisions) (*QueryImportableRevisionsResponse, error) {
	panic("baseApiSoap.QueryImportableRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableRevisionsContext(ctx context.Context, request *QueryImportableRevisions) (*QueryImportableRevisionsResponse, error) {
	panic("baseApiSoap.QueryImportableRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportLayerRevisions(request *ImportLayerRevisions) (*ImportLayerRevisionsResponse, error) {
	panic("baseApiSoap.ImportLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportLayerRevisionsContext(ctx context.Context, request *ImportLayerRevisions) (*ImportLayerRevisionsResponse, error) {
	panic("baseApiSoap.ImportLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportExportFileShare(request *QueryImportExportFileShare) (*QueryImportExportFileShareResponse, error) {
	panic("baseApiSoap.QueryImportExportFileShare not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportExportFileShareContext(ctx context.Context, request *QueryImportExportFileShare) (*QueryImportExportFileShareResponse, error) {
	panic("baseApiSoap.QueryImportExportFileShareContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableRevisionDetails(request *QueryImportableRevisionDetails) (*QueryImportableRevisionDetailsResponse, error) {
	panic("baseApiSoap.QueryImportableRevisionDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableRevisionDetailsContext(ctx context.Context, request *QueryImportableRevisionDetails) (*QueryImportableRevisionDetailsResponse, error) {
	panic("baseApiSoap.QueryImportableRevisionDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryExportableRevisionDetails(request *QueryExportableRevisionDetails) (*QueryExportableRevisionDetailsResponse, error) {
	panic("baseApiSoap.QueryExportableRevisionDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryExportableRevisionDetailsContext(ctx context.Context, request *QueryExportableRevisionDetails) (*QueryExportableRevisionDetailsResponse, error) {
	panic("baseApiSoap.QueryExportableRevisionDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RegisterCompositingEngine(request *RegisterCompositingEngine) (*RegisterCompositingEngineResponse, error) {
	panic("baseApiSoap.RegisterCompositingEngine not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RegisterCompositingEngineContext(ctx context.Context, request *RegisterCompositingEngine) (*RegisterCompositingEngineResponse, error) {
	panic("baseApiSoap.RegisterCompositingEngineContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ConfigureCompositingEngineRegistration(request *ConfigureCompositingEngineRegistration) (*ConfigureCompositingEngineRegistrationResponse, error) {
	panic("baseApiSoap.ConfigureCompositingEngineRegistration not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ConfigureCompositingEngineRegistrationContext(ctx context.Context, request *ConfigureCompositingEngineRegistration) (*ConfigureCompositingEngineRegistrationResponse, error) {
	panic("baseApiSoap.ConfigureCompositingEngineRegistrationContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatform(request *QueryPlatform) (*QueryPlatformResponse, error) {
	panic("baseApiSoap.QueryPlatform not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformContext(ctx context.Context, request *QueryPlatform) (*QueryPlatformResponse, error) {
	panic("baseApiSoap.QueryPlatformContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConfig(request *QueryPlatformConfig) (*QueryPlatformConfigResponse, error) {
	panic("baseApiSoap.QueryPlatformConfig not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConfigContext(ctx context.Context, request *QueryPlatformConfig) (*QueryPlatformConfigResponse, error) {
	panic("baseApiSoap.QueryPlatformConfigContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateOrUpdatePlatformConfig(request *CreateOrUpdatePlatformConfig) (*CreateOrUpdatePlatformConfigResponse, error) {
	panic("baseApiSoap.CreateOrUpdatePlatformConfig not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateOrUpdatePlatformConfigContext(ctx context.Context, request *CreateOrUpdatePlatformConfig) (*CreateOrUpdatePlatformConfigResponse, error) {
	panic("baseApiSoap.CreateOrUpdatePlatformConfigContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) MigrateAppliance(request *MigrateAppliance) (*MigrateApplianceResponse, error) {
	panic("baseApiSoap.MigrateAppliance not implemented in test — embed and override this method")
}

func (b *baseApiSoap) MigrateApplianceContext(ctx context.Context, request *MigrateAppliance) (*MigrateApplianceResponse, error) {
	panic("baseApiSoap.MigrateApplianceContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PauseResumeMigrateAppliance(request *PauseResumeMigrateAppliance) (*PauseResumeMigrateApplianceResponse, error) {
	panic("baseApiSoap.PauseResumeMigrateAppliance not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PauseResumeMigrateApplianceContext(ctx context.Context, request *PauseResumeMigrateAppliance) (*PauseResumeMigrateApplianceResponse, error) {
	panic("baseApiSoap.PauseResumeMigrateApplianceContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) FinalizeMigrateAppliance(request *FinalizeMigrateAppliance) (*FinalizeMigrateApplianceResponse, error) {
	panic("baseApiSoap.FinalizeMigrateAppliance not implemented in test — embed and override this method")
}

func (b *baseApiSoap) FinalizeMigrateApplianceContext(ctx context.Context, request *FinalizeMigrateAppliance) (*FinalizeMigrateApplianceResponse, error) {
	panic("baseApiSoap.FinalizeMigrateApplianceContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePoints(request *QueryCachePoints) (*QueryCachePointsResponse, error) {
	panic("baseApiSoap.QueryCachePoints not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePointsContext(ctx context.Context, request *QueryCachePoints) (*QueryCachePointsResponse, error) {
	panic("baseApiSoap.QueryCachePointsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePointDetails(request *QueryCachePointDetails) (*QueryCachePointDetailsResponse, error) {
	panic("baseApiSoap.QueryCachePointDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePointDetailsContext(ctx context.Context, request *QueryCachePointDetails) (*QueryCachePointDetailsResponse, error) {
	panic("baseApiSoap.QueryCachePointDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePoint(request *QueryCachePoint) (*QueryCachePointResponse, error) {
	panic("baseApiSoap.QueryCachePoint not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePointContext(ctx context.Context, request *QueryCachePoint) (*QueryCachePointResponse, error) {
	panic("baseApiSoap.QueryCachePointContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateImage(request *CreateImage) (*CreateImageResponse, error) {
	panic("baseApiSoap.CreateImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateImageContext(ctx context.Context, request *CreateImage) (*CreateImageResponse, error) {
	panic("baseApiSoap.CreateImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CloneImage(request *CloneImage) (*CloneImageResponse, error) {
	panic("baseApiSoap.CloneImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CloneImageContext(ctx context.Context, request *CloneImage) (*CloneImageResponse, error) {
	panic("baseApiSoap.CloneImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditImage(request *EditImage) (*EditImageResponse, error) {
	panic("baseApiSoap.EditImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditImageContext(ctx context.Context, request *EditImage) (*EditImageResponse, error) {
	panic("baseApiSoap.EditImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExportImage(request *ExportImage) (*ExportImageResponse, error) {
	panic("baseApiSoap.ExportImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExportImageContext(ctx context.Context, request *ExportImage) (*ExportImageResponse, error) {
	panic("baseApiSoap.ExportImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteImage(request *DeleteImage) (*DeleteImageResponse, error) {
	panic("baseApiSoap.DeleteImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteImageContext(ctx context.Context, request *DeleteImage) (*DeleteImageResponse, error) {
	panic("baseApiSoap.DeleteImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionFolders(request *QueryDirectoryJunctionFolders) (*QueryDirectoryJunctionFoldersResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionFolders not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionFoldersContext(ctx context.Context, request *QueryDirectoryJunctionFolders) (*QueryDirectoryJunctionFoldersResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionFoldersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionDetails(request *QueryDirectoryJunctionDetails) (*QueryDirectoryJunctionDetailsResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionDetailsContext(ctx context.Context, request *QueryDirectoryJunctionDetails) (*QueryDirectoryJunctionDetailsResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestDirectoryJunction(request *TestDirectoryJunction) (*TestDirectoryJunctionResponse, error) {
	panic("baseApiSoap.TestDirectoryJunction not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestDirectoryJunctionContext(ctx context.Context, request *TestDirectoryJunction) (*TestDirectoryJunctionResponse, error) {
	panic("baseApiSoap.TestDirectoryJunctionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateDirectoryJunction(request *CreateDirectoryJunction) (*CreateDirectoryJunctionResponse, error) {
	panic("baseApiSoap.CreateDirectoryJunction not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateDirectoryJunctionContext(ctx context.Context, request *CreateDirectoryJunction) (*CreateDirectoryJunctionResponse, error) {
	panic("baseApiSoap.CreateDirectoryJunctionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditDirectoryJunction(request *EditDirectoryJunction) (*EditDirectoryJunctionResponse, error) {
	panic("baseApiSoap.EditDirectoryJunction not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditDirectoryJunctionContext(ctx context.Context, request *EditDirectoryJunction) (*EditDirectoryJunctionResponse, error) {
	panic("baseApiSoap.EditDirectoryJunctionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionAttributes(request *QueryDirectoryJunctionAttributes) (*QueryDirectoryJunctionAttributesResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionAttributes not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionAttributesContext(ctx context.Context, request *QueryDirectoryJunctionAttributes) (*QueryDirectoryJunctionAttributesResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionAttributesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteCloudController(request *DeleteCloudController) (*DeleteCloudControllerResponse, error) {
	panic("baseApiSoap.DeleteCloudController not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteCloudControllerContext(ctx context.Context, request *DeleteCloudController) (*DeleteCloudControllerResponse, error) {
	panic("baseApiSoap.DeleteCloudControllerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestNetBiosName(request *TestNetBiosName) (*TestNetBiosNameResponse, error) {
	panic("baseApiSoap.TestNetBiosName not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestNetBiosNameContext(ctx context.Context, request *TestNetBiosName) (*TestNetBiosNameResponse, error) {
	panic("baseApiSoap.TestNetBiosNameContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryMaintenanceSchedules(request *QueryMaintenanceSchedules) (*QueryMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.QueryMaintenanceSchedules not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryMaintenanceSchedulesContext(ctx context.Context, request *QueryMaintenanceSchedules) (*QueryMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.QueryMaintenanceSchedulesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateMaintenanceSchedules(request *CreateMaintenanceSchedules) (*CreateMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.CreateMaintenanceSchedules not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateMaintenanceSchedulesContext(ctx context.Context, request *CreateMaintenanceSchedules) (*CreateMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.CreateMaintenanceSchedulesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateMaintenanceSchedules(request *UpdateMaintenanceSchedules) (*UpdateMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.UpdateMaintenanceSchedules not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateMaintenanceSchedulesContext(ctx context.Context, request *UpdateMaintenanceSchedules) (*UpdateMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.UpdateMaintenanceSchedulesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteMaintenanceSchedules(request *DeleteMaintenanceSchedules) (*DeleteMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.DeleteMaintenanceSchedules not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteMaintenanceSchedulesContext(ctx context.Context, request *DeleteMaintenanceSchedules) (*DeleteMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.DeleteMaintenanceSchedulesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeletePlatformConnectorConfiguration(request *DeletePlatformConnectorConfiguration) (*DeletePlatformConnectorConfigurationResponse, error) {
	panic("baseApiSoap.DeletePlatformConnectorConfiguration not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeletePlatformConnectorConfigurationContext(ctx context.Context, request *DeletePlatformConnectorConfiguration) (*DeletePlatformConnectorConfigurationResponse, error) {
	panic("baseApiSoap.DeletePlatformConnectorConfigurationContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QuerySystemSettings(request *QuerySystemSettings) (*QuerySystemSettingsResponse, error) {
	panic("baseApiSoap.QuerySystemSettings not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QuerySystemSettingsContext(ctx context.Context, request *QuerySystemSettings) (*QuerySystemSettingsResponse, error) {
	panic("baseApiSoap.QuerySystemSettingsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetSystemSettings(request *SetSystemSettings) (*SetSystemSettingsResponse, error) {
	panic("baseApiSoap.SetSystemSettings not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetSystemSettingsContext(ctx context.Context, request *SetSystemSettings) (*SetSystemSettingsResponse, error) {
	panic("baseApiSoap.SetSystemSettingsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteSystemSettings(request *DeleteSystemSettings) (*DeleteSystemSettingsResponse, error) {
	panic("baseApiSoap.DeleteSystemSettings not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteSystemSettingsContext(ctx context.Context, request *DeleteSystemSettings) (*DeleteSystemSettingsResponse, error) {
	panic("baseApiSoap.DeleteSystemSettingsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryNotificationSettings(request *QueryNotificationSettings) (*QueryNotificationSettingsResponse, error) {
	panic("baseApiSoap.QueryNotificationSettings not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryNotificationSettingsContext(ctx context.Context, request *QueryNotificationSettings) (*QueryNotificationSettingsResponse, error) {
	panic("baseApiSoap.QueryNotificationSettingsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetNotificationSettings(request *SetNotificationSettings) (*SetNotificationSettingsResponse, error) {
	panic("baseApiSoap.SetNotificationSettings not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetNotificationSettingsContext(ctx context.Context, request *SetNotificationSettings) (*SetNotificationSettingsResponse, error) {
	panic("baseApiSoap.SetNotificationSettingsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestEmailServer(request *TestEmailServer) (*TestEmailServerResponse, error) {
	panic("baseApiSoap.TestEmailServer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestEmailServerContext(ctx context.Context, request *TestEmailServer) (*TestEmailServerResponse, error) {
	panic("baseApiSoap.TestEmailServerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetSoapClient() *soap.Client {
	panic("baseApiSoap.GetSoapClient not implemented in test — embed and override this method")
}
