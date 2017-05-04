package main

import (
    "fmt"
    "math/rand"
    "strconv"
)

const (
    PLAY = iota
    QUIT = iota
    RETRY = iota
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

// CopyBoard returns a deep copy
func CopyBoard(o *Board) *Board {
    b := new(Board)
    b.size = Dimensions{width: o.size.width, height: o.size.height}
    b.cells = make([]int, b.size.width * b.size.height)
    copy(b.cells, o.cells)
    return b
}

// Equals returns if the two objects are deeply equal
func (b *Board) Equals(o *Board) bool {
    if b.size.width != o.size.width {
        return false
    }

    if b.size.height != o.size.height {
        return false
    }

    for i := 0; i < len(b.cells); i++ {
        if b.cells[i] != o.cells[i] {
            return false
        }
    }

    return true
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

        if (rand.Float64() < chanceFor4) {
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

// SlideDown slides all cells to the bottom
func (b *Board) SlideDown() *Board {
    for x := 0; x < b.size.width; x++ {
        limit := b.size.height - 1

        for y := b.size.height - 2; y >= 0; y-- {
            if (b.GetCell(x, y) == 0) {
                continue
            }

            for z := y; z + 1 <= limit; z++ {
                curr := b.GetCell(x, z)
                next := b.GetCell(x, z + 1)

                if next == 0 {
                    b.SetCell(x, z + 1, curr)
                    b.SetCell(x, z, 0)
                    continue
                }
                
                if next == curr {
                    b.SetCell(x, z + 1, curr * 2)
                    b.SetCell(x, z, 0)
                    limit = z
                    break
                }

                if next != curr {
                    break
                }
            }
        }
    }

    return b
}

// SlideLeft slides all cells to the left
func (b *Board) SlideLeft() *Board {
    for y := 0; y < b.size.height; y++ {
        limit := 0

        for x := 1; x < b.size.width; x++ {
            if (b.GetCell(x, y) == 0) {
                continue
            }

            for z := x; z - 1 >= limit; z-- {
                curr := b.GetCell(z, y)
                next := b.GetCell(z - 1, y)

                if next == 0 {
                    b.SetCell(z - 1, y, curr)
                    b.SetCell(z, y, 0)
                    continue
                }

                if next == curr {
                    b.SetCell(z - 1, y, curr * 2)
                    b.SetCell(z, y, 0)
                    limit = z
                    break
                }

                if next != curr {
                    break
                }
            }
        }
    }

    return b
}

// SlideRight slides all cells to the right
func (b *Board) SlideRight() *Board {
    for y := 0; y < b.size.height; y++ {
        limit := b.size.width - 1

        for x := b.size.width - 2; x >= 0; x-- {
            if (b.GetCell(x, y) == 0) {
                continue
            }

            for z := x; z + 1 <= limit; z++ {
                curr := b.GetCell(z, y)
                next := b.GetCell(z + 1, y)

                if next == 0 {
                    b.SetCell(z + 1, y, curr)
                    b.SetCell(z, y, 0)
                    continue
                }

                if next == curr {
                    b.SetCell(z + 1, y, curr * 2)
                    b.SetCell(z, y, 0)
                    limit = z
                    break
                }

                if next != curr {
                    break
                }
            }
        }
    }

    return b
}

// SlideUp slides all cells to the top
func (b *Board) SlideUp() *Board {
    for x := 0; x < b.size.width; x++ {
        limit := 0

        for y := 1; y < b.size.height; y++ {
            if (b.GetCell(x, y) == 0) {
                continue
            }

            for z := y; z - 1 >= limit; z-- {
                curr := b.GetCell(x, z)
                next := b.GetCell(x, z - 1)

                if next == 0 {
                    b.SetCell(x, z - 1, curr)
                    b.SetCell(x, z, 0)
                    continue
                }

                if next == curr {
                    b.SetCell(x, z - 1, curr * 2)
                    b.SetCell(x, z, 0)
                    limit = z
                    break
                }

                if next != curr {
                    break
                }
            }
        }
    }

    return b
}

// Spawn add a new cell to the board
func (b *Board) Spawn() {
    available := make([]int, 0, len(b.cells))

    for i, v := range b.cells {
        if (v == 0) {
            available = append(available, i);
        }
    }

    index := available[rand.Intn(len(available))]
    b.cells[index] = 2
    chanceFor4 := 0.25

    if (rand.Float64() < chanceFor4) {
        b.cells[index] = 4
    }  
}

// String returns the string representation of the board
func (b *Board) String() string {
    board := ""
    width := b.size.width

    for i, v := range b.cells {
        if (i % width == 0) {
            board += "\n|"

            for j := 0; j < width; j++ {
                board += "----|"
            }

            board += "\n|"
        }
        
        value := strconv.Itoa(v)

        if (v == 0) {
            value = ""
        }

        value = fmt.Sprintf("%4s", value);
        board += value + "|";
    }
    
    board += "\n|"

    for w := 0; w < width; w++ {
        board += "¯¯¯¯|"
    }

    return board;
}

// Game contains information regarding the state
type Game struct {
    board *Board
    state int
}

// NewGame returns an initialized Game
func NewGame(width, height int) *Game {
    g := new(Game)
    g.board = NewBoard(width, height)
    g.board.Populate(2)
    g.board.Shuffle()
    g.state = PLAY
    return g
}

// MovesLeft returns whether or not there are possible moves
func (g *Game) MovesLeft() bool {
    base := CopyBoard(g.board)
    return !base.Equals(CopyBoard(g.board).SlideDown()) ||
        !base.Equals(CopyBoard(g.board).SlideLeft()) ||
        !base.Equals(CopyBoard(g.board).SlideRight()) ||
        !base.Equals(CopyBoard(g.board).SlideUp())
}
