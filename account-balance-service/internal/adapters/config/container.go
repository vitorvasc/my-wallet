package config

import "sync"

type ContainerService struct {
	services map[string]interface{}
	lock     sync.RWMutex
}

func NewContainer() *ContainerService {
	return &ContainerService{
		services: make(map[string]interface{}),
	}
}

func (c *ContainerService) Register(name string, service interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.services[name] = service
}

func (c *ContainerService) Get(name string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()

	service, ok := c.services[name]
	if !ok {
		return nil
	}

	return service
}
