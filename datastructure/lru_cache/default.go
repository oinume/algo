package lru_cache

import (
	"fmt"
	"io"
	"math"
	"strings"
	"sync"
)

type defaultLRUCache struct {
	capacity   int
	values     map[int]*item
	currentAge int
	mutex      *sync.Mutex
}

func NewDefault(capacity int) LRUCache {
	return &defaultLRUCache{
		capacity:   capacity,
		values:     make(map[int]*item, capacity),
		currentAge: 0,
		mutex:      new(sync.Mutex),
	}
}

func (c *defaultLRUCache) Get(key int) int {
	i, ok := c.values[key]
	if !ok {
		return -1
	}
	c.mutex.Lock()
	i.age = c.currentAge
	// `Get` also increment current age
	c.currentAge++
	c.mutex.Unlock()
	return i.value
}

func (c *defaultLRUCache) Put(key int, value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if i, ok := c.values[key]; ok {
		// If the key exists, update its value and increment its age for this key
		i.value = value
		i.age = c.currentAge
		c.currentAge++
	} else {
		if len(c.values) >= c.capacity {
			// Search key with least age when over capacity before setting key and value
			leastAge := math.MaxInt32
			leastAgeKey := 0
			for key, item := range c.values {
				if item.age < leastAge {
					leastAge = item.age
					leastAgeKey = key
				}
			}
			if leastAgeKey != 0 {
				// Evict least age key from cache
				delete(c.values, leastAgeKey)
			}
		}
		// Set key and value to cache
		c.values[key] = &item{
			value: value,
			age:   c.currentAge,
		}
		c.currentAge++
	}
}

func (c *defaultLRUCache) Dump(w io.Writer) error {
	t := `
{
  capacity: %v
  currentAge: %v
  values: { %v }
}
`
	t = strings.TrimSpace(t)
	a := make([]string, 0, len(c.values))
	for k, i := range c.values {
		a = append(a, fmt.Sprintf(`"%v":%+v`, k, i))
	}
	values := "{" + strings.Join(a, ", ") + "}"
	if _, err := fmt.Fprintf(w, t, c.capacity, c.currentAge, values); err != nil {
		return err
	}
	return nil
}
