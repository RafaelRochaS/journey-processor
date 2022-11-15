package processor

type EventHandler struct {
}

func NewEventHandlerProcessor() Processor {
	return &EventHandler{}
}

func (eh *EventHandler) ProcessEvent(event []byte) error {
	// Custom processing would happen here  .
	// Since this is just an exercise, no processing is needed
	// (aka no creativity / lazy to create some random processing stuff)
	return nil
}
