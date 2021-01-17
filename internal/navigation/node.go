package navigation

import (
	"github.com/faiface/pixel"
)

type node struct {
	pos        pixel.Vec
	isWall     bool
	neighbours []*node
	parent     *node
}

func newNode(x, y int) *node {
	return &node{
		pos: pixel.V(float64(x), float64(y)),
	}
}

func (n *node) setWall() {
	n.isWall = true
}

func (n *node) clearWall() {
	n.isWall = false
}

func (n *node) addNeighbour(neighbour *node) {
	n.neighbours = append(n.neighbours, neighbour)
}

func (n *node) setParent(parent *node) {
	n.parent = parent
}
