package navigation

import "github.com/faiface/pixel"

type Navigation struct {
	hexmap HexMap
}

func NewNavigation(cols, rows int) Navigation {
	return Navigation{
		hexmap: NewHexMap(cols, rows),
	}
}

func (n Navigation) Find(fromX, fromY, toX, toY int) []pixel.Vec {
	start := n.hexmap.node[fromX][fromY]
	end := n.hexmap.node[toX][toY]
	as := NewAstar(start, end)

	var p *node = nil
	for p == nil {
		p = as.Search()
	}

	path := make([]pixel.Vec, 0)
	for ; !p.pos.Eq(start.pos); p = p.parent {
		path = append(path, p.pos)
	}
	return path
}
