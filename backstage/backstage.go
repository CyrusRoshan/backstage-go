package backstage

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/cyrusroshan/backstage-go/backstage/chart"
	"github.com/cyrusroshan/qli/utils"
)

type Info struct {
	GlobalDefaults      string `json:"globalDefaults"`
	ClientLogs          bool   `json:"clientLogs"`
	ServerLogs          bool   `json:"serverLogs"`
	DisableTerminalLogs bool
	Port                int `json:"-"`
}

type stageFloor struct {
	Name   string
	Info   *Info
	Charts []*chart.Chart
}

func Start(name string, info *Info, charts []*chart.Chart) {
	if info.Port == 0 {
		info.Port = 9999
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", info.Port))
	if err != nil {
		l, err := net.Listen("tcp", ":0")
		utils.PanicIf(err)

		info.Port = l.Addr().(*net.TCPAddr).Port
	}
	l.Close()

	stage := stageFloor{
		Name:   name,
		Info:   info,
		Charts: charts,
	}

	backstage := http.NewServeMux()
	backstage.HandleFunc("/ping", ping(&stage))
	backstage.HandleFunc("/info", sendInfo(&stage))
	backstage.HandleFunc("/data", sendData(&stage))
	backstage.HandleFunc("/", giveInstructions)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", stage.Info.Port), backstage)
		utils.PanicIf(err)
	}()

	if !stage.Info.DisableTerminalLogs {
		log.Printf("backstage running on port %d\n", stage.Info.Port)
	}

	return
}
