package main

import (
	"fmt"
	"bufio"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)

	S := make([][]byte, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		S[i] = scanner.Bytes()
	}

	Q := readNum(scanner)

	R := make([]int, Q)
	P := make([][]byte, Q)

	for i := 0; i < Q; i++ {
		scanner.Scan()
		pos := readInt(scanner.Bytes(), 0, &R[i])
		P[i] = scanner.Bytes()[pos+1:]
	}

	res := solve(n, S, Q, R, P)
	for i := 0; i < Q; i++ {
		fmt.Println(res[i])
	}
}

func solve(n int, S [][]byte, Q int, R []int, P [][]byte) []string {
	nodes := make([]*TrieNode, n)

	nodes[0] = Insert(nil, S[0])
	for i := 1; i < n; i++ {
		nodes[i] = Insert(nodes[i-1], S[i])
		//debug(nodes[i])
	}

	res := make([]string, Q)
	for i := 0; i < Q; i++ {
		r := R[i] - 1
		res[i] = Query(nodes[r], P[i])
	}
	return res
}

type TrieNode struct {
	children []*TrieNode
	leaf     bool
}

func copyTree(src *TrieNode) *TrieNode {
	dst := new(TrieNode)
	dst.children = make([]*TrieNode, 26)
	dst.leaf = src.leaf
	copy(dst.children, src.children)
	return dst
}

func Insert(src *TrieNode, S []byte) *TrieNode {
	if src != nil {
		src = copyTree(src)
	} else {
		src = new(TrieNode)
		src.children = make([]*TrieNode, 26)
	}
	if len(S) == 0 {
		src.leaf = true
		return src
	}
	j := int(S[0] - 'a')
	src.children[j] = Insert(src.children[j], S[1:])
	return src
}

func Query(root *TrieNode, P []byte) string {
	buf := make([]byte, 20)

	var loop func(node *TrieNode, i int, level int) string

	loop = func(node *TrieNode, i int, level int) string {
		if node.leaf && i == len(P) {
			return string(buf[:level])
		}
		if i < len(P) {
			j := int(P[i] - 'a')
			if node.children[j] != nil {
				buf[level] = P[i]
				return loop(node.children[j], i+1, level+1)
			}
		}

		// not a prefix any more, just find the first valid child
		x := -1
		for i := 0; i < 26; i++ {
			if node.children[i] != nil {
				x = i
				break
			}
		}
		if x < 0 {
			return string(buf[:level])
		}
		buf[level] = byte(x + 'a')

		return loop(node.children[x], len(P), level+1)
	}

	return loop(root, 0, 0)
}

func debug(root *TrieNode) {
	fmt.Println("debug----")
	var dfs func(node *TrieNode, level int)

	buf := make([]byte, 11)

	dfs = func(node *TrieNode, level int) {
		if node.leaf {
			fmt.Printf("%s\n", buf[:level])
		}

		for i := 0; i < 26; i++ {
			if node.children[i] != nil {
				buf[level] = byte('a' + i)
				dfs(node.children[i], level+1)
			}
		}
	}

	dfs(root, 0)
}
