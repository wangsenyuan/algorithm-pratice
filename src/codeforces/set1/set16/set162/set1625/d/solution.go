package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, K := readTwoNums(reader)
	X := readNNums(reader, N)

	res := solve(X, K)

	if len(res) == 0 {
		fmt.Println("-1")
	} else {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
	}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const D = 31

func solve(X []int, K int) []int {
	n := len(X)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	sort.Slice(arr, func(i, j int) bool {
		return X[arr[i]] < X[arr[j]]
	})

	prev := make([]int, n)

	root := new(Node)
	root.children = make([]*Node, 2)

	var best int
	var bestAt int

	for i := 0; i < n; i++ {
		prev[i] = -1
		var dp int
		j := arr[i]
		num := X[j]
		cur := root
		for d := D - 1; d >= 0 && cur != nil; d-- {
			u := (K >> d) & 1
			v := (num >> d) & 1
			tmp := cur.children[1^v]

			if u == 1 {
				cur = tmp
			} else {
				// u == 0
				if tmp != nil {
					// record another path
					if tmp.val+1 > dp {
						dp = tmp.val + 1
						prev[i] = tmp.idx
					}
				}
				// but follow the safe path
				cur = cur.children[v]
			}
		}

		if cur != nil && cur.val+1 > dp {
			dp = cur.val + 1
			prev[i] = cur.idx
		}

		best = max(best, dp)

		if best == dp {
			bestAt = i
		}

		root = root.AddNum(num, dp, i, D-1)
	}

	if best == 0 {
		return nil
	}

	var res []int

	for bestAt >= 0 {
		res = append(res, arr[bestAt]+1)
		bestAt = prev[bestAt]
	}

	sort.Ints(res)

	return res
}

type Node struct {
	children []*Node
	val      int
	idx      int
}

func (node *Node) AddNum(num int, v int, i int, d int) *Node {
	if node == nil {
		node = new(Node)
		node.val = -1;
		node.children = make([]*Node, 2)
	}
	if d < 0 {
		if node.val < v {
			node.val = v
			node.idx = i
		}
		return node
	}

	x := (num >> d) & 1

	node.children[x] = node.children[x].AddNum(num, v, i, d-1)

	if node.children[x^1] == nil {
		node.val = node.children[x].val
		node.idx = node.children[x].idx
	} else {
		c := x
		if node.children[x].val < node.children[x^1].val {
			c = x ^ 1
		}
		node.val = node.children[c].val
		node.idx = node.children[c].idx
	}

	return node
}

func (node *Node) GetValue() int {
	if node == nil {
		return -1
	}
	return node.val
}

func (node *Node) GetIndex() int {
	if node == nil {
		return -1
	}
	return node.idx
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
