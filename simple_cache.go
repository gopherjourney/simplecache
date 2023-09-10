package simplecache

import (
	"sync"
	"time"
)

const (
	NoExpiration = -1
)

type Entry struct {
	key string

	value interface{}

	expiration int64
}

func (e *Entry) Expired() bool {
	if e.expiration > 0 {
		return time.Now().UnixNano() > e.expiration
	}

	return false
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
	c.SetWithTTL(key, value, NoExpiration)
}

func (c *Cache) SetWithTTL(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	exp := time.Now().Add(ttl).UnixNano()
	if ttl == NoExpiration {
		exp = NoExpiration
	}

	c.entries[key] = &Entry{
		key:        key,
		value:      value,
		expiration: exp,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	v, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	if v.Expired() {
		c.Delete(key)
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
