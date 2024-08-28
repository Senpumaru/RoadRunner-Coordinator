package activity

import (
	"context"
	"fmt"
)

func GreetingActivity(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
