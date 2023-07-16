package p2778

func longestValidSubstring(word string, forbidden []string) int {
	t := NewTrie()

	for i, w := range forbidden {
		t.Add(w, i)
	}

	n := len(word)
	var best int
	x := n
	for i := n - 1; i >= 0; i-- {
		// 考虑以i开始的最长的子串的长度
		// 考察字符长[i...j] 如果[k..j] 时某个forbidden的前缀
		// 假设[k...x]是forbidden串
		// 那么 dp[i] = x - i
		// 对于任何i来说，就是后面最小的x

		cur := t
		j := i
		for ; j < n && cur != nil && cur.id < 0; j++ {
			c := int(word[j] - 'a')
			if cur.children[c] == nil {
				break
			}
			cur = cur.children[c]
		}
		if cur.id >= 0 {
			x = min(x, j-1)
		}
		tmp := x - i
		if best < tmp {
			best = tmp
		}
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Trie struct {
	children [26]*Trie
	id       int
}

func NewTrie() *Trie {
	t := new(Trie)
	t.id = -1
	return t
}

func (t *Trie) Add(s string, id int) {
	if len(s) == 0 {
		t.id = id
		return
	}
	x := int(s[0] - 'a')
	if t.children[x] == nil {
		t.children[x] = NewTrie()
	}
	t.children[x].Add(s[1:], id)
}
