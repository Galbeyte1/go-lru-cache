package lru

import "container/list"

// set element, can I get the element
// map

// max capacity -- maintain the most recently used elements -- reserve LRU
//	track least recently used -- order

type Cache interface {
	Get (key string) string
	Set (key string, value string)
}

type cacheMapElement struct {
	el *list.Element
	value string
}

type LRUCache struct {
	m map[string]*cacheMapElement
	cap int
	l list.List
}

func newLRU(cap int) LRUCache {
	return LRUCache{
		m: map[string]*cacheMapElement{},
		cap: cap,
		l: list.List{},
	}
}

func (c *LRUCache) Get(key string) string {
	v, ok := c.m[key]
	if !ok {
		return ""
	}

	c.l.MoveToFront(v.el)
	return v.value
}

func (c *LRUCache) Set(key string, value string) {
	v, ok := c.m[key]
	if !ok {
		el := c.l.PushFront(key)
		c.m[key] = &cacheMapElement{
			el: el,
			value: value,
		}

		if c.l.Len() > c.cap {
			backEl := c.l.Back()
			backElementKey := backEl.Value.(string)
			c.l.Remove(backEl)
			delete(c.m, backElementKey)
		}

	} else {
		v.value = value
		c.l.MoveToFront(v.el)
	}
}