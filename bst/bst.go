package main

type Bst struct {
	val      int
	direita  *Bst
	esquerda *Bst
}

func Add(bst *Bst, val int) {
	if bst == nil {
		bst = &Bst{val: val}
	} else if val > bst.val {
		if bst.direita == nil {
			bst.direita = &Bst{val: val}
		} else {
			Add(bst.direita, val)
		}

	} else {
		if bst.esquerda == nil {
			bst.esquerda = &Bst{val: val}
		} else {
			Add(bst.esquerda, val)
		}
	}
}

func search(bst *Bst, val int) bool {
	if bst == nil {
		return false
	}

	if bst.val == val {
		return true
	}

	if bst.val > val {
		return search(bst.esquerda, val)
	}

	return search(bst.direita, val)

}

// func (bst *Bst) PreOrder(val int) {}

func main() {
	var bst Bst

	Add(&bst, 13)
	Add(&bst, 15)
	Add(&bst, 7)
	Add(&bst, 6)
	Add(&bst, 8)
	Add(&bst, 20)
	Add(&bst, 14)
	println(search(&bst, 13))
}
