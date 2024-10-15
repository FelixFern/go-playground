package counter

func (c *Counter) IncMutex() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Value++
}

func (c *Counter) Inc() {
	c.Value++
}

func (c *Counter) GetValue() int {
	return c.Value
}
