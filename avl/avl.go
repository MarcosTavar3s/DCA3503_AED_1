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

	avl.UpdateProperties()
	node.UpdateProperties()
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
		return
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

	avl.bf = hright - hleft
}

func (avl *AVL) RebalanceLeftLeft() *AVL {
	return avl.RotRight()
}

func (avl *AVL) RebalanceLeftNeutral() *AVL {
	return avl.RotRight()
}

func (avl *AVL) RebalanceLeftRight() *AVL {
	avl.left = avl.left.RotLeft()
	return avl.RotRight()
}

func (avl *AVL) RebalanceRightRight() *AVL {
	return avl.RotLeft()
}

func (avl *AVL) RebalanceRightNeutral() *AVL {
	return avl.RotLeft()
}

func (avl *AVL) RebalanceRightLeft() *AVL {
	avl.right = avl.right.RotRight()
	return avl.RotLeft()
}

func (avl *AVL) Rebalance() *AVL {
	if avl.bf == -2 {
		if avl.left.bf == -1 {
			avl = avl.RebalanceLeftLeft()
		} else if avl.left.bf == 1 {
			avl = avl.RebalanceLeftRight()
		} else {
			avl = avl.RebalanceLeftNeutral()
		}

	}
	if avl.bf == 2 {
		if avl.right.bf == 1 {
			avl = avl.RebalanceRightRight()
		} else if avl.right.bf == -1 {
			avl = avl.RebalanceRightLeft()
		} else {
			avl = avl.RebalanceRightNeutral()
		}

	}
	return avl
}

func (avl *AVL) Min() int {
	if avl.left == nil {
		return avl.val
	}

	return avl.left.Min()
}

func (avl *AVL) Max() int {
	if avl.right == nil {
		return avl.val
	}

	return avl.right.Max()
}

func (avl *AVL) Add(val int) *AVL {
	if avl == nil {
		return &AVL{val: val}
	}

	if val <= avl.val {
		avl.left = avl.left.Add(val)
	} else {
		avl.right = avl.right.Add(val)
	}

	avl.UpdateProperties()
	return avl.Rebalance()
}

func (avl *AVL) Remove(val int) *AVL {

	if avl.val == val {
		if avl.left == nil && avl.right == nil { // no folha
			return nil
		}
		if avl.left == nil && avl.right != nil {
			return avl.right
		}
		if avl.left != nil && avl.right == nil {
			return avl.left
		}
		if avl.left != nil && avl.right != nil {
			min := avl.right.Min()
			avl.val = min
			avl = avl.Remove(min)

		}
	} else {
		if avl.val < val {
			avl = avl.left.Remove(val)
		} else {
			avl = avl.right.Remove(val)
		}
	}

	avl.UpdateProperties()
	return avl.Rebalance()
}

func (avl *AVL) LevelOrder() {
	queue := make([]*AVL, 0)
	queue = append(queue, avl)

	for len(queue) != 0 {
		curNode := queue[0]

		print(curNode.val)
		print(" ")

		if curNode.left != nil {
			queue = append(queue, curNode.left)
		}
		if curNode.right != nil {
			queue = append(queue, curNode.right)
		}

		queue = queue[1:]
	}
}

func main() {
	avl := &AVL{val: 1}
	avl = avl.Add(2)
	print("Level: ")
	avl.LevelOrder()
	avl = avl.Add(3)
	print("\nLevel: ")
	avl.LevelOrder()
	avl = avl.Add(4)
	print("\nLevel: ")
	avl.LevelOrder()
	avl = avl.Add(5)
	print("\nLevel: ")
	avl.LevelOrder()
	avl = avl.Add(6)
	print("\nLevel: ")
	avl.LevelOrder()
	avl = avl.Add(7)
	print("\nLevel: ")
	avl.LevelOrder()
	avl = avl.Add(8)
	print("\nLevel: ")
	avl.LevelOrder()
	avl = avl.Add(9)
	print("\nLevel: ")
	avl.LevelOrder()
}
