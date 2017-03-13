package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cyrusroshan/backstage-go/backstage"
	"github.com/cyrusroshan/backstage-go/backstage/chart"
)

func main() {
	randomIntChart := chart.NewChart("Random Int Chart", chart.BAR, "")
	backstage.Start(
		"Sample Application",
		&backstage.Info{
			Port:        9999,
			RefreshRate: 1000,
		},
		[]*chart.Chart{
			randomIntChart,
		},
	)

	for {
		randInt := rand.Intn(100)

		fmt.Println("Pushing", randInt)
		randomIntChart.Push(randInt)

		time.Sleep(500 * time.Millisecond)
	}
}
