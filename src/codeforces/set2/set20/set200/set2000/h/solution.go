package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	for tc > 0 {
		tc--
		res := process(reader)

		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	m := readNum(reader)
	queries := make([]string, m)
	for i := range m {
		queries[i] = readString(reader)
	}
	return solve(a, queries)
}

func solve(a []int, queries []string) []int {
	var tr *Node

	for _, x := range a {
		tr = tr.Update(1, X, x, 1)
	}

	ans := make([]int, 0, len(queries))

	for _, cur := range queries {
		var v int
		readInt([]byte(cur), 2, &v)
		if cur[0] == '+' {
			tr = tr.Update(1, X, v, 1)
		} else if cur[0] == '-' {
			tr = tr.Update(1, X, v, -1)
		} else {
			if tr == nil {
				ans = append(ans, 1)
			} else if tr.GetValue() < v {
				// 最大的数+1
				ans = append(ans, tr.rv+1)
			} else {
				ans = append(ans, tr.Find(1, X, v))
			}
		}
	}

	return ans
}

const X = 2_000_000

type Node struct {
	left, right *Node
	val         int
	lv, rv      int
}

func (node *Node) GetValue() int {
	if node == nil {
		return 0
	}
	return node.val
}

func (node *Node) Update(l int, r int, p int, v int) *Node {
	if node == nil {
		node = new(Node)
	}
	if l == r {
		if v < 0 {
			// negative v is for remove
			return nil
		}
		node.val = 0
		node.lv = l
		node.rv = r
		return node
	}
	mid := (l + r) / 2
	if p <= mid {
		node.left = node.left.Update(l, mid, p, v)
	} else {
		node.right = node.right.Update(mid+1, r, p, v)
	}

	if node.left == nil && node.right == nil {
		// 删除完子节点后，它也得变成nil
		return nil
	}
	// merge
	node.val = max(node.left.GetValue(), node.right.GetValue())
	if node.left != nil && node.right != nil {
		// 空洞的大小
		node.val = max(node.val, node.right.lv-node.left.rv-1)
		node.lv = node.left.lv
		node.rv = node.right.rv
	} else if node.right != nil {
		// 只有右节点时，比如它的最小值时5,那么空洞的大小就是4
		node.val = max(node.val, node.right.lv-l)
		node.lv = node.right.lv
		node.rv = node.right.rv
	} else {
		// node.right == nil  r也不存在，所以不需要-1
		node.val = max(node.val, r-node.left.rv)
		node.lv = node.left.lv
		node.rv = node.left.rv
	}

	return node
}

func (node *Node) Find(l int, r int, k int) int {
	if node == nil {
		return l
	}
	// node.val >= k holdes 而且不会到叶子节点
	mid := (l + r) / 2
	if node.left.GetValue() >= k {
		return node.left.Find(l, mid, k)
	}

	if node.left == nil {
		if node.right.lv-l >= k {
			// 这个是这个区间能达到的最小值
			return l
		}
		// else
		return node.right.Find(mid+1, r, k)
	}
	// node.left != nil
	x := r
	if node.right != nil {
		x = node.right.lv - 1
	}
	if x-node.left.rv >= k {
		return node.left.rv + 1
	}
	return node.right.Find(mid+1, r, k)
}
