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
	Name    string
	Options string
}

type infoContainer struct {
	Name       string
	GlobalInfo *Info
	ChartInfo  []chartInfo
}

func headerModificationWrapper(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

func sendInfo(s *stageFloor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sentInfo := infoContainer{
			Name:       s.Name,
			GlobalInfo: s.Info,
			ChartInfo:  make([]chartInfo, len(s.Charts)),
		}

		for i, chart := range s.Charts {
			sentInfo.ChartInfo[i] = chartInfo{
				Name:    chart.Name,
				Options: chart.Options,
			}
		}

		w.WriteHeader(200)
		w.Write(utils.MustMarshal(&sentInfo))
	}
}

func ping() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(""))
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
