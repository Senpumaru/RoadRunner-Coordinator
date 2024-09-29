// cmd/worker/main.go

package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/Senpumaru/RoadRunner-Coordinator/internal/activities"
	"github.com/Senpumaru/RoadRunner-Coordinator/internal/workflows"
)

func main() {
	log.Println("Starting Temporal worker...")

	// Create the client object just once per process
	c, err := client.NewClient(client.Options{
		HostPort: "temporal:7233",
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both workflow and activity functions
	w := worker.New(c, "kafka-to-iceberg-task-queue", worker.Options{})

	w.RegisterWorkflow(workflows.KafkaToIcebergWorkflow)
	w.RegisterActivity(activities.TriggerSparkJob)

	log.Println("Worker registered. Starting to listen on task queue: kafka-to-iceberg-task-queue")

	// Start listening to the task queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
