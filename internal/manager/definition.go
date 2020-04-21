package manager

import "context"

// Manager represents the manager interface
type Manager interface {
	Options() *Options

	Run() error
	Stop() error
}

// Component represents a component
type Component interface {
	Start(ctx context.Context) error
	Stop() error
}
