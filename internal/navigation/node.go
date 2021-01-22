package navigation

type Node struct {
	X          int
	Y          int
	neighbours []*Node
	parent     *Node
}

func newNode(x, y int) *Node {
	return &Node{
		X: x,
		Y: y,
	}
}

func (n *Node) addNeighbour(neighbour *Node) {
	n.neighbours = append(n.neighbours, neighbour)
}

func (n *Node) setParent(parent *Node) {
	n.parent = parent
}
