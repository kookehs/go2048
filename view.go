package main

import (
    "fmt"
)

func (b *Board) Display() {
    fmt.Println(b);
}

func (g *Game) Display() {
    fmt.Println(g.board);
}
