package navigation

import "math"

type Node struct {
	X          int
	Y          int
	neighbours []*Node
	parent     *Node
	dfs        int
	dte        int
	score      int
}

func newNode(x, y int) *Node {
	return &Node{
		X: x,
		Y: y,
	}
}

func (n *Node) addNeighbour(neighbour *Node) {
	n.neighbours = append(n.neighbours, neighbour)
}

func (n *Node) setParent(parent *Node) {
	n.parent = parent
}

func (n Node) calcScore(prev *Node, end *Node) int {
	return n.calcDfs(prev) + n.guessDistTo(end)
}

func (n Node) calcDfs(prev *Node) int {
	return prev.dfs + 10
}

func (n Node) guessDistTo(end *Node) int {
	posX := 10*n.X + 5*n.Y%2
	endX := 10*end.X + 5*end.Y%2
	xd := math.Abs(float64(posX - endX))
	yd := math.Abs(float64(10*n.Y - 10*end.Y))
	return int(xd + yd)
}

func (n *Node) update(prev *Node, end *Node) {
	n.dfs = n.calcDfs(prev)
	n.dte = n.guessDistTo(end)
	n.score = n.dfs + n.dte
}
