package nvmenamespace

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	. "stash.us.cray.com/rabsw/nnf-ec/internal/api"
	. "stash.us.cray.com/rabsw/nnf-ec/internal/events"

	log "github.com/sirupsen/logrus"

	"stash.us.cray.com/rabsw/ec"
	sf "stash.us.cray.com/rabsw/rfsf-openapi/pkg/models"
	"stash.us.cray.com/rabsw/switchtec-fabric/pkg/nvme"
)

const (
	ResourceBlockId = "Rabbit"
)

const (
	defaultStoragePoolId = "0"
)

// Manager -
type Manager struct {
	id   string
	ctrl NvmeControllerInterface

	config *ConfigFile

	storage []Storage
}

// Storage -
type Storage struct {
	id      string
	address string

	config *ControllerConfig

	virtManagementEnabled bool
	controllers           []StorageController
	volumes               []Volume

	// Capacity in bytes of the storage device. This value is read once and is fixed for
	// the life of the object.
	capacityBytes uint64

	// Unallocted capacity in bytes. This value is updated for any namespaces create or
	// delete operation that might shrink or grow the byte count as expected.
	unallocatedBytes uint64

	// Namespace Properties - Read using the Common Namespace Identifier (0xffffffff)
	// These are properties common to all namespaces for this controller (we use controller
	// zero as the basis for all other controllers - technically the spec supports uinque
	// LBA Formats per controller, but this is not done in practice by drive vendors.)
	lbaFormatIndex uint8
	blockSizeBytes uint32

	// These values allow us to communicate a storage device with its corresponding
	// Fabric Controller. Read once during Port Up Events and remain fixed thereafter.
	fabricId string
	switchId string
	portId   string

	device NvmeDeviceInterface
}

// StorageController -
type StorageController struct {
	id             string
	controllerId   uint16
	functionNumber uint16

	deviceCtrl NvmeDeviceControllerInterface
}

// Volumes -
type Volume struct {
	id          string
	namespaceId nvme.NamespaceIdentifier

	storage             *Storage
	AttachedControllers []*StorageController
}

// TODO: We may want to put this manager under a resource block
//   /​redfish/​v1/​ResourceBlocks/​{ResourceBlockId} // <- Rabbit
//   /​redfish/​v1/​ResourceBlocks/​{ResourceBlockId}/​Systems/{​ComputerSystemId} // <- Also Rabbit & Computes
//   /​redfish/​v1/​ResourceBlocks/​{ResourceBlockId}/​Systems/{​ComputerSystemId}/​PCIeDevices/​{PCIeDeviceId}
//   /​redfish/​v1/​ResourceBlocks/​{ResourceBlockId}/​Systems/{​ComputerSystemId}/​PCIeDevices/​{PCIeDeviceId}/​PCIeFunctions/{​PCIeFunctionId}
//
//   /​redfish/​v1/​ResourceBlocks/{​ResourceBlockId}/​Systems/{​ComputerSystemId}/​Storage/​{StorageId}/​Controllers/​{ControllerId}

var mgr Manager

func init() {
	RegisterNvmeInterface(&mgr)
}

func findStorage(storageId string) *Storage {
	id, err := strconv.Atoi(storageId)
	if err != nil {
		return nil
	}

	if !(id < len(mgr.storage)) {
		return nil
	}

	return &mgr.storage[id]
}

func findStorageController(storageId, controllerId string) (*Storage, *StorageController) {
	s := findStorage(storageId)
	if s == nil {
		return nil, nil
	}

	return s, s.findController(controllerId)
}

func findStorageVolume(storageId, volumeId string) (*Storage, *Volume) {
	s := findStorage(storageId)
	if s == nil {
		return nil, nil
	}

	return s, s.findVolume(volumeId)
}

func findStoragePool(storageId, storagePoolId string) (*Storage, *interface{}) {
	return nil, nil
}

func (m *Manager) fmt(format string, a ...interface{}) string {
	return fmt.Sprintf("/redfish/v1") + fmt.Sprintf(format, a...)
}

// GetVolumes -
func (m *Manager) GetVolumes(controllerId string) ([]string, error) {
	volumes := []string{}
	for _, s := range m.storage {
		c := s.findController(controllerId)
		if c == nil {
			return volumes, ec.ErrNotFound
		}

		nsids, err := s.device.ListNamespaces(c.functionNumber)
		if err != nil {
			return volumes, err
		}

		for _, nsid := range nsids {
			for _, v := range s.volumes {
				if v.namespaceId == nsid {
					volumes = append(volumes, fmt.Sprintf("/redfish/v1/Storage/%s/Volumes/%s", s.id, v.id))
				}
			}
		}

	}

	return volumes, nil
}

func ConvertRelativePortIndexToControllerIndex(index int) (uint16, error) {
	if !(index < mgr.config.Storage.Controller.Functions) {
		return 0, fmt.Errorf("Port Index %d is beyond supported controller count (%d)",
			index, mgr.config.Storage.Controller.Functions)
	}

	return uint16(index + 1), nil
}

func GetStorage() []*Storage {
	storage := make([]*Storage, len(mgr.storage))
	for idx := range storage {
		storage[idx] = &mgr.storage[idx]
	}

	return storage
}

func CreateVolume(s *Storage, capacityBytes uint64) (*Volume, error) {
	return s.createVolume(capacityBytes)
}

func DeleteVolume(v *Volume) error {
	return v.storage.deleteVolume(v.id)
}

func AttachControllers(volume *Volume, controllers []uint16) error {
	return fmt.Errorf("Not Yet Implemented")
}

func DetachControllers(volume *Volume, controllers []uint16) error {
	return fmt.Errorf("Not Yet Implemented")
}

func (s *Storage) UnallocatedBytes() uint64 { return s.unallocatedBytes }
func (s *Storage) IsEnabled() bool          { return s.getStatus().State == sf.ENABLED_RST }

func (s *Storage) fmt(format string, a ...interface{}) string {
	return fmt.Sprintf("/redfish/v1/Storage/%s", s.id) + fmt.Sprintf(format, a...)
}

func (s *Storage) initialize(conf *ControllerConfig, device string) error {
	s.address = device
	s.config = conf

	return nil
}

func (s *Storage) initializeController() error {
	ctrl, err := s.device.IdentifyController()
	if err != nil {
		return err
	}

	capacityToUint64s := func(c [16]byte) (lo uint64, hi uint64) {
		lo, hi = 0, 0
		for i := 0; i < 8; i++ {
			lo, hi = lo<<8, hi<<8
			lo += uint64(c[7-i])
			hi += uint64(c[15-i])
		}

		return lo, hi
	}

	totalCapBytesLo, totalCapBytesHi := capacityToUint64s(ctrl.TotalNVMCapacity)

	s.capacityBytes = totalCapBytesLo
	if totalCapBytesHi != 0 {
		return fmt.Errorf("Unsupported capacity 0x%x_%x: will overflow uint64 definition", totalCapBytesHi, totalCapBytesLo)
	}

	unallocatedCapBytesLo, unallocatedCapBytesHi := capacityToUint64s(ctrl.UnallocatedNVMCapacity)

	s.unallocatedBytes = unallocatedCapBytesLo
	if unallocatedCapBytesHi != 0 {
		return fmt.Errorf("Unsupported unallocated 0x%x_%x, will overflow uint64 definition", unallocatedCapBytesHi, unallocatedCapBytesLo)
	}

	s.virtManagementEnabled = ctrl.GetCapability(nvme.VirtualiztionManagementSupport)

	ns, err := s.device.IdentifyNamespace()
	if err != nil {
		return err
	}

	bestIndex := 0
	bestRelativePerformance := ^uint8(0)
	for i := 0; i < int(ns.NumberOfLBAFormats); i++ {
		if ns.LBAFormats[i].MetadataSize == 0 &&
			ns.LBAFormats[i].RelativePerformance < bestRelativePerformance {
			bestIndex = i
		}
	}
	s.lbaFormatIndex = uint8(bestIndex)
	s.blockSizeBytes = 1 << ns.LBAFormats[bestIndex].LBADataSize

	return nil
}

func (s *Storage) findController(controllerId string) *StorageController {
	for idx, ctrl := range s.controllers {
		if ctrl.id == controllerId {
			return &s.controllers[idx]
		}
	}

	return nil
}

func (s *Storage) getStatus() (stat sf.ResourceStatus) {
	if len(s.controllers) == 0 {
		stat.State = sf.UNAVAILABLE_OFFLINE_RST
	} else {
		stat.Health = sf.OK_RH
		stat.State = sf.ENABLED_RST // TODO: Need some different Unavailable / Offline state maybe?
	}

	return stat
}

func (s *Storage) createVolume(capacityBytes uint64) (*Volume, error) {
	namespaceId, err := s.device.CreateNamespace(capacityBytes)
	// TODO: CreateNamespace can round up the requested capacity
	// Need to pass in a pointer here and then get the updated capacity
	// bytes programmed into the volume.
	if err != nil {
		return nil, err
	}

	id := strconv.Itoa(int(namespaceId))
	s.volumes = append(s.volumes, Volume{
		id:          id,
		namespaceId: namespaceId,
		storage:     s,
	})

	return &s.volumes[len(s.volumes)-1], nil
}

func (s *Storage) deleteVolume(volumeId string) error {
	for idx, volume := range s.volumes {
		if volume.id == volumeId {
			if err := s.device.DeleteNamespace(volume.namespaceId); err != nil {
				return err
			}

			// remove the volume from the array
			copy(s.volumes[idx:], s.volumes[idx+1:]) // shift left 1 at idx
			s.volumes = s.volumes[:len(s.volumes)-1] // truncate tail

			return nil
		}
	}

	return ec.ErrNotFound
}

func (s *Storage) findVolume(volumeId string) *Volume {
	for idx, v := range s.volumes {
		if v.id == volumeId {
			return &s.volumes[idx]
		}
	}

	return nil
}

func (v *Volume) GetOdataId() string {
	return v.storage.fmt("/Volumes/%s", v.id)
}

func (v *Volume) GetCapaityBytes() uint64 {
	return 0
}

// Initialize
func Initialize(ctrl NvmeControllerInterface) error {

	mgr = Manager{
		id:   ResourceBlockId,
		ctrl: ctrl,
	}

	log.SetLevel(log.DebugLevel)

	log.Infof("Initialize %s NVMe Namespace Manager", mgr.id)

	conf, err := loadConfig()
	if err != nil {
		log.WithError(err).Errorf("Failed to load %s configuration", mgr.id)
		return err
	}

	mgr.config = conf

	log.Debugf("NVMe Configuration '%s' Loaded...", conf.Metadata.Name)
	log.Debugf("  Controller Config:")
	log.Debugf("    Virtual Functions: %d", conf.Storage.Controller.Functions)
	log.Debugf("    Num Resources: %d", conf.Storage.Controller.Resources)
	log.Debugf("  Device List: %+v", conf.Storage.Devices)

	mgr.storage = make([]Storage, len(conf.Storage.Devices))
	for storageIdx, storageConfig := range conf.Storage.Devices {
		storage := &mgr.storage[storageIdx]

		storage.id = strconv.Itoa(storageIdx)
		if err := storage.initialize(&conf.Storage.Controller, storageConfig); err != nil {
			log.WithError(err).Errorf("Failed to initialize storage device %s", storage.id)
		}
	}

	PortEventManager.Subscribe(PortEventSubscriber{
		HandlerFunc: PortEventHandler,
		Data:        &mgr,
	})

	return nil
}

// PortEventHandler - Receives port events from the event manager
func PortEventHandler(event PortEvent, data interface{}) {
	m := data.(*Manager)

	log.Infof("%s Port Event Received %+v", m.id, event)

	if event.PortType != PORT_TYPE_DSP {
		return
	}

	// Storage ID is related to the fabric controller that is running. Convert the
	// PortEvent to our storage index.
	idx, err := FabricController.ConvertPortEventToRelativePortIndex(event)
	if err != nil {
		log.WithError(err).Errorf("Unable to find port index for event %+v", event)
		return
	}

	if !(idx < len(m.storage)) {
		log.Errorf("No storage device exists for index %d", idx)
		return
	}

	s := &m.storage[idx]

	switch event.EventType {
	case PORT_EVENT_UP:

		// Connect
		device, err := m.ctrl.NewNvmeDevice(event.FabricId, event.SwitchId, event.PortId)
		if err != nil {
			log.WithError(err).Errorf("Could not allocate storage controller")
			return
		}

		s.device = device

		err = s.initializeController()
		if err != nil {
			log.WithError(err).Errorf("Failied to initialize device controller")
			return
		}

		// Initialize a Storage object with the required number of controllers
		initFunc := func(s *Storage) SecondaryControllersInitFunc {
			return func(count uint8) {
				if count > uint8(s.config.Functions) {
					count = uint8(s.config.Functions)
				}

				s.controllers = make([]StorageController, 1 /*PF*/ +count)

				// Initialize the PF
				s.controllers[0] = StorageController{
					id:             "0",
					controllerId:   0,
					functionNumber: 0,
				}
			}
		}

		handlerFunc := func(s *Storage) SecondaryControllerHandlerFunc {
			return func(controllerId uint16, controllerOnline bool, virtualFunctionNumber uint16, numVQResourcesAssinged, numVIResourcesAssigned uint16) error {

				if controllerId == s.controllers[0].controllerId {
					return fmt.Errorf("Controller ID overlaps with PF")
				}

				if !(controllerId < uint16(len(s.controllers))) {
					return nil
				}

				s.controllers[int(controllerId)] = StorageController{
					id:             strconv.Itoa(int(controllerId)),
					controllerId:   controllerId,
					functionNumber: virtualFunctionNumber,
				}

				if !s.virtManagementEnabled {
					return nil
				}

				if numVQResourcesAssinged != uint16(s.config.Resources) {
					if err := s.device.AssignControllerResources(controllerId, VQResourceType, uint32(int(numVQResourcesAssinged)-s.config.Resources)); err != nil {
						return err
					}
				}

				if numVIResourcesAssigned != uint16(s.config.Resources) {
					if err := s.device.AssignControllerResources(controllerId, VIResourceType, uint32(int(numVIResourcesAssigned)-s.config.Resources)); err != nil {
						return err
					}
				}

				if !controllerOnline {
					if err := s.device.OnlineController(controllerId); err != nil {
						return err
					}
				}

				return nil
			}
		}

		if err := device.EnumerateSecondaryControllers(initFunc(s), handlerFunc(s)); err != nil {
			log.WithError(err).Errorf("Failed to enumerate %s storage controllers ", s.id)
		}

		// Port is ready to make connections
		event.EventType = PORT_EVENT_READY
		PortEventManager.Publish(event)

	case PORT_EVENT_DOWN:
		// TODO: Set this and all controllers down
	}
}

// Get -
func Get(model *sf.StorageCollectionStorageCollection) error {
	model.MembersodataCount = int64(len(mgr.storage))
	model.Members = make([]sf.OdataV4IdRef, int(model.MembersodataCount))
	for idx, s := range mgr.storage {
		model.Members[idx].OdataId = fmt.Sprintf("/redfish/v1/Storage/%s", s.id)
	}
	return nil
}

// StorageIdGet -
func StorageIdGet(storageId string, model *sf.StorageV190Storage) error {
	s := findStorage(storageId)
	if s == nil {
		return ec.ErrNotFound
	}

	model.Status = s.getStatus()

	// TODO: The model is missing a bunch of stuff
	// Manufacturer, Model, PartNumber, SerialNumber, etc.

	model.Controllers.OdataId = fmt.Sprintf("/redfish/v1/Storage/%s/Controllers", storageId)
	model.StoragePools.OdataId = fmt.Sprintf("/redfish/v1/Storage/%s/StoragePools", storageId)
	model.Volumes.OdataId = fmt.Sprintf("/redfish/v1/Storage/%s/Volumes", storageId)

	return nil
}

// StorageIdStoragePoolsGet -
func StorageIdStoragePoolsGet(storageId string, model *sf.StoragePoolCollectionStoragePoolCollection) error {
	s := findStorage(storageId)
	if s == nil {
		return ec.ErrNotFound
	}

	model.MembersodataCount = 1
	model.Members = make([]sf.OdataV4IdRef, model.MembersodataCount)
	model.Members[0].OdataId = fmt.Sprintf("/redfish/v1/Storage/%s/StoragePool/%s", storageId, defaultStoragePoolId)

	return nil
}

// StorageIdStoragePoolIdGet -
func StorageIdStoragePoolIdGet(storageId, storagePoolId string, model *sf.StoragePoolV150StoragePool) error {
	_, sp := findStoragePool(storageId, storagePoolId)
	if sp == nil {
		return ec.ErrNotFound
	}

	// TODO: This should reflect the total namespaces allocated over the drive
	model.Capacity = sf.CapacityV100Capacity{
		Data: sf.CapacityV100CapacityInfo{
			AllocatedBytes:   0,
			ConsumedBytes:    0,
			GuaranteedBytes:  0,
			ProvisionedBytes: 0,
		},
	}

	// TODO
	model.RemainingCapacityPercent = 0

	return nil
}

// StorageIdControllersGet -
func StorageIdControllersGet(storageId string, model *sf.StorageControllerCollectionStorageControllerCollection) error {
	s := findStorage(storageId)
	if s == nil {
		return ec.ErrNotFound
	}

	model.MembersodataCount = int64(len(s.controllers))
	model.Members = make([]sf.OdataV4IdRef, model.MembersodataCount)
	for idx, c := range s.controllers {
		model.Members[idx].OdataId = fmt.Sprintf("/redfish/v1/Storage/%s/Controllers/%s", storageId, c.id)
	}

	return nil
}

// StorageIdControllerIdGet -
func StorageIdControllerIdGet(storageId, controllerId string, model *sf.StorageControllerV100StorageController) error {
	_, c := findStorageController(storageId, controllerId)
	if c == nil {
		return ec.ErrNotFound
	}

	// Fill in the relative endpoint for this storage controller
	endpointId, err := FabricController.FindDownstreamEndpoint(storageId, controllerId)
	if err != nil {
		return err
	}

	model.Links.EndpointsodataCount = 1
	model.Links.Endpoints = make([]sf.OdataV4IdRef, model.Links.EndpointsodataCount)
	model.Links.Endpoints[0].OdataId = endpointId

	// model.Links.PCIeFunctions

	/*
		f := sf.PcIeFunctionV123PcIeFunction{
			ClassCode: "",
			DeviceClass: "",
			DeviceId: "",
			VendorId: "",
			SubsystemId: "",
			SubsystemVendorId: "",
			FunctionId: 0,
			FunctionType: sf.PHYSICAL_PCIFV123FT, // or sf.VIRTUAL_PCIFV123FT
			Links: sf.PcIeFunctionV123Links {
				StorageControllersodataCount: 1,
				StorageControllers: make([]sf.StorageStorageController, 1),
			},
		}
	*/

	model.NVMeControllerProperties = sf.StorageControllerV100NvMeControllerProperties{
		ControllerType: sf.IO_SCV100NVMCT, // OR ADMIN IF PF
	}

	return nil
}

// StorageIdVolumesGet -
func StorageIdVolumesGet(storageId string, model *sf.VolumeCollectionVolumeCollection) error {
	s := findStorage(storageId)
	if s == nil {
		return ec.ErrNotFound
	}

	// TODO: If s.ctrl is down - fail

	model.MembersodataCount = int64(len(s.volumes))
	model.Members = make([]sf.OdataV4IdRef, model.MembersodataCount)
	for idx, volume := range s.volumes {
		model.Members[idx].OdataId = s.fmt("/Volumes/%s", volume.id)
	}

	return nil
}

// StorageIdVolumeIdGet -
func StorageIdVolumeIdGet(storageId, volumeId string, model *sf.VolumeV161Volume) error {
	s, v := findStorageVolume(storageId, volumeId)
	if v == nil {
		return ec.ErrNotFound
	}

	// TODO: If s.ctrl is down - fail

	ns, err := s.device.GetNamespace(nvme.NamespaceIdentifier(v.namespaceId))
	if err != nil {
		return ec.ErrNotFound
	}

	formatGUID := func(guid []byte) string {
		var b strings.Builder
		for _, byt := range guid {
			b.WriteString(fmt.Sprintf("%02x", byt))
		}
		return b.String()
	}

	lbaFormat := ns.LBAFormats[ns.FormattedLBASize.Format]
	blockSizeInBytes := int64(math.Pow(2, float64(lbaFormat.LBADataSize)))

	model.BlockSizeBytes = blockSizeInBytes
	model.CapacityBytes = int64(ns.Capacity) * blockSizeInBytes
	model.Id = v.id
	model.Identifiers = make([]sf.ResourceIdentifier, 2)
	model.Identifiers = []sf.ResourceIdentifier{
		{
			DurableNameFormat: sf.NSID_RV1100DNF,
			DurableName:       fmt.Sprintf("%d", v.namespaceId),
		},
		{
			DurableNameFormat: sf.NGUID_RV1100DNF,
			DurableName:       formatGUID(ns.GloballyUniqueIdentifier[:]),
		},
	}

	model.Capacity = sf.CapacityV100Capacity{
		IsThinProvisioned: ns.Features.Thinp == 1,
		Data: sf.CapacityV100CapacityInfo{
			AllocatedBytes: int64(ns.Capacity) * blockSizeInBytes,
			ConsumedBytes:  int64(ns.Utilization) * blockSizeInBytes,
		},
	}

	model.NVMeNamespaceProperties = sf.VolumeV161NvMeNamespaceProperties{
		FormattedLBASize:                  fmt.Sprintf("%d", model.BlockSizeBytes),
		IsShareable:                       ns.MultiPathIOSharingCapabilities.Sharing == 1,
		MetadataTransferredAtEndOfDataLBA: lbaFormat.MetadataSize != 0,
		NamespaceId:                       fmt.Sprintf("%d", v.namespaceId),
		NumberLBAFormats:                  int64(ns.NumberOfLBAFormats),
	}

	model.VolumeType = sf.RAW_DEVICE_VVT

	// TODO: Find the attached status of the volume - if it is attached via a connection
	// to an endpoint that should go in model.Links.ClientEndpoints or model.Links.ServerEndpoints

	// TODO: Maybe StorageGroups??? An array of references to Storage Groups that includes this volume.
	// Storage Groups could be the Rabbit Slice

	// TODO: Should reference the Storage Pool

	return nil
}

// StorageIdVolumePost -
func StorageIdVolumePost(storageId string, model *sf.VolumeV161Volume) error {
	s := findStorage(storageId)
	if s == nil {
		return ec.ErrNotFound
	}

	volume, err := s.createVolume(uint64(model.CapacityBytes))

	// TODO: We should parse the error and make it more obvious (404, 405, etc)
	if err != nil {
		return err
	}

	return StorageIdVolumeIdGet(storageId, volume.id, model)
}

// StorageIdVolumeIdDelete -
func StorageIdVolumeIdDelete(storageId, volumeId string) error {
	s, v := findStorageVolume(storageId, volumeId)
	if v == nil {
		return ec.ErrBadRequest
	}

	return s.deleteVolume(volumeId)
}
