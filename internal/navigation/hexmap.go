package navigation

type HexMap struct {
	cols int
	rows int
	node [][]*node
}

func NewHexMap(cols, rows int) HexMap {
	nhm := HexMap{
		cols: cols,
		rows: rows,
	}

	nhm.node = make([][]*node, cols)
	for c := 0; c < cols; c++ {
		nhm.node[c] = make([]*node, rows)
		for r := 0; r < rows; r++ {
			nhm.node[c][r] = newNode(c, r)
		}
	}

	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			nhm.joinNeighbours(c, r)
		}
	}
	return nhm
}

func (hm HexMap) SetWall(c, r int) {
	hm.node[c][r].setWall()
}

func (hm HexMap) joinNeighbours(c, r int) {
	evenLine := r%2 == 0
	// bottom neighbours
	if r > 0 {
		hm.node[c][r].neighbours = append(hm.node[c][r].neighbours, hm.node[c][r-1])

		if evenLine {
			if c > 0 {
				hm.node[c][r].neighbours = append(hm.node[c][r].neighbours, hm.node[c-1][r-1])
			}
		} else {
			if c < hm.cols-1 {
				hm.node[c][r].neighbours = append(hm.node[c][r].neighbours, hm.node[c+1][r-1])
			}
		}
	}

	// middle neighbours
	if c > 0 {
		hm.node[c][r].neighbours = append(hm.node[c][r].neighbours, hm.node[c-1][r])
	}

	if c < hm.cols-1 {
		hm.node[c][r].neighbours = append(hm.node[c][r].neighbours, hm.node[c+1][r])
	}

	// top neighbours
	if r < hm.rows-1 {
		hm.node[c][r].neighbours = append(hm.node[c][r].neighbours, hm.node[c][r+1])

		if evenLine {
			if c > 0 {
				hm.node[c][r].neighbours = append(hm.node[c][r].neighbours, hm.node[c-1][r+1])
			}
		} else {
			if c < hm.cols-1 {
				hm.node[c][r].neighbours = append(hm.node[c][r].neighbours, hm.node[c+1][r+1])
			}
		}
	}
}
