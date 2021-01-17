package navigation

import (
	"math"
)

type candidate struct {
	node          *node
	distFromStart int
	score         int
}

func newCandidate(n *node) *candidate {
	return &candidate{
		node:          n,
		distFromStart: 0,
		score:         0,
	}
}

func (c candidate) calcScore(prev *candidate, end *node) int {
	return c.guessDistTo(end) + c.calcDFS(prev)
}

func (c candidate) guessDistTo(end *node) int {
	posX := 10*c.node.pos.X + 5*float64(int(c.node.pos.Y)%2)
	endX := 10*end.pos.X + 5*float64(int(end.pos.Y)%2)
	xd := math.Abs(posX - endX)
	yd := math.Abs(10*c.node.pos.Y - 10*end.pos.Y)
	return int(xd + yd)
}

func (c candidate) calcDFS(prev *candidate) int {
	return prev.distFromStart + 10
}

func (c *candidate) update(prev *candidate, score int) {
	c.distFromStart = c.calcDFS(prev)
	c.score = score
}
