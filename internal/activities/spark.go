// internal/activities/spark.go

package activities

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func TriggerSparkJob(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "spark-submit",
		"--packages",
		"org.apache.iceberg:iceberg-spark-runtime-3.5_2.12:1.4.2,org.apache.spark:spark-sql-kafka-0-10_2.12:3.5.2,org.apache.hadoop:hadoop-aws:3.3.4,com.amazonaws:aws-java-sdk-bundle:1.12.262",
		"./src/python/kafka/kafka_to_iceberg_iot.py")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute Spark job: %w\nOutput: %s", err, string(output))
	}

	return strings.TrimSpace(string(output)), nil
}
