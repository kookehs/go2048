package main

import (
    "strconv"
)

// dimensions contains width and height
type Dimensions struct {
    width, height int
}

// Game contains information regarding the state
type Game struct {
    board []int
    size Dimensions
}

// NewGame returns an initialized Game
func NewGame(width, height int) *Game {
    g := new(Game)
    g.size = Dimensions{width: width, height: height}
    g.board = make([]int, g.size.width * g.size.height)
    return g
}

// GetDimensions returns the dimensions of the board
func (g *Game) GetDimensions() Dimensions {
    return g.size
}

// GetBoard returns the value specified at x, y
func (g *Game) GetBoard(x, y int) int {
    x += g.size.width
    x %= g.size.width
    y += g.size.height
    y %= g.size.height
    return g.board[g.size.width * y + x]
}

// SetBoard sets the specified value at x, y to supplied value
func (g *Game) SetBoard(x, y, value int) {
    x += g.size.width
    x %= g.size.width
    y += g.size.height
    y %= g.size.height
    g.board[g.size.width * y + x] = value
}

func (g *Game) String() string {
    board := "|"

    for i, v := range g.board {
        if (i % g.size.width == 0 && i != 0) {
            board += "\n|"

            for j := 0; j < g.size.width; j++ {
                board += "----|"
            }

            board += "\n|"
        }
        
        // Pad based on length
        board += strconv.Itoa(v) + "   |";
    }
    
    return board;
}
