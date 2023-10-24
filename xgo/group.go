package xgo

import (
	"context"
	"sync"
)

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and does not cancel on error.
type Group struct {
	cancel func()

	wg sync.WaitGroup

	errOnce sync.Once
	err     error

	initOnce sync.Once

	restrict chan struct{}

	fns []func() error
}

// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first timeUtil a function passed to Go
// returns a non-nil error or the first timeUtil Wait returns, whichever occurs
// first.
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel}, ctx
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
// @param restrict: concurrent restrict number
func (g *Group) Wait(restrict int) error {
	if restrict > 0 {
		g.restrict = make(chan struct{}, restrict)
	}
	for _, fn := range g.fns {
		g.do(fn)
	}
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the Group; its error will be
// returned by Wait.
func (g *Group) Go(fn func() error) {
	g.initOnce.Do(func() {
		g.fns = make([]func() error, 0)
	})

	g.fns = append(g.fns, fn)
}

func (g *Group) do(f func() error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		if g.restrict != nil {
			g.restrict <- struct{}{}
			defer func() {
				<-g.restrict
			}()
		}

		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}
