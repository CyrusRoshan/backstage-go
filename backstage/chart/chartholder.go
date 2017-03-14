package chart

type ChartHolder struct {
	charts []*Chart
}

func (c *ChartHolder) Add(chart *Chart) {
	c.charts = append(c.charts, chart)
}
