package nnf

import (
	"stash.us.cray.com/rabsw/ec"
)

type DefaultApiRouter struct {
	servicer   Api
	controller NnfControllerInterface
}

func NewDefaultApiRouter(servicer Api, ctrl NnfControllerInterface) ec.Router {
	return &DefaultApiRouter{servicer: servicer, controller: ctrl}
}

// Name -
func (*DefaultApiRouter) Name() string {
	return "NNF Storage Service Manager"
}

func (r *DefaultApiRouter) Init() error {
	return nil //Initialize(r.controller)
}

func (r *DefaultApiRouter) Start() error {
	return nil
}

func (r *DefaultApiRouter) Routes() ec.Routes {
	s := r.servicer
	return ec.Routes{
		{
			Name:        "RedfishV1StorageServicesGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices",
			HandlerFunc: s.RedfishV1StorageServicesGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdGet,
		},

		/* ------------------------- STORAGE POOLS ------------------------- */
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStoragePoolsGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StoragePools",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStoragePoolsGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStoragePoolsPost",
			Method:      ec.POST_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StoragePools",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStoragePoolsPost,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStoragePoolsStoragePoolIdGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StoragePools/{StoragePoolId}",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStoragePoolsStoragePoolIdGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStoragePoolsStoragePoolIdDelete",
			Method:      ec.DELETE_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StoragePools/{StoragePoolId}",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStoragePoolsStoragePoolIdDelete,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStoragePoolsStoragePoolIdAllocatedVolumesGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StoragePools/{StoragePoolId}/AllocatedVolumes",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStoragePoolsStoragePoolIdAllocatedVolumesGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStoragePoolsStoragePoolIdProvidingVolumesGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StoragePools/{StoragePoolId}/ProvidingVolumes",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStoragePoolsStoragePoolIdProvidingVolumesGet,
		},

		/* ----------------------------- VOLUMES ---------------------------- */

		{
			Name:        "RedfishV1StorageServicesStorageServiceIdVolumesGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/Volumes",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdVolumesGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdVolumesVolumeIdGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/Volumes/{VolumeId}",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdVolumesVolumeIdGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdVolumesVolumeIdProvidingPoolsGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/Volumes/{VolumeId}/ProvidingPools",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdVolumesVolumeIdProvidingPoolsGet,
		},

		/* ------------------------- STORAGE GROUPS ------------------------ */

		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStorageGroupsGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StorageGroups",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStorageGroupsGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStorageGroupsPost",
			Method:      ec.POST_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StorageGroups",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStorageGroupsPost,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStorageGroupsStorageGroupIdGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StorageGroups/{StorageGroupId}",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStorageGroupsStorageGroupIdGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStorageGroupsStorageGroupIdDelete",
			Method:      ec.DELETE_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StorageGroups/{StorageGroupId}",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStorageGroupsStorageGroupIdDelete,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStorageGroupsStorageGroupIdActionsStorageGroupExposeVolumesPost",
			Method:      ec.POST_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StorageGroups/{StorageGroupId}/Actions/StorageGroup.ExposeVolumes",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStorageGroupsStorageGroupIdActionsStorageGroupExposeVolumesPost,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdStorageGroupsStorageGroupIdActionsStorageGroupHideVolumesPost",
			Method:      ec.POST_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/StorageGroups/{StorageGroupId}/Actions/StorageGroup.HideVolumes",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdStorageGroupsStorageGroupIdActionsStorageGroupHideVolumesPost,
		},

		/* --------------------------- ENDPOINTS --------------------------- */

		{
			Name:        "RedfishV1StorageServicesStorageServiceIdEndpointsGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/Endpoints",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdEndpointsGet,
		},

		/* -------------------------- FILE SYSTEMS ------------------------- */

		{
			Name:        "RedfishV1StorageServicesStorageServiceIdFileSystemsGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/FileSystems",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdFileSystemsGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdFileSystemsPost",
			Method:      ec.POST_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/FileSystems",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdFileSystemsPost,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemIdGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/FileSystems/{FileSystemId}",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemIdGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemIdDelete",
			Method:      ec.DELETE_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/FileSystems/{FileSystemId}",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemIdDelete,
		},

		/* --------------------------- FILE SHARES ------------------------- */

		{
			Name:        "RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemsIdExportedFileSharesGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/FileSystems/{FileSystemsId}/ExportedFileShares",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemsIdExportedFileSharesGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemsIdExportedFileSharesPost",
			Method:      ec.POST_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/FileSystems/{FileSystemsId}/ExportedFileShares",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemsIdExportedFileSharesPost,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemsIdExportedFileSharesExportedFileSharesIdGet",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/FileSystems/{FileSystemsId}/ExportedFileShares/{ExportedFileSharesId}",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemsIdExportedFileSharesExportedFileSharesIdGet,
		},
		{
			Name:        "RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemsIdExportedFileSharesExportedFileSharesIdDelete",
			Method:      ec.GET_METHOD,
			Path:        "/redfish/v1/StorageServices/{StorageServiceId}/FileSystems/{FileSystemsId}/ExportedFileShares/{ExportedFileSharesId}",
			HandlerFunc: s.RedfishV1StorageServicesStorageServiceIdFileSystemsFileSystemsIdExportedFileSharesExportedFileSharesIdDelete,
		},
	}
}
