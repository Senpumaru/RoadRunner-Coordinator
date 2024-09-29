// cmd/client/main.go

package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"github.com/Senpumaru/RoadRunner-Coordinator/internal/workflows"
)

func main() {
	// Create the client object
	c, err := client.NewClient(client.Options{
		HostPort: "temporal:7233",
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "kafka-to-iceberg-workflow",
		TaskQueue: "kafka-to-iceberg-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.KafkaToIcebergWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
