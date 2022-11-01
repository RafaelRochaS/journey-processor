package main

import (
	"log"

	"github.com/RafaelRochaS/journey-processor/consumers"
)

func main() {
	log.Println("Starting event consumption...")

	consumers.HandleKafkaEvents()
}
