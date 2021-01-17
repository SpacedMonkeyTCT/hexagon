package navigation

import "github.com/SpacedMonkeyTCT/hexagon/internal/hexmap"

type Mesh struct {
	node [][]*node
}

func newMesh(hm *hexmap.HexMap) Mesh {
	m := Mesh{}

	m.node = make([][]*node, hm.MapWidth)
	for c := 0; c < hm.MapWidth; c++ {
		m.node[c] = make([]*node, hm.MapHeight)
		for r := 0; r < hm.MapHeight; r++ {
			m.node[c][r] = newNode(c, r)
		}
	}

	for c := 0; c < hm.MapWidth; c++ {
		for r := 0; r < hm.MapHeight; r++ {
			if !hm.IsWall(c, r) {
				m.joinNeighbours(hm, c, r)
			}
		}
	}
	return m
}

func (m Mesh) joinNeighbours(hm *hexmap.HexMap, c, r int) {
	evenLine := r%2 == 0
	// bottom neighbours
	if r > 0 {
		m.addNeighbourIfNotWall(hm, m.node[c][r], c, r-1)

		if evenLine {
			if c > 0 {
				m.addNeighbourIfNotWall(hm, m.node[c][r], c-1, r-1)
			}
		} else {
			if c < hm.MapWidth-1 {
				m.addNeighbourIfNotWall(hm, m.node[c][r], c+1, r-1)
			}
		}
	}

	// middle neighbours
	if c > 0 {
		m.addNeighbourIfNotWall(hm, m.node[c][r], c-1, r)
	}

	if c < hm.MapWidth-1 {
		m.addNeighbourIfNotWall(hm, m.node[c][r], c+1, r)
	}

	// top neighbours
	if r < hm.MapHeight-1 {
		m.addNeighbourIfNotWall(hm, m.node[c][r], c, r+1)

		if evenLine {
			if c > 0 {
				m.addNeighbourIfNotWall(hm, m.node[c][r], c-1, r+1)
			}
		} else {
			if c < hm.MapWidth-1 {
				m.addNeighbourIfNotWall(hm, m.node[c][r], c+1, r+1)
			}
		}
	}
}

func (m Mesh) addNeighbourIfNotWall(hm *hexmap.HexMap, n *node, c, r int) {
	if !hm.IsWall(c, r) {
		n.addNeighbour(m.node[c][r])
	}
}
