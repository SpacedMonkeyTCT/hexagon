package navigation

import (
	"github.com/SpacedMonkeyTCT/hexagon/internal/hexmap"
)

type Navigation struct {
	m mesh
}

func NewNavigation(hm *hexmap.HexMap) Navigation {
	return Navigation{
		m: newMesh(hm),
	}
}

func (n Navigation) Find(fromX, fromY, toX, toY int) []*Node {
	start := n.m.node[fromX][fromY]
	end := n.m.node[toX][toY]
	as := newAstar(end, start)

	var p *Node = nil
	for i := 0; p == nil && i < len(n.m.node[0]) * len(n.m.node); i++ {
		p = as.search()
	}

	if p == nil {
		return nil
	}

	path := make([]*Node, 0)
	for ; p.X != end.X || p.Y != end.Y; p = p.parent {
		path = append(path, p)
	}
	return append(path, end)
}
