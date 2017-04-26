package main

import (
    "fmt"
)

func (b *Board) Display() {
    fmt.Printf("%s\n", b);
}

func (g *Game) Display() {
    fmt.Printf("%s\n", g.board);
}
