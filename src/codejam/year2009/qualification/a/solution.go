package main

import (
	"fmt"
	"os"
)

func main() {
	var L, D, N int
	fmt.Scanf("%d %d %d", &L, &D, &N)
	words := make([]string, D)
	for i := 0; i < D; i++ {
		fmt.Scanf("%s", &words[i])
	}
	puzzles := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &puzzles[i])
	}
	res := solve(L, D, N, words, puzzles)
	for i := 0; i < N; i++ {
		fmt.Printf("Case #%d: %d\n", i+1, res[i])
	}
}

func solve(L int, D int, N int, words []string, puzzles []string) []int {
	fmt.Fprintf(os.Stderr, "[debug]%d %d %d %v %v\n", L, D, N, words, puzzles)
	var root *Trie

	for i := 0; i < D; i++ {
		root = root.AddWord(words[i])
	}
	res := make([]int, N)
	for i := 0; i < N; i++ {
		res[i] = root.Solve(puzzles[i])
	}

	return res
}

type Trie struct {
	leaf     bool
	children []*Trie
}

func (root *Trie) AddWord(word string) *Trie {
	n := len(word)
	var loop func(i int, node *Trie) *Trie

	loop = func(i int, node *Trie) *Trie {
		if node == nil {
			node = new(Trie)
			node.children = make([]*Trie, 26)
		}

		if i == n {
			node.leaf = true
			return node
		}
		x := int(word[i] - 'a')
		node.children[x] = loop(i+1, node.children[x])
		return node
	}

	return loop(0, root)
}

func (root *Trie) Solve(puzzle string) int {
	var res int

	var loop func(node *Trie, i int)

	loop = func(node *Trie, i int) {
		if i == len(puzzle) {
			if node.leaf {
				res++
			}
			return
		} else if node.leaf {
			return
		}

		if puzzle[i] != '(' {
			x := int(puzzle[i] - 'a')
			if node.children[x] != nil {
				loop(node.children[x], i+1)
			}
			return
		}
		var j = i + 1
		for puzzle[j] != ')' {
			j++
		}
		for k := i + 1; k < j; k++ {
			x := int(puzzle[k] - 'a')
			if node.children[x] != nil {
				loop(node.children[x], j+1)
			}
		}
	}

	loop(root, 0)

	return res
}
