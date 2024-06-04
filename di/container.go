package di

import "go.uber.org/dig"

type Container struct {
	*dig.Container
}

func (c *Container) Register(constructor any, opts ...dig.ProvideOption) error {
	return c.Provide(constructor, opts...)
}

func (c *Container) Call(function any, opts ...dig.InvokeOption) error {
	return c.Invoke(function, opts...)
}

func NewContainer() *Container {
	return &Container{dig.New()}
}
