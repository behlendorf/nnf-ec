package benchmarks

import (
	"testing"

	ec "stash.us.cray.com/rabsw/nnf-ec/pkg"
	nnf "stash.us.cray.com/rabsw/nnf-ec/pkg/manager-nnf"
	server "stash.us.cray.com/rabsw/nnf-ec/pkg/manager-server"

	openapi "stash.us.cray.com/rabsw/nnf-ec/pkg/rfsf/pkg/common"
	sf "stash.us.cray.com/rabsw/nnf-ec/pkg/rfsf/pkg/models"
)

func TestStoragePools(t *testing.T) {
	c := ec.NewController(ec.NewMockOptions())
	defer c.Close()

	if err := c.Init(nil); err != nil {
		t.Fatalf("Failed to start nnf controller")
	}

	ss := nnf.NewDefaultStorageService()

	cs := &sf.CapacityCapacitySource{}
	if err := ss.StorageServiceIdCapacitySourceGet(ss.Id(), cs); err != nil {
		t.Errorf("Failed to retrieve capacity source: %v", err)
	}

	pools := make([]*sf.StoragePoolV150StoragePool, 0)

	for j := 0; j < 32; j++ {

		sp := &sf.StoragePoolV150StoragePool{
			CapacityBytes: 1024 * 1024,
			Oem: openapi.MarshalOem(nnf.AllocationPolicyOem{
				Policy:     nnf.SpareAllocationPolicyType,
				Compliance: nnf.RelaxedAllocationComplianceType,
			}),
		}

		if err := ss.StorageServiceIdStoragePoolsPost(ss.Id(), sp); err != nil {
			t.Fatalf("Failed to create storage pool %d Error: %+v", j, err)
		}

		pools = append(pools, sp)

		rabbitEndpointId := "0"
		ep := &sf.EndpointV150Endpoint{}
		if err := ss.StorageServiceIdEndpointIdGet(ss.Id(), rabbitEndpointId, ep); err != nil {
			t.Fatalf("Failed to get endpoint ID: %s Error: %+v", rabbitEndpointId, err)
		}

		sg := &sf.StorageGroupV150StorageGroup{
			Links: sf.StorageGroupV150Links{
				StoragePool:    sf.OdataV4IdRef{OdataId: sp.OdataId},
				ServerEndpoint: sf.OdataV4IdRef{OdataId: ep.OdataId},
			},
		}

		if err := ss.StorageServiceIdStorageGroupPost(ss.Id(), sg); err != nil {
			t.Fatalf("Failed to create storage group Pool ID: %s Error: %+v", sp.Id, err)
		}

		fs := &sf.FileSystemV122FileSystem{
			Links: sf.FileSystemV122Links{
				StoragePool: sf.OdataV4IdRef{OdataId: sp.OdataId},
			},
			Oem: openapi.MarshalOem(server.FileSystemOem{
				Type: "zfs",
				Name: "zfs",
			}),
		}

		if err := ss.StorageServiceIdFileSystemsPost(ss.Id(), fs); err != nil {
			t.Fatalf("Failed to create file system Pool ID: %s Error: %+v", sp.Id, err)
		}

		sh := &sf.FileShareV120FileShare{
			FileSharePath: "/mnt/test",
			Links: sf.FileShareV120Links{
				FileSystem: sf.OdataV4IdRef{OdataId: fs.OdataId},
				Endpoint:   sf.OdataV4IdRef{OdataId: ep.OdataId},
			},
		}

		if err := ss.StorageServiceIdFileSystemIdExportedSharesPost(ss.Id(), fs.Id, sh); err != nil {
			t.Fatalf("Failed to create file share Pool ID: %s Error: %+v", sp.Id, err)
		}
	}

	t.Logf("Created %d Storage Pools. Starting Delete...", len(pools))

	for _, pool := range pools {
		if err := ss.StorageServiceIdStoragePoolIdDelete(ss.Id(), pool.Id); err != nil {
			t.Fatalf("Failed to delete storage pool ID %s Error: %v", pool.Id, err)
		}
	}
}
