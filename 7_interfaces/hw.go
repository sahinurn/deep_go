package main

import "errors"

type UserService struct {
	// not need to implement
	NotEmptyStruct bool
}
type MessageService struct {
	// not need to implement
	NotEmptyStruct bool
}

type Container struct {
	constructors map[string]func() any
}

func NewContainer() *Container {
	return &Container{
		constructors: make(map[string]func() any),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	callableConstructor, ok := constructor.(func() any)
	if !ok {
		return
	}

	c.constructors[name] = callableConstructor
}

func (c *Container) Resolve(name string) (interface{}, error) {
	constructor, ok := c.constructors[name]
	if !ok {
		return nil, errors.New("constructor not found")
	}

	return constructor(), nil
}
