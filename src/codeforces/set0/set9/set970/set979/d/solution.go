package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	queries := make([][]int, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')

		if s[0] == '1' {
			queries[i] = make([]int, 2)
			queries[i][0] = 1
		} else {
			queries[i] = make([]int, 4)
			queries[i][0] = 2
		}
		pos := 2
		for j := 1; j < len(queries[i]); j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}
	return solve(queries)
}

type Node struct {
	children [2]*Node
	val      int
}

func NewNode() *Node {
	node := new(Node)
	node.val = inf
	return node
}

func (root *Node) add(num int) {
	node := root
	node.val = min(node.val, num)
	for i := D - 1; i >= 0; i-- {
		x := (num >> i) & 1
		if node.children[x] == nil {
			node.children[x] = NewNode()
		}
		node = node.children[x]
		node.val = min(node.val, num)
	}
}

const inf = 1 << 60
const X = 1e5 + 5
const D = 19

func solve(queries [][]int) []int {

	var fs [X][]int
	var nodes [X]*Node

	for i := 1; i < X; i++ {
		for j := i; j < X; j += i {
			fs[j] = append(fs[j], i)
		}
		nodes[i] = NewNode()
	}

	query := func(x, k, s int) int {
		if x%k != 0 || nodes[k].val+x > s {
			return -1
		}
		var res int
		node := nodes[k]
		for i := D - 1; i >= 0; i-- {
			d := (x >> i) & 1
			if node.children[d^1] != nil && node.children[d^1].val+x <= s {
				res ^= (d ^ 1) << i
				node = node.children[d^1]
			} else {
				res ^= d << i
				node = node.children[d]
			}
		}
		return res
	}

	marked := make([]bool, X)

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			u := cur[1]
			if !marked[u] {
				marked[u] = true
				for _, k := range fs[u] {
					nodes[k].add(u)
				}
			}
		} else {
			ans = append(ans, query(cur[1], cur[2], cur[3]))
		}
	}

	return ans
}
