package agent

import "github.com/julienbreux/exotic/internal/manager"

// Client represents the agent interface
type Agent interface {
	manager.Component
}
