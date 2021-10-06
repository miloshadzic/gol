package life_test

import (
	"fmt"
	"gol/life"
	"reflect"
	"testing"
)

func TestNeighbourCount(t *testing.T) {
	game := life.Init(5, 5)

	game.AddFloater(1, 1)
	// 4 . . . . .
	// 3 . . O . .
	// 2 . O . . .
	// 1 . O O O .
	// 0 . . . . .
	//   0 1 2 3 4

	if actual := game.NeighbourCount(1, 1); actual != 2 {
		t.Fatal("Wrong neighbor count for 1,1. Expected 2, got", actual)
	}

	if actual := game.NeighbourCount(2, 1); actual != 3 {
		t.Fatal("Wrong neighbor count for 2,1. Expected 3, got", actual)
	}

	if actual := game.NeighbourCount(3, 2); actual != 3 {
		t.Fatal("Wrong neighbor count for 3,2. Expected 3, got", actual)
	}

	if actual := game.NeighbourCount(3, 3); actual != 1 {
		t.Fatal("Wrong neighbor count for 3,3. Expected 1, got", actual)
	}

	if actual := game.NeighbourCount(4, 1); actual != 1 {
		t.Fatal("Wrong neighbor count for 4,1. Expected 1, got", actual)
	}
}

func TestNextFloater(t *testing.T) {
	game := life.Init(5, 5)

	game.AddFloater(1, 1)
	printBoard(game.State())

	//   Initial
	// 4 . . . . .
	// 3 . . O . .
	// 2 . O . . .
	// 1 . O O O .
	// 0 . . . . .
	//   0 1 2 3 4

	expected := life.CreateBoard(5, 5)

	// Manually add the next floater iteration
	expected[2][0] = true
	expected[2][1] = true
	expected[1][1] = true
	expected[3][2] = true
	expected[1][2] = true

	printBoard(expected)

	//   Expected
	// 4 . . . . .
	// 3 . . . . .
	// 2 . O . O .
	// 1 . . O O .
	// 0 . . O . .
	//   0 1 2 3 4

	game.Next()

	printBoard(game.State())

	if !reflect.DeepEqual(game.State(), expected) {
		t.Fatal("Next step is incorrect.")
	}
}

func printBoard(board [][]bool) {
	for y := len(board) - 1; y >= 0; y-- {
		for x := range board[y] {
			if board[x][y] {
				fmt.Print("O ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
