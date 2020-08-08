package main

import "fmt"

func main() {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	fmt.Println(palindromePairs(words))
}

func palindromePairs(words []string) [][]int {
	// xxxw + xxx
	var res [][]int
	find(words, func(i, j int) {
		if len(words[i]) >= len(words[j]) {
			res = append(res, []int{i, j})
		}
	})

	rw := make([]string, len(words))

	for i := 0; i < len(words); i++ {
		rw[i] = reverse(words[i])
	}

	find(rw, func(i, j int) {
		if len(rw[i]) > len(rw[j]) {
			res = append(res, []int{j, i})
		}
	})

	return res
}

func find(words []string, f func(i, j int)) {
	root := new(Node)

	for i, word := range words {
		pal := calcPalindrome(word)
		root.Add(i, word, 0, pal)
	}

	for i, w := range words {
		node := root
		j := len(w) - 1
		for node != nil && j >= 0 {
			node = node.Get(w[j])
			j--
		}
		if node == nil {
			continue
		}
		// j < 0
		for _, k := range node.pals {
			if i == k {
				continue
			}
			f(k, i)
		}
	}
}

func calcPalindrome(s string) []bool {
	n := len(s)
	p := make([]bool, n+1)
	p[n] = true
	for i := n - 1; i >= 0; i-- {
		j, k := i, n-1
		for j <= k && s[j] == s[k] {
			j++
			k--
		}
		p[i] = j > k
	}
	return p
}

type Node struct {
	pals     []int
	children [26]*Node
}

func (node *Node) Add(index int, s string, pos int, pal []bool) {

	// palindrome after pos
	if pal[pos] {
		node.pals = append(node.pals, index)
	}

	if pos == len(s) {
		return
	}

	x := int(s[pos] - 'a')

	if node.children[x] == nil {
		node.children[x] = new(Node)
	}

	node.children[x].Add(index, s, pos+1, pal)
}

func (node *Node) Get(x byte) *Node {
	y := int(x - 'a')
	return node.children[y]
}

func palindromePairs1(words []string) [][]int {
	wordMap := make(map[string]int)
	for i, word := range words {
		wordMap[word] = i
	}

	var ans [][]int
	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			if isPalindrome(word[:j]) {
				if k, ok := wordMap[reverse(word[j:])]; ok && k != i {
					ans = append(ans, []int{k, i})
				}
			}

			if j < len(word) && isPalindrome(word[j:]) {
				if k, ok := wordMap[reverse(word[:j])]; ok && k != i {
					ans = append(ans, []int{i, k})
				}
			}
		}
	}

	return ans
}

func reverse(word string) string {
	rev := make([]byte, len(word))
	n := len(word)
	for i := range word {
		rev[n-i-1] = word[i]
	}
	return string(rev)
}

func isPalindrome(word string) bool {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}
