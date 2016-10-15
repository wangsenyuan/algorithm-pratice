package main

import "fmt"

func main() {
	//fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
	fmt.Println(findMaximumXOR([]int{10, 23, 20, 18, 28}))
}

func findMaximumXOR(nums []int) int {
	t := &Trie{children: make([]*Trie, 2)}

	for _, num := range nums {
		tmp := t
		for i := 31; i >= 0; i-- {
			x := (num >> uint(i)) & 1

			if tmp.children[x] == nil {
				tmp.children[x] = &Trie{children: make([]*Trie, 2)}
			}

			tmp = tmp.children[x]
		}
	}

	res := 0

	for _, num := range nums {
		var tmp = t
		cur := 0
		for i := 31; i >= 0; i-- {
			x := uint(num >> uint(i) & 1)
			if tmp.children[x^1] != nil {
				cur = cur | (1 << uint(i))
				tmp = tmp.children[x^1]
			} else {
				tmp = tmp.children[x]
			}
		}
		if cur > res {
			res = cur
		}
	}

	return res
}

type Trie struct {
	children []*Trie
}
