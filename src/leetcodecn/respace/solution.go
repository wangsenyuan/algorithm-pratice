package respace

func respace(dictionary []string, sentence string) int {
	mem := make(map[string]bool)

	for _, dict := range dictionary {
		mem[dict] = true
	}
	n := len(sentence)
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = n + 1
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if mem[sentence[j:i]] {
				dp[i] = min(dp[i], dp[j])
			} else {
				dp[i] = min(dp[i], dp[j]+i-j)
			}
		}
	}

	return dp[n]
}

func respace1(dictionary []string, sentence string) int {
	root := NewTrie()

	for _, dict := range dictionary {
		root = root.Add(reverse(dict), 0)
	}

	n := len(sentence)

	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = n + 1
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		cur := root

		for j := i; j >= 0 && cur != nil; j-- {
			dp[i+1] = min(dp[i+1], dp[j]+i-j+1)

			cur = cur.Get(sentence[j])
			if cur != nil && cur.leaf {
				dp[i+1] = min(dp[i+1], dp[j])
			}
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func reverse(s string) string {
	x := []byte(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}

	return string(x)
}

type Trie struct {
	children []*Trie
	leaf     bool
}

func NewTrie() *Trie {
	node := new(Trie)
	node.children = make([]*Trie, 26)
	node.leaf = false
	return node
}

func (node *Trie) Add(s string, i int) *Trie {
	if node == nil {
		node = NewTrie()
	}

	if i == len(s) {
		node.leaf = true
		return node
	}

	x := int(s[i] - 'a')

	node.children[x] = node.children[x].Add(s, i+1)

	return node
}

func (node *Trie) Get(x byte) *Trie {
	return node.children[int(x-'a')]
}
