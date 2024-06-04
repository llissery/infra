package di

import (
	"sync"
)

var (
	// Global 提供全局注册能力.
	Global = NewContainer()
	// containerMap module 容器
	containerMap = &sync.Map{}
	lockMap      = &sync.Map{}
)

// for test
func clearContainers() {
	Global = NewContainer()
	containerMap = &sync.Map{}
}

func containerName(c any) string {
	switch v := c.(type) {
	case string:
		return v
	case ModuleOption:
		return v.name()
	default:
		return ""
	}
}

// LoadOrStoreContainer 用来保证一个容器是单例的
func LoadOrStoreContainer(c any) *Container {
	name := containerName(c)
	val, ok := containerMap.Load(name)
	if ok {
		if container, ok := val.(*Container); ok {
			return container
		}
	}

	lockVal, _ := lockMap.LoadOrStore(name, &sync.Mutex{})
	lock := lockVal.(*sync.Mutex)
	lock.Lock()
	defer lock.Unlock()

	val, ok = containerMap.Load(name)
	if ok {
		if container, ok := val.(*Container); ok {
			return container
		}
	}

	container := NewContainer()
	containerMap.Store(name, container)
	return container
}

func LoadContainer(c any) (*Container, bool) {
	name := containerName(c)
	val, ok := containerMap.Load(name)
	return val.(*Container), ok
}
