package mongo

import (
	"context"
	"time"
)

const DefaultTimeout = 5 * time.Second

func WithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), DefaultTimeout)
}
