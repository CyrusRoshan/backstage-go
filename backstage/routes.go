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

func sendData(c *stageFloor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, chart := range c.Charts {
			chart.readAndClear()
		}

		w.WriteHeader(200)
		w.Write(utils.MustMarshal(c.Charts))
	}
}
