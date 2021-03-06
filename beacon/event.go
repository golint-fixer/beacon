package beacon

// Action is the thing that's happening to the container.
type Action string

// Available event action values.
const (
	Start  Action = "start"  // Container started.
	Stop          = "stop"   // Container stopped.
	Update        = "update" // Container updated.
)

// Event indicates when the status of a container changes.
type Event struct {
	// The action that triggered this event.
	Action Action

	// The container affected by this event.
	Container *Container
}

// Copy allocates a copy of the Event.
func (e *Event) Copy() *Event {
	if e == nil {
		return nil
	}
	return &Event{
		Action:    e.Action,
		Container: e.Container.Copy(),
	}
}
