package p3295

func reportSpam(message []string, bannedWords []string) bool {
	root := new(Trie)
	for _, w := range bannedWords {
		root.AddWord(w)
	}
	var cnt int
	for _, s := range message {
		node := root
		for i := 0; i < len(s) && node != nil; i++ {
			x := int(s[i] - 'a')
			node = node.children[x]
		}
		if node != nil && node.leaf {
			cnt++
		}
		if cnt == 2 {
			return true
		}
	}

	return false
}

type Trie struct {
	children [26]*Trie
	leaf     bool
}

func (root *Trie) AddWord(s string) {
	node := root
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		if node.children[x] == nil {
			node.children[x] = new(Trie)
		}
		node = node.children[x]
	}
	node.leaf = true
}
