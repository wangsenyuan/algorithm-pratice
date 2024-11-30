package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}

	fmt.Print(buf.String())
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

const inf = 1 << 60

func solve(a []int) [][]int {
	// 需要知道最大值和最小值
	var b []int32
	n := len(a)
	for i := 0; i < n; i++ {
		var sum int32
		for j := i; j < n; j++ {
			sum += int32(a[j])
			b = append(b, sum)
		}
	}
	slices.SortFunc(b, func(x, y int32) int {
		return int(x - y)
	})

	var m int

	for i := 1; i <= len(b); i++ {
		if i == len(b) || b[i] > b[i-1] {
			b[m] = b[i-1]
			m++
		}
	}
	b = b[:m]

	findPos := func(v int32) int32 {
		ans := sort.Search(m, func(i int) bool {
			return b[i] >= v
		})
		return int32(ans)
	}

	dp := make([]*Node, n+1)

	for i := n - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		var sum int32
		for j := i; j < n; j++ {
			sum += int32(a[j])
			p := findPos(sum)
			v := dp[j+1].Get(p, 0, int32(m-1))
			if dp[i].Get(p, 0, int32(m-1)) <= v {
				dp[i] = dp[i].Set(p, v+1, 0, int32(m-1))
			}
		}
	}
	// 也就是说最多有这么多的区间
	mv := dp[0].val

	node := dp[0]

	var l, r int32 = 0, int32(m - 1)

	for l < r {
		mid := (l + r) / 2
		if node.left.Value() == mv {
			r = mid
			node = node.left
		} else {
			l = mid + 1
			node = node.right
		}
	}

	s := b[r]

	var res [][]int
	for i := 0; i < n; {
		sum := int32(a[i])
		j := i + 1
		for j < n && sum != s {
			sum += int32(a[j])
			j++
		}
		if sum == s && dp[j].Get(r, 0, int32(m-1)) == mv-1 {
			res = append(res, []int{i + 1, j})
			i = j
			mv--
		} else {
			i++
		}
	}
	return res
}

type Node struct {
	val         int32
	left, right *Node
}

func (node *Node) Value() int32 {
	if node == nil {
		return 0
	}
	return node.val
}

func (node *Node) Set(p int32, v int32, l int32, r int32) *Node {
	res := new(Node)
	if node != nil {
		res.left = node.left
		res.right = node.right
		res.val = node.val
	}
	if l == r {
		res.val = max(res.val, v)
	} else {
		mid := (l + r) / 2
		if p <= mid {
			res.left = res.left.Set(p, v, l, mid)
		} else {
			res.right = res.right.Set(p, v, mid+1, r)
		}
		res.val = max(res.left.Value(), res.right.Value())
	}
	return res
}

func (node *Node) Get(p int32, l int32, r int32) int32 {
	if node == nil {
		return 0
	}
	if l == r {
		return node.val
	}
	mid := (l + r) / 2
	if p <= mid {
		return node.left.Get(p, l, mid)
	}
	return node.right.Get(p, mid+1, r)
}
