package cache

type Cache struct {
	cache []*node
	size  int
}

type node struct {
	key   string
	value interface{}
}

func New() Cache {
	return Cache{}
}

func (c *Cache) Set(key string, value interface{}) {
	n := node{
		key:   key,
		value: value,
	}
	if v, ok := c.check(key); ok != true {
		c.cache = append(c.cache, &n)
		c.size += 1
	} else {
		c.cache[v] = &n
	}

}
func (c *Cache) Size() int {
	return c.size
}
func (c *Cache) check(key string) (int, bool) {
	if len(c.cache) == 0 {
		return 0, false
	} else {
		for i := range c.cache {
			if c.cache[i].key == key {
				return i, true
			}
		}
		return 0, false
	}
}

func (c *Cache) Get(key string) interface{} {
	if len(c.cache) == 0 {
		return "cache is empty"
	}
	if i, ok := c.check(key); ok != true {
		return ""
	} else {
		return c.cache[i].value
	}
}
func (c *Cache) Delete(key string) {
	if i, ok := c.check(key); ok == true {
		lastNode := c.cache[len(c.cache)-1]
		c.cache[i] = lastNode
		c.cache = c.cache[:len(c.cache)-1]
		c.size -= 1
	}
}
