// internal/workflows/kafka_to_iceberg.go

package workflows

import (
	"time"

	"github.com/Senpumaru/RoadRunner-Coordinator/internal/activities"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func KafkaToIcebergWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting Kafka to Iceberg workflow")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 2 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Minute,
			BackoffCoefficient: 2.0,
			MaximumInterval:    30 * time.Minute,
			MaximumAttempts:    5,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	for {
		var result string
		err := workflow.ExecuteActivity(ctx, activities.TriggerSparkJob).Get(ctx, &result)
		if err != nil {
			logger.Error("Failed to execute Spark job", "error", err)
			// Depending on the error, you might want to break the loop or continue
			continue
		}

		logger.Info("Spark job completed", "result", result)

		// Wait for some time before triggering the next job
		if err := workflow.Sleep(ctx, 5*time.Minute); err != nil {
			return err
		}
	}
}
