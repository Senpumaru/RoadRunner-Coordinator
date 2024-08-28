package main

import (
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/Senpumaru/RoadRunner-Coordinator/internal/activity"
	"github.com/Senpumaru/RoadRunner-Coordinator/internal/workflow"
)

func main() {
	address := os.Getenv("TEMPORAL_HOST_ADDRESS")
	if address == "" {
		address = "localhost:7233"
	}
	c, err := client.Dial(client.Options{
		HostPort: address,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "greeting-task-queue", worker.Options{})

	w.RegisterWorkflow(workflow.GreetingWorkflow)
	w.RegisterActivity(activity.GreetingActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
