package di

import "fmt"

type ModuleOption interface {
	applyModuleOption(*moduleOptions)
	name() string
}

type moduleOptions struct {
	Name string
}

func Module(module string) ModuleOption {
	return moduleOption(module)
}

type moduleOption string

func (o moduleOption) name() string {
	return string(o)
}

func (o moduleOption) String() string {
	return fmt.Sprintf("Module(%q)", string(o))
}

func (o moduleOption) applyModuleOption(opt *moduleOptions) {
	opt.Name = string(o)
}
