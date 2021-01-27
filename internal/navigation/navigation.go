package navigation

import (
	"github.com/SpacedMonkeyTCT/hexagon/internal/hexmap"
)

type Navigation struct {
	mesh Mesh
}

func NewNavigation(hm *hexmap.HexMap) Navigation {
	return Navigation{
		mesh: newMesh(hm),
	}
}

func (n Navigation) Find(fromX, fromY, toX, toY int) []*Node {
	start := n.mesh.node[fromX][fromY]
	end := n.mesh.node[toX][toY]
	as := newAstar(end, start)

	var p *Node = nil
	for p == nil {
		p = as.search()
	}

	path := make([]*Node, 0)
	for ; p.X != end.X || p.Y != end.Y; p = p.parent {
		path = append(path, p)
	}
	return append(path, end)
}
