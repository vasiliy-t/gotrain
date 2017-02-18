package tree

type NodeInterface interface {
	Left() NodeInterface
	Right() NodeInterface
	SetLeft(n NodeInterface)
	SetRight(n NodeInterface)
	Value() int
}

func Insert(tree, node NodeInterface) NodeInterface {
	if tree == nil {
		return node
	}

	if node == nil {
		return tree
	}

	if tree.Value() > node.Value() {
		tree.SetLeft(Insert(tree.Left(), node))
	}

	if tree.Value() < node.Value() {
		tree.SetRight(Insert(tree.Right(), node))
	}

	return tree
}

func Lookup(tree NodeInterface, val int) NodeInterface {
	if tree == nil {
		return nil
	}

	if val > tree.Value() {
		return Lookup(tree.Right(), val)
	}

	if val < tree.Value() {
		return Lookup(tree.Left(), val)
	}

	return tree
}

func ApplyInOrder(tree NodeInterface, fn func(n NodeInterface)) {
	if tree == nil {
		return
	}
	ApplyInOrder(tree.Left(), fn)
	fn(tree)
	ApplyInOrder(tree.Right(), fn)
}
