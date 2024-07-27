package binary_search_tree

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

type BinarySearchTree struct {
	root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Insert(data int) {
	bst.root = insert(bst.root, data)
}

func (bst *BinarySearchTree) Search(data int) bool {
	return search(bst.root, data) != nil
}

func insert(node *Node, data int) *Node {
	if node == nil {
		return NewNode(data)
	}

	if data < node.data {
		node.left = insert(node.left, data)
	} else {
		node.right = insert(node.right, data)
	}

	return node
}

func search(node *Node, data int) *int {
	if node == nil {
		return nil
	}

	if data == node.data {
		return &node.data
	} else if data < node.data {
		return search(node.left, data)
	} else {
		return search(node.right, data)
	}
}
