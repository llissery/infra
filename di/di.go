package di

import (
	"errors"
)

func Register(constructor any, opts ...ModuleOption) error {
	if len(opts) <= 0 {
		return Global.Register(constructor)
	}

	c := LoadOrStoreContainer(opts[0])

	err := c.Register(constructor)

	if err != nil {
		return err
	}

	return nil
}

func MustRegister(constructor any, opts ...ModuleOption) {
	if len(opts) <= 0 {
		err := Global.Register(constructor)
		if err != nil {
			panic(err)
		}
		return
	}

	c := LoadOrStoreContainer(opts[0])

	err := c.Register(constructor)

	if err != nil {
		panic(err) // nolint: byted_s_panic_detect
	}
}

func Call(function any, opts ...ModuleOption) error {
	if len(opts) <= 0 {
		return Global.Call(function)
	}

	c, ok := LoadContainer(opts[0])

	if !ok {
		return errors.New("module container not exist")
	}

	err := c.Call(function)

	if err != nil {
		return err
	}

	return nil
}
