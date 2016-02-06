// Package calmly is example of golang package that should not exist.
// See calmly_test.go for usage examples
package calmly

import (
	"reflect"
)

type E interface{}

type exceptionHandler struct {
	e E
	f func(E)
}

type calm struct {
	try      func()
	catch    []exceptionHandler
	catchAny func(E)
	finally  func()
}

func Try(f func()) *calm {
	return &calm{f, make([]exceptionHandler, 0, 4), nil, nil}
}

func (c *calm) Catch(except E, f func(e E)) *calm {
	c.catch = append(c.catch, exceptionHandler{except, f})
	return c
}

func (c *calm) CatchAny(f func(e E)) *calm {
	c.catchAny = f
	return c
}

func (c *calm) Finally(f func()) {
	c.finally = f
	c.run()
}

func (c *calm) run() {
	if c.finally != nil {
		defer c.finally()
	}
	defer func() {
		if r := recover(); r != nil {
			for _, catch := range c.catch {
				if reflect.TypeOf(catch.e) == reflect.TypeOf(r) {
					catch.f(r)
					return
				}
			}
			if c.catchAny != nil {
				c.catchAny(r)
				return
			}
			panic(r)
		}
	}()
	c.try()
}
