package simplecache

import (
	"sync"
)

type Entry struct {
	key string

	value interface{}
}

type Cache struct {
	mutex sync.RWMutex

	entries map[string]*Entry
}

func New() *Cache {
	return &Cache{
		mutex:   sync.RWMutex{},
		entries: make(map[string]*Entry),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = &Entry{
		key:   key,
		value: value,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	v, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return v.value, true
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, ok := c.entries[key]
	if ok {
		delete(c.entries, key)
	}
}
