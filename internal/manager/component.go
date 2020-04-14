package manager

import "context"

// Component represents a component
type Component interface {
	Start(ctx context.Context) error
	Stop() error
}
