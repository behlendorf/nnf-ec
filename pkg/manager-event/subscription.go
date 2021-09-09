package event

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	sf "stash.us.cray.com/rabsw/rfsf-openapi/pkg/models"
)

type Subscription interface {
	EventHandler(e Event) error
}

type RedfishSubscription struct {
	// A client-supplied string that is stored with the event destination subscription.
	Context string

	// The URI of the destination event receiver.
	Destination string

	// The subscription delivery retry policy for events, where the subscription type is RedfishEvent.
	DeliveryRetryPolicy sf.EventDestinationV190DeliveryRetryPolicy
}

func (s RedfishSubscription) EventHandler(e Event) error {

	event := RedfishEvent{
		Context: s.Context,
		Events: []Event{
			e,
		},
	}

	attempts := 0

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// TODO: This is a quick way to implement event notification, but is short of a full
	// fledged system. Full functionality would include the subscription re-written as
	// its own service with an pending event buffer. The service would post any
	// accumulated events, and would perform retries as a global policy (not per-event, as
	// is done here). We should also consider the context headers a subscription might
	// want to use, as well as any security context.
	go func() {

		for {
			rsp, err := http.Post(s.Destination, "application/json", bytes.NewBuffer(body))
			if err == nil {
				switch rsp.StatusCode {
				case http.StatusOK, http.StatusAccepted:
					return
				}
			}

			if s.DeliveryRetryPolicy == sf.RETRY_FOREVER_EDV190DRP {
				continue
			}

			attempts++
			if attempts <= EventManager.deliveryRetryAttempts {
				time.Sleep(time.Duration(EventManager.deliveryRetryInvervalSeconds) * time.Second)
				continue
			}

			// TODO: Implement subscription suspension algorithm if needed
			// if s.DeliveryRetryPolicy == sf.SUSPEND_RETRIES_EDV190DRP

			// Assumes s.DeliveryRetryPolicy == sf.TERMINATE_AFTER_RETRIES_EDV190DRP
			return
		}
	}()

	return nil
}

type RedfishEvent struct {
	Context string `json:"Context,omitempty"`

	Events []Event `json:"Events,omitempty"`
}
