package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		P := readNNums(reader, n)
		res := solve(k, P)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

const H = 31

func solve(k int, A []int) int {
	// a[l...r] xor >= k
	// 假设迭代r，到r时，从高位一次处理
	// 如果到j时， k[j] = 0, 这时如果存在靠近右边的l，有 a[l...r] xor = 1,
	// 那么 r - l + 1是一个结果
	// 然后继续进行
	// 如果 k[j] = 1, 但是在当前的答案范围内，不存在 xor[j] = 1
	// 那么停止
	// 现在剩下的就是，如何找到左边，最近的l， 可以用bit
	n := len(A)

	best := n + 1

	var x int

	root := new(Node)

	root.Add(0, H-1, -1)

	for r := 0; r < n; r++ {
		// xor(A[l..r]) >= k
		//if A[r] >= k {
		//	return 1
		//}
		// xor(A[0..r]) ^ xor(A[0..l-1]) >= k
		x ^= A[r]

		cur := root

		for i := H - 1; i >= 0 && cur != nil; i-- {
			a := (x >> i) & 1
			b := (k >> i) & 1
			if b == 0 {
				// if we can find a 1,
				if cur.children[1^a] != nil {
					best = min(best, r-cur.children[1^a].id)
				}
				cur = cur.children[a]
			} else {
				// b == 1
				cur = cur.children[1^a]
			}
		}

		//if x == k {
		//	best = min(best, r+1)
		//}

		if cur != nil {
			best = min(best, r-cur.id)
		}

		root.Add(x, H-1, r)
	}

	if best > n {
		return -1
	}

	return best
}

type Node struct {
	children [2]*Node
	id       int
}

func (node *Node) Add(v int, d int, id int) {
	node.id = id
	if d < 0 {
		return
	}
	//node.id = id
	x := (v >> d) & 1
	if node.children[x] == nil {
		node.children[x] = new(Node)
	}
	node.children[x].Add(v, d-1, id)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
