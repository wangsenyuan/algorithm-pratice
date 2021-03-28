package p1805

func numDifferentIntegers(word string) int {
	root := new(Trie)
	var i int
	for i < len(word) {
		j := i
		for j < len(word) && word[j] >= '0' && word[j] <= '9' {
			j++
		}

		if j == i {
			i++
		} else {
			AddNum(root, word[i:j])
			i = j
		}
	}

	return CountLeaf(root)
}

type Trie struct {
	children [10]*Trie
	leaf     bool
}

func AddNum(root *Trie, num string) {
	var i int
	for i < len(num) && num[i] == '0' {
		i++
	}
	if i == len(num) {
		i--
	}
	node := root
	for i < len(num) {
		x := int(num[i] - '0')
		if node.children[x] == nil {
			node.children[x] = new(Trie)
		}
		node = node.children[x]
		i++
	}
	node.leaf = true
}

func CountLeaf(node *Trie) int {
	var res int
	for i := 0; i < 10; i++ {
		if node.children[i] != nil {
			res += CountLeaf(node.children[i])
		}
	}
	if node.leaf {
		res++
	}
	return res
}
