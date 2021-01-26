package nnf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	ec "stash.us.cray.com/rabsw/nnf-ec/ec"
	"stash.us.cray.com/rabsw/nnf-ec/internal/fabric-manager"
	nvmenamespace "stash.us.cray.com/rabsw/nnf-ec/internal/nvme-namespace-manager"
	sf "stash.us.cray.com/sp/rfsf-openapi/pkg/models"
)

// DefaultApiService -
type DefaultApiService struct {
}

// NewDefaultApiService -
func NewDefaultApiService() Api {
	return &DefaultApiService{}
}

// Params -
func Params(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func UnmarshalRequest(r *http.Request, model interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &model)
}

// RedfishV1FabricsGet -
func (*DefaultApiService) RedfishV1FabricsGet(w http.ResponseWriter, r *http.Request) {

	model := sf.FabricCollectionFabricCollection{
		OdataId:   "/redfish/v1/Fabrics",
		OdataType: "#FabricCollection.v1_0_0.FabricCollection",
		Name:      "Fabric Collection",
	}

	err := fabric.Get(&model)

	log.WithError(err).Info("RedfishV1FabricsGet")

	encodeResponse(model, err, w)
}

// RedfishV1FabricsFabricIdGet -
func (*DefaultApiService) RedfishV1FabricsFabricIdGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	fabricId := params["FabricId"]

	model := sf.FabricV120Fabric{
		OdataId:   fmt.Sprintf("/redfish/v1/Fabrics/%s", fabricId),
		OdataType: "#Fabric.v1_0_0.Fabric",
		Id:        fabricId,
		Name:      "Fabric",
	}

	err := fabric.FabricIdGet(fabricId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1FabricsFabricIdSwitchesGet -
func (*DefaultApiService) RedfishV1FabricsFabricIdSwitchesGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	fabricId := params["FabricId"]

	model := sf.SwitchCollectionSwitchCollection{
		OdataId:   fmt.Sprintf("/redfish/v1/Fabrics/%s/Switches", fabricId),
		OdataType: "#SwitchCollection.v1_0_0.SwitchCollection",
		Name:      "Switch Collection",
	}

	err := fabric.FabricIdSwitchesGet(fabricId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1FabricsFabricIdSwitchesSwitchIdGet -
func (*DefaultApiService) RedfishV1FabricsFabricIdSwitchesSwitchIdGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	fabricId := params["FabricId"]
	switchId := params["SwitchId"]

	model := sf.SwitchV140Switch{
		OdataId:   fmt.Sprintf("/redfish/v1/Fabrics/%s/Switches/%s", fabricId, switchId),
		OdataType: "#Switch.v1_4_0.Switch",
		Name:      "Swtich",
	}

	err := fabric.FabricIdSwitchesSwitchIdGet(fabricId, switchId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1FabricsFabricIdSwitchesSwitchIdPortsGet -
func (*DefaultApiService) RedfishV1FabricsFabricIdSwitchesSwitchIdPortsGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	fabricId := params["FabricId"]
	switchId := params["SwitchId"]

	model := sf.PortCollectionPortCollection{
		OdataId:   fmt.Sprintf("/redfish/v1/Fabrics/%s/Switches/%s/Ports", fabricId, switchId),
		OdataType: "#PortCollection.v1_0_0.PortCollection",
		Name:      "Port Collection",
	}

	err := fabric.FabricIdSwitchesSwitchIdPortsGet(fabricId, switchId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1FabricsFabricIdSwitchesSwitchIdPortsPortIdGet -
func (*DefaultApiService) RedfishV1FabricsFabricIdSwitchesSwitchIdPortsPortIdGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	fabricId := params["FabricId"]
	switchId := params["SwitchId"]
	portId := params["PortId"]

	model := sf.PortV130Port{
		OdataId:   fmt.Sprintf("/redfish/v1/Fabrics/%s/Switches/%s/Ports/%s", fabricId, switchId, portId),
		OdataType: "#Port.v1_3_0.Port",
		Name:      "Port",
	}

	err := fabric.FabricIdSwitchesSwitchIdPortsPortIdGet(fabricId, switchId, portId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1FabricsFabricIdEndpointsGet -
func (*DefaultApiService) RedfishV1FabricsFabricIdEndpointsGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	fabricId := params["FabricId"]

	model := sf.EndpointCollectionEndpointCollection{
		OdataId:   fmt.Sprintf("/redfish/v1/Fabrics/%s/Endpoints", fabricId),
		OdataType: "#EndpointCollection.v1_0_0.EndpointCollection",
		Name:      "Endpoint Collection",
	}

	err := fabric.FabricIdEndpointsGet(fabricId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1FabricsFabricIdEndpointsEndpointIdGet -
func (*DefaultApiService) RedfishV1FabricsFabricIdEndpointsEndpointIdGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	fabricId := params["FabricId"]
	endpointId := params["EndpointId"]

	model := sf.EndpointV150Endpoint{
		OdataId:   fmt.Sprintf("/redfish/v1/Fabrics/%s/Endpoints/%s", fabricId, endpointId),
		OdataType: "#Endpoint.v1_5_0.Endpoint",
		Name:      "Endpoint",
	}

	err := fabric.FabricIdEndpointsEndpointIdGet(fabricId, endpointId, &model)

	encodeResponse(model, err, w)
}

func (*DefaultApiService) RedfishV1FabricsFabricIdEndpointGroupsGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	fabricId := params["FabricId"]

	model := sf.EndpointGroupCollectionEndpointGroupCollection{
		OdataId:   fmt.Sprintf("/redfish/v1/Fabrics/%s/EndpointGroups", fabricId),
		OdataType: "#EndpointGroupCollection.v1_0_0.EndpointGroupCollection",
		Name:      "Endpoint Group Collection",
	}

	err := fabric.FabricIdEndpointGroupsGet(fabricId, &model)

	encodeResponse(model, err, w)
}

func (*DefaultApiService) RedfishV1FabricsFabricIdEndpointGroupsEndpointGroupIdGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	fabricId := params["FabricId"]
	groupId := params["EndpointGroupId"]

	model := sf.EndpointGroupV130EndpointGroup{
		OdataId:   fmt.Sprintf("/redfish/v1/Fabrics/%s/EndpointGroups/%s", fabricId, groupId),
		OdataType: "#EndpointGroup.v1_3_0.EndpointGroup",
		Name:      "Endpoint Group",
	}

	err := fabric.FabricIdEndpointGroupsEndpointIdGet(fabricId, groupId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1FabricsFabricIdConnectionsGet -
func (*DefaultApiService) RedfishV1FabricsFabricIdConnectionsGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	fabricId := params["FabricId"]

	model := sf.ConnectionCollectionConnectionCollection{
		OdataId:   fmt.Sprintf("/redfish/v1/Fabrics/%s/Connections", fabricId),
		OdataType: "#ConnectionCollection.v1_0_0.ConnectionCollection",
		Name:      "Connection Collection",
	}

	err := fabric.FabricIdConnectionsGet(fabricId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1FabricsFabricIdConnectionsConnectionIdGet -
func (*DefaultApiService) RedfishV1FabricsFabricIdConnectionsConnectionIdGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	fabricId := params["FabricId"]
	connectionId := params["ConnectionId"]

	model := sf.ConnectionV100Connection{
		OdataId:   fmt.Sprintf("/redfish/v1/Fabrics/%s/Connections/%s", fabricId, connectionId),
		OdataType: "#Connection.v1_0_0.Connection",
		Name:      "Connection",
	}

	err := fabric.FabricIdConnectionsConnectionIdGet(fabricId, connectionId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1StorageGet
func (*DefaultApiService) RedfishV1StorageGet(w http.ResponseWriter, r *http.Request) {

	model := sf.StorageCollectionStorageCollection{
		OdataId:   fmt.Sprintf("/redfish/v1/Storage"),
		OdataType: "#StorageCollection.v1_0_0.StorageCollection",
		Name:      "Storage Collection",
	}

	err := nvmenamespace.Get(&model)

	encodeResponse(model, err, w)
}

// RedfishV1StorageStorageIdGet
func (*DefaultApiService) RedfishV1StorageStorageIdGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	storageId := params["StorageId"]

	model := sf.StorageV190Storage{
		OdataId:   fmt.Sprintf("/redfish/v1/Storage/%s", storageId),
		OdataType: "#Storage.v1_9_0.Storage",
		Name:      "Storage",
	}

	err := nvmenamespace.StorageIdGet(storageId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1StorageStorageIdStoragePoolsGet
func (*DefaultApiService) RedfishV1StorageStorageIdStoragePoolsGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	storageId := params["StorageId"]

	model := sf.StoragePoolCollectionStoragePoolCollection{
		OdataId:   fmt.Sprintf("/redfish/v1/Storage/%s/StoragePools", storageId),
		OdataType: "#StoragePoolCollection.v1_0_0.StoragePoolCollection",
		Name:      "Storage Pool Collection",
	}

	err := nvmenamespace.StorageIdStoragePoolsGet(storageId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1StorageStorageIdStoragePoolsStoragePoolIdGet
func (*DefaultApiService) RedfishV1StorageStorageIdStoragePoolsStoragePoolIdGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	storageId := params["StorageId"]
	storagePoolId := params["StoragePoolId"]

	model := sf.StoragePoolV150StoragePool{
		OdataId:   fmt.Sprintf("/redfish/v1/Storage/%s/StoragePools/%s", storageId, storagePoolId),
		OdataType: "#StoragePool.v1_5_0.StoragePool",
		Name:      "Storage Pool",
	}

	err := nvmenamespace.StorageIdStoragePoolIdGet(storageId, storagePoolId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1StorageStorageIdControllersGet
func (*DefaultApiService) RedfishV1StorageStorageIdControllersGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	storageId := params["StorageId"]

	model := sf.StorageControllerCollectionStorageControllerCollection{
		OdataId:   fmt.Sprintf("/redfish/v1/Storage/%s/StorageControllers", storageId),
		OdataType: "#StorageControllerCollection.v1_0_0.StorageControllerCollection",
		Name:      "Storage Controller Collection",
	}

	err := nvmenamespace.StorageIdControllersGet(storageId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1StorageStorageIdControllersControllerIdGet
func (*DefaultApiService) RedfishV1StorageStorageIdControllersControllerIdGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	storageId := params["StorageId"]
	controllerId := params["ControllerId"]

	model := sf.StorageControllerV100StorageController{
		OdataId:   fmt.Sprintf("/redfish/v1/Storage/%s/StorageControllers/%s", storageId, controllerId),
		OdataType: "#StorageController.v1_0_0.StorageController",
		Name:      "Storage Controller",
	}

	err := nvmenamespace.StorageIdControllerIdGet(storageId, controllerId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1StorageStorageIdVolumesGet
func (*DefaultApiService) RedfishV1StorageStorageIdVolumesGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	storageId := params["StorageId"]

	model := sf.VolumeCollectionVolumeCollection{
		OdataId:   fmt.Sprintf("/redfish/v1/Storage/%s/Volumes", storageId),
		OdataType: "#VolumeCollection.v1_0_0.VolumeCollection",
		Name:      "Volume Collection",
	}

	err := nvmenamespace.StorageIdVolumesGet(storageId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1StorageStorageIdVolumesPost
func (*DefaultApiService) RedfishV1StorageStorageIdVolumesPost(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	storageId := params["StorageId"]

	var model sf.VolumeV161Volume

	if err := UnmarshalRequest(r, &model); err != nil {
		encodeResponse(model, err, w)
	}

	_, err := nvmenamespace.StorageIdVolumePost(storageId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1StorageStorageIdVolumesVolumeIdGet
func (*DefaultApiService) RedfishV1StorageStorageIdVolumesVolumeIdGet(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	storageId := params["StorageId"]
	volumeId := params["VolumeId"]

	model := sf.VolumeV161Volume{
		OdataId:   fmt.Sprintf("/redfish/v1/Storage/%s/Volumes/%s", storageId, volumeId),
		OdataType: "Volume.v1_6_1.Volume",
		Name:      "Volume",
	}

	err := nvmenamespace.StorageIdVolumeIdGet(storageId, volumeId, &model)

	encodeResponse(model, err, w)
}

// RedfishV1StorageStorageIdVolumesVolumeIdDelete
func (*DefaultApiService) RedfishV1StorageStorageIdVolumesVolumeIdDelete(w http.ResponseWriter, r *http.Request) {
	params := Params(r)
	storageId := params["StorageId"]
	volumeId := params["VolumeId"]

	err := nvmenamespace.StorageIdVolumeIdDelete(storageId, volumeId)

	encodeResponse(nil, err, w)
}

func encodeResponse(s interface{}, err error, w http.ResponseWriter) {
	if err != nil {
		log.WithError(err).Warn("Element Controller Error")
	}

	var e *ec.ControllerError
	if errors.As(err, &e) {
		w.WriteHeader(e.StatusCode)
		return
	}

	if s != nil {
		response, err := json.Marshal(s)
		if err != nil {
			log.WithError(err).Error("Failed to marshal json response")
			// TODO
		}
		_, err = w.Write(response)
		if err != nil {
			log.WithError(err).Error("Failed to write json response")
			// TODO
		}
	}
}
