package client

import (
	"github.com/pkg/errors"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type client struct {
	opt *Options

	// lgr zerolog.Logger
	cli *kubernetes.Clientset
}

// New creates a new client instance
func New(opts ...Option) (Client, error) {
	opt, err := newOptions(opts...)
	if err != nil {
		return nil, err
	}

	cli, err := clientset(opt.InCluster, opt.ConfigPath.String())
	if err != nil {
		return nil, err
	}

	return &client{
		opt: opt,
		// lgr: lgr,
		cli: cli,
	}, nil
}

// Options returns the list of options
func (c *client) Options() *Options {
	return c.opt
}

// Nodes returns a list of nodes
func (c *client) Nodes() (*corev1.NodeList, error) {
	return c.cli.CoreV1().Nodes().List(metav1.ListOptions{
		TimeoutSeconds: &c.opt.TimeoutSeconds,
	})
}

// Namespaces returns a list of namespaces
func (c *client) Namespaces() (*corev1.NamespaceList, error) {
	return c.cli.CoreV1().Namespaces().List(metav1.ListOptions{
		TimeoutSeconds: &c.opt.TimeoutSeconds,
	})
}

// Deployments returns a list of deployments
func (c *client) Deployments(ns string) (*appsv1.DeploymentList, error) {
	return c.cli.AppsV1().Deployments(ns).List(metav1.ListOptions{
		TimeoutSeconds: &c.opt.TimeoutSeconds,
	})
}

// Jobs returns a list of jobs
func (c *client) Jobs(ns string) (*batchv1.JobList, error) {
	return c.cli.BatchV1().Jobs(ns).List(metav1.ListOptions{
		TimeoutSeconds: &c.opt.TimeoutSeconds,
	})
}

// Pods returns a list of pods
func (c *client) Pods(ns string) (*corev1.PodList, error) {
	return c.cli.CoreV1().Pods(ns).List(metav1.ListOptions{
		TimeoutSeconds: &c.opt.TimeoutSeconds,
	})
}

// NodesWatch return a nodes watcher
func (c *client) NodesWatch() (watch.Interface, error) {
	return c.cli.CoreV1().Nodes().Watch(metav1.ListOptions{
		TimeoutSeconds: &c.opt.TimeoutSeconds,
	})
}

// NamespacesWatch return a namespaces watcher
func (c *client) NamespacesWatch() (watch.Interface, error) {
	return c.cli.CoreV1().Namespaces().Watch(metav1.ListOptions{
		TimeoutSeconds: &c.opt.TimeoutSeconds,
	})
}

// DeploymentsWatch return a deployments watcher
func (c *client) DeploymentsWatch(ns string) (watch.Interface, error) {
	return c.cli.AppsV1().Deployments(ns).Watch(metav1.ListOptions{
		TimeoutSeconds: &c.opt.TimeoutSeconds,
	})
}

// JobsWatch return a jobs watcher
func (c *client) JobsWatch(ns string) (watch.Interface, error) {
	return c.cli.BatchV1().Jobs(ns).Watch(metav1.ListOptions{
		TimeoutSeconds: &c.opt.TimeoutSeconds,
	})
}

// PodsWatch return a pods watcher
func (c *client) PodsWatch(ns string) (watch.Interface, error) {
	return c.cli.CoreV1().Pods(ns).Watch(metav1.ListOptions{
		TimeoutSeconds: &c.opt.TimeoutSeconds,
	})
}

func clientset(inCluster bool, configPath string) (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error

	if inCluster {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", configPath)
	}
	if err != nil {
		// TODO: Change error management and cause
		return nil, errors.Wrap(err, "unable to configure client")
	}

	return kubernetes.NewForConfig(config)
}
