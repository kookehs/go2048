package main

import (
    "math"
    "math/rand"
    "sort"
)

// Nodes is an array of pointers of type node
type Nodes []*Node

// Node is a MCTS node used to simulate moves in a game
type Node struct {
    children Nodes
    move Direction
    moves []Direction
    parent *Node
    score int
    visits int
}

// NewNode returns an initialized node
func NewNode(move Direction, parent *Node, game *Game) *Node {
    n := new(Node)
    n.children = make(Nodes, 0)
    n.move = move
    n.moves = game.Moves()
    n.parent = parent
    n.score = 0
    n.visits = 0
    return n
}

// AddChild appends a child to children and returns it
func (n *Node) AddChild(move Direction, game *Game) *Node {
    c := NewNode(move, n, game)
    n.RemoveMove(move)
    n.children = append(n.children, c)
    return c
}

// BestChild returns the child with the highest UCT score
func (n *Node) BestChild() *Node {
    sort.Sort(n.children)
    return n.children[len(n.children) - 1]
}

// MostSimulatedChild returns the child with the most visits
func (n *Node) MostSimulatedChild() *Node {
    index := 0
    max := 0

    for i, v := range n.children {
        if v.visits > max {
            max = v.visits
            index = i
        }
    }

    return n.children[index]
}

// RemoveMove removes the specified element
func (n *Node) RemoveMove(move Direction) {
    index := -1

    for i, v := range n.moves {
        if v == move {
            index = i
            break
        }
    }

    lastIndex := len(n.moves) - 1
    n.moves[lastIndex], n.moves[index] = n.moves[index], n.moves[lastIndex]
    n.moves = n.moves[:lastIndex]
}

// UCT returns the calculated UCT score
func (n *Node) UCT() float64 {
    // TODO: Heuristic needs to be tweaked
    return float64(n.score) / float64(n.visits) + math.Sqrt(2 * math.Log(float64(n.parent.visits)) / float64(n.visits))
}

// Update propagates changes
func (n *Node) Update(score int) {
    n.visits += 1
    n.score += score
}

// Len returns the number of elements
func (n Nodes) Len() int {
    return len(n)
}

// Less returns whether the element at i is less than the element at j
func (n Nodes) Less(i, j int) bool {
    return n[i].UCT() < n[j].UCT()
}

// Swap swaps the elements at indexes i and j
func (n Nodes) Swap(i, j int) {
    n[i], n[j] = n[j], n[i]
}

// MCTS runs Monte Carlo tree search from the given state.
// Returns the best move after processing.
func MCTS(game *Game) Direction {
    root := NewNode(NONE, nil, game)
    iterations := 0

    for iterations < 250 {
        node := root
        state := CopyGame(game)

        // Selection
        for len(node.moves) == 0 && len(node.children) > 0 {
            node = node.BestChild()
            state.ApplyMove(node.move)
        }

        // Expansion
        if len(node.moves) > 0 {
            move := RandomMove(node.moves)
            state.ApplyMove(move)
            node = node.AddChild(move, game)
        }

        // Simulation
        for len(state.Moves()) > 0 {
            state.ApplyMove(RandomMove(state.Moves()))
        }

        // Backpropagation
        for node != nil {
            node.Update(state.score)
            node = node.parent
        }

        iterations++
    }

    return root.MostSimulatedChild().move
}

// RandomMove returns a random move
func RandomMove(moves []Direction) Direction {
    index := rand.Intn(len(moves))
    return moves[index]
}

