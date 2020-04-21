package client

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
)

// Client represents the client interface
type Client interface {
	Options() *Options

	Listable
	Watchable
}

// Listable represents a listable client interface
type Listable interface {
	Nodes() (*corev1.NodeList, error)
	Namespaces() (*corev1.NamespaceList, error)
	Deployments(ns string) (*appsv1.DeploymentList, error)
	Jobs(ns string) (*batchv1.JobList, error)
	Pods(ns string) (*corev1.PodList, error)
}

// Watchable represents a watchable client interface
type Watchable interface {
	NodesWatch() (watch.Interface, error)
	NamespacesWatch() (watch.Interface, error)
	DeploymentsWatch(ns string) (watch.Interface, error)
	JobsWatch(ns string) (watch.Interface, error)
	PodsWatch(ns string) (watch.Interface, error)
}
