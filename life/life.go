package life

import "fmt"

type Life struct {
	state [][]bool
	buf   [][]bool
}

func Init(width, height int) *Life {
	state := CreateBoard(width, height)
	buf := CreateBoard(width, height)

	return &Life{state, buf}
}

func CreateBoard(width, height int) [][]bool {
	// 2D slice single-allocation trick from
	// https://golang.org/doc/effective_go#two_dimensional_slices
	board := make([][]bool, height)
	cells := make([]bool, width*height)

	for i := range board {
		board[i], cells = cells[:width], cells[width:]
	}

	return board
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

	return (this.state[x][y] && (count == 2 || count == 3)) ||
		(!this.state[x][y] && count == 3)
}

func (this *Life) NeighbourCount(x, y int) int {
	count := 0

	xMin := x - 1
	if xMin < 0 {
		xMin = 0
	}

	yMin := y - 1
	if yMin < 0 {
		yMin = 0
	}

	yMax := y + 2
	if yMax > len(this.state[0]) {
		yMax = len(this.state[0])
	}

	xMax := x + 2
	if xMax > len(this.state) {
		xMax = len(this.state)
	}

	for col := xMin; col < xMax; col++ {
		for row := yMin; row < yMax; row++ {
			if !(row == y && col == x) && this.state[col][row] {
				count++
			}
		}
	}

	return count
}

func (this *Life) State() [][]bool {
	return this.state
}

func drawCell(cell bool) string {
	if cell {
		return "O "
	} else {
		return ". "
	}
}

func (this *Life) Draw() {
	// Render in reverse because terminal renders from the top
	for y := len(this.state) - 1; y >= 0; y-- {
		for x := range this.state[y] {
			fmt.Print(drawCell(this.state[x][y]))
		}
		fmt.Print("\n")
	}
}

func (this *Life) AddFloater(x, y int) {
	// I don't want to deal with overflow here.
	if x+2 >= len(this.state[0]) || y+2 >= len(this.state[0]) {
		return
	}

	this.state[x][y] = true
	this.state[x+1][y] = true
	this.state[x+2][y] = true
	this.state[x][y+1] = true
	this.state[x+1][y+2] = true
}
