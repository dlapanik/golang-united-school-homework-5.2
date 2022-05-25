package cache

import "time"

type cacheRecord struct {
	value    string
	deadline time.Time
}

type Cache struct {
	data map[string]cacheRecord
}

func NewCache() Cache {
	return Cache{data: make(map[string]cacheRecord)}
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.data[key]

	if !ok {
		return "", false
	}

	if !value.deadline.IsZero() && value.deadline.Before(time.Now()) {
		delete(c.data, key)
		return "", false
	}

	return value.value, true
}

func (c *Cache) Put(key, value string) {
	c.data[key] = cacheRecord{
		value:    value,
		deadline: time.Time{}, // zero value for Time
	}
}

func (c *Cache) Keys() []string {
	keys := make([]string, len(c.data))
	now := time.Now()

	for key, value := range c.data {
		if !value.deadline.IsZero() && value.deadline.Before(now) {
			delete(c.data, key)
		} else {
			keys = append(keys, key)
		}
	}

	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = cacheRecord{
		value:    value,
		deadline: deadline,
	}
}
