package p1032

type StreamChecker struct {
	root  *Trie
	nodes []*Trie
}

func Constructor(words []string) StreamChecker {
	root := new(Trie)
	root.children = make([]*Trie, 26)
	for _, word := range words {
		root.add(word)
	}
	var nodes []*Trie
	return StreamChecker{root: root, nodes: nodes}
}

func (this *StreamChecker) Query(letter byte) bool {
	var found bool
	nodes := this.nodes
	ns := make([]*Trie, 0, len(nodes))
	for _, node := range nodes {
		node = node.query(letter)
		if node != nil {
			ns = append(ns, node)
			found = found || node.leaf
		}
	}
	next := this.root.query(letter)

	if next != nil {
		ns = append(ns, next)
		found = found || next.leaf
	}

	this.nodes = ns

	return found
}

type Trie struct {
	children []*Trie
	leaf     bool
}

func (node *Trie) add(word string) {
	if len(word) != 0 {
		i := int(word[0] - 'a')
		if node.children[i] == nil {
			child := new(Trie)
			child.children = make([]*Trie, 26)
			node.children[i] = child
		}
		node.children[i].add(word[1:])
	} else {
		node.leaf = true
	}
}

func (node *Trie) query(letter byte) *Trie {
	i := int(letter - 'a')
	return node.children[i]
}
