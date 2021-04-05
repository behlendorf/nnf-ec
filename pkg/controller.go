/*
 * Near Node Flash
 *
 * This file contains the declaration of the Near-Node Flash
 * Element Controller.
 *
 * Author: Nate Roiger
 *
 * Copyright 2020 Hewlett Packard Enterprise Development LP
 */

package nnf

import (
	ec "stash.us.cray.com/rabsw/ec"

	fabric "stash.us.cray.com/rabsw/nnf-ec/internal/manager-fabric"
	nnf "stash.us.cray.com/rabsw/nnf-ec/internal/manager-nnf"
	nvme "stash.us.cray.com/rabsw/nnf-ec/internal/manager-nvme"
)

// NewController - Create a new NNF Element Controller with the desired mocking behavior
func NewController(cli, mock bool) *ec.Controller {
	switchCtrl := fabric.NewSwitchtecController()
	nvmeCtrl := nvme.NewSwitchtecNvmeController()
	nnfCtrl := nnf.NewNnfController()

	if cli {
		switchCtrl = fabric.NewSwitchtecCliController()
		nvmeCtrl = nvme.NewCliNvmeController()
	}

	if mock {
		switchCtrl = fabric.NewMockSwitchtecController()
		nvmeCtrl = nvme.NewMockNvmeController()
		nnfCtrl = nnf.NewMockNnfController()
	}

	return &ec.Controller{
		Name:    "Near Node Flash",
		Port:    50057,
		Version: "v2",
		Routers: NewDefaultApiRouters(switchCtrl, nvmeCtrl, nnfCtrl),
	}
}

// NewDefaultApiRouters -
func NewDefaultApiRouters(switchCtrl fabric.SwitchtecControllerInterface, nvmeCtrl nvme.NvmeController, nnfCtrl nnf.NnfControllerInterface) ec.Routers {

	routers := []ec.Router{
		fabric.NewDefaultApiRouter(fabric.NewDefaultApiService(), switchCtrl),
		nvme.NewDefaultApiRouter(nvme.NewDefaultApiService(), nvmeCtrl),
		nnf.NewDefaultApiRouter(nnf.NewDefaultApiService(), nnfCtrl),
	}

	return routers
}
