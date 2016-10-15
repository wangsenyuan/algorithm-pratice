package main

import "fmt"

func main() {
	//fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
	fmt.Println(findMaximumXOR([]int{10, 23, 20, 18, 28}))
}

func findMaximumXOR(nums []int) int {
	t := &Trie{children: make([]*Trie, 2)}

	for _, num := range nums {
		t.add(num, 31)
	}

	res := 0

	for _, num := range nums {
		at := t
		for i := 31; i >= 0 && at != nil; i-- {
			if (num>>uint(i))&1 == 0 {
				if at.children[1] != nil {
					at = at.children[1]
				} else {
					at = at.children[0]
				}
			} else {
				if at.children[0] != nil {
					at = at.children[0]
				} else {
					at = at.children[1]
				}
			}
		}

		if at != nil && at.num^num > res {
			res = at.num ^ num
		}
	}

	return res
}

type Trie struct {
	children []*Trie
	num      int
}

func (t *Trie) add(num int, i uint) {
	t.num = num
	if i == 0 {
		return
	}
	x := (num >> i) & 1

	if t.children[x] == nil {
		t.children[x] = &Trie{children: make([]*Trie, 2)}
	}

	t.children[x].add(num, i-1)
}
