package navigation

import (
	"sort"
)

type orderedSet struct {
	array byScore
	dict  map[*node]*candidate
}

type byScore []*candidate

func (bs byScore) Len() int {
	return len(bs)
}

func (bs byScore) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func (bs byScore) Less(i, j int) bool {
	return bs[j].score < bs[i].score
}

func newOrderedSet(cs ...*candidate) orderedSet {
	dict := make(map[*node]*candidate, len(cs))
	for _, c := range cs {
		dict[c.node] = c
	}
	return orderedSet{
		array: cs,
		dict:  dict,
	}
}

func (os *orderedSet) push(c *candidate) {
	os.array = append(os.array, c)
	os.dict[c.node] = c
}

func (os *orderedSet) pop() *candidate {
	c := os.array[len(os.array)-1]
	os.array = os.array[:len(os.array)-1]
	delete(os.dict, c.node)
	return c
}

func (os orderedSet) get(n *node) *candidate {
	return os.dict[n]
}

func (os orderedSet) includes(c *candidate) bool {
	_, exists := os.dict[c.node]
	return exists
}

func (os orderedSet) sort() {
	sort.Sort(os.array)
}
