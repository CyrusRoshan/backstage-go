package backstage

import (
	"container/ring"
	"fmt"
)

type Chart struct {
	Name       string
	DataBuffer *ring.Ring
}

func (c *Chart) Push(data interface{}) {
	c.DataBuffer.Value = data
	c.DataBuffer = c.DataBuffer.Next()
}

func (c *Chart) readAndClear() []interface{} {
	length := c.DataBuffer.Len()
	dataArr := make([]interface{}, length)

	for i := 0; i < length; i++ {
		if c.DataBuffer.Value == nil {
			c.DataBuffer = c.DataBuffer.Next()
		}
	}

	i := 0
	for i < length {
		data := c.DataBuffer.Value
		fmt.Println(data)
		if data == nil {
			break
		}

		dataArr[i] = data
		c.DataBuffer.Value = nil
		i++
		c.DataBuffer = c.DataBuffer.Next()
	}

	return dataArr[:i]
}
