package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) string {
	n, m, k := readThreeNums(reader)
	p := readNNums(reader, n)
	a := readNNums(reader, m)
	queries := make([][]int, k)
	for i := range k {
		queries[i] = readNNums(reader, 2)
	}
	return solve(p, a, queries)
}

func solve(p []int, a []int, queries [][]int) string {
	n := len(p)
	pos := make([]int, n+1)
	for i, x := range p {
		pos[x] = i
	}
	m := len(a)

	for i := 0; i < m; i++ {
		a[i] = pos[a[i]]
	}

	nodes := make([]*Node, m+1)

	next := make([]int, m+1)
	last := make([]int, n+1)
	for i := 0; i <= n; i++ {
		last[i] = m
	}

	dp := make([]int, m+1)
	fp := make([]int, m+1)
	dp[m] = m
	next[m] = m

	for i := 0; i <= m; i++ {
		fp[i] = -1
	}

	for i := m - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		v := a[i]
		if v == n-1 {
			fp[i] = i
		} else if last[v+1] > i {
			fp[i] = fp[last[v+1]]
		}
		if fp[i] >= i {
			// j是离i最近的连续递增的n-1的位置
			j := fp[i]
			if v == 0 {
				dp[i] = min(dp[i], j)
			} else if j < m {
				// v > 0
				k := nodes[j+1].Find(v-1, 0, n-1)
				if j < k {
					dp[i] = min(dp[i], k)
				}
			}
		}
		nodes[i] = nodes[i+1]
		next[i] = last[v+1]
		if v == 0 {
			j := i
			node := nodes[i]
			for j < m {
				k := node.Find(a[j], 0, n-1)
				if k == j {
					// 如果j在原来的位置没有变，不需要更新
					break
				}
				node = node.Set(a[j], j, 0, n-1)
				j = next[j]
			}
			nodes[i] = node
		}
		// else 没有从i开始的v, v+1, v +2.... n-1序列
		last[v] = i
	}

	buf := make([]byte, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		l--
		r--
		if dp[l] <= r {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}
	}

	return string(buf)
}

type Node struct {
	l, r        int
	val         int
	left, right *Node
}

func (node *Node) Set(p int, v int, l int, r int) *Node {
	res := new(Node)
	if node != nil {
		res.left = node.left
		res.right = node.right
		res.val = node.val
	} else {
		res.l = l
		res.r = r
		res.val = -1
	}
	if l == r {
		res.val = v
	} else {
		mid := (l + r) / 2
		if p <= mid {
			res.left = res.left.Set(p, v, l, mid)
		} else {
			res.right = res.right.Set(p, v, mid+1, r)
		}
	}

	return res
}

func (node *Node) Find(p int, l int, r int) int {
	if node == nil {
		return -1
	}

	if l == r {
		return node.val
	}

	mid := (l + r) / 2
	if p <= mid {
		return node.left.Find(p, l, mid)
	}
	return node.right.Find(p, mid+1, r)
}
