package tapdance

import "sync"

type counter_uint struct {
	sync.Mutex
	value uint
}

func (c *counter_uint) inc() (uint) {
	c.Lock()
	c.value++
	c.Unlock()
	return c.value
}

func (c *counter_uint) dec() (uint) {
	c.Lock()
	c.value--
	c.Unlock()
	return c.value
}

func (c *counter_uint) get() (value uint) {
	c.Lock()
	value = c.value
	c.Unlock()
	return
}
