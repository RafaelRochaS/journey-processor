package processor

type Processor interface {
	ProcessEvent([]byte) error
}
