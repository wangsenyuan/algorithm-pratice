package p2227

import "bytes"

type Encrypter struct {
	keyPos []int
	keys   []byte
	valPos map[string][]int
	vals   []string
	root   *Node
}

type Node struct {
	children []*Node
	leaf     bool
}

func AddWord(root *Node, word string) {
	node := root
	for i := 0; i < len(word); i++ {
		x := int(word[i] - 'a')
		if node.children[x] == nil {
			node.children[x] = new(Node)
			node.children[x].children = make([]*Node, 26)
		}
		node = node.children[x]
	}
	node.leaf = true
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	root := new(Node)
	root.children = make([]*Node, 26)

	for _, w := range dictionary {
		AddWord(root, w)
	}

	keyPos := make([]int, 26)
	for i := 0; i < len(keys); i++ {
		x := int(keys[i] - 'a')
		keyPos[x] = i
	}

	valPos := make(map[string][]int)

	for i := 0; i < len(values); i++ {
		v := values[i]
		if _, ok := valPos[v]; !ok {
			valPos[v] = make([]int, 0, 2)
		}
		valPos[v] = append(valPos[v], i)
	}

	return Encrypter{keyPos, keys, valPos, values, root}
}

func (this *Encrypter) Encrypt(word1 string) string {
	var buf bytes.Buffer

	for i := 0; i < len(word1); i++ {
		x := int(word1[i] - 'a')
		j := this.keyPos[x]
		buf.WriteString(this.vals[j])
	}

	return buf.String()
}

func (this *Encrypter) Decrypt(word2 string) int {
	n := len(word2)

	var dfs func(i int, node *Node)

	var cnt int

	dfs = func(i int, node *Node) {
		if i == n {
			if node.leaf {
				cnt++
			}
			return
		}
		s := word2[i : i+2]
		if js, ok := this.valPos[s]; !ok {
			return
		} else {
			for _, j := range js {
				x := int(this.keys[j] - 'a')
				if node.children[x] != nil {
					dfs(i+2, node.children[x])
				}
			}
		}
	}

	dfs(0, this.root)

	return cnt
}

/**
 * Your Encrypter object will be instantiated and called as such:
 * obj := Constructor(keys, values, dictionary);
 * param_1 := obj.Encrypt(word1);
 * param_2 := obj.Decrypt(word2);
 */
