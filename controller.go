package main

import (
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    g := NewGame(4, 4)
    g.Display()
}
