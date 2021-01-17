package navigation

import (
	"math"
)

type candidate struct {
	node  *node
	dfs   int
	dte   int
	score int
}

func newCandidate(n *node) *candidate {
	return &candidate{
		node: n,
		dfs:  0,
		dte:  0,
	}
}

func (c candidate) calcScore(prev *candidate, end *node) int {
	return c.calcDfs(prev) + c.guessDistTo(end)
}

func (c candidate) calcDfs(prev *candidate) int {
	return prev.dfs + 10
}

func (c candidate) guessDistTo(end *node) int {
	posX := 10*c.node.pos.X + 5*float64(int(c.node.pos.Y)%2)
	endX := 10*end.pos.X + 5*float64(int(end.pos.Y)%2)
	xd := math.Abs(posX - endX)
	yd := math.Abs(10*c.node.pos.Y - 10*end.pos.Y)
	return int(xd + yd)
}

func (c *candidate) update(prev *candidate, end *node) {
	c.dfs = c.calcDfs(prev)
	c.dte = c.guessDistTo(end)
	c.score = c.dfs + c.dte
}
