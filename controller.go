package main

import (
    "bufio"
    "fmt"
    "log"
    "math/rand"
    "os"
    "strings"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    g := NewGame(4, 4)
    g.Display()

    r := bufio.NewReader(os.Stdin)

    for {
        input, err :=  r.ReadString('\n')

        if err != nil {
            log.Fatal(err)
        }

        update(strings.ToLower(input), g)
        g.Display()
        lateUpdate(g)

        if g.state == RETRY {
            g.state = QUIT

            if retry(r) {
                g = NewGame(4, 4)
                g.Display()
            }
        }

        if g.state == QUIT {
            break
        }

    }
}

// lateUpdate does additional checks after game has been updated
func lateUpdate(g *Game) {
    if (!g.MovesLeft()) {
        g.state = RETRY
    }
}

// retry prompts the player if they would like to play again
func retry(r *bufio.Reader) bool {
    fmt.Println("Game over. Do you want to play again?")
    input, err := r.ReadString('\n');
            
    if err != nil {
        log.Fatal(err)
    }

    trimmed := strings.TrimRight(input, "\n")

    switch strings.ToLower(trimmed) {
    case "y":
        fallthrough
    case "yes":
        return true
    case "n":
        fallthrough
    case "no":
        fallthrough
    default:
        return false
    }

    return false
}

// update takes action based on input
func update(i string, g *Game) {
    trimmed := strings.TrimRight(i, "\n")
    old := CopyBoard(g.board)

    switch trimmed {
    case "w":
        _, score := g.board.SlideUp()
        g.AddScore(score)
    case "a":
        _, score := g.board.SlideLeft()
        g.AddScore(score)
    case "s":
        _, score := g.board.SlideDown()
        g.AddScore(score)
    case "d":
        _, score := g.board.SlideRight()
        g.AddScore(score)
    case "q":
        g.state = QUIT
    default:
        return
    }

    if !g.board.Equals(old) {
        g.board.Spawn()
    }
}
