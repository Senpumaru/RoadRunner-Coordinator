package workflow

import (
	"time"

	"github.com/Senpumaru/RoadRunner-Coordinator/internal/activity"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, activity.GreetingActivity, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}
