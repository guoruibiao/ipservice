package cache
// cache based on LRU algorithm

import (
	"container/list"
	"github.com/pkg/errors"
	"sync"
	"github.com/guoruibiao/ipservice/constants"
)



type LRUCache struct {
	sortedlist *list.List
	container map[string]*list.Element
	capacity int
	size int
	lock *sync.RWMutex
}

func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity <=0 {
		return nil, errors.New("capacity for LRUCache can not less than Zero.")
	}
	return &LRUCache{
		size: 0,
		capacity: capacity,
		container: make(map[string]*list.Element, capacity),
		sortedlist: list.New(),
		lock: new(sync.RWMutex),
	}, nil
}

func (c *LRUCache)Cache(key string, e *constants.Entry) (bool, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.size == c.capacity {
		// remove sorted list front
		oldest := c.sortedlist.Back()
		if oldest != nil {
			c.sortedlist.Remove(oldest)
			// container 会惰性更新，不用理会
		}
	}
	element := &list.Element{Value:e}
	if e, exists := c.container[key]; exists == false {
		c.sortedlist.PushFront(element)
	}else{
		element = e
		c.sortedlist.MoveToFront(e)
	}
	c.container[key] = element
	c.size += 1
	return true, nil
}

func (c *LRUCache)Get(key string)(*constants.Entry, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if e, exists := c.container[key]; exists == true {
		c.sortedlist.MoveToFront(e)
		return e.Value.(*constants.Entry), nil
	}
	return nil, errors.New("no such cache for " + key + " ")
}

// todo remove after this line.
func (c *LRUCache)PrintLRUCacheList() *list.List {
	return c.sortedlist
}

func (c *LRUCache)PrintLRUCacheContainer() map[string]*list.Element {
	return c.container
}

