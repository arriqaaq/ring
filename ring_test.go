package ring

import (
	"testing"
)

func TestNew(t *testing.T) {
	nodes := []string{"a", "b", "c"}
	hashRing := NewRing(nodes, 1)

	expectNodesABC(t, hashRing)
	// expectNodeRangesABC(t, hashRing)
}

func expectNodesABC(t *testing.T, hashRing *Ring) {
	// Python hash_ring module test case
	expectNode(t, hashRing, "test", "a")
	expectNode(t, hashRing, "test", "a")
	expectNode(t, hashRing, "test1", "c")
	expectNode(t, hashRing, "test2", "c")
	expectNode(t, hashRing, "test3", "c")
	expectNode(t, hashRing, "test4", "c")
	expectNode(t, hashRing, "test5", "b")
	expectNode(t, hashRing, "aaaa", "c")
	expectNode(t, hashRing, "bbbb", "a")
}

func expectNode(t *testing.T, hashRing *Ring, key string, expected string) {
	node, err := hashRing.Get(key)
	if err != nil || node != expected {
		t.Error("GetNode(", key, ") expected", expected, "but got", node, err)
	}
}
