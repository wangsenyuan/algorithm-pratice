package longestword

import "sort"

func longestWord(words []string) string {
	root := NewNode()

	for _, word := range words {
		root.Add(word)
	}

	var dfs func(s string, cnt int) bool

	dfs = func(s string, cnt int) bool {
		if len(s) == 0 {
			return cnt > 1
		}

		tmp := root

		for i := 0; i < len(s); i++ {
			x := int(s[i] - 'a')
			if tmp.children[x] == nil {
				break
			}
			tmp = tmp.children[x]
			if tmp.leaf {
				res := dfs(s[i+1:], cnt+1)
				if res {
					return true
				}
			}
		}
		return false
	}

	sort.Sort(Strings(words))

	for _, word := range words {
		if dfs(word, 0) {
			return word
		}
	}
	return ""
}

type Node struct {
	leaf     bool
	children []*Node
}

func NewNode() *Node {
	node := new(Node)
	node.children = make([]*Node, 26)
	return node
}

func (node *Node) Add(word string) {
	if len(word) == 0 {
		node.leaf = true
		return
	}
	x := int(word[0] - 'a')
	if node.children[x] == nil {
		node.children[x] = NewNode()
	}
	node.children[x].Add(word[1:])
}

type Strings []string

func (this Strings) Len() int {
	return len(this)
}

func (this Strings) Less(i, j int) bool {
	if len(this[i]) > len(this[j]) {
		return true
	}
	if len(this[i]) == len(this[j]) {
		return this[i] < this[j]
	}
	return false
}

func (this Strings) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
