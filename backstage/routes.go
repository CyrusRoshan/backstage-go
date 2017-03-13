package backstage

import (
	"net/http"

	"github.com/cyrusroshan/backstage-go/utils"
)

const instructionString = "Send backstage GET requests to /ping, /info, and /data"

func giveInstructions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(instructionString))
}

type chartInfo struct {
	Name string
	Info string
}

type infoContainer struct {
	GlobalInfo *Info
	ChartInfo  []chartInfo
}

func sendInfo(s *stageFloor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sentInfo := infoContainer{
			GlobalInfo: s.Info,
			ChartInfo:  make([]chartInfo, len(s.Charts)),
		}

		for i, chart := range s.Charts {
			sentInfo.ChartInfo[i] = chartInfo{
				Name: chart.Name,
				Info: chart.Info,
			}
		}

		w.WriteHeader(200)
		w.Write(utils.MustMarshal(&sentInfo))
	}
}

func ping(s *stageFloor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(s.Name))
	}
}

func sendData(s *stageFloor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, chart := range s.Charts {
			chart.ReadAndClear()
		}

		w.WriteHeader(200)
		w.Write(utils.MustMarshal(s.Charts))
	}
}
