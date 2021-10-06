package main

import (
	"fmt"
	"gol/life"
	"time"
)

func main() {
	life := life.Init(30, 30)

	life.AddFloater(15, 15)

	for {
		fmt.Print("\033[H\033[2J")

		life.Draw()
		time.Sleep(500 * time.Millisecond)
		life.Next()
	}
}
