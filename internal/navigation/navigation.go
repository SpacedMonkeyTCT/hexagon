package navigation

import (
	"github.com/SpacedMonkeyTCT/hexagon/internal/hexmap"
	"github.com/faiface/pixel"
)

type Navigation struct {
	mesh      Mesh
	MapWidth  int
	MapHeight int
}

func NewNavigation(hm *hexmap.HexMap) Navigation {
	return Navigation{
		mesh:      newMesh(hm),
		MapWidth:  hm.MapWidth,
		MapHeight: hm.MapHeight,
	}
}

func (n Navigation) SetWall(c, r int) {
	n.mesh.setWall(c, r)
}

func (n Navigation) IsWall(c, r int) bool {
	return n.mesh.isWall(c, r)
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
