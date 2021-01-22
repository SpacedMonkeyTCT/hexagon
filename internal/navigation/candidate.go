package navigation

import (
	"math"
)

type candidate struct {
	node  *Node
	dfs   int
	dte   int
	score int
}

func newCandidate(n *Node) *candidate {
	return &candidate{
		node: n,
		dfs:  0,
		dte:  0,
	}
}

func (c candidate) calcScore(prev *candidate, end *Node) int {
	return c.calcDfs(prev) + c.guessDistTo(end)
}

func (c candidate) calcDfs(prev *candidate) int {
	return prev.dfs + 10
}

func (c candidate) guessDistTo(end *Node) int {
	posX := 10*c.node.X + 5*c.node.Y%2
	endX := 10*end.X + 5*end.Y%2
	xd := math.Abs(float64(posX - endX))
	yd := math.Abs(float64(10*c.node.Y - 10*end.Y))
	return int(xd + yd)
}

func (c *candidate) update(prev *candidate, end *Node) {
	c.dfs = c.calcDfs(prev)
	c.dte = c.guessDistTo(end)
	c.score = c.dfs + c.dte
}
