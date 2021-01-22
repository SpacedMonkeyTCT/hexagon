package creature

import (
	"math/rand"

	"github.com/SpacedMonkeyTCT/hexagon/internal/hexagon"
	"github.com/SpacedMonkeyTCT/hexagon/internal/hexmap"
	"github.com/SpacedMonkeyTCT/hexagon/internal/navigation"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

// Creature moves around the hexmap
type Creature struct {
	hex     *hexagon.Hexagon
	hm      *hexmap.HexMap
	n       navigation.Navigation
	x       int
	y       int
	targetX int
	targetY int
	path    []pixel.Vec
	steps   int
	pos     pixel.Vec
	nextPos pixel.Vec
}

const (
	stepsPerTile = 20
)

// New creates a creature that lives on a HexMap
func New(hm *hexmap.HexMap, n navigation.Navigation) *Creature {
	c := &Creature{
		hex:     hexagon.New(hm.Size / 2),
		hm:      hm,
		n:       n,
		steps:   0,
		x:       0,
		y:       hm.MapHeight - 2,
		targetX: hm.MapWidth - 1,
		targetY: hm.MapHeight - 1,
		path:    n.Find(0, hm.MapHeight-2, hm.MapWidth-1, hm.MapHeight-1),
	}
	c.startWalk()
	return c
}

func (c *Creature) Update() {
	if c.step(); c.steps < stepsPerTile {
		return
	}
	// when you reach the target, pick another
	if len(c.path) == 0 {
		c.pickTarget()
	}
	c.startWalk()
}

func (c *Creature) startWalk() {
	thisTile := c.path[len(c.path)-1]
	c.x = int(thisTile.X)
	c.y = int(thisTile.Y)
	sx, sy := c.hm.ToScreen(c.x, c.y)
	c.pos = pixel.V(float64(sx), float64(sy))
	c.path = c.path[:len(c.path)-1]

	if len(c.path) > 0 {
		nextTile := c.path[len(c.path)-1]
		ex, ey := c.hm.ToScreen(int(nextTile.X), int(nextTile.Y))
		c.nextPos = pixel.V(float64(ex), float64(ey))
	}
}

func (c *Creature) step() {
	if c.steps < stepsPerTile {
		c.steps++
		return
	}
	c.steps = 0
}

func (c *Creature) pickTarget() {
	for len(c.path) == 0 {
		c.targetX = rand.Intn(c.hm.MapWidth)
		c.targetY = rand.Intn(c.hm.MapHeight)
		if c.hm.IsWall(c.targetX, c.targetY) {
			continue
		}

		c.path = c.n.Find(c.x, c.y, c.targetX, c.targetY)
		if len(c.path) > 0 {
			return
		}
	}
}

// DrawTo draws the creature to an IMDraw
func (c Creature) DrawTo(imd *imdraw.IMDraw) {
	c.drawPath(imd)
	c.drawTarget(imd)

	// draw creature
	imd.Color = colornames.Cornflowerblue

	if c.steps < stepsPerTile && !c.pos.Eq(c.nextPos) {
		ratio := float64(c.steps) / float64(stepsPerTile)
		pos := pixel.Lerp(c.pos, c.nextPos, ratio)
		c.hex.MoveTo(int(pos.X), int(pos.Y))
	} else {
		c.hex.MoveTo(int(c.pos.X), int(c.pos.Y))
	}
	c.hex.DrawTo(imd)
}

func (c Creature) drawPath(imd *imdraw.IMDraw) {
	imd.Color = colornames.Violet
	for _, step := range c.path {
		sx, sy := c.hm.ToScreen(int(step.X), int(step.Y))
		c.hex.MoveTo(sx, sy)
		c.hex.DrawTo(imd)
	}
}

func (c Creature) drawTarget(imd *imdraw.IMDraw) {
	imd.Color = colornames.Red
	sx, sy := c.hm.ToScreen(c.targetX, c.targetY)
	c.hex.MoveTo(sx, sy)
	c.hex.DrawTo(imd)
}
