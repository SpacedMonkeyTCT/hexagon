package navigation

type Astar struct {
	start  *Node
	end    *Node
	open   orderedSet
	closed map[*Node]struct{}
}

func newAstar(start, end *Node) Astar {
	return Astar{
		start:  start,
		end:    end,
		open:   newOrderedSet(newCandidate(start)),
		closed: make(map[*Node]struct{}),
	}
}

func (a *Astar) search() *Node {
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
