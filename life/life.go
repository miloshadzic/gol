package life

import "fmt"

// Consists of two boards, one is current state, and the other is used
// to build the next generation
type Life struct {
	state [][]bool
	buf   [][]bool
}

// Initiates a new Game
func Init(width, height int) *Life {
	state := CreateBoard(width, height)
	buf := CreateBoard(width, height)

	return &Life{state, buf}
}

// Factory to build an empty board.
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

// Updates the state a generation.
func (this *Life) Next() {
	for x := range this.state {
		for y := range this.state[x] {
			this.buf[x][y] = this.isAlive(x, y)
		}
	}

	this.state, this.buf = this.buf, this.state
}

// Returns the neighbour count for a given pair of coordinates in the
// current game state.
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

// Is the cell at coordinate pair (x, y) going to live in the next
// generation.
func (this *Life) isAlive(x, y int) bool {
	count := this.NeighbourCount(x, y)

	return (this.state[x][y] && (count == 2 || count == 3)) ||
		(!this.state[x][y] && count == 3)
}

// Returns the current state slice.
func (this *Life) State() [][]bool {
	return this.state
}

// Draw outputs the current state of the board.
func (this *Life) Draw() {
	// Render in reverse because terminal renders from the top
	for y := len(this.state) - 1; y >= 0; y-- {
		for x := range this.state[y] {
			fmt.Print(drawCell(this.state[x][y]))
		}
		fmt.Print("\n")
	}
}

func drawCell(cell bool) string {
	if cell {
		return "O "
	} else {
		return ". "
	}
}

// Utility function to add a floater at coordinates (x, y)
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
