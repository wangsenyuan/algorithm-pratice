package main

import "fmt"

func main() {
	words := []string{
		"area", "lead", "wall", "lady", "ball",
	}

	res := wordSquares(words)

	for _, x := range res {
		fmt.Println(x)
	}
}

func wordSquares(words []string) [][]string {
	m := len(words)
	if m == 0 {
		return nil
	}

	trie := &Trie{children: make([]*Trie, 26)}
	for _, word := range words {
		trie.add([]byte(word))
	}

	n := len(words[0])

	var res [][]string
	path := make([]string, n)
	var dfs func(level int)

	dfs = func(level int) {
		if level == n {
			res = append(res, copySlice(path))
			return
		}

		var p []byte
		for i := 0; i < level; i++ {
			p = append(p, path[i][level])
		}
		strs := trie.findRep(p, []byte{})
		for _, x := range strs {
			path[level] = x
			dfs(level + 1)
			//prefix[p] = append([]string{x}, prefix[p]...)
		}
	}

	dfs(0)

	return res
}

func copySlice(slice []string) []string {
	cp := make([]string, len(slice))
	copy(cp, slice)
	return cp
}

type Trie struct {
	children []*Trie
	leaf     bool
}

func (t *Trie) add(bs []byte) {
	if len(bs) == 0 {
		t.leaf = true
		return
	}

	if t.children[bs[0]-'a'] == nil {
		t.children[bs[0]-'a'] = &Trie{children: make([]*Trie, 26)}
	}

	t.children[bs[0]-'a'].add(bs[1:])
}

func (t *Trie) findRep(bs []byte, path []byte) []string {
	if len(bs) == 0 {
		return t.stringRep(path)
	}
	if t.children[bs[0]-'a'] != nil {
		return t.children[bs[0]-'a'].findRep(bs[1:], append(path, bs[0]))
	}
	return nil
}

func (t *Trie) stringRep(path []byte) []string {
	if t.leaf {
		return []string{string(path)}
	}
	var res []string
	for i, c := range t.children {
		if c == nil {
			continue
		}
		path = append(path, byte(i+'a'))
		res = append(res, c.stringRep(path)...)
		path = path[:len(path)-1]
	}
	return res
}
