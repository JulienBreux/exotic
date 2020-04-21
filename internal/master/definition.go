package master

import "github.com/julienbreux/exotic/internal/manager"

// Client represents the master interface
type Master interface {
	manager.Component
}
