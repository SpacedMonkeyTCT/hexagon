package navigation

import (
	"github.com/SpacedMonkeyTCT/hexagon/internal/hexmap"
	"github.com/faiface/pixel"
)

type Navigation struct {
	mesh Mesh
}

func NewNavigation(hm *hexmap.HexMap) Navigation {
	return Navigation{
		mesh: newMesh(hm),
	}
}

func (n Navigation) Find(fromX, fromY, toX, toY int) []pixel.Vec {
	start := n.mesh.node[fromX][fromY]
	end := n.mesh.node[toX][toY]
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
