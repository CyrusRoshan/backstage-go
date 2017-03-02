package backstage

import "container/ring"

type Chart struct {
	Name       string
	Info       string
	ReadData   []interface{}
	DataBuffer *ring.Ring `json:"-"`
}

func (c *Chart) Push(data interface{}) {
	c.DataBuffer.Value = data
	c.DataBuffer = c.DataBuffer.Next()
}

func (c *Chart) readAndClear() *[]interface{} {
	length := c.DataBuffer.Len()
	c.ReadData = make([]interface{}, length)

	for i := 0; i < length; i++ {
		if c.DataBuffer.Value == nil {
			c.DataBuffer = c.DataBuffer.Next()
		}
	}

	i := 0
	for i < length {
		data := c.DataBuffer.Value
		if data == nil {
			break
		}

		c.ReadData[i] = data
		c.DataBuffer.Value = nil
		i++
		c.DataBuffer = c.DataBuffer.Next()
	}

	c.ReadData = c.ReadData[:i]
	return &(c.ReadData)
}
