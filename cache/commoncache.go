package cache

import (
	"sync"
	"github.com/pkg/errors"
	)

type Entry struct {
	RegionName string
	CityName string
}

type CommonCache struct {
	container map[string]Entry
	lock *sync.RWMutex
}

// common cache
func NewCommonCacher() *CommonCache {
	return &CommonCache{
		container:make(map[string]Entry),
		lock: new(sync.RWMutex),
	}
}

func (c *CommonCache)Cache(key string, e Entry) (bool, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.container[key] = e
	return true, nil
}

func (c *CommonCache)Get(key string) (*Entry, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	value, exists := c.container[key]
	if exists == false {
		return nil, errors.New("no cache for " + key + " ")
	}
	return &value, nil
}

// todo remove this.
func (c *CommonCache)PrintContainer() map[string]Entry{
	return c.container
}