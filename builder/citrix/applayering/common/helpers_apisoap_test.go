// Copyright (c) Citrix, Inc.

package common

// baseApiSoap satisfies the very large elmsoap.ApiSoap interface by panicking
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

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// Compile-time assertion that *baseApiSoap satisfies elmsoap.ApiSoap.
var _ elmsoap.ApiSoap = (*baseApiSoap)(nil)

// baseApiSoap is the empty receiver that holds all panic stubs below.
type baseApiSoap struct{}

// Suppress "imported and not used" if the generated stub set ever shrinks.
// context.Context and *soap.Client both appear in ApiSoap method signatures.
var (
	_ = context.TODO
	_ *soap.Client
)

func (b *baseApiSoap) Login(request *elmsoap.Login) (*elmsoap.LoginResponse, error) {
	panic("baseApiSoap.Login not implemented in test — embed and override this method")
}

func (b *baseApiSoap) LoginContext(ctx context.Context, request *elmsoap.Login) (*elmsoap.LoginResponse, error) {
	panic("baseApiSoap.LoginContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) Logout(request *elmsoap.Logout) (*elmsoap.LogoutResponse, error) {
	panic("baseApiSoap.Logout not implemented in test — embed and override this method")
}

func (b *baseApiSoap) LogoutContext(ctx context.Context, request *elmsoap.Logout) (*elmsoap.LogoutResponse, error) {
	panic("baseApiSoap.LogoutContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) IsAuthenticated(request *elmsoap.IsAuthenticated) (*elmsoap.IsAuthenticatedResponse, error) {
	panic("baseApiSoap.IsAuthenticated not implemented in test — embed and override this method")
}

func (b *baseApiSoap) IsAuthenticatedContext(ctx context.Context, request *elmsoap.IsAuthenticated) (*elmsoap.IsAuthenticatedResponse, error) {
	panic("baseApiSoap.IsAuthenticatedContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) Log(request *elmsoap.Log) (*elmsoap.LogResponse, error) {
	panic("baseApiSoap.Log not implemented in test — embed and override this method")
}

func (b *baseApiSoap) LogContext(ctx context.Context, request *elmsoap.Log) (*elmsoap.LogResponse, error) {
	panic("baseApiSoap.LogContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementAppliance(request *elmsoap.QueryManagementAppliance) (*elmsoap.QueryManagementApplianceResponse, error) {
	panic("baseApiSoap.QueryManagementAppliance not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementApplianceContext(ctx context.Context, request *elmsoap.QueryManagementAppliance) (*elmsoap.QueryManagementApplianceResponse, error) {
	panic("baseApiSoap.QueryManagementApplianceContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementApplianceDetails(request *elmsoap.QueryManagementApplianceDetails) (*elmsoap.QueryManagementApplianceDetailsResponse, error) {
	panic("baseApiSoap.QueryManagementApplianceDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementApplianceDetailsContext(ctx context.Context, request *elmsoap.QueryManagementApplianceDetails) (*elmsoap.QueryManagementApplianceDetailsResponse, error) {
	panic("baseApiSoap.QueryManagementApplianceDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementApplianceDetailsWithParam(request *elmsoap.QueryManagementApplianceDetailsWithParam) (*elmsoap.QueryManagementApplianceDetailsWithParamResponse, error) {
	panic("baseApiSoap.QueryManagementApplianceDetailsWithParam not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryManagementApplianceDetailsWithParamContext(ctx context.Context, request *elmsoap.QueryManagementApplianceDetailsWithParam) (*elmsoap.QueryManagementApplianceDetailsWithParamResponse, error) {
	panic("baseApiSoap.QueryManagementApplianceDetailsWithParamContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryActivity(request *elmsoap.QueryActivity) (*elmsoap.QueryActivityResponse, error) {
	panic("baseApiSoap.QueryActivity not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryActivityContext(ctx context.Context, request *elmsoap.QueryActivity) (*elmsoap.QueryActivityResponse, error) {
	panic("baseApiSoap.QueryActivityContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) BrowseContainer(request *elmsoap.BrowseContainer) (*elmsoap.BrowseContainerResponse, error) {
	panic("baseApiSoap.BrowseContainer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) BrowseContainerContext(ctx context.Context, request *elmsoap.BrowseContainer) (*elmsoap.BrowseContainerResponse, error) {
	panic("baseApiSoap.BrowseContainerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RegisterImportableVM(request *elmsoap.RegisterImportableVM) (*elmsoap.RegisterImportableVMResponse, error) {
	panic("baseApiSoap.RegisterImportableVM not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RegisterImportableVMContext(ctx context.Context, request *elmsoap.RegisterImportableVM) (*elmsoap.RegisterImportableVMResponse, error) {
	panic("baseApiSoap.RegisterImportableVMContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteGoldImage(request *elmsoap.DeleteGoldImage) (*elmsoap.DeleteGoldImageResponse, error) {
	panic("baseApiSoap.DeleteGoldImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteGoldImageContext(ctx context.Context, request *elmsoap.DeleteGoldImage) (*elmsoap.DeleteGoldImageResponse, error) {
	panic("baseApiSoap.DeleteGoldImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUsers(request *elmsoap.QueryUsers) (*elmsoap.QueryUsersResponse, error) {
	panic("baseApiSoap.QueryUsers not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUsersContext(ctx context.Context, request *elmsoap.QueryUsers) (*elmsoap.QueryUsersResponse, error) {
	panic("baseApiSoap.QueryUsersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryGroups(request *elmsoap.QueryGroups) (*elmsoap.QueryGroupsResponse, error) {
	panic("baseApiSoap.QueryGroups not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryGroupsContext(ctx context.Context, request *elmsoap.QueryGroups) (*elmsoap.QueryGroupsResponse, error) {
	panic("baseApiSoap.QueryGroupsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUpgradeAvailability(request *elmsoap.QueryUpgradeAvailability) (*elmsoap.QueryUpgradeAvailabilityResponse, error) {
	panic("baseApiSoap.QueryUpgradeAvailability not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUpgradeAvailabilityContext(ctx context.Context, request *elmsoap.QueryUpgradeAvailability) (*elmsoap.QueryUpgradeAvailabilityResponse, error) {
	panic("baseApiSoap.QueryUpgradeAvailabilityContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ClearUpgradeStatus(request *elmsoap.ClearUpgradeStatus) (*elmsoap.ClearUpgradeStatusResponse, error) {
	panic("baseApiSoap.ClearUpgradeStatus not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ClearUpgradeStatusContext(ctx context.Context, request *elmsoap.ClearUpgradeStatus) (*elmsoap.ClearUpgradeStatusResponse, error) {
	panic("baseApiSoap.ClearUpgradeStatusContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemAppAssignments(request *elmsoap.QueryDirectoryItemAppAssignments) (*elmsoap.QueryDirectoryItemAppAssignmentsResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemAppAssignments not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemAppAssignmentsContext(ctx context.Context, request *elmsoap.QueryDirectoryItemAppAssignments) (*elmsoap.QueryDirectoryItemAppAssignmentsResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemAppAssignmentsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemDetails(request *elmsoap.QueryDirectoryItemDetails) (*elmsoap.QueryDirectoryItemDetailsResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemDetailsContext(ctx context.Context, request *elmsoap.QueryDirectoryItemDetails) (*elmsoap.QueryDirectoryItemDetailsResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SearchDirectoryItem(request *elmsoap.SearchDirectoryItem) (*elmsoap.SearchDirectoryItemResponse, error) {
	panic("baseApiSoap.SearchDirectoryItem not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SearchDirectoryItemContext(ctx context.Context, request *elmsoap.SearchDirectoryItem) (*elmsoap.SearchDirectoryItemResponse, error) {
	panic("baseApiSoap.SearchDirectoryItemContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SearchDirectoryItemPendingOp(request *elmsoap.SearchDirectoryItemPendingOp) (*elmsoap.SearchDirectoryItemPendingOpResponse, error) {
	panic("baseApiSoap.SearchDirectoryItemPendingOp not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SearchDirectoryItemPendingOpContext(ctx context.Context, request *elmsoap.SearchDirectoryItemPendingOp) (*elmsoap.SearchDirectoryItemPendingOpResponse, error) {
	panic("baseApiSoap.SearchDirectoryItemPendingOpContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryWorkTicketsAsPendingOp(request *elmsoap.QueryWorkTicketsAsPendingOp) (*elmsoap.QueryWorkTicketsAsPendingOpResponse, error) {
	panic("baseApiSoap.QueryWorkTicketsAsPendingOp not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryWorkTicketsAsPendingOpContext(ctx context.Context, request *elmsoap.QueryWorkTicketsAsPendingOp) (*elmsoap.QueryWorkTicketsAsPendingOpResponse, error) {
	panic("baseApiSoap.QueryWorkTicketsAsPendingOpContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateDirectoryItem(request *elmsoap.CreateDirectoryItem) (*elmsoap.CreateDirectoryItemResponse, error) {
	panic("baseApiSoap.CreateDirectoryItem not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateDirectoryItemContext(ctx context.Context, request *elmsoap.CreateDirectoryItem) (*elmsoap.CreateDirectoryItemResponse, error) {
	panic("baseApiSoap.CreateDirectoryItemContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryAuditLog(request *elmsoap.QueryAuditLog) (*elmsoap.QueryAuditLogResponse, error) {
	panic("baseApiSoap.QueryAuditLog not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryAuditLogContext(ctx context.Context, request *elmsoap.QueryAuditLog) (*elmsoap.QueryAuditLogResponse, error) {
	panic("baseApiSoap.QueryAuditLogContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryAuditLogDetail(request *elmsoap.QueryAuditLogDetail) (*elmsoap.QueryAuditLogDetailResponse, error) {
	panic("baseApiSoap.QueryAuditLogDetail not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryAuditLogDetailContext(ctx context.Context, request *elmsoap.QueryAuditLogDetail) (*elmsoap.QueryAuditLogDetailResponse, error) {
	panic("baseApiSoap.QueryAuditLogDetailContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditDirectoryItem(request *elmsoap.EditDirectoryItem) (*elmsoap.EditDirectoryItemResponse, error) {
	panic("baseApiSoap.EditDirectoryItem not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditDirectoryItemContext(ctx context.Context, request *elmsoap.EditDirectoryItem) (*elmsoap.EditDirectoryItemResponse, error) {
	panic("baseApiSoap.EditDirectoryItemContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteDirectoryItem(request *elmsoap.DeleteDirectoryItem) (*elmsoap.DeleteDirectoryItemResponse, error) {
	panic("baseApiSoap.DeleteDirectoryItem not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteDirectoryItemContext(ctx context.Context, request *elmsoap.DeleteDirectoryItem) (*elmsoap.DeleteDirectoryItemResponse, error) {
	panic("baseApiSoap.DeleteDirectoryItemContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteDirectoryJunction(request *elmsoap.DeleteDirectoryJunction) (*elmsoap.DeleteDirectoryJunctionResponse, error) {
	panic("baseApiSoap.DeleteDirectoryJunction not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteDirectoryJunctionContext(ctx context.Context, request *elmsoap.DeleteDirectoryJunction) (*elmsoap.DeleteDirectoryJunctionResponse, error) {
	panic("baseApiSoap.DeleteDirectoryJunctionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryInfrastructure(request *elmsoap.QueryInfrastructure) (*elmsoap.QueryInfrastructureResponse, error) {
	panic("baseApiSoap.QueryInfrastructure not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryInfrastructureContext(ctx context.Context, request *elmsoap.QueryInfrastructure) (*elmsoap.QueryInfrastructureResponse, error) {
	panic("baseApiSoap.QueryInfrastructureContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateInfrastructure(request *elmsoap.CreateInfrastructure) (*elmsoap.CreateInfrastructureResponse, error) {
	panic("baseApiSoap.CreateInfrastructure not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateInfrastructureContext(ctx context.Context, request *elmsoap.CreateInfrastructure) (*elmsoap.CreateInfrastructureResponse, error) {
	panic("baseApiSoap.CreateInfrastructureContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateInfrastructure(request *elmsoap.UpdateInfrastructure) (*elmsoap.UpdateInfrastructureResponse, error) {
	panic("baseApiSoap.UpdateInfrastructure not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateInfrastructureContext(ctx context.Context, request *elmsoap.UpdateInfrastructure) (*elmsoap.UpdateInfrastructureResponse, error) {
	panic("baseApiSoap.UpdateInfrastructureContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestInfrastructure(request *elmsoap.TestInfrastructure) (*elmsoap.TestInfrastructureResponse, error) {
	panic("baseApiSoap.TestInfrastructure not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestInfrastructureContext(ctx context.Context, request *elmsoap.TestInfrastructure) (*elmsoap.TestInfrastructureResponse, error) {
	panic("baseApiSoap.TestInfrastructureContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDataCenters(request *elmsoap.QueryDataCenters) (*elmsoap.QueryDataCentersResponse, error) {
	panic("baseApiSoap.QueryDataCenters not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDataCentersContext(ctx context.Context, request *elmsoap.QueryDataCenters) (*elmsoap.QueryDataCentersResponse, error) {
	panic("baseApiSoap.QueryDataCentersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryComputeResources(request *elmsoap.QueryComputeResources) (*elmsoap.QueryComputeResourcesResponse, error) {
	panic("baseApiSoap.QueryComputeResources not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryComputeResourcesContext(ctx context.Context, request *elmsoap.QueryComputeResources) (*elmsoap.QueryComputeResourcesResponse, error) {
	panic("baseApiSoap.QueryComputeResourcesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryEnvironment(request *elmsoap.QueryEnvironment) (*elmsoap.QueryEnvironmentResponse, error) {
	panic("baseApiSoap.QueryEnvironment not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryEnvironmentContext(ctx context.Context, request *elmsoap.QueryEnvironment) (*elmsoap.QueryEnvironmentResponse, error) {
	panic("baseApiSoap.QueryEnvironmentContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryWorkTickets(request *elmsoap.QueryWorkTickets) (*elmsoap.QueryWorkTicketsResponse, error) {
	panic("baseApiSoap.QueryWorkTickets not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryWorkTicketsContext(ctx context.Context, request *elmsoap.QueryWorkTickets) (*elmsoap.QueryWorkTicketsResponse, error) {
	panic("baseApiSoap.QueryWorkTicketsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableVMs(request *elmsoap.QueryImportableVMs) (*elmsoap.QueryImportableVMsResponse, error) {
	panic("baseApiSoap.QueryImportableVMs not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableVMsContext(ctx context.Context, request *elmsoap.QueryImportableVMs) (*elmsoap.QueryImportableVMsResponse, error) {
	panic("baseApiSoap.QueryImportableVMsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerIcons(request *elmsoap.QueryLayerIcons) (*elmsoap.QueryLayerIconsResponse, error) {
	panic("baseApiSoap.QueryLayerIcons not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerIconsContext(ctx context.Context, request *elmsoap.QueryLayerIcons) (*elmsoap.QueryLayerIconsResponse, error) {
	panic("baseApiSoap.QueryLayerIconsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelWorkTickets(request *elmsoap.CancelWorkTickets) (*elmsoap.CancelWorkTicketsResponse, error) {
	panic("baseApiSoap.CancelWorkTickets not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelWorkTicketsContext(ctx context.Context, request *elmsoap.CancelWorkTickets) (*elmsoap.CancelWorkTicketsResponse, error) {
	panic("baseApiSoap.CancelWorkTicketsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelWorkItems(request *elmsoap.CancelWorkItems) (*elmsoap.CancelWorkItemsResponse, error) {
	panic("baseApiSoap.CancelWorkItems not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelWorkItemsContext(ctx context.Context, request *elmsoap.CancelWorkItems) (*elmsoap.CancelWorkItemsResponse, error) {
	panic("baseApiSoap.CancelWorkItemsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPendingOperation(request *elmsoap.QueryPendingOperation) (*elmsoap.QueryPendingOperationResponse, error) {
	panic("baseApiSoap.QueryPendingOperation not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPendingOperationContext(ctx context.Context, request *elmsoap.QueryPendingOperation) (*elmsoap.QueryPendingOperationResponse, error) {
	panic("baseApiSoap.QueryPendingOperationContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelPendingOperation(request *elmsoap.CancelPendingOperation) (*elmsoap.CancelPendingOperationResponse, error) {
	panic("baseApiSoap.CancelPendingOperation not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CancelPendingOperationContext(ctx context.Context, request *elmsoap.CancelPendingOperation) (*elmsoap.CancelPendingOperationResponse, error) {
	panic("baseApiSoap.CancelPendingOperationContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportOs(request *elmsoap.ImportOs) (*elmsoap.ImportOsResponse, error) {
	panic("baseApiSoap.ImportOs not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportOsContext(ctx context.Context, request *elmsoap.ImportOs) (*elmsoap.ImportOsResponse, error) {
	panic("baseApiSoap.ImportOsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryOsLayers(request *elmsoap.QueryOsLayers) (*elmsoap.QueryOsLayersResponse, error) {
	panic("baseApiSoap.QueryOsLayers not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryOsLayersContext(ctx context.Context, request *elmsoap.QueryOsLayers) (*elmsoap.QueryOsLayersResponse, error) {
	panic("baseApiSoap.QueryOsLayersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryOsLayerDetails(request *elmsoap.QueryOsLayerDetails) (*elmsoap.QueryOsLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryOsLayerDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryOsLayerDetailsContext(ctx context.Context, request *elmsoap.QueryOsLayerDetails) (*elmsoap.QueryOsLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryOsLayerDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerRevisions(request *elmsoap.QueryLayerRevisions) (*elmsoap.QueryLayerRevisionsResponse, error) {
	panic("baseApiSoap.QueryLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerRevisionsContext(ctx context.Context, request *elmsoap.QueryLayerRevisions) (*elmsoap.QueryLayerRevisionsResponse, error) {
	panic("baseApiSoap.QueryLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteOsLayerRevisions(request *elmsoap.DeleteOsLayerRevisions) (*elmsoap.DeleteOsLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeleteOsLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteOsLayerRevisionsContext(ctx context.Context, request *elmsoap.DeleteOsLayerRevisions) (*elmsoap.DeleteOsLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeleteOsLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteAppLayerRevisions(request *elmsoap.DeleteAppLayerRevisions) (*elmsoap.DeleteAppLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeleteAppLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteAppLayerRevisionsContext(ctx context.Context, request *elmsoap.DeleteAppLayerRevisions) (*elmsoap.DeleteAppLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeleteAppLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeletePlatformLayerRevisions(request *elmsoap.DeletePlatformLayerRevisions) (*elmsoap.DeletePlatformLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeletePlatformLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeletePlatformLayerRevisionsContext(ctx context.Context, request *elmsoap.DeletePlatformLayerRevisions) (*elmsoap.DeletePlatformLayerRevisionsResponse, error) {
	panic("baseApiSoap.DeletePlatformLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateApplicationLayer(request *elmsoap.CreateApplicationLayer) (*elmsoap.CreateApplicationLayerResponse, error) {
	panic("baseApiSoap.CreateApplicationLayer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateApplicationLayerContext(ctx context.Context, request *elmsoap.CreateApplicationLayer) (*elmsoap.CreateApplicationLayerResponse, error) {
	panic("baseApiSoap.CreateApplicationLayerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CloneLayer(request *elmsoap.CloneLayer) (*elmsoap.CloneLayerResponse, error) {
	panic("baseApiSoap.CloneLayer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CloneLayerContext(ctx context.Context, request *elmsoap.CloneLayer) (*elmsoap.CloneLayerResponse, error) {
	panic("baseApiSoap.CloneLayerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreatePlatformLayer(request *elmsoap.CreatePlatformLayer) (*elmsoap.CreatePlatformLayerResponse, error) {
	panic("baseApiSoap.CreatePlatformLayer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreatePlatformLayerContext(ctx context.Context, request *elmsoap.CreatePlatformLayer) (*elmsoap.CreatePlatformLayerResponse, error) {
	panic("baseApiSoap.CreatePlatformLayerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) FinalizeLayerRevision(request *elmsoap.FinalizeLayerRevision) (*elmsoap.FinalizeLayerRevisionResponse, error) {
	panic("baseApiSoap.FinalizeLayerRevision not implemented in test — embed and override this method")
}

func (b *baseApiSoap) FinalizeLayerRevisionContext(ctx context.Context, request *elmsoap.FinalizeLayerRevision) (*elmsoap.FinalizeLayerRevisionResponse, error) {
	panic("baseApiSoap.FinalizeLayerRevisionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerInstallDisk(request *elmsoap.QueryLayerInstallDisk) (*elmsoap.QueryLayerInstallDiskResponse, error) {
	panic("baseApiSoap.QueryLayerInstallDisk not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryLayerInstallDiskContext(ctx context.Context, request *elmsoap.QueryLayerInstallDisk) (*elmsoap.QueryLayerInstallDiskResponse, error) {
	panic("baseApiSoap.QueryLayerInstallDiskContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryApplicationLayers(request *elmsoap.QueryApplicationLayers) (*elmsoap.QueryApplicationLayersResponse, error) {
	panic("baseApiSoap.QueryApplicationLayers not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryApplicationLayersContext(ctx context.Context, request *elmsoap.QueryApplicationLayers) (*elmsoap.QueryApplicationLayersResponse, error) {
	panic("baseApiSoap.QueryApplicationLayersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryApplicationLayerDetails(request *elmsoap.QueryApplicationLayerDetails) (*elmsoap.QueryApplicationLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryApplicationLayerDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryApplicationLayerDetailsContext(ctx context.Context, request *elmsoap.QueryApplicationLayerDetails) (*elmsoap.QueryApplicationLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryApplicationLayerDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformLayers(request *elmsoap.QueryPlatformLayers) (*elmsoap.QueryPlatformLayersResponse, error) {
	panic("baseApiSoap.QueryPlatformLayers not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformLayersContext(ctx context.Context, request *elmsoap.QueryPlatformLayers) (*elmsoap.QueryPlatformLayersResponse, error) {
	panic("baseApiSoap.QueryPlatformLayersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformLayerDetails(request *elmsoap.QueryPlatformLayerDetails) (*elmsoap.QueryPlatformLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryPlatformLayerDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformLayerDetailsContext(ctx context.Context, request *elmsoap.QueryPlatformLayerDetails) (*elmsoap.QueryPlatformLayerDetailsResponse, error) {
	panic("baseApiSoap.QueryPlatformLayerDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateIcon(request *elmsoap.CreateIcon) (*elmsoap.CreateIconResponse, error) {
	panic("baseApiSoap.CreateIcon not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateIconContext(ctx context.Context, request *elmsoap.CreateIcon) (*elmsoap.CreateIconResponse, error) {
	panic("baseApiSoap.CreateIconContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetItemsAssociatedWithIcon(request *elmsoap.GetItemsAssociatedWithIcon) (*elmsoap.GetItemsAssociatedWithIconResponse, error) {
	panic("baseApiSoap.GetItemsAssociatedWithIcon not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetItemsAssociatedWithIconContext(ctx context.Context, request *elmsoap.GetItemsAssociatedWithIcon) (*elmsoap.GetItemsAssociatedWithIconResponse, error) {
	panic("baseApiSoap.GetItemsAssociatedWithIconContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteIcon(request *elmsoap.DeleteIcon) (*elmsoap.DeleteIconResponse, error) {
	panic("baseApiSoap.DeleteIcon not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteIconContext(ctx context.Context, request *elmsoap.DeleteIcon) (*elmsoap.DeleteIconResponse, error) {
	panic("baseApiSoap.DeleteIconContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateAppLayerRevision(request *elmsoap.CreateAppLayerRevision) (*elmsoap.CreateAppLayerRevisionResponse, error) {
	panic("baseApiSoap.CreateAppLayerRevision not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateAppLayerRevisionContext(ctx context.Context, request *elmsoap.CreateAppLayerRevision) (*elmsoap.CreateAppLayerRevisionResponse, error) {
	panic("baseApiSoap.CreateAppLayerRevisionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreatePlatformLayerRevision(request *elmsoap.CreatePlatformLayerRevision) (*elmsoap.CreatePlatformLayerRevisionResponse, error) {
	panic("baseApiSoap.CreatePlatformLayerRevision not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreatePlatformLayerRevisionContext(ctx context.Context, request *elmsoap.CreatePlatformLayerRevision) (*elmsoap.CreatePlatformLayerRevisionResponse, error) {
	panic("baseApiSoap.CreatePlatformLayerRevisionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateOsLayerRevision(request *elmsoap.CreateOsLayerRevision) (*elmsoap.CreateOsLayerRevisionResponse, error) {
	panic("baseApiSoap.CreateOsLayerRevision not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateOsLayerRevisionContext(ctx context.Context, request *elmsoap.CreateOsLayerRevision) (*elmsoap.CreateOsLayerRevisionResponse, error) {
	panic("baseApiSoap.CreateOsLayerRevisionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpgradeElm(request *elmsoap.UpgradeElm) (*elmsoap.UpgradeElmResponse, error) {
	panic("baseApiSoap.UpgradeElm not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpgradeElmContext(ctx context.Context, request *elmsoap.UpgradeElm) (*elmsoap.UpgradeElmResponse, error) {
	panic("baseApiSoap.UpgradeElmContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUpgradeStatus(request *elmsoap.QueryUpgradeStatus) (*elmsoap.QueryUpgradeStatusResponse, error) {
	panic("baseApiSoap.QueryUpgradeStatus not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUpgradeStatusContext(ctx context.Context, request *elmsoap.QueryUpgradeStatus) (*elmsoap.QueryUpgradeStatusResponse, error) {
	panic("baseApiSoap.QueryUpgradeStatusContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetWelcomeWindowState(request *elmsoap.GetWelcomeWindowState) (*elmsoap.GetWelcomeWindowStateResponse, error) {
	panic("baseApiSoap.GetWelcomeWindowState not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetWelcomeWindowStateContext(ctx context.Context, request *elmsoap.GetWelcomeWindowState) (*elmsoap.GetWelcomeWindowStateResponse, error) {
	panic("baseApiSoap.GetWelcomeWindowStateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GatherDiagnostics(request *elmsoap.GatherDiagnostics) (*elmsoap.GatherDiagnosticsResponse, error) {
	panic("baseApiSoap.GatherDiagnostics not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GatherDiagnosticsContext(ctx context.Context, request *elmsoap.GatherDiagnostics) (*elmsoap.GatherDiagnosticsResponse, error) {
	panic("baseApiSoap.GatherDiagnosticsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportLdapItem(request *elmsoap.ImportLdapItem) (*elmsoap.ImportLdapItemResponse, error) {
	panic("baseApiSoap.ImportLdapItem not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportLdapItemContext(ctx context.Context, request *elmsoap.ImportLdapItem) (*elmsoap.ImportLdapItemResponse, error) {
	panic("baseApiSoap.ImportLdapItemContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetNotifications(request *elmsoap.GetNotifications) (*elmsoap.GetNotificationsResponse, error) {
	panic("baseApiSoap.GetNotifications not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetNotificationsContext(ctx context.Context, request *elmsoap.GetNotifications) (*elmsoap.GetNotificationsResponse, error) {
	panic("baseApiSoap.GetNotificationsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditLayer(request *elmsoap.EditLayer) (*elmsoap.EditLayerResponse, error) {
	panic("baseApiSoap.EditLayer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditLayerContext(ctx context.Context, request *elmsoap.EditLayer) (*elmsoap.EditLayerResponse, error) {
	panic("baseApiSoap.EditLayerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PreviewLicense(request *elmsoap.PreviewLicense) (*elmsoap.PreviewLicenseResponse, error) {
	panic("baseApiSoap.PreviewLicense not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PreviewLicenseContext(ctx context.Context, request *elmsoap.PreviewLicense) (*elmsoap.PreviewLicenseResponse, error) {
	panic("baseApiSoap.PreviewLicenseContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetLicense(request *elmsoap.GetLicense) (*elmsoap.GetLicenseResponse, error) {
	panic("baseApiSoap.GetLicense not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetLicenseContext(ctx context.Context, request *elmsoap.GetLicense) (*elmsoap.GetLicenseResponse, error) {
	panic("baseApiSoap.GetLicenseContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetLicense(request *elmsoap.SetLicense) (*elmsoap.SetLicenseResponse, error) {
	panic("baseApiSoap.SetLicense not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetLicenseContext(ctx context.Context, request *elmsoap.SetLicense) (*elmsoap.SetLicenseResponse, error) {
	panic("baseApiSoap.SetLicenseContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CurrentHttpCertificate(request *elmsoap.CurrentHttpCertificate) (*elmsoap.CurrentHttpCertificateResponse, error) {
	panic("baseApiSoap.CurrentHttpCertificate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CurrentHttpCertificateContext(ctx context.Context, request *elmsoap.CurrentHttpCertificate) (*elmsoap.CurrentHttpCertificateResponse, error) {
	panic("baseApiSoap.CurrentHttpCertificateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) StoreHttpCertificate(request *elmsoap.StoreHttpCertificate) (*elmsoap.StoreHttpCertificateResponse, error) {
	panic("baseApiSoap.StoreHttpCertificate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) StoreHttpCertificateContext(ctx context.Context, request *elmsoap.StoreHttpCertificate) (*elmsoap.StoreHttpCertificateResponse, error) {
	panic("baseApiSoap.StoreHttpCertificateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PreviewHttpCertificate(request *elmsoap.PreviewHttpCertificate) (*elmsoap.PreviewHttpCertificateResponse, error) {
	panic("baseApiSoap.PreviewHttpCertificate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PreviewHttpCertificateContext(ctx context.Context, request *elmsoap.PreviewHttpCertificate) (*elmsoap.PreviewHttpCertificateResponse, error) {
	panic("baseApiSoap.PreviewHttpCertificateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GenerateHttpCertificate(request *elmsoap.GenerateHttpCertificate) (*elmsoap.GenerateHttpCertificateResponse, error) {
	panic("baseApiSoap.GenerateHttpCertificate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GenerateHttpCertificateContext(ctx context.Context, request *elmsoap.GenerateHttpCertificate) (*elmsoap.GenerateHttpCertificateResponse, error) {
	panic("baseApiSoap.GenerateHttpCertificateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetStorageRecords(request *elmsoap.GetStorageRecords) (*elmsoap.GetStorageRecordsResponse, error) {
	panic("baseApiSoap.GetStorageRecords not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetStorageRecordsContext(ctx context.Context, request *elmsoap.GetStorageRecords) (*elmsoap.GetStorageRecordsResponse, error) {
	panic("baseApiSoap.GetStorageRecordsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageSummary(request *elmsoap.QueryImageSummary) (*elmsoap.QueryImageSummaryResponse, error) {
	panic("baseApiSoap.QueryImageSummary not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageSummaryContext(ctx context.Context, request *elmsoap.QueryImageSummary) (*elmsoap.QueryImageSummaryResponse, error) {
	panic("baseApiSoap.QueryImageSummaryContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageDetail(request *elmsoap.QueryImageDetail) (*elmsoap.QueryImageDetailResponse, error) {
	panic("baseApiSoap.QueryImageDetail not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageDetailContext(ctx context.Context, request *elmsoap.QueryImageDetail) (*elmsoap.QueryImageDetailResponse, error) {
	panic("baseApiSoap.QueryImageDetailContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageByLayerRevision(request *elmsoap.QueryImageByLayerRevision) (*elmsoap.QueryImageByLayerRevisionResponse, error) {
	panic("baseApiSoap.QueryImageByLayerRevision not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImageByLayerRevisionContext(ctx context.Context, request *elmsoap.QueryImageByLayerRevision) (*elmsoap.QueryImageByLayerRevisionResponse, error) {
	panic("baseApiSoap.QueryImageByLayerRevisionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRemoteFileShares(request *elmsoap.QueryRemoteFileShares) (*elmsoap.QueryRemoteFileSharesResponse, error) {
	panic("baseApiSoap.QueryRemoteFileShares not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRemoteFileSharesContext(ctx context.Context, request *elmsoap.QueryRemoteFileShares) (*elmsoap.QueryRemoteFileSharesResponse, error) {
	panic("baseApiSoap.QueryRemoteFileSharesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteRemoteFileShares(request *elmsoap.DeleteRemoteFileShares) (*elmsoap.DeleteRemoteFileSharesResponse, error) {
	panic("baseApiSoap.DeleteRemoteFileShares not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteRemoteFileSharesContext(ctx context.Context, request *elmsoap.DeleteRemoteFileShares) (*elmsoap.DeleteRemoteFileSharesResponse, error) {
	panic("baseApiSoap.DeleteRemoteFileSharesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectors(request *elmsoap.QueryPlatformConnectors) (*elmsoap.QueryPlatformConnectorsResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectors not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorsContext(ctx context.Context, request *elmsoap.QueryPlatformConnectors) (*elmsoap.QueryPlatformConnectorsResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfig(request *elmsoap.QueryPlatformConnectorConfig) (*elmsoap.QueryPlatformConnectorConfigResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfig not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfigContext(ctx context.Context, request *elmsoap.QueryPlatformConnectorConfig) (*elmsoap.QueryPlatformConnectorConfigResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfigContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfigSummary(request *elmsoap.QueryPlatformConnectorConfigSummary) (*elmsoap.QueryPlatformConnectorConfigSummaryResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfigSummary not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfigSummaryContext(ctx context.Context, request *elmsoap.QueryPlatformConnectorConfigSummary) (*elmsoap.QueryPlatformConnectorConfigSummaryResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfigSummaryContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfigDetails(request *elmsoap.QueryPlatformConnectorConfigDetails) (*elmsoap.QueryPlatformConnectorConfigDetailsResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfigDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConnectorConfigDetailsContext(ctx context.Context, request *elmsoap.QueryPlatformConnectorConfigDetails) (*elmsoap.QueryPlatformConnectorConfigDetailsResponse, error) {
	panic("baseApiSoap.QueryPlatformConnectorConfigDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectory(request *elmsoap.QueryDirectory) (*elmsoap.QueryDirectoryResponse, error) {
	panic("baseApiSoap.QueryDirectory not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryContext(ctx context.Context, request *elmsoap.QueryDirectory) (*elmsoap.QueryDirectoryResponse, error) {
	panic("baseApiSoap.QueryDirectoryContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetRemoteFileShare(request *elmsoap.SetRemoteFileShare) (*elmsoap.SetRemoteFileShareResponse, error) {
	panic("baseApiSoap.SetRemoteFileShare not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetRemoteFileShareContext(ctx context.Context, request *elmsoap.SetRemoteFileShare) (*elmsoap.SetRemoteFileShareResponse, error) {
	panic("baseApiSoap.SetRemoteFileShareContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestRemoteFileShare(request *elmsoap.TestRemoteFileShare) (*elmsoap.TestRemoteFileShareResponse, error) {
	panic("baseApiSoap.TestRemoteFileShare not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestRemoteFileShareContext(ctx context.Context, request *elmsoap.TestRemoteFileShare) (*elmsoap.TestRemoteFileShareResponse, error) {
	panic("baseApiSoap.TestRemoteFileShareContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetAvailableDrives(request *elmsoap.GetAvailableDrives) (*elmsoap.GetAvailableDrivesResponse, error) {
	panic("baseApiSoap.GetAvailableDrives not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetAvailableDrivesContext(ctx context.Context, request *elmsoap.GetAvailableDrives) (*elmsoap.GetAvailableDrivesResponse, error) {
	panic("baseApiSoap.GetAvailableDrivesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryCreate(request *elmsoap.DirectoryCreate) (*elmsoap.DirectoryCreateResponse, error) {
	panic("baseApiSoap.DirectoryCreate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryCreateContext(ctx context.Context, request *elmsoap.DirectoryCreate) (*elmsoap.DirectoryCreateResponse, error) {
	panic("baseApiSoap.DirectoryCreateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryUpdate(request *elmsoap.DirectoryUpdate) (*elmsoap.DirectoryUpdateResponse, error) {
	panic("baseApiSoap.DirectoryUpdate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryUpdateContext(ctx context.Context, request *elmsoap.DirectoryUpdate) (*elmsoap.DirectoryUpdateResponse, error) {
	panic("baseApiSoap.DirectoryUpdateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryDelete(request *elmsoap.DirectoryDelete) (*elmsoap.DirectoryDeleteResponse, error) {
	panic("baseApiSoap.DirectoryDelete not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DirectoryDeleteContext(ctx context.Context, request *elmsoap.DirectoryDelete) (*elmsoap.DirectoryDeleteResponse, error) {
	panic("baseApiSoap.DirectoryDeleteContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRecipes(request *elmsoap.QueryRecipes) (*elmsoap.QueryRecipesResponse, error) {
	panic("baseApiSoap.QueryRecipes not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRecipesContext(ctx context.Context, request *elmsoap.QueryRecipes) (*elmsoap.QueryRecipesResponse, error) {
	panic("baseApiSoap.QueryRecipesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypes(request *elmsoap.QueryPlatformTypes) (*elmsoap.QueryPlatformTypesResponse, error) {
	panic("baseApiSoap.QueryPlatformTypes not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypesContext(ctx context.Context, request *elmsoap.QueryPlatformTypes) (*elmsoap.QueryPlatformTypesResponse, error) {
	panic("baseApiSoap.QueryPlatformTypesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditPlatformTypesAssociations(request *elmsoap.EditPlatformTypesAssociations) (*elmsoap.EditPlatformTypesAssociationsResponse, error) {
	panic("baseApiSoap.EditPlatformTypesAssociations not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditPlatformTypesAssociationsContext(ctx context.Context, request *elmsoap.EditPlatformTypesAssociations) (*elmsoap.EditPlatformTypesAssociationsResponse, error) {
	panic("baseApiSoap.EditPlatformTypesAssociationsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateAppLayerAssignment(request *elmsoap.UpdateAppLayerAssignment) (*elmsoap.UpdateAppLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdateAppLayerAssignment not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateAppLayerAssignmentContext(ctx context.Context, request *elmsoap.UpdateAppLayerAssignment) (*elmsoap.UpdateAppLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdateAppLayerAssignmentContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RemoveAppLayerAssignment(request *elmsoap.RemoveAppLayerAssignment) (*elmsoap.RemoveAppLayerAssignmentResponse, error) {
	panic("baseApiSoap.RemoveAppLayerAssignment not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RemoveAppLayerAssignmentContext(ctx context.Context, request *elmsoap.RemoveAppLayerAssignment) (*elmsoap.RemoveAppLayerAssignmentResponse, error) {
	panic("baseApiSoap.RemoveAppLayerAssignmentContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateOsLayerAssignment(request *elmsoap.UpdateOsLayerAssignment) (*elmsoap.UpdateOsLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdateOsLayerAssignment not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateOsLayerAssignmentContext(ctx context.Context, request *elmsoap.UpdateOsLayerAssignment) (*elmsoap.UpdateOsLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdateOsLayerAssignmentContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdatePlatformLayerAssignment(request *elmsoap.UpdatePlatformLayerAssignment) (*elmsoap.UpdatePlatformLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdatePlatformLayerAssignment not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdatePlatformLayerAssignmentContext(ctx context.Context, request *elmsoap.UpdatePlatformLayerAssignment) (*elmsoap.UpdatePlatformLayerAssignmentResponse, error) {
	panic("baseApiSoap.UpdatePlatformLayerAssignmentContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypeHelpUrl(request *elmsoap.QueryPlatformTypeHelpUrl) (*elmsoap.QueryPlatformTypeHelpUrlResponse, error) {
	panic("baseApiSoap.QueryPlatformTypeHelpUrl not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypeHelpUrlContext(ctx context.Context, request *elmsoap.QueryPlatformTypeHelpUrl) (*elmsoap.QueryPlatformTypeHelpUrlResponse, error) {
	panic("baseApiSoap.QueryPlatformTypeHelpUrlContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypeLayerHelpUrl(request *elmsoap.QueryPlatformTypeLayerHelpUrl) (*elmsoap.QueryPlatformTypeLayerHelpUrlResponse, error) {
	panic("baseApiSoap.QueryPlatformTypeLayerHelpUrl not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformTypeLayerHelpUrlContext(ctx context.Context, request *elmsoap.QueryPlatformTypeLayerHelpUrl) (*elmsoap.QueryPlatformTypeLayerHelpUrlResponse, error) {
	panic("baseApiSoap.QueryPlatformTypeLayerHelpUrlContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFreeDisks(request *elmsoap.QueryFreeDisks) (*elmsoap.QueryFreeDisksResponse, error) {
	panic("baseApiSoap.QueryFreeDisks not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFreeDisksContext(ctx context.Context, request *elmsoap.QueryFreeDisks) (*elmsoap.QueryFreeDisksResponse, error) {
	panic("baseApiSoap.QueryFreeDisksContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExpandStorage(request *elmsoap.ExpandStorage) (*elmsoap.ExpandStorageResponse, error) {
	panic("baseApiSoap.ExpandStorage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExpandStorageContext(ctx context.Context, request *elmsoap.ExpandStorage) (*elmsoap.ExpandStorageResponse, error) {
	panic("baseApiSoap.ExpandStorageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateLayerFolder(request *elmsoap.UpdateLayerFolder) (*elmsoap.UpdateLayerFolderResponse, error) {
	panic("baseApiSoap.UpdateLayerFolder not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateLayerFolderContext(ctx context.Context, request *elmsoap.UpdateLayerFolder) (*elmsoap.UpdateLayerFolderResponse, error) {
	panic("baseApiSoap.UpdateLayerFolderContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) AnalyzeLayerRevisions(request *elmsoap.AnalyzeLayerRevisions) (*elmsoap.AnalyzeLayerRevisionsResponse, error) {
	panic("baseApiSoap.AnalyzeLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) AnalyzeLayerRevisionsContext(ctx context.Context, request *elmsoap.AnalyzeLayerRevisions) (*elmsoap.AnalyzeLayerRevisionsResponse, error) {
	panic("baseApiSoap.AnalyzeLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDiskUsageEstimate(request *elmsoap.QueryDiskUsageEstimate) (*elmsoap.QueryDiskUsageEstimateResponse, error) {
	panic("baseApiSoap.QueryDiskUsageEstimate not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDiskUsageEstimateContext(ctx context.Context, request *elmsoap.QueryDiskUsageEstimate) (*elmsoap.QueryDiskUsageEstimateResponse, error) {
	panic("baseApiSoap.QueryDiskUsageEstimateContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUserRoles(request *elmsoap.QueryUserRoles) (*elmsoap.QueryUserRolesResponse, error) {
	panic("baseApiSoap.QueryUserRoles not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryUserRolesContext(ctx context.Context, request *elmsoap.QueryUserRoles) (*elmsoap.QueryUserRolesResponse, error) {
	panic("baseApiSoap.QueryUserRolesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemGroupMembership(request *elmsoap.QueryDirectoryItemGroupMembership) (*elmsoap.QueryDirectoryItemGroupMembershipResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemGroupMembership not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryItemGroupMembershipContext(ctx context.Context, request *elmsoap.QueryDirectoryItemGroupMembership) (*elmsoap.QueryDirectoryItemGroupMembershipResponse, error) {
	panic("baseApiSoap.QueryDirectoryItemGroupMembershipContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetUserSetting(request *elmsoap.SetUserSetting) (*elmsoap.SetUserSettingResponse, error) {
	panic("baseApiSoap.SetUserSetting not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetUserSettingContext(ctx context.Context, request *elmsoap.SetUserSetting) (*elmsoap.SetUserSettingResponse, error) {
	panic("baseApiSoap.SetUserSettingContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateServerStorageSlot(request *elmsoap.CreateServerStorageSlot) (*elmsoap.CreateServerStorageSlotResponse, error) {
	panic("baseApiSoap.CreateServerStorageSlot not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateServerStorageSlotContext(ctx context.Context, request *elmsoap.CreateServerStorageSlot) (*elmsoap.CreateServerStorageSlotResponse, error) {
	panic("baseApiSoap.CreateServerStorageSlotContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ReadServerStorageSlot(request *elmsoap.ReadServerStorageSlot) (*elmsoap.ReadServerStorageSlotResponse, error) {
	panic("baseApiSoap.ReadServerStorageSlot not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ReadServerStorageSlotContext(ctx context.Context, request *elmsoap.ReadServerStorageSlot) (*elmsoap.ReadServerStorageSlotResponse, error) {
	panic("baseApiSoap.ReadServerStorageSlotContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateServerStorageSlot(request *elmsoap.UpdateServerStorageSlot) (*elmsoap.UpdateServerStorageSlotResponse, error) {
	panic("baseApiSoap.UpdateServerStorageSlot not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateServerStorageSlotContext(ctx context.Context, request *elmsoap.UpdateServerStorageSlot) (*elmsoap.UpdateServerStorageSlotResponse, error) {
	panic("baseApiSoap.UpdateServerStorageSlotContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteServerStorageSlot(request *elmsoap.DeleteServerStorageSlot) (*elmsoap.DeleteServerStorageSlotResponse, error) {
	panic("baseApiSoap.DeleteServerStorageSlot not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteServerStorageSlotContext(ctx context.Context, request *elmsoap.DeleteServerStorageSlot) (*elmsoap.DeleteServerStorageSlotResponse, error) {
	panic("baseApiSoap.DeleteServerStorageSlotContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RunCacheCommand(request *elmsoap.RunCacheCommand) (*elmsoap.RunCacheCommandResponse, error) {
	panic("baseApiSoap.RunCacheCommand not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RunCacheCommandContext(ctx context.Context, request *elmsoap.RunCacheCommand) (*elmsoap.RunCacheCommandResponse, error) {
	panic("baseApiSoap.RunCacheCommandContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DefineFileShare(request *elmsoap.DefineFileShare) (*elmsoap.DefineFileShareResponse, error) {
	panic("baseApiSoap.DefineFileShare not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DefineFileShareContext(ctx context.Context, request *elmsoap.DefineFileShare) (*elmsoap.DefineFileShareResponse, error) {
	panic("baseApiSoap.DefineFileShareContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditFileShare(request *elmsoap.EditFileShare) (*elmsoap.EditFileShareResponse, error) {
	panic("baseApiSoap.EditFileShare not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditFileShareContext(ctx context.Context, request *elmsoap.EditFileShare) (*elmsoap.EditFileShareResponse, error) {
	panic("baseApiSoap.EditFileShareContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFileShareSummary(request *elmsoap.QueryFileShareSummary) (*elmsoap.QueryFileShareSummaryResponse, error) {
	panic("baseApiSoap.QueryFileShareSummary not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFileShareSummaryContext(ctx context.Context, request *elmsoap.QueryFileShareSummary) (*elmsoap.QueryFileShareSummaryResponse, error) {
	panic("baseApiSoap.QueryFileShareSummaryContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFileShareDetails(request *elmsoap.QueryFileShareDetails) (*elmsoap.QueryFileShareDetailsResponse, error) {
	panic("baseApiSoap.QueryFileShareDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryFileShareDetailsContext(ctx context.Context, request *elmsoap.QueryFileShareDetails) (*elmsoap.QueryFileShareDetailsResponse, error) {
	panic("baseApiSoap.QueryFileShareDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteFileShares(request *elmsoap.DeleteFileShares) (*elmsoap.DeleteFileSharesResponse, error) {
	panic("baseApiSoap.DeleteFileShares not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteFileSharesContext(ctx context.Context, request *elmsoap.DeleteFileShares) (*elmsoap.DeleteFileSharesResponse, error) {
	panic("baseApiSoap.DeleteFileSharesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ReorderFileShares(request *elmsoap.ReorderFileShares) (*elmsoap.ReorderFileSharesResponse, error) {
	panic("baseApiSoap.ReorderFileShares not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ReorderFileSharesContext(ctx context.Context, request *elmsoap.ReorderFileShares) (*elmsoap.ReorderFileSharesResponse, error) {
	panic("baseApiSoap.ReorderFileSharesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdatePasswords(request *elmsoap.UpdatePasswords) (*elmsoap.UpdatePasswordsResponse, error) {
	panic("baseApiSoap.UpdatePasswords not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdatePasswordsContext(ctx context.Context, request *elmsoap.UpdatePasswords) (*elmsoap.UpdatePasswordsResponse, error) {
	panic("baseApiSoap.UpdatePasswordsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DefaultFileShareMessages(request *elmsoap.DefaultFileShareMessages) (*elmsoap.DefaultFileShareMessagesResponse, error) {
	panic("baseApiSoap.DefaultFileShareMessages not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DefaultFileShareMessagesContext(ctx context.Context, request *elmsoap.DefaultFileShareMessages) (*elmsoap.DefaultFileShareMessagesResponse, error) {
	panic("baseApiSoap.DefaultFileShareMessagesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryExportableRevisions(request *elmsoap.QueryExportableRevisions) (*elmsoap.QueryExportableRevisionsResponse, error) {
	panic("baseApiSoap.QueryExportableRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryExportableRevisionsContext(ctx context.Context, request *elmsoap.QueryExportableRevisions) (*elmsoap.QueryExportableRevisionsResponse, error) {
	panic("baseApiSoap.QueryExportableRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRemoteFileShareAvailableSpace(request *elmsoap.QueryRemoteFileShareAvailableSpace) (*elmsoap.QueryRemoteFileShareAvailableSpaceResponse, error) {
	panic("baseApiSoap.QueryRemoteFileShareAvailableSpace not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryRemoteFileShareAvailableSpaceContext(ctx context.Context, request *elmsoap.QueryRemoteFileShareAvailableSpace) (*elmsoap.QueryRemoteFileShareAvailableSpaceResponse, error) {
	panic("baseApiSoap.QueryRemoteFileShareAvailableSpaceContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExportLayerRevisions(request *elmsoap.ExportLayerRevisions) (*elmsoap.ExportLayerRevisionsResponse, error) {
	panic("baseApiSoap.ExportLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExportLayerRevisionsContext(ctx context.Context, request *elmsoap.ExportLayerRevisions) (*elmsoap.ExportLayerRevisionsResponse, error) {
	panic("baseApiSoap.ExportLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableRevisions(request *elmsoap.QueryImportableRevisions) (*elmsoap.QueryImportableRevisionsResponse, error) {
	panic("baseApiSoap.QueryImportableRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableRevisionsContext(ctx context.Context, request *elmsoap.QueryImportableRevisions) (*elmsoap.QueryImportableRevisionsResponse, error) {
	panic("baseApiSoap.QueryImportableRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportLayerRevisions(request *elmsoap.ImportLayerRevisions) (*elmsoap.ImportLayerRevisionsResponse, error) {
	panic("baseApiSoap.ImportLayerRevisions not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ImportLayerRevisionsContext(ctx context.Context, request *elmsoap.ImportLayerRevisions) (*elmsoap.ImportLayerRevisionsResponse, error) {
	panic("baseApiSoap.ImportLayerRevisionsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportExportFileShare(request *elmsoap.QueryImportExportFileShare) (*elmsoap.QueryImportExportFileShareResponse, error) {
	panic("baseApiSoap.QueryImportExportFileShare not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportExportFileShareContext(ctx context.Context, request *elmsoap.QueryImportExportFileShare) (*elmsoap.QueryImportExportFileShareResponse, error) {
	panic("baseApiSoap.QueryImportExportFileShareContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableRevisionDetails(request *elmsoap.QueryImportableRevisionDetails) (*elmsoap.QueryImportableRevisionDetailsResponse, error) {
	panic("baseApiSoap.QueryImportableRevisionDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryImportableRevisionDetailsContext(ctx context.Context, request *elmsoap.QueryImportableRevisionDetails) (*elmsoap.QueryImportableRevisionDetailsResponse, error) {
	panic("baseApiSoap.QueryImportableRevisionDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryExportableRevisionDetails(request *elmsoap.QueryExportableRevisionDetails) (*elmsoap.QueryExportableRevisionDetailsResponse, error) {
	panic("baseApiSoap.QueryExportableRevisionDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryExportableRevisionDetailsContext(ctx context.Context, request *elmsoap.QueryExportableRevisionDetails) (*elmsoap.QueryExportableRevisionDetailsResponse, error) {
	panic("baseApiSoap.QueryExportableRevisionDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RegisterCompositingEngine(request *elmsoap.RegisterCompositingEngine) (*elmsoap.RegisterCompositingEngineResponse, error) {
	panic("baseApiSoap.RegisterCompositingEngine not implemented in test — embed and override this method")
}

func (b *baseApiSoap) RegisterCompositingEngineContext(ctx context.Context, request *elmsoap.RegisterCompositingEngine) (*elmsoap.RegisterCompositingEngineResponse, error) {
	panic("baseApiSoap.RegisterCompositingEngineContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ConfigureCompositingEngineRegistration(request *elmsoap.ConfigureCompositingEngineRegistration) (*elmsoap.ConfigureCompositingEngineRegistrationResponse, error) {
	panic("baseApiSoap.ConfigureCompositingEngineRegistration not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ConfigureCompositingEngineRegistrationContext(ctx context.Context, request *elmsoap.ConfigureCompositingEngineRegistration) (*elmsoap.ConfigureCompositingEngineRegistrationResponse, error) {
	panic("baseApiSoap.ConfigureCompositingEngineRegistrationContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatform(request *elmsoap.QueryPlatform) (*elmsoap.QueryPlatformResponse, error) {
	panic("baseApiSoap.QueryPlatform not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformContext(ctx context.Context, request *elmsoap.QueryPlatform) (*elmsoap.QueryPlatformResponse, error) {
	panic("baseApiSoap.QueryPlatformContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConfig(request *elmsoap.QueryPlatformConfig) (*elmsoap.QueryPlatformConfigResponse, error) {
	panic("baseApiSoap.QueryPlatformConfig not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryPlatformConfigContext(ctx context.Context, request *elmsoap.QueryPlatformConfig) (*elmsoap.QueryPlatformConfigResponse, error) {
	panic("baseApiSoap.QueryPlatformConfigContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateOrUpdatePlatformConfig(request *elmsoap.CreateOrUpdatePlatformConfig) (*elmsoap.CreateOrUpdatePlatformConfigResponse, error) {
	panic("baseApiSoap.CreateOrUpdatePlatformConfig not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateOrUpdatePlatformConfigContext(ctx context.Context, request *elmsoap.CreateOrUpdatePlatformConfig) (*elmsoap.CreateOrUpdatePlatformConfigResponse, error) {
	panic("baseApiSoap.CreateOrUpdatePlatformConfigContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) MigrateAppliance(request *elmsoap.MigrateAppliance) (*elmsoap.MigrateApplianceResponse, error) {
	panic("baseApiSoap.MigrateAppliance not implemented in test — embed and override this method")
}

func (b *baseApiSoap) MigrateApplianceContext(ctx context.Context, request *elmsoap.MigrateAppliance) (*elmsoap.MigrateApplianceResponse, error) {
	panic("baseApiSoap.MigrateApplianceContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PauseResumeMigrateAppliance(request *elmsoap.PauseResumeMigrateAppliance) (*elmsoap.PauseResumeMigrateApplianceResponse, error) {
	panic("baseApiSoap.PauseResumeMigrateAppliance not implemented in test — embed and override this method")
}

func (b *baseApiSoap) PauseResumeMigrateApplianceContext(ctx context.Context, request *elmsoap.PauseResumeMigrateAppliance) (*elmsoap.PauseResumeMigrateApplianceResponse, error) {
	panic("baseApiSoap.PauseResumeMigrateApplianceContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) FinalizeMigrateAppliance(request *elmsoap.FinalizeMigrateAppliance) (*elmsoap.FinalizeMigrateApplianceResponse, error) {
	panic("baseApiSoap.FinalizeMigrateAppliance not implemented in test — embed and override this method")
}

func (b *baseApiSoap) FinalizeMigrateApplianceContext(ctx context.Context, request *elmsoap.FinalizeMigrateAppliance) (*elmsoap.FinalizeMigrateApplianceResponse, error) {
	panic("baseApiSoap.FinalizeMigrateApplianceContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePoints(request *elmsoap.QueryCachePoints) (*elmsoap.QueryCachePointsResponse, error) {
	panic("baseApiSoap.QueryCachePoints not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePointsContext(ctx context.Context, request *elmsoap.QueryCachePoints) (*elmsoap.QueryCachePointsResponse, error) {
	panic("baseApiSoap.QueryCachePointsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePointDetails(request *elmsoap.QueryCachePointDetails) (*elmsoap.QueryCachePointDetailsResponse, error) {
	panic("baseApiSoap.QueryCachePointDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePointDetailsContext(ctx context.Context, request *elmsoap.QueryCachePointDetails) (*elmsoap.QueryCachePointDetailsResponse, error) {
	panic("baseApiSoap.QueryCachePointDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePoint(request *elmsoap.QueryCachePoint) (*elmsoap.QueryCachePointResponse, error) {
	panic("baseApiSoap.QueryCachePoint not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryCachePointContext(ctx context.Context, request *elmsoap.QueryCachePoint) (*elmsoap.QueryCachePointResponse, error) {
	panic("baseApiSoap.QueryCachePointContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateImage(request *elmsoap.CreateImage) (*elmsoap.CreateImageResponse, error) {
	panic("baseApiSoap.CreateImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateImageContext(ctx context.Context, request *elmsoap.CreateImage) (*elmsoap.CreateImageResponse, error) {
	panic("baseApiSoap.CreateImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CloneImage(request *elmsoap.CloneImage) (*elmsoap.CloneImageResponse, error) {
	panic("baseApiSoap.CloneImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CloneImageContext(ctx context.Context, request *elmsoap.CloneImage) (*elmsoap.CloneImageResponse, error) {
	panic("baseApiSoap.CloneImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditImage(request *elmsoap.EditImage) (*elmsoap.EditImageResponse, error) {
	panic("baseApiSoap.EditImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditImageContext(ctx context.Context, request *elmsoap.EditImage) (*elmsoap.EditImageResponse, error) {
	panic("baseApiSoap.EditImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExportImage(request *elmsoap.ExportImage) (*elmsoap.ExportImageResponse, error) {
	panic("baseApiSoap.ExportImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) ExportImageContext(ctx context.Context, request *elmsoap.ExportImage) (*elmsoap.ExportImageResponse, error) {
	panic("baseApiSoap.ExportImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteImage(request *elmsoap.DeleteImage) (*elmsoap.DeleteImageResponse, error) {
	panic("baseApiSoap.DeleteImage not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteImageContext(ctx context.Context, request *elmsoap.DeleteImage) (*elmsoap.DeleteImageResponse, error) {
	panic("baseApiSoap.DeleteImageContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionFolders(request *elmsoap.QueryDirectoryJunctionFolders) (*elmsoap.QueryDirectoryJunctionFoldersResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionFolders not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionFoldersContext(ctx context.Context, request *elmsoap.QueryDirectoryJunctionFolders) (*elmsoap.QueryDirectoryJunctionFoldersResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionFoldersContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionDetails(request *elmsoap.QueryDirectoryJunctionDetails) (*elmsoap.QueryDirectoryJunctionDetailsResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionDetails not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionDetailsContext(ctx context.Context, request *elmsoap.QueryDirectoryJunctionDetails) (*elmsoap.QueryDirectoryJunctionDetailsResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionDetailsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestDirectoryJunction(request *elmsoap.TestDirectoryJunction) (*elmsoap.TestDirectoryJunctionResponse, error) {
	panic("baseApiSoap.TestDirectoryJunction not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestDirectoryJunctionContext(ctx context.Context, request *elmsoap.TestDirectoryJunction) (*elmsoap.TestDirectoryJunctionResponse, error) {
	panic("baseApiSoap.TestDirectoryJunctionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateDirectoryJunction(request *elmsoap.CreateDirectoryJunction) (*elmsoap.CreateDirectoryJunctionResponse, error) {
	panic("baseApiSoap.CreateDirectoryJunction not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateDirectoryJunctionContext(ctx context.Context, request *elmsoap.CreateDirectoryJunction) (*elmsoap.CreateDirectoryJunctionResponse, error) {
	panic("baseApiSoap.CreateDirectoryJunctionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditDirectoryJunction(request *elmsoap.EditDirectoryJunction) (*elmsoap.EditDirectoryJunctionResponse, error) {
	panic("baseApiSoap.EditDirectoryJunction not implemented in test — embed and override this method")
}

func (b *baseApiSoap) EditDirectoryJunctionContext(ctx context.Context, request *elmsoap.EditDirectoryJunction) (*elmsoap.EditDirectoryJunctionResponse, error) {
	panic("baseApiSoap.EditDirectoryJunctionContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionAttributes(request *elmsoap.QueryDirectoryJunctionAttributes) (*elmsoap.QueryDirectoryJunctionAttributesResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionAttributes not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryDirectoryJunctionAttributesContext(ctx context.Context, request *elmsoap.QueryDirectoryJunctionAttributes) (*elmsoap.QueryDirectoryJunctionAttributesResponse, error) {
	panic("baseApiSoap.QueryDirectoryJunctionAttributesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteCloudController(request *elmsoap.DeleteCloudController) (*elmsoap.DeleteCloudControllerResponse, error) {
	panic("baseApiSoap.DeleteCloudController not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteCloudControllerContext(ctx context.Context, request *elmsoap.DeleteCloudController) (*elmsoap.DeleteCloudControllerResponse, error) {
	panic("baseApiSoap.DeleteCloudControllerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestNetBiosName(request *elmsoap.TestNetBiosName) (*elmsoap.TestNetBiosNameResponse, error) {
	panic("baseApiSoap.TestNetBiosName not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestNetBiosNameContext(ctx context.Context, request *elmsoap.TestNetBiosName) (*elmsoap.TestNetBiosNameResponse, error) {
	panic("baseApiSoap.TestNetBiosNameContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryMaintenanceSchedules(request *elmsoap.QueryMaintenanceSchedules) (*elmsoap.QueryMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.QueryMaintenanceSchedules not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryMaintenanceSchedulesContext(ctx context.Context, request *elmsoap.QueryMaintenanceSchedules) (*elmsoap.QueryMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.QueryMaintenanceSchedulesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateMaintenanceSchedules(request *elmsoap.CreateMaintenanceSchedules) (*elmsoap.CreateMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.CreateMaintenanceSchedules not implemented in test — embed and override this method")
}

func (b *baseApiSoap) CreateMaintenanceSchedulesContext(ctx context.Context, request *elmsoap.CreateMaintenanceSchedules) (*elmsoap.CreateMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.CreateMaintenanceSchedulesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateMaintenanceSchedules(request *elmsoap.UpdateMaintenanceSchedules) (*elmsoap.UpdateMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.UpdateMaintenanceSchedules not implemented in test — embed and override this method")
}

func (b *baseApiSoap) UpdateMaintenanceSchedulesContext(ctx context.Context, request *elmsoap.UpdateMaintenanceSchedules) (*elmsoap.UpdateMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.UpdateMaintenanceSchedulesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteMaintenanceSchedules(request *elmsoap.DeleteMaintenanceSchedules) (*elmsoap.DeleteMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.DeleteMaintenanceSchedules not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteMaintenanceSchedulesContext(ctx context.Context, request *elmsoap.DeleteMaintenanceSchedules) (*elmsoap.DeleteMaintenanceSchedulesResponse, error) {
	panic("baseApiSoap.DeleteMaintenanceSchedulesContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeletePlatformConnectorConfiguration(request *elmsoap.DeletePlatformConnectorConfiguration) (*elmsoap.DeletePlatformConnectorConfigurationResponse, error) {
	panic("baseApiSoap.DeletePlatformConnectorConfiguration not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeletePlatformConnectorConfigurationContext(ctx context.Context, request *elmsoap.DeletePlatformConnectorConfiguration) (*elmsoap.DeletePlatformConnectorConfigurationResponse, error) {
	panic("baseApiSoap.DeletePlatformConnectorConfigurationContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QuerySystemSettings(request *elmsoap.QuerySystemSettings) (*elmsoap.QuerySystemSettingsResponse, error) {
	panic("baseApiSoap.QuerySystemSettings not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QuerySystemSettingsContext(ctx context.Context, request *elmsoap.QuerySystemSettings) (*elmsoap.QuerySystemSettingsResponse, error) {
	panic("baseApiSoap.QuerySystemSettingsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetSystemSettings(request *elmsoap.SetSystemSettings) (*elmsoap.SetSystemSettingsResponse, error) {
	panic("baseApiSoap.SetSystemSettings not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetSystemSettingsContext(ctx context.Context, request *elmsoap.SetSystemSettings) (*elmsoap.SetSystemSettingsResponse, error) {
	panic("baseApiSoap.SetSystemSettingsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteSystemSettings(request *elmsoap.DeleteSystemSettings) (*elmsoap.DeleteSystemSettingsResponse, error) {
	panic("baseApiSoap.DeleteSystemSettings not implemented in test — embed and override this method")
}

func (b *baseApiSoap) DeleteSystemSettingsContext(ctx context.Context, request *elmsoap.DeleteSystemSettings) (*elmsoap.DeleteSystemSettingsResponse, error) {
	panic("baseApiSoap.DeleteSystemSettingsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryNotificationSettings(request *elmsoap.QueryNotificationSettings) (*elmsoap.QueryNotificationSettingsResponse, error) {
	panic("baseApiSoap.QueryNotificationSettings not implemented in test — embed and override this method")
}

func (b *baseApiSoap) QueryNotificationSettingsContext(ctx context.Context, request *elmsoap.QueryNotificationSettings) (*elmsoap.QueryNotificationSettingsResponse, error) {
	panic("baseApiSoap.QueryNotificationSettingsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetNotificationSettings(request *elmsoap.SetNotificationSettings) (*elmsoap.SetNotificationSettingsResponse, error) {
	panic("baseApiSoap.SetNotificationSettings not implemented in test — embed and override this method")
}

func (b *baseApiSoap) SetNotificationSettingsContext(ctx context.Context, request *elmsoap.SetNotificationSettings) (*elmsoap.SetNotificationSettingsResponse, error) {
	panic("baseApiSoap.SetNotificationSettingsContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestEmailServer(request *elmsoap.TestEmailServer) (*elmsoap.TestEmailServerResponse, error) {
	panic("baseApiSoap.TestEmailServer not implemented in test — embed and override this method")
}

func (b *baseApiSoap) TestEmailServerContext(ctx context.Context, request *elmsoap.TestEmailServer) (*elmsoap.TestEmailServerResponse, error) {
	panic("baseApiSoap.TestEmailServerContext not implemented in test — embed and override this method")
}

func (b *baseApiSoap) GetSoapClient() *soap.Client {
	panic("baseApiSoap.GetSoapClient not implemented in test — embed and override this method")
}

