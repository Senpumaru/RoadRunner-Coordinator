package activity

import (
	"context"
	"fmt"
)

func GreetingActivity(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hola, %s!", name), nil
}
