package component

import (
	"fmt"

	"github.com/base-org/pessimism/internal/core"
)

// egressHandler ... Used to route transit data from a component to it's respective edge components.
// Also used to manage egresses or "edge routes" for some component.
type egressHandler struct {
	egresses map[core.ComponentPID]chan core.TransitData
}

// newEgress ... Initializer
func newEgressHandler() *egressHandler {
	return &egressHandler{
		egresses: make(map[core.ComponentPID]chan core.TransitData),
	}
}

// Send ... Sends single piece of transitData to all innner mapping value channels
func (eh *egressHandler) Send(td core.TransitData) error {
	if len(eh.egresses) == 0 {
		return fmt.Errorf(egressNotExistErr)
	}

	// NOTE - Consider introducing a fail safe timeout to ensure that freezing on clogged chanel buffers is recognized
	for _, channel := range eh.egresses {
		channel <- td
	}

	return nil
}

// SendBatch ... Sends slice of transitData to all innner mapping value channels
func (eh *egressHandler) SendBatch(dataSlice []core.TransitData) error {
	// NOTE - Consider introducing a fail safe timeout to ensure that freezing on clogged chanel buffers is recognized
	for _, data := range dataSlice {
		if err := eh.Send(data); err != nil {
			return err
		}
	}

	return nil
}

// AddEgress ... Inserts a new egress given an ID and channel; fail on key collision
func (eh *egressHandler) AddEgress(componentID core.ComponentUUID, outChan chan core.TransitData) error {
	if _, found := eh.egresses[componentID.PID]; found {
		return fmt.Errorf(egressAlreadyExistsErr, componentID.String())
	}

	eh.egresses[componentID.PID] = outChan
	return nil
}

// RemoveEgress ... Removes an egress given an ID; fail if no key found
func (eh *egressHandler) RemoveEgress(componentID core.ComponentUUID) error {
	if _, found := eh.egresses[componentID.PID]; !found {
		return fmt.Errorf(egressNotFoundErr, componentID.PID.String())
	}

	delete(eh.egresses, componentID.PID)

	return nil
}
