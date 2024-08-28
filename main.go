package main

import (
	"context"
	"log"
	"os"

	"go.temporal.io/sdk/client"

	"github.com/Senpumaru/RoadRunner-Coordinator/internal/workflow"
)

func main() {
	address := os.Getenv("TEMPORAL_HOST_ADDRESS")
	if address == "" {
		address = "temporal:7233"
	}
	c, err := client.Dial(client.Options{
		HostPort: address,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: "greeting-task-queue",
	}

	name := "Temporal"
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflow.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
