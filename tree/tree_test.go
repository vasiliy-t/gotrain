package tree

import (
	"fmt"
	"testing"
)

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
	n3 := &node{
		Name: "d",
		v:    0,
	}
	n4 := Insert(n0, n3)
	n5 := Insert(n4, n1)
	n6 := Insert(n5, n2)

	if n6.Right().Right() != n2 {
		t.Errorf("Tree composition failed expected %+v not equals actual %+v", n6.Right().Right(), n2)
	}

	if n6.Left() != n3 {
		t.Errorf("Tree composition failed expected %+v not equals actual %+v", n6.Left(), n3)
	}
}

func TestLookup(t *testing.T) {
	node := &node{
		Name: "a",
		v:    1,
		r: &node{
			Name: "b",
			v:    2,
			r: &node{
				Name: "c",
				v:    3,
			},
		},
	}

	act := Lookup(node, 3)

	if act != node.Right().Right() {
		t.Errorf("Failed to assert: expected leaf %+v not equals actual %+v", node.Right().Right(), act)
	}
}

func TestApplyInOrder(t *testing.T) {
	node := &node{
		Name: "a",
		v:    1,
		l: &node{
			Name: "d",
			v:    0,
		},
		r: &node{
			Name: "b",
			v:    2,
			r: &node{
				Name: "c",
				v:    3,
			},
		},
	}
	var res string

	ApplyInOrder(node, func(n NodeInterface) {
		res = fmt.Sprintf("%s%d", res, n.Value())
	})

	if res != "0123" {
		t.Errorf("Failed to assert: expected %s not equals actual %s", "0123", res)
	}
}

func TestApplyPostOrder(t *testing.T) {
	node := &node{
		Name: "a",
		v:    1,
		l: &node{
			Name: "d",
			v:    0,
		},
		r: &node{
			Name: "b",
			v:    2,
			r: &node{
				Name: "c",
				v:    3,
			},
		},
	}

	var res string

	ApplyPostOrder(node, func(n NodeInterface) {
		res = fmt.Sprintf("%s%d", res, n.Value())
	})

	if res != "0321" {
		t.Errorf("Failed to assert: expected %s not equals %s", "0321", res)
	}
}
