package navigation

import (
	"math/rand"
	"testing"
)

func TestPush(t *testing.T) {
	ons := newOrderedSet()
	ons.push(newNode(1, 2))

	if len(ons.array) != 1 {
		t.Error("Push did not add to array")
	}

	if len(ons.dict) != 1 {
		t.Error("Push did not add to dict")
	}

	if ons.array[0].X != 1 || ons.array[0].Y != 2 {
		t.Error("Pushed item not as expected")
	}
}

func TestPop(t *testing.T) {
	ons := newOrderedSet()
	ons.push(newNode(1, 2))
	n := ons.pop()

	if len(ons.array) != 0 {
		t.Error("Pop did not remove from array")
	}

	if len(ons.dict) != 0 {
		t.Error("Pop did not remove from dict")
	}

	if n.X != 1 || n.Y != 2 {
		t.Error("Popped item not as expected")
	}
}

func TestIncludes(t *testing.T) {
	ons := newOrderedSet()
	n := newNode(1, 2)
	ons.push(n)

	if !ons.includes(n) {
		t.Error("Includes didn't find added candidate")
	}
}

func TestSort(t *testing.T) {
	ons := newOrderedSet()
	for i := 0; i < 10; i++ {
		n := newNode(1, 2)
		n.score = rand.Intn(100)
		ons.push(n)
	}
	ons.sort()

	prevScore := ons.array[0].score
	for _, n := range ons.array[1:] {
		if n.score > prevScore {
			t.Errorf("Sort order not as expected, %v after %v", n.score, prevScore)
		}
		prevScore = n.score
	}
}
