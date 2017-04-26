package main

import (
    "bufio"
    "log"
    "math/rand"
    "os"
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

        trimmed := input[:len(input) - 1]

        if (handleInput(trimmed, g)) {
            break;
        }

        g.Display()
    }
}

// handleInput takes action based on input.
// Returns true if application is to close.
func handleInput(i string, g *Game) bool {
    switch i {
    case "w":
        g.board.SlideUp();
        break;
    case "a":
        g.board.SlideLeft();
        break;
    case "s":
        g.board.SlideDown();
        break;
    case "d":
        g.board.SlideRight();
        break;
    case "q":
        return true;
    }

    return false;
}
