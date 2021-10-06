package main

import (
	"fmt"
	"time"
)

func main() {
	life := Init()

	for {
		fmt.Print("\033[H\033[2J")

		life.Draw()
		time.Sleep(1 * time.Second)
		life.Next()
	}
}
