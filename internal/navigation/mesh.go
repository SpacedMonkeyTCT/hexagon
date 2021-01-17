package navigation

import "github.com/SpacedMonkeyTCT/hexagon/internal/hexmap"

type Mesh struct {
	cols int
	rows int
	node [][]*node
}

func newMesh(hm *hexmap.HexMap) Mesh {
	nm := Mesh{
		cols: hm.MapWidth,
		rows: hm.MapHeight,
	}

	nm.node = make([][]*node, nm.cols)
	for c := 0; c < nm.cols; c++ {
		nm.node[c] = make([]*node, nm.rows)
		for r := 0; r < nm.rows; r++ {
			nm.node[c][r] = newNode(c, r)
		}
	}

	for c := 0; c < nm.cols; c++ {
		for r := 0; r < nm.rows; r++ {
			nm.joinNeighbours(c, r)
		}
	}
	return nm
}

func (m Mesh) setWall(c, r int) {
	m.node[c][r].setWall()
}

func (m Mesh) isWall(c, r int) bool {
	return m.node[c][r].isWall
}

func (m Mesh) joinNeighbours(c, r int) {
	evenLine := r%2 == 0
	// bottom neighbours
	if r > 0 {
		m.node[c][r].neighbours = append(m.node[c][r].neighbours, m.node[c][r-1])

		if evenLine {
			if c > 0 {
				m.node[c][r].neighbours = append(m.node[c][r].neighbours, m.node[c-1][r-1])
			}
		} else {
			if c < m.cols-1 {
				m.node[c][r].neighbours = append(m.node[c][r].neighbours, m.node[c+1][r-1])
			}
		}
	}

	// middle neighbours
	if c > 0 {
		m.node[c][r].neighbours = append(m.node[c][r].neighbours, m.node[c-1][r])
	}

	if c < m.cols-1 {
		m.node[c][r].neighbours = append(m.node[c][r].neighbours, m.node[c+1][r])
	}

	// top neighbours
	if r < m.rows-1 {
		m.node[c][r].neighbours = append(m.node[c][r].neighbours, m.node[c][r+1])

		if evenLine {
			if c > 0 {
				m.node[c][r].neighbours = append(m.node[c][r].neighbours, m.node[c-1][r+1])
			}
		} else {
			if c < m.cols-1 {
				m.node[c][r].neighbours = append(m.node[c][r].neighbours, m.node[c+1][r+1])
			}
		}
	}
}