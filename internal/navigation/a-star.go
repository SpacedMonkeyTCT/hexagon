package navigation

type Astar struct {
	start  *node
	end    *node
	open   orderedSet
	closed map[*node]struct{}
}

func newAstar(start, end *node) Astar {
	return Astar{
		start:  start,
		end:    end,
		open:   newOrderedSet(newCandidate(start)),
		closed: make(map[*node]struct{}),
	}
}

func (a *Astar) search() *node {
	candidate := a.open.pop()

	if candidate.node == a.end {
		return candidate.node
	}
	a.closed[candidate.node] = struct{}{}

	for _, neighbour := range candidate.node.neighbours {
		if _, closed := a.closed[neighbour]; closed {
			continue
		}
		nextCandidate := a.open.get(neighbour)
		if nextCandidate == nil {
			nextCandidate = newCandidate(neighbour)
		}

		score := nextCandidate.calcScore(candidate, a.end)
		if score < nextCandidate.score || !a.open.includes(nextCandidate) {
			nextCandidate.update(candidate, a.end)
			neighbour.setParent(candidate.node)

			if !a.open.includes(nextCandidate) {
				a.open.push(nextCandidate)
			}
		}
	}
	a.open.sort()
	return nil
}
