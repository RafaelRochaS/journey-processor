package main

import (
	"log"

	"github.com/RafaelRochaS/journey-processor/consumers"
	"github.com/RafaelRochaS/journey-processor/processor"
	"github.com/RafaelRochaS/journey-processor/utils"
)

func main() {
	log.Println("Starting event consumption...")
	utils.LoadEnvs()
	processor := processor.NewEventHandlerProcessor()

	consumers.HandleKafkaEvents(processor)
}
