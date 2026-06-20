package main

type AVL struct {
	val    int
	left   *AVL
	right  *AVL
	height int
	bf     int
}

func (avl *AVL) RotRight() *AVL {
	node := avl.left
	avl.left = node.right
	node.right = avl

	return node

}

func (avl *AVL) RotLeft() *AVL {
	node := avl.right
	avl.right = node.left
	node.left = avl

	avl.UpdateProperties()
	node.UpdateProperties()

	return node
}

func (avl *AVL) UpdateProperties() {
	if avl.left == nil && avl.right == nil {
		avl.bf = 0
		avl.height = 0
	}

	hleft := 0
	hright := 0

	if avl.left != nil {
		hleft = avl.left.height
	}
	if avl.right != nil {
		hright = avl.right.height
	}

	if hleft >= hright {
		avl.height = hleft + 1
	} else {
		avl.height = hright + 1
	}
}

func main() {
	println("hello")
}
