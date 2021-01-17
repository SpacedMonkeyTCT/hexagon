package navigation

import (
	"math/rand"
	"testing"

	"github.com/faiface/pixel"
)

func TestPush(t *testing.T) {
	ons := newOrderedSet()
	ons.push(newCandidate(newNode(1, 2)))

	if len(ons.array) != 1 {
		t.Error("Push did not add to array")
	}

	if len(ons.dict) != 1 {
		t.Error("Push did not add to dict")
	}

	if !ons.array[0].node.pos.Eq(pixel.V(1, 2)) {
		t.Error("Pushed item not as expected")
	}
}

func TestPop(t *testing.T) {
	ons := newOrderedSet()
	ons.push(newCandidate(newNode(1, 2)))
	c := ons.pop()

	if len(ons.array) != 0 {
		t.Error("Pop did not remove from array")
	}

	if len(ons.dict) != 0 {
		t.Error("Pop did not remove from dict")
	}

	if !c.node.pos.Eq(pixel.V(1, 2)) {
		t.Error("Popped item not as expected")
	}
}

func TestIncludes(t *testing.T) {
	ons := newOrderedSet()
	c := newCandidate(newNode(1, 2))
	ons.push(c)

	if !ons.includes(c) {
		t.Error("Includes didn't find added candidate")
	}
}

func TestSort(t *testing.T) {
	ons := newOrderedSet()
	for i := 0; i < 10; i++ {
		c := newCandidate(newNode(1, 2))
		c.score = rand.Intn(100)
		ons.push(c)
	}
	ons.sort()

	prevScore := ons.array[0].score
	for _, c := range ons.array[1:] {
		if c.score < prevScore {
			t.Error("Sort order not as expected")
			return
		}
		prevScore = c.score
	}
}
