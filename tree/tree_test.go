package tree

import "testing"

type node struct {
	Name string
	v    int
	l    NodeInterface
	r    NodeInterface
}

func (n *node) Left() NodeInterface {
	return n.l
}

func (n *node) Right() NodeInterface {
	return n.r
}

func (n *node) SetLeft(ln NodeInterface) {
	n.l = ln
}

func (n *node) SetRight(rn NodeInterface) {
	n.r = rn
}

func (n *node) Value() int {
	return n.v
}

func TestInsert(t *testing.T) {
	n0 := &node{
		Name: "a",
		v:    1,
	}

	n1 := &node{
		Name: "b",
		v:    2,
	}

	n2 := &node{
		Name: "c",
		v:    3,
	}
	n3 := Insert(n0, n1)
	n4 := Insert(n3, n2)

	if n4.Right().Right() != n2 {
		t.Errorf("Failed to assert")
	}
}
