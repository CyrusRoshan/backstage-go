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
	GlobalDefaults      string
	ClientLogs          bool
	ServerLogs          bool
	DisableTerminalLogs bool
	Port                int `json:"-"`
	RefreshRate         int
}

type stageFloor struct {
	Name        string
	Info        *Info
	ChartHolder *chart.ChartHolder
}

func Start(name string, info *Info, chartHolder *chart.ChartHolder) {
	setDefaults(info)

	stage := stageFloor{
		Name:        name,
		Info:        info,
		ChartHolder: chartHolder,
	}

	backstage := http.NewServeMux()
	backstage.HandleFunc("/ping", headerModificationWrapper(ping()))
	backstage.HandleFunc("/info", headerModificationWrapper(sendInfo(&stage)))
	backstage.HandleFunc("/data", headerModificationWrapper(sendData(&stage)))
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

func setDefaults(info *Info) {
	if info.Port == 0 {
		info.Port = 9999
	}

	if info.RefreshRate < 1 {
		info.RefreshRate = 1000
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", info.Port))
	if err != nil {
		l, err = net.Listen("tcp", ":0")
		utils.PanicIf(err)

		info.Port = l.Addr().(*net.TCPAddr).Port
	}
	l.Close()

	return
}
