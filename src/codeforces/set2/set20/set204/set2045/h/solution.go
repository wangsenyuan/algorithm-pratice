package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	s := readString(reader)

	res := solve(s)
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

const inf = 1 << 15

func solve(s string) []string {
	n := len(s)
	lcp := make([][]int16, n+1)
	for i := 0; i <= n; i++ {
		lcp[i] = make([]int16, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if s[i] == s[j] {
				lcp[i][j] = 1 + lcp[i+1][j+1]
			}
		}
	}

	// dp[i][j] = 如果以s[i...j]为第一段时，能够形成的最大段数量
	dp := make([][]int16, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int16, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -(1 << 15)
		}
	}

	dp[n][n] = 0

	for i := n - 1; i >= 0; i-- {
		for j := n; j > i; j-- {
			k := min(int(lcp[i][j]), j-i)
			if j == n || j+k < n && (i+k == j || s[i+k] < s[j+k]) {
				// i + k < j
				dp[i][j] = dp[j][min(j+k+1, n)] + 1
			}
			if j+1 <= n {
				dp[i][j] = max(dp[i][j], dp[i][j+1])
			}
		}
		dp[i][i] = dp[i][i+1]
	}

	var res []string
	best := dp[0][0]

	for i := 0; i < n; {
		for j := n; j > i; j-- {
			k := min(int(lcp[i][j]), j-i)
			if dp[i][j] == best && (j == n || j+k < n && (i+k == j || s[i+k] < s[j+k])) {
				res = append(res, s[i:j])
				i = j
				best--
				break
			}
		}
	}

	return res
}

// this wont work, but my own idea, keep it
func solve1(s string) []string {
	n := len(s)

	tr := NewNode()
	N := n * (n + 1) / 2
	tr.cnt = int32(N)

	for i := 0; i < n; i++ {
		node := tr
		for j := i; j < n; j++ {
			v := int(s[j] - 'A')
			node = node.add(v)
			node.cnt += int32(n - j)
		}
	}
	tr.calcPref()

	roots := make([]*Tree, n+1)

	for i := n - 1; i >= 0; i-- {
		node := tr
		var rk int32

		for j := i; j < n; j++ {
			rk += int32(node.cnt - node.pref[25])

			v := int(s[j] - 'A')
			if v > 0 {
				rk += int32(node.pref[v-1])
			}
			y := roots[j+1].Get(int(rk+1), N-1, 0, N-1)
			if y > 0 || j == n-1 {
				roots[i] = roots[i].Set(int(rk), y+1, 0, N-1)
			}

			node = node.children[v]
		}
	}

	best := roots[0].Get(0, N-1, 0, N-1)
	// 最多这么多段
	var res []string

	var prev int32 = -1

	for i := 0; i < n; {
		var rk int32
		node := tr
		for j := i; j < n; j++ {
			rk += int32(node.cnt - node.pref[25])

			v := int(s[j] - 'A')
			if v > 0 {
				rk += int32(node.pref[v-1])
			}
			if rk > prev {
				tmp := roots[i].Get(int(rk), int(rk), 0, N-1)
				if tmp == best {
					res = append(res, s[i:j+1])
					best--
					i = j + 1
					prev = rk
					break
				}
			}
			node = node.children[v]
		}
	}

	return res
}

type Tree struct {
	left, right *Tree
	val         int32
}

func (tr *Tree) Set(p int, v int32, l int, r int) *Tree {
	if tr == nil {
		tr = new(Tree)
	}
	tr.val = max(tr.val, v)

	if l == r {
		return tr
	}

	mid := (l + r) / 2
	if p <= mid {
		tr.left = tr.left.Set(p, v, l, mid)
	} else {
		tr.right = tr.right.Set(p, v, mid+1, r)
	}
	return tr
}

func (tr *Tree) Get(L int, R int, l int, r int) int32 {
	if tr == nil || R < l || r < L {
		return 0
	}
	if L == l && R == r {
		return tr.val
	}
	var ans int32
	mid := (l + r) / 2
	if L <= mid {
		ans = tr.left.Get(L, min(R, mid), l, mid)
	}
	if mid < R {
		ans = max(ans, tr.right.Get(max(L, mid+1), R, mid+1, r))
	}

	return ans
}

type Node struct {
	children []*Node
	cnt      int32
	pref     []int32
}

func NewNode() *Node {
	node := new(Node)
	node.children = make([]*Node, 26)
	node.pref = make([]int32, 26)
	return node
}
func (node *Node) add(v int) *Node {
	if node.children[v] == nil {
		node.children[v] = NewNode()
	}
	return node.children[v]
}

func (node *Node) calcPref() int32 {
	if node == nil {
		return 0
	}
	for i := 0; i < 26; i++ {
		node.pref[i] = node.children[i].calcPref()
		if i > 0 {
			node.pref[i] += node.pref[i-1]
		}
	}
	return node.cnt
}
