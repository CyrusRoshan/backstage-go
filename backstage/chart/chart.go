package chart

import "container/ring"

type chartType string

const (
	DOUGHNUT      = chartType("Doughnut")
	PIE           = chartType("Pie")
	LINE          = chartType("Line")
	BAR           = chartType("Bar")
	HORIZONTALBAR = chartType("HorizontalBar")
	RADAR         = chartType("Radar")
	POLARAREA     = chartType("PolarArea")
	BUBBLE        = chartType("Bubble")
)

type Chart struct {
	Name     string
	Type     chartType
	Labels   []string
	Options  string `json:"-"`
	Datasets []*Dataset
}

type Dataset struct {
	Options    string
	Data       []interface{}
	DataBuffer *ring.Ring `json:"-"`
}

func (c *Dataset) Push(data interface{}) {
	d.DataBuffer.Value = data
	d.DataBuffer = d.DataBuffer.Next()
}

func (d *Dataset) readAndClear() *[]interface{} {
	length := d.DataBuffer.Len()
	d.ReadData = make([]interface{}, length)

	for i := 0; i < length; i++ {
		if d.DataBuffer.Value == nil {
			d.DataBuffer = d.DataBuffer.Next()
		}
	}

	i := 0
	for i < length {
		data := d.DataBuffer.Value
		if data == nil {
			break
		}

		d.ReadData[i] = data
		d.DataBuffer.Value = nil
		i++
		d.DataBuffer = d.DataBuffer.Next()
	}

	d.ReadData = d.ReadData[:i]
	return &(d.ReadData)
}
