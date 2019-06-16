package main

import (
	"fmt"
	"math/rand"
)

func main() {
	A := []int{15777, 7355, 6475, 15448, 18412}
	B := []int{986, 13574, 14234, 18412, 19893}
	res := advantageCount(A, B)
	fmt.Printf("%v", res)
}

func advantageCount(A []int, B []int) []int {
	n := len(A)

	var root *Treap

	for i := 0; i < n; i++ {
		root = root.Insert(A[i])
	}

	// sort.Sort(BB)

	res := make([]int, n)

	for i := 0; i < n; i++ {
		b := B[i]
		at := root.Find(b)
		if at == nil {
			at = root.FindMin()
		}
		res[i] = at.val
		root = root.Remove(at)
	}

	return res
}

type Treap struct {
	val         int
	priority    int
	left, right *Treap
}

func (root *Treap) Insert(val int) *Treap {
	priotiry := rand.Intn(100000)
	var loop func(node *Treap) *Treap

	loop = func(node *Treap) *Treap {
		if node == nil {
			node = new(Treap)
			node.val = val
			node.priority = priotiry
			return node
		}
		if val <= node.val {
			node.left = loop(node.left)
		} else {
			node.right = loop(node.right)
		}

		return balance(node)
	}

	return loop(root)
}

func (root *Treap) Find(val int) *Treap {
	var loop func(node *Treap)
	var res *Treap
	loop = func(node *Treap) {
		if node == nil {
			return
		}
		if val < node.val {
			res = node
			loop(node.left)
			return
		}
		loop(node.right)
	}
	loop(root)
	return res
}

func (root *Treap) FindMin() *Treap {
	node := root
	for node.left != nil {
		node = node.left
	}
	return node
}

func (root *Treap) Remove(at *Treap) *Treap {
	var loop func(node *Treap) *Treap
	loop = func(node *Treap) *Treap {
		if at.val < node.val {
			node.left = loop(node.left)
		} else if at.val > node.val {
			node.right = loop(node.right)
		} else if node.left == nil {
			node = node.right
		} else if node.right == nil {
			node = node.left
		} else if node.left.priority > node.right.priority {
			node = rotateRight(node)
			node.right = loop(node.right)
		} else {
			node = rotateLeft(node)
			node.left = loop(node.left)
		}

		return node
	}

	return loop(root)
}

func balance(node *Treap) *Treap {
	if node.left != nil && node.left.priority > node.priority {
		node = rotateRight(node)
	}

	if node.right != nil && node.right.priority > node.priority {
		node = rotateLeft(node)
	}
	return node
}

func rotateRight(node *Treap) *Treap {
	tmp := node.left
	node.left = tmp.right
	tmp.right = node
	return tmp
}

func rotateLeft(node *Treap) *Treap {
	tmp := node.right
	node.right = tmp.left
	tmp.left = node
	return tmp
}
