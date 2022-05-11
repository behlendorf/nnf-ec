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

// Code generated by msgenerator ./registries/ResourceEvent.1.0.3.json; DO NOT EDIT

package messageregistry

import events "github.com/nearnodeflash/nnf-ec/pkg/manager-event"

// arg0: The relevant resource. This argument shall contain the name of the relevant Redfish resource.
// arg1: The state of the resource. This argument shall contain the value of the `Health` property for the relevant Redfish resource.  The values shall be used from the definition of the `Health` enumeration in the `Resource` schema.
func ResourceStatusChangedWarningResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "The health of resource `%1` has changed to %2.",
		MessageSeverity: "Warning",
		MessageId:       "ResourceEvent.1.0.3.ResourceStatusChangedWarning",
		MessageArgs:     []string{arg0, arg1},
	}
}

func ResourceChangedResourceEvent() events.Event {
	return events.Event{
		Message:         "One or more resource properties have changed.",
		MessageSeverity: "OK",
		MessageId:       "ResourceEvent.1.0.3.ResourceChanged",
	}
}

// arg0: The name of the property. This argument shall contain the name of the relevant property from a Redfish resource.
// arg1: The value of the threshold. This argument shall contain the value of the relevant error threshold.
func ResourceErrorThresholdClearedResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "The resource property %1 has cleared the error threshold of value %2.",
		MessageSeverity: "OK",
		MessageId:       "ResourceEvent.1.0.3.ResourceErrorThresholdCleared",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The name of the property. This argument shall contain the name of the relevant property from a Redfish resource.
// arg1: The value of the threshold. This argument shall contain the value of the relevant error threshold.
func ResourceWarningThresholdClearedResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "The resource property %1 has cleared the warning threshold of value %2.",
		MessageSeverity: "OK",
		MessageId:       "ResourceEvent.1.0.3.ResourceWarningThresholdCleared",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The relevant resource. This argument shall contain the name of the relevant resource or service affected by the software license.
// arg1: The message returned from the license validation process. This argument shall contain the message returned from the license validation process or software.
func LicenseExpiredResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "A license for '%1' has expired.  The following message was returned: '%2'.",
		MessageSeverity: "Warning",
		MessageId:       "ResourceEvent.1.0.3.LicenseExpired",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The name of the property. This argument shall contain the name of the relevant property from a Redfish resource.
// arg1: The value of the threshold. This argument shall contain the value of the relevant error threshold.
func ResourceErrorThresholdExceededResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "The resource property %1 has exceeded error threshold of value %2.",
		MessageSeverity: "Critical",
		MessageId:       "ResourceEvent.1.0.3.ResourceErrorThresholdExceeded",
		MessageArgs:     []string{arg0, arg1},
	}
}

func URIForResourceChangedResourceEvent() events.Event {
	return events.Event{
		Message:         "The URI for the resource has changed.",
		MessageSeverity: "OK",
		MessageId:       "ResourceEvent.1.0.3.URIForResourceChanged",
	}
}

// arg0: The name of the property. This argument shall contain the name of the relevant property from a Redfish resource.
// arg1: The type of error encountered. This argument shall contain a description of the type of error encountered.
func ResourceErrorsCorrectedResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "The resource property %1 has corrected errors of type '%2'.",
		MessageSeverity: "OK",
		MessageId:       "ResourceEvent.1.0.3.ResourceErrorsCorrected",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The relevant resource. This argument shall contain the name of the relevant Redfish resource.
// arg1: The state of the resource. This argument shall contain the value of the `Health` property for the relevant Redfish resource.  The values shall be used from the definition of the `Health` enumeration in the `Resource` schema.
func ResourceStatusChangedCriticalResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "The health of resource `%1` has changed to %2.",
		MessageSeverity: "Critical",
		MessageId:       "ResourceEvent.1.0.3.ResourceStatusChangedCritical",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The name of the software component. This argument shall contain the name of the relevant software component or package.  This might include both name and version information.
func ResourceVersionIncompatibleResourceEvent(arg0 string) events.Event {
	return events.Event{
		Message:         "An incompatible version of software '%1' has been detected.",
		MessageSeverity: "Warning",
		MessageId:       "ResourceEvent.1.0.3.ResourceVersionIncompatible",
		MessageArgs:     []string{arg0},
	}
}

// arg0: The self-test error message. This argument shall contain the error message received as a result from the self-test.
func ResourceSelfTestFailedResourceEvent(arg0 string) events.Event {
	return events.Event{
		Message:         "A self-test has failed.  The following message was returned: '%1'.",
		MessageSeverity: "Critical",
		MessageId:       "ResourceEvent.1.0.3.ResourceSelfTestFailed",
		MessageArgs:     []string{arg0},
	}
}

func ResourceSelfTestCompletedResourceEvent() events.Event {
	return events.Event{
		Message:         "A self-test has completed.",
		MessageSeverity: "OK",
		MessageId:       "ResourceEvent.1.0.3.ResourceSelfTestCompleted",
	}
}

func LicenseChangedResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "A license for '%1' has changed.  The following message was returned: '%2'.",
		MessageSeverity: "Warning",
		MessageId:       "ResourceEvent.1.0.3.LicenseChanged",
		MessageArgs:     []string{arg0, arg1},
	}
}

func ResourceRemovedResourceEvent() events.Event {
	return events.Event{
		Message:         "The resource has been removed successfully.",
		MessageSeverity: "OK",
		MessageId:       "ResourceEvent.1.0.3.ResourceRemoved",
	}
}

// arg0: The name of the property. This argument shall contain the name of the relevant property from a Redfish resource.
// arg1: The type of error encountered. This argument shall contain a description of the type of error encountered.
func ResourceErrorsDetectedResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "The resource property %1 has detected errors of type '%2'.",
		MessageSeverity: "Warning",
		MessageId:       "ResourceEvent.1.0.3.ResourceErrorsDetected",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The relevant resource. This argument shall contain the name of the relevant resource or service affected by the software license.
// arg1: The message returned from the license validation process. This argument shall contain the message returned from the license validation process or software.
func LicenseAddedResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "A license for '%1' has been added.  The following message was returned: '%2'.",
		MessageSeverity: "OK",
		MessageId:       "ResourceEvent.1.0.3.LicenseAdded",
		MessageArgs:     []string{arg0, arg1},
	}
}

// arg0: The relevant resource. This argument shall contain the name of the relevant Redfish resource.
// arg1: The state of the resource. This argument shall contain the value of the `Health` property for the relevant Redfish resource.  The values shall be used from the definition of the `Health` enumeration in the `Resource` schema.
func ResourceStatusChangedOKResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "The health of resource '%1' has changed to %2.",
		MessageSeverity: "OK",
		MessageId:       "ResourceEvent.1.0.3.ResourceStatusChangedOK",
		MessageArgs:     []string{arg0, arg1},
	}
}

func ResourceCreatedResourceEvent() events.Event {
	return events.Event{
		Message:         "The resource has been created successfully.",
		MessageSeverity: "OK",
		MessageId:       "ResourceEvent.1.0.3.ResourceCreated",
	}
}

// arg0: The name of the property. This argument shall contain the name of the relevant property from a Redfish resource.
// arg1: The value of the threshold. This argument shall contain the value of the relevant error threshold.
func ResourceWarningThresholdExceededResourceEvent(arg0, arg1 string) events.Event {
	return events.Event{
		Message:         "The resource property %1 has exceeded its warning threshold of value %2.",
		MessageSeverity: "Warning",
		MessageId:       "ResourceEvent.1.0.3.ResourceWarningThresholdExceeded",
		MessageArgs:     []string{arg0, arg1},
	}
}
