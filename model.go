package main

import (
    "fmt"
    "math/rand"
    "strconv"
)

// dimensions contains width and height
type Dimensions struct {
    width, height int
}

// Board contains information regarding the cells
type Board struct {
    cells []int
    size Dimensions
}

// NewBoard returns an initialized Board
func NewBoard(width, height int) *Board {
    b := new(Board)
    b.size = Dimensions{width: width, height: height}
    b.cells = make([]int, b.size.width * b.size.height)
    return b
}

// GetDimensions returns the dimensions of the board
func (b *Board) GetDimensions() Dimensions {
    return b.size
}

// GetBoard returns the value specified at x, y
func (b *Board) GetCell(x, y int) int {
    x += b.size.width
    x %= b.size.width
    y += b.size.height
    y %= b.size.height
    return b.cells[b.size.width * y + x]
}

// Populate randomly populates the board with n cells
func (b *Board) Populate(n int) {
    chanceFor4 := 0.25

    for i := 0; i < n; i++ {
        b.cells[i] = 2
        r := rand.Float64()

        if (r < chanceFor4) {
            b.cells[i] = 4
        }  
    }
}

// SetBoard sets the specified value at x, y to supplied value
func (b *Board) SetCell(x, y, value int) {
    x += b.size.width
    x %= b.size.width
    y += b.size.height
    y %= b.size.height
    b.cells[b.size.width * y + x] = value
}

// Shuffle randomly shuffles board using Fisher-Yates' algorithm
func (b *Board) Shuffle() {
    c := b.cells

    for i := len(c) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        c[i], c[j] = c[j], c[i]
    }
}

// String returns the string representation of the board
func (b *Board) String() string {
    // TODO: Padding needs to be dynamic
    board := "|"
    width := b.size.width

    for i, v := range b.cells {
        if (i % width == 0 && i != 0) {
            board += "\n|"

            for j := 0; j < width; j++ {
                board += "----|"
            }

            board += "\n|"
        }
        
        value := fmt.Sprintf("%4s", strconv.Itoa(v))
        board += value + "|";
    }
    
    return board;
}

// Game contains information regarding the state
type Game struct {
    board *Board
}

// NewGame returns an initialized Game
func NewGame(width, height int) *Game {
    g := new(Game)
    g.board = NewBoard(width, height)
    g.board.Populate(2)
    g.board.Shuffle()
    return g
}
