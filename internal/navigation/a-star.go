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
		open:   newOrderedSet(start),
		closed: make(map[*Node]struct{}),
	}
}

func (a *Astar) search() *Node {
	candidate := a.open.pop()

	if candidate == a.end {
		return candidate
	}
	a.closed[candidate] = struct{}{}

	for _, nextCandidate := range candidate.neighbours {
		if _, closed := a.closed[nextCandidate]; closed {
			continue
		}

		score := nextCandidate.calcScore(candidate, a.end)
		if score < nextCandidate.score || !a.open.includes(nextCandidate) {
			nextCandidate.update(candidate, a.end)
			nextCandidate.setParent(candidate)

			if !a.open.includes(nextCandidate) {
				a.open.push(nextCandidate)
			}
		}
	}
	a.open.sort()
	return nil
}
