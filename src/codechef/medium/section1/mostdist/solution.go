package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func readNNums2(bs []byte, n int) []int {
	res := make([]int, n)
	x := -1
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Q := readNum(scanner)
	A := make([]byte, Q)
	B := make([]int, Q)
	C := make([]int, Q)

	for i := 0; i < Q; i++ {
		scanner.Scan()
		A[i] = scanner.Bytes()[0]
		pos := readInt(scanner.Bytes(), 2, &B[i])
		if A[i] != '-' {
			readInt(scanner.Bytes(), pos+1, &C[i])
		}
	}
	// fmt.Println("Result===")
	res := solve(Q, A, B, C)
	for _, ans := range res {
		fmt.Printf("%d\n", ans)
	}
}

func solve(Q int, A []byte, B []int, C []int) []int {
	var f1, f2, f3, f4 *Node
	res := make([]int, Q)
	var j int

	adds := make([][]int, Q)
	var k int
	var ans int
	for i := 0; i < Q; i++ {
		if A[i] == '+' {
			// Add
			x, y := B[i]^ans, C[i]^ans
			// fmt.Printf("+ %d %d\n", B[i], C[i])

			adds[k] = []int{x, y}
			k++

			f1 = f1.Insert(x + y)
			f2 = f2.Insert(x - y)
			f3 = f3.Insert(-x + y)
			f4 = f4.Insert(-x - y)
		} else if A[i] == '-' {
			N := (B[i] ^ ans)
			// fmt.Printf("- %d\n", B[i])
			x, y := adds[N-1][0], adds[N-1][1]
			f1 = f1.Remove(x + y)
			f2 = f2.Remove(x - y)
			f3 = f3.Remove(-x + y)
			f4 = f4.Remove(-x - y)
		} else {
			// Query
			x, y := B[i]^ans, C[i]^ans
			// fmt.Printf("? %d %d\n", B[i], C[i])

			tmp := f1.GetMax() - (x + y)
			ans = tmp
			tmp = f2.GetMax() - (x - y)
			if tmp > ans {
				ans = tmp
			}

			tmp = f3.GetMax() - (-x + y)
			if tmp > ans {
				ans = tmp
			}

			tmp = f4.GetMax() - (-x - y)
			if tmp > ans {
				ans = tmp
			}

			res[j] = ans
			j++
		}
	}

	return res[:j]
}

type Node struct {
	val      int
	priority int
	left     *Node
	right    *Node
}

func (root *Node) Insert(val int) *Node {
	priority := rand.Int()
	var loop func(node *Node) *Node

	loop = func(node *Node) *Node {
		if node == nil {
			node = new(Node)
			node.val = val
			node.priority = priority
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

func balance(node *Node) *Node {
	if node.left != nil && node.left.priority > node.priority {
		node = rotateRight(node)
	}

	if node.right != nil && node.right.priority > node.priority {
		node = rotateLeft(node)
	}

	return node
}

func rotateRight(node *Node) *Node {
	tmp := node.left
	node.left = tmp.right
	tmp.right = node
	return tmp
}

func rotateLeft(node *Node) *Node {
	tmp := node.right
	node.right = tmp.left
	tmp.left = node
	return tmp
}

func (root *Node) Remove(val int) *Node {
	var loop func(node *Node) *Node

	loop = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		if val < node.val {
			node.left = loop(node.left)
		} else if val > node.val {
			node.right = loop(node.right)
		} else if node.left == nil {
			node = node.right
		} else if node.right == nil {
			node = node.right
		} else if node.left.priority < node.right.priority {
			// current node will go down to the left child of new node
			node = rotateLeft(node)
			node.left = loop(node.left)
		} else {
			node = rotateRight(node)
			node.right = loop(node.right)
		}
		return node
	}

	return loop(root)
}

func (root *Node) GetMax() int {
	if root == nil {
		return 0
	}

	node := root
	for node.right != nil {
		node = node.right
	}

	return node.val
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
