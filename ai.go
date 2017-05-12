package main

import (
    "math"
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

    n.children = append(n.children, c)
    return c
}

// BestChild returns the child with the highest UCT score
func (n *Node) BestChild() *Node {
    sort.Sort(n.children)
    return n.children[len(n.children) - 1]
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
    return float64(n.score) / float64(n.visits) + math.Sqrt(2 * math.Log(float64(n.parent.visits)) / float64(n.visits))
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
