package main

import "fmt"

const (
	WIDTH  int = 50
	HEIGHT int = 50
)

type Life struct {
	state *[WIDTH][HEIGHT]bool
	buf   *[WIDTH][HEIGHT]bool
}

func Init() *Life {
	var state [WIDTH][HEIGHT]bool
	var buf [WIDTH][HEIGHT]bool

	state[15][15] = true
	state[15][16] = true
	state[15][17] = true
	state[16][17] = true
	state[17][16] = true

	return &Life{&state, &buf}
}

func (this *Life) Next() {
	for x := range this.state {
		for y := range this.state[x] {
			this.buf[x][y] = this.isAlive(x, y)
		}
	}

	this.state, this.buf = this.buf, this.state
}

func (this *Life) isAlive(x, y int) bool {
	count := this.NeighbourCount(x, y)

	switch {
	case this.state[x][y] && count < 2:
		return false
	case this.state[x][y] && (count == 2 || count == 3):
		return true
	case !this.state[x][y] && count > 3:
		return true
	default:
		return false
	}
}

func (this *Life) NeighbourCount(x, y int) int {
	count := 0

	min_y := y - 1
	if min_y < 0 {
		min_y = 0
	}

	min_x := x - 1
	if min_x < 0 {
		min_x = 0
	}

	max_y := y + 1
	if max_y < HEIGHT {
		max_y = HEIGHT
	}

	max_x := x + 1
	if max_x > WIDTH {
		max_x = WIDTH
	}

	for row := min_x; row < max_x; row++ {
		for col := min_y; col < max_y; col++ {
			if !(row == x && col == y) && this.state[row][col] {
				count++
			}
		}
	}

	return count
}

func (this *Life) State() *[WIDTH][HEIGHT]bool {
	return this.state
}

func drawCell(cell bool) string {
	if cell {
		return "O"
	} else {
		return "."
	}
}

func (this *Life) Draw() {
	for col := range this.state {
		for row := range this.state[col] {
			fmt.Print(drawCell(this.state[col][row]))
		}
		fmt.Print("\n")
	}
}
