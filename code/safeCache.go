import "sync"

type safeCache struct {
	m map[string]interface{}
	sync.RWMutex
}

func (c *safeCache) Set(key string, value interface{}) {
	c.Lock()
	c.m[key] = value
	c.Unlock()
}

func (c *safeCache) Get(key) (interface{}, bool) {
	c.RLock()
	v, ok := c.m[key]
	c.RUnlock()
	return v, ok
}
