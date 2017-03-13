package chart

import "container/ring"

type chartType string

const (
	DOUGHNUT      = chartType("doughnut")
	PIE           = chartType("pie")
	LINE          = chartType("line")
	BAR           = chartType("bar")
	HORIZONTALBAR = chartType("horizontalBar")
	RADAR         = chartType("radar")
	POLARAREA     = chartType("polarArea")
	BUBBLE        = chartType("bubble")
)

type Chart struct {
	Name       string
	Type       chartType `json:"-"`
	Info       string    `json:"-"`
	ReadData   []interface{}
	DataBuffer *ring.Ring `json:"-"`
}

func NewChart(name string, chartType chartType, info string) *Chart {
	return &Chart{
		Name:       name,
		Info:       info,
		Type:       chartType,
		DataBuffer: ring.New(30),
	}
}

func (c *Chart) Push(data interface{}) {
	c.DataBuffer.Value = data
	c.DataBuffer = c.DataBuffer.Next()
}

func (c *Chart) ReadAndClear() *[]interface{} {
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
