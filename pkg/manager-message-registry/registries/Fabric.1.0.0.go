/*
 * Copyright 2020, 2021, 2022 Hewlett Packard Enterprise Development LP
 * Other additional copyright holders may be indicated within.
 *
 * The entirety of this work is licensed under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 *
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by msgenerator ./registries/Fabric.1.0.0.json; DO NOT EDIT

package messageregistry

import events "github.com/NearNodeFlash/nnf-ec/pkg/manager-event"

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the downstream port.
func DownstreamLinkEstablishedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' downstream link is established on port '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.DownstreamLinkEstablished",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the interswitch port.
func InterswitchLinkEstablishedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' interswitch link is established on port '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.InterswitchLinkEstablished",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch that has returned to a functional state.
func SwitchOKFabric(arg0 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' has returned to a functional state.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.SwitchOK",
		MessageArgs:     []string{arg0},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch with the failed cable.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port with the failed cable.
func CableFailedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "The cable in switch '%1' port '%2' has failed.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.CableFailed",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the endpoint. This argument shall contain the value of the `Id` property of the endpoint that was removed.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the endpoint was removed.
func EndpointRemovedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Endpoint '%1' has been removed from fabric '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.EndpointRemoved",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the address pool. This argument shall contain the value of the `Id` property of the address pool that was modified.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the address pool was modified.
func AddressPoolModifiedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Address pool '%1' in fabric '%2' has been modified.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.AddressPoolModified",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the media controller. This argument shall contain the value of the `Id` property of the media controller that was modified.
// arg1: The `Id` of the chassis. This argument shall contain the value of the `Id` property of the chassis in which the media controller was modified.
func MediaControllerModifiedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Media controller '%1' in chassis '%2' has been modified.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.MediaControllerModified",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the upstream port.
func UpstreamLinkEstablishedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' upstream link is established on port '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.UpstreamLinkEstablished",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the upstream port.
func DegradedUpstreamLinkEstablishedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' upstream link is established on port '%2', but is running in a degraded state.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.DegradedUpstreamLinkEstablished",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the connection. This argument shall contain the value of the `Id` property of the connection that was created.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the connection was created.
func ConnectionCreatedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Connection '%1' has been created in fabric '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.ConnectionCreated",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the interswitch port.
func InterswitchLinkDroppedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' interswitch link has gone down on port '%2'.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.InterswitchLinkDropped",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch with the disabled port.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port that has been disabled.
func PortManuallyDisabledFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' port '%2' has been manually disabled.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.PortManuallyDisabled",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch with the cable that returned to a functional state.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port with the cable that returned to a functional state.
func CableOKFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "The cable in switch '%1' port '%2' has returned to working condition.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.CableOK",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the zone. This argument shall contain the value of the `Id` property of the zone that was removed.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the zone was removed.
func ZoneRemovedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Zone '%1' has been removed from fabric '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.ZoneRemoved",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the downstream port that is flapping.
// arg2: The number of times the link has flapped. This argument shall contain the count of uplink establishment/disconnection cycles.
// arg3: The number of minutes over which the link has flapped. This argument shall contain the number of minutes over which link flapping activity has been detected.
func DownstreamLinkFlapDetectedFabric(arg0, arg1, arg2, arg3 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' downstream link on port '%2' has been established and dropped %3 times in the last %4 minutes.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.DownstreamLinkFlapDetected",
		MessageArgs:     []string{arg0, arg1, arg2, arg3},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch that has failed.
func SwitchFailedFabric(arg0 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' has failed and is inoperative.",
		MessageSeverity: "Critical",
		MessageId:       "Fabric.1.0.0.SwitchFailed",
		MessageArgs:     []string{arg0},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch that has entered a degraded state.
func SwitchDegradedFabric(arg0 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' is in a degraded state.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.SwitchDegraded",
		MessageArgs:     []string{arg0},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch with the port whose cable was inserted.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port whose cable was inserted.
func CableInsertedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "A cable has been inserted into switch '%1' port '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.CableInserted",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the upstream port.
func UpstreamLinkDroppedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' upstream link has gone down on port '%2'.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.UpstreamLinkDropped",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the downstream port.
func DownstreamLinkDroppedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' downstream link has gone down on port '%2'.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.DownstreamLinkDropped",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch with the enabled port.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port that has been enabled.
func PortManuallyEnabledFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' port '%2' has been manually enabled.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.PortManuallyEnabled",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch with the functional port.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port that has returned to a functional state.
func PortOKFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' port '%2' has returned to a functional state.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.PortOK",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the zone. This argument shall contain the value of the `Id` property of the zone that was modified.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the zone was modified.
func ZoneModifiedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Zone '%1' in fabric '%2' has been modified.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.ZoneModified",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the endpoint. This argument shall contain the value of the `Id` property of the endpoint that was created.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the endpoint was created.
func EndpointCreatedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Endpoint '%1' has been created in fabric '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.EndpointCreated",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the endpoint. This argument shall contain the value of the `Id` property of the endpoint that was modified.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the endpoint was modified.
func EndpointModifiedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Endpoint '%1' in fabric '%2' has been modified.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.EndpointModified",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the downstream port.
func DegradedDownstreamLinkEstablishedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' downstream link is established on port '%2', but is running in a degraded state.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.DegradedDownstreamLinkEstablished",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch with the enabled port.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port that has been enabled.
func PortAutomaticallyEnabledFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' port '%2' has been automatically enabled.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.PortAutomaticallyEnabled",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch with the failed port.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port that has failed.
func PortFailedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' port '%2' has failed and is inoperative.",
		MessageSeverity: "Critical",
		MessageId:       "Fabric.1.0.0.PortFailed",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch with the degraded port.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port that has entered a degraded state.
func PortDegradedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' port '%2' is in a degraded state.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.PortDegraded",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the connection. This argument shall contain the value of the `Id` property of the connection that was removed.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the connection was removed.
func ConnectionRemovedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Connection '%1' has been removed from fabric '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.ConnectionRemoved",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the connection. This argument shall contain the value of the `Id` property of the connection that was modified.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the connection was modified.
func ConnectionModifiedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Connection '%1' in fabric '%2' has been modified.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.ConnectionModified",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the zone. This argument shall contain the value of the `Id` property of the zone that was created.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the zone was created.
func ZoneCreatedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Zone '%1' has been created in fabric '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.ZoneCreated",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the media controller. This argument shall contain the value of the `Id` property of the media controller that was added.
// arg1: The `Id` of the chassis. This argument shall contain the value of the `Id` property of the chassis in which the media controller was added.
func MediaControllerAddedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Media controller '%1' has been added to chassis '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.MediaControllerAdded",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch in which one or more packets have been dropped.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port in which one or more oversize packets were received.
// arg2: The MTU size. This argument shall contain the MTU size.
func MaxFrameSizeExceededFabric(arg0, arg1, arg2 string) events.Event {
	return events.Event{
		Message:         "MTU size on switch '%1' port '%2' is set to %3.  One or more packets with a larger size have been dropped.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.MaxFrameSizeExceeded",
		MessageArgs:     []string{arg0, arg1, arg2},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the upstream port that is flapping.
// arg2: The number of times the link has flapped. This argument shall contain the count of uplink establishment/disconnection cycles.
// arg3: The number of minutes over which the link has flapped. This argument shall contain the number of minutes over which link flapping activity has been detected.
func UpstreamLinkFlapDetectedFabric(arg0, arg1, arg2, arg3 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' upstream link on port '%2' has been established and dropped %3 times in the last %4 minutes.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.UpstreamLinkFlapDetected",
		MessageArgs:     []string{arg0, arg1, arg2, arg3},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the interswitch port.
func DegradedInterswitchLinkEstablishedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' interswitch link is established on port '%2', but is running in a degraded state.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.DegradedInterswitchLinkEstablished",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch with the disabled port.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port that has been disabled.
func PortAutomaticallyDisabledFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' port '%2' has been automatically disabled.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.PortAutomaticallyDisabled",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch with the port whose cable was removed.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the port whose cable was removed.
func CableRemovedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "A cable has been removed from switch '%1' port '%2'.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.CableRemoved",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the address pool. This argument shall contain the value of the `Id` property of the address pool that was removed.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the address pool was removed.
func AddressPoolRemovedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Address pool '%1' has been removed from fabric '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.AddressPoolRemoved",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the media controller. This argument shall contain the value of the `Id` property of the media controller that was removed.
// arg1: The `Id` of the chassis. This argument shall contain the value of the `Id` property of the chassis in which the media controller was removed.
func MediaControllerRemovedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Media controller '%1' has been removed from chassis '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.MediaControllerRemoved",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The `Id` of the switch. This argument shall contain the value of the `Id` property of the switch.
// arg1: The `Id` of the port. This argument shall contain the value of the `Id` property of the interswitch port that is flapping.
// arg2: The number of times the link has flapped. This argument shall contain the count of uplink establishment/disconnection cycles.
// arg3: The number of minutes over which the link has flapped. This argument shall contain the number of minutes over which link flapping activity has been detected.
func InterswitchLinkFlapDetectedFabric(arg0, arg1, arg2, arg3 string) events.Event {
	return events.Event{
		Message:         "Switch '%1' interswitch link on port '%2' has been established and dropped %3 times in the last %4 minutes.",
		MessageSeverity: "Warning",
		MessageId:       "Fabric.1.0.0.InterswitchLinkFlapDetected",
		MessageArgs:     []string{arg0, arg1, arg2, arg3},
	}
}

// arg0: The `Id` of the address pool. This argument shall contain the value of the `Id` property of the address pool that was created.
// arg1: The `Id` of the fabric. This argument shall contain the value of the `Id` property of the fabric in which the address pool was created.
func AddressPoolCreatedFabric(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "Address pool '%1' has been created in fabric '%2'.",
		MessageSeverity: "OK",
		MessageId:       "Fabric.1.0.0.AddressPoolCreated",
		MessageArgs:     []string{arg0, arg1},
	}
}
