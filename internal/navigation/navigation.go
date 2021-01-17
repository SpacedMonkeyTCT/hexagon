package navigation

import "github.com/faiface/pixel"

type Navigation struct {
	hexmap HexMap
}

func NewNavigation(cols, rows int) Navigation {
	return Navigation{
		hexmap: newHexMap(cols, rows),
	}
}

func (n Navigation) SetWall(c, r int) {
	n.hexmap.setWall(c, r)
}

func (n Navigation) IsWall(c, r int) bool {
	return n.hexmap.isWall(c, r)
}

func (n Navigation) Find(fromX, fromY, toX, toY int) []pixel.Vec {
	start := n.hexmap.node[fromX][fromY]
	end := n.hexmap.node[toX][toY]
	as := newAstar(start, end)

	var p *node = nil
	for p == nil {
		p = as.search()
	}

	path := make([]pixel.Vec, 0)
	for ; !p.pos.Eq(start.pos); p = p.parent {
		path = append(path, p.pos)
	}
	return path
}
