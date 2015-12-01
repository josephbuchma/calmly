// Package calmly is example of golang package that should not exist.
// See calmly_test.go for usage examples
package calmly

import (
	"reflect"
)

type E interface{}

type exception struct {
	e interface{}
	f func(E)
}

type calm struct {
	try     func()
	catch   []exception
	finally func()
}

func Try(f func()) *calm {
	return &calm{f, make([]exception, 1), nil}
}

func (c *calm) Catch(except interface{}, f func(e E)) *calm {
	c.catch = append(c.catch, exception{except, f})
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
					break
				}
			}
		}
	}()
	c.try()
}
