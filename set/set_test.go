package set_test

import (
	"math/rand"
	"testing"

	"github.com/wdlea/SaveSystem/set"
)

func TestSet(t *testing.T) {
	s := set.MakeSet[Item](1024)

	i := Item(rand.Int63())

	s.Push(i)
	if !s.Has(i) {
		t.Fatalf("Push, set did not have item i in it")
	}
	if !s.Pop(i) {
		t.Fatalf("Pop, set did not have item i in it")
	}
	if s.Has(i) {
		t.Fatalf("Set contains deleted element")
	}
}

type Item uint64

func (i Item) Hash(size uint64) uint64 {
	return uint64(i % Item(size))
}
