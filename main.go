package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cyrusroshan/backstage-go/backstage"
)

func main() {
	info := backstage.Info{}
	stage := backstage.Create(info)

	randomInt := stage.NewChart("randomInt")

	for {
		randInt := rand.Intn(100)

		fmt.Println("Pushing", randInt)
		randomInt.Push(randInt)

		time.Sleep(500 * time.Millisecond)
	}
}
