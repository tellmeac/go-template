package commands

import (
	"context"
	"errors"
	"time"
)

var (
	ErrCommandTimeout = errors.New("command timeout")
)

type Command interface {
	Init(ctx context.Context) error
	Check() error
	Prepare() error
	Run() error
	Cleanup() error
	Debug() bool
	ShutdownTimeout() time.Duration
	Context() context.Context
	Cancel(ctx context.Context)
}

var _ Command = (*BaseCommand)(nil)

type BaseCommand struct {
	Timeout time.Duration

	debug  bool
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *BaseCommand) Init(ctx context.Context) error {
	if c.Timeout == 0 {
		c.Timeout = 3 * time.Second
	}

	c.ctx, c.cancel = context.WithCancel(ctx)
	return nil
}

func (c *BaseCommand) Check() error {
	return nil
}

func (c *BaseCommand) Prepare() error {
	return nil
}

func (c *BaseCommand) Run() error {
	return nil
}

func (c *BaseCommand) Cleanup() error {
	return nil
}

func (c *BaseCommand) Debug() bool {
	return c.debug
}

func (c *BaseCommand) ShutdownTimeout() time.Duration {
	return c.Timeout
}

func (c *BaseCommand) Context() context.Context {
	return c.ctx
}

func (c *BaseCommand) Cancel(_ context.Context) {
	c.cancel()
}
