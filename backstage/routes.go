package backstage

import (
	"net/http"

	"github.com/cyrusroshan/backstage-go/utils"
)

func sendInfo(info *Info) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(utils.MustMarshal(&info))
	}
}

type chartRead struct {
	Name string
	Data []interface{}
}

func sendData(c *stageFloor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		chartData := make([]chartRead, len(c.Charts))
		for i, chart := range c.Charts {
			chartData[i] = chartRead{
				Name: chart.Name,
				Data: chart.readAndClear(),
			}
		}

		w.WriteHeader(200)
		w.Write(utils.MustMarshal(chartData))
	}
}
