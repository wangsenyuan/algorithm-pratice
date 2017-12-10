package p745

type Trie struct {
	children []*Trie
	ii       []int
}

func (trie *Trie) add(word string, idx int) {
	trie.ii = append(trie.ii, idx)

	if len(word) == 0 {
		return
	}

	h := word[0] - 'a'

	if trie.children[h] == nil {
		trie.children[h] = &Trie{make([]*Trie, 26), make([]int, 0, 10)}
	}
	trie.children[h].add(word[1:], idx)
}

func (trie Trie) find(word string) []int {
	if len(word) == 0 {
		return trie.ii
	}
	h := word[0] - 'a'
	if trie.children[h] == nil {
		return nil
	}
	return trie.children[h].find(word[1:])
}

type WordFilter struct {
	pre *Trie
	suf *Trie
}

func Constructor(words []string) WordFilter {
	pre := &Trie{make([]*Trie, 26), make([]int, 0, 10)}
	suf := &Trie{make([]*Trie, 26), make([]int, 0, 10)}
	for i, word := range words {
		pre.add(word, i)
		suf.add(reverse(word), i)
	}
	//pre.ii = nil
	//suf.ii = nil
	return WordFilter{pre, suf}
}

func reverse(str string) string {
	res := []byte(str)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}

func (this *WordFilter) F(prefix string, suffix string) int {
	x := this.pre.find(prefix)
	y := this.suf.find(reverse(suffix))
	if len(x) == 0 || len(y) == 0 {
		return -1
	}
	for a, b := len(x)-1, len(y)-1; a >= 0 && b >= 0; {
		if x[a] > y[b] {
			a--
		} else if x[a] < y[b] {
			b--
		} else {
			return x[a]
		}
	}
	return -1
}
