package agent

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/julienbreux/exotic/internal/agent/client"
)

const defaultProcess = "agent"

type agent struct {
	opt *Options

	lgr zerolog.Logger
	cli client.Client
}

// New creates a new agent instance
func New(opts ...Option) (Agent, error) {
	opt, err := newOptions(opts...)
	if err != nil {
		return nil, err
	}

	cli, err := client.New(
		client.Logger(opt.Logger),
	)
	if err != nil {
		return nil, errors.Wrap(err, "unable to configure agent")
	}

	return &agent{
		opt: opt,
		lgr: opt.Logger.Log().With().Str("process", defaultProcess).Logger(),
		cli: cli,
	}, nil
}

// Options returns the list of options
func (a *agent) Options() *Options {
	return a.opt
}

// Start starts the agent
func (a *agent) Start(ctx context.Context) (err error) {
	a.lgr.Info().Msg("Started")

	nodeTicker := time.NewTicker(a.opt.Frequency)
	podTicker := time.NewTicker(a.opt.Frequency)

	for {
		select {
		case <-ctx.Done():
			nodeTicker.Stop()
			podTicker.Stop()
			return
		case <-nodeTicker.C:
			a.sendNodesChangesToMaster()
		case <-podTicker.C:
			a.sendPodsChangesToMaster()
		}
	}
}

// Stop stops the agent
func (a *agent) Stop() (err error) {
	a.lgr.Info().Msg("Stopped")
	return
}

func (a *agent) sendNodesChangesToMaster() {
	nl, _ := a.cli.Nodes()
	for _, node := range nl.Items {
		a.lgr.Info().
			Str("node", node.Name).
			Msgf("Send node %s changes to master", node.Name)
		// FIXME: Write server part code
	}
}

func (a *agent) sendPodsChangesToMaster() {
	for _, ns := range a.opt.Namespaces {
		pl, _ := a.cli.Pods(ns)
		for _, pod := range pl.Items {
			a.lgr.Info().
				Str("node", pod.Spec.NodeName).
				Str("namespace", ns).
				Str("pod", pod.Name).
				Str("phase", string(pod.Status.Phase)).
				Msgf("Send pod %s changes to master", pod.Name)
			// FIXME: Write server part code
		}
	}
}
