package navigation

import "math"

type Astar struct {
	start  *Node
	end    *Node
	open   map[*Node]struct{}
	closed map[*Node]struct{}
}

func newAstar(start, end *Node) Astar {
	open := make(map[*Node]struct{})
	open[start] = struct{}{}
	return Astar{
		start:  start,
		end:    end,
		open:   open,
		closed: make(map[*Node]struct{}),
	}
}

func (a *Astar) search() *Node {
	candidate := popLowestScore(a.open)

	if candidate == a.end {
		return candidate
	}
	a.closed[candidate] = struct{}{}

	for _, nextCandidate := range candidate.neighbours {
		if _, closed := a.closed[nextCandidate]; closed {
			continue
		}

		score := nextCandidate.calcScore(candidate, a.end)
		if _, open := a.open[nextCandidate]; !open || score < nextCandidate.score {
			nextCandidate.update(candidate, a.end)
			nextCandidate.setParent(candidate)

			if !open {
				a.open[nextCandidate] = struct{}{}
			}
		}
	}
	return nil
}

func popLowestScore(m map[*Node]struct{}) *Node {
	lowest := &Node{score: math.MaxInt32}
	for n := range m {
		if n.score < lowest.score {
			lowest = n
		}
	}
	delete(m, lowest)
	return lowest
}
