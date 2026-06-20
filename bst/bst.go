package main

import "math"

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

func (bst *Bst) PreOrder() {

	if bst == nil {
		return
	}

	println(bst.val)
	bst.esquerda.PreOrder()
	bst.direita.PreOrder()

}

func (bst *Bst) InOrder() {

	if bst == nil {
		return
	}
	bst.esquerda.InOrder()
	println(bst.val)
	bst.direita.InOrder()

}

func (bst *Bst) PostOrder() {
	if bst == nil {
		return
	}

	bst.esquerda.PostOrder()
	bst.direita.PostOrder()
	println(bst.val)
}

func (bst *Bst) LevelOrder() {
	queue := make([]*Bst, 0)
	queue = append(queue, bst)

	for len(queue) != 0 {
		curNode := queue[0]

		println(curNode.val)

		if curNode.esquerda != nil {
			queue = append(queue, curNode.esquerda)
		}
		if curNode.direita != nil {
			queue = append(queue, curNode.direita)
		}

		queue = queue[1:]
	}
}

func (bst *Bst) Min() int {
	if bst.esquerda == nil {
		return bst.val
	}

	return bst.esquerda.Min()
}

func (bst *Bst) Max() int {
	if bst.direita == nil {
		return bst.val
	}

	return bst.direita.Max()
}

func (bst *Bst) Size() int {
	if bst == nil {
		return 0
	} else {
		return 1 + bst.direita.Size() + bst.esquerda.Size()
	}

}

func (bst *Bst) Par() int {
	if bst == nil {
		return 0
	} else {
		if bst.val%2 == 0 {
			return 1 + bst.direita.Par() + bst.esquerda.Par()
		}
		return bst.direita.Par() + bst.esquerda.Par()
	}
}

func (bst *Bst) Impar() int {
	if bst == nil {
		return 0
	} else {
		if bst.val%2 == 1 {
			return 1 + bst.direita.Impar() + bst.esquerda.Impar()
		}
		return bst.direita.Impar() + bst.esquerda.Impar()
	}
}

func convertToBalancedBst(v []int, ini int, fim int) *Bst {
	if ini > fim {
		return nil
	}

	meio := (ini + fim) / 2

	novoBst := Bst{val: v[meio]}

	novoBst.direita = convertToBalancedBst(v, meio+1, fim)
	novoBst.esquerda = convertToBalancedBst(v, ini, meio-1)

	return &novoBst
}

func (bst *Bst) Height() int {
	heightLeft := 0
	heightRight := 0

	if bst == nil {
		return 0
	}

	if bst.direita != nil {
		heightRight = 1 + bst.direita.Height()
	}

	if bst.esquerda != nil {
		heightLeft = 1 + bst.esquerda.Height()
	}

	if heightLeft > heightRight {
		return heightLeft
	}
	return heightRight
}

func (bst *Bst) Remove(val int) *Bst {

	// achou o valor
	if bst.val == val {
		if bst.direita == nil && bst.esquerda == nil { // caso1: folha
			return nil
		}
		if bst.direita == nil && bst.esquerda != nil { // caso2: tem 1 filho à esquerda
			return bst.esquerda
		}
		if bst.direita != nil && bst.esquerda == nil { // caso2: tem 1 filho à direita
			return bst.direita
		}
		if bst.direita != nil && bst.esquerda != nil { // caso3: tem 2 filhos
			min := bst.direita.Min()
			bst.val = min
			bst.direita = bst.direita.Remove(min)
			return bst
		}
	} else {
		if bst.val > val {
			bst.esquerda = bst.esquerda.Remove(val)
		} else {
			bst.direita = bst.direita.Remove(val)
		}
	}
	return bst

}

func (bstNode *Bst) IsBst() bool {
	// Iniciamos com os limites teóricos de mínimo e máximo
	return validate(bstNode, math.MinInt, math.MaxInt)
}

// Função auxiliar recursiva que checa os limites de cada nó
func validate(node *Bst, min, max int) bool {
	// Uma árvore vazia (ou folha alcançada) é uma BST válida
	if node == nil {
		return true
	}

	// O valor do nó atual deve estar estritamente dentro do intervalo permitido
	if node.val <= min || node.val >= max {
		return false
	}

	// Para a subárvore esquerda: o valor máximo passa a ser o valor do nó atual
	// Para a subárvore direita: o valor mínimo passa a ser o valor do nó atual
	return validate(node.esquerda, min, node.val) && validate(node.direita, node.esquerda.val, max)
}

func main() {
	// cria a primeira raiz
	bst := &Bst{val: 13}

	// Add(&bst, 13)
	Add(bst, 15)
	Add(bst, 7)
	Add(bst, 6)
	Add(bst, 8)
	Add(bst, 20)
	Add(bst, 14)
	Add(bst, 9)
	Add(bst, 10)

	// bst.InOrder()
	// bst.PreOrder()
	// bst.PostOrder()
	// bst.LevelOrder()

	println("Minimo:", bst.Min())
	println("Maximo:", bst.Max())
	println("Tamanho:", bst.Size())
	println("Pares:", bst.Par())
	println("Impares:", bst.Impar())
	bst = bst.Remove(13)
	println("Altura:", bst.Height())
	bst.LevelOrder()

	// a := make([]int, 10)
	// for i := 0; i < 10; i++ {
	// 	a[i] = i * 10
	// }

	// novoBst := convertToBalancedBst(a, 0, len(a)-1)
	// novoBst.PreOrder()
}
