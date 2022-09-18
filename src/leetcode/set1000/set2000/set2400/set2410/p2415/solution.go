package p2415

func sumPrefixScores(words []string) []int {
	root := new(Trie)

	for _, word := range words {
		root.Add(word)
	}

	ans := make([]int, len(words))

	for i, word := range words {
		ans[i] = root.Query(word, 0)
	}

	return ans
}

type Trie struct {
	children [26]*Trie
	cnt      int
}

func (node *Trie) Add(s string) {
	node.cnt++
	if len(s) == 0 {
		return
	}
	x := int(s[0] - 'a')
	if node.children[x] == nil {
		node.children[x] = new(Trie)
	}
	node.children[x].Add(s[1:])
}

func (node *Trie) Query(s string, sum int) int {
	if len(s) == 0 {
		return sum
	}
	x := int(s[0] - 'a')
	if node.children[x] != nil {
		return node.children[x].Query(s[1:], sum+node.children[x].cnt)
	}
	return sum
}
