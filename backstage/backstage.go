package backstage

import (
	"container/ring"
	"fmt"
	"log"
	"net/http"
)

type Info struct {
	ClientLogs     bool   `json:"clientLogs"`
	ServerLogs     bool   `json:"serverLogs"`
	GlobalDefaults string `json:"globalDefaults"`
	Port           int
}

type stageFloor struct {
	Info   *Info
	Charts []*Chart
}

func Create(info Info) *stageFloor {
	if info.Port == 0 {
		info.Port = 8080
	}

	stage := stageFloor{
		Info: &info,
	}

	http.HandleFunc("/info", sendInfo(&info))
	http.HandleFunc("/", sendData(&stage))
	go http.ListenAndServe(fmt.Sprintf(":%d", info.Port), nil)
	log.Printf("backstage running on port %d\n", info.Port)

	return &stage
}

func (s *stageFloor) NewChart(name string) *Chart {
	chart := Chart{
		Name:       name,
		DataBuffer: ring.New(30),
	}
	s.Charts = append(s.Charts, &chart)

	return &chart
}
