package navigation

import (
	"sort"
)

type orderedSet struct {
	array byScore
	dict  map[*Node]struct{}
}

type byScore []*Node

func (bs byScore) Len() int {
	return len(bs)
}

func (bs byScore) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func (bs byScore) Less(i, j int) bool {
	return bs[j].score < bs[i].score
}

func newOrderedSet(ns ...*Node) orderedSet {
	dict := make(map[*Node]struct{}, len(ns))
	for _, n := range ns {
		dict[n] = struct{}{}
	}
	return orderedSet{
		array: ns,
		dict:  dict,
	}
}

func (os *orderedSet) push(n *Node) {
	os.array = append(os.array, n)
	os.dict[n] = struct{}{}
}

func (os *orderedSet) pop() *Node {
	n := os.array[len(os.array)-1]
	os.array = os.array[:len(os.array)-1]
	delete(os.dict, n)
	return n
}

func (os orderedSet) includes(n *Node) bool {
	_, exists := os.dict[n]
	return exists
}

func (os orderedSet) sort() {
	sort.Sort(os.array)
}
