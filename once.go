package once

import (
	"sync"

	"github.com/go-zoox/core-utils/safe"
)

// Once ...
type Once interface {
	Do(name string, fn func())
	Func(name string, fn func()) func()
	Get(name string, fn func() any) any
}

type once struct {
	onces *safe.Map[string, any]
	datas *safe.Map[string, any]
	sync.RWMutex
}

// New creates once.
func New() Once {
	return &once{
		onces: safe.NewMap[string, any](),
		datas: safe.NewMap[string, any](),
	}
}

// Do ...
func (o *once) Do(name string, fn func()) {
	if !o.onces.Has(name) {
		o.onces.Set(name, &sync.Once{})
	}

	o.onces.Get(name).(*sync.Once).Do(fn)
}

// Func ...
func (o *once) Func(name string, fn func()) func() {
	return func() {
		o.Do(name, fn)
	}
}

// Get ...
func (o *once) Get(name string, fn func() any) any {
	o.Do(name, func() {
		o.datas.Set(name, fn())
	})

	return o.datas.Get(name)
}
