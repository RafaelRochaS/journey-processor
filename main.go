package main

import (
	"log"

	"github.com/RafaelRochaS/journey-processor/consumers"
	"github.com/RafaelRochaS/journey-processor/utils"
)

func main() {
	utils.LoadEnvs()
	log.Println("Starting event consumption...")

	consumers.HandleKafkaEvents()
}
