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
		s, _ := reader.ReadString('\n')
		res := solve(s)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Println(buf.String())
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

func solve(S string) bool {
	vs := make(map[byte]int)
	var n int
	for i := 0; i < len(S); i++ {
		if S[i] == '\n' {
			continue
		}
		n++
		if S[i] >= 'a' && S[i] <= 'z' {
			if _, ok := vs[S[i]]; !ok {
				vs[S[i]] = len(vs)
			}
		}
	}

	root := ParseExpr(S[:n])

	check := func(state int) bool {
		fn := func(x byte) bool {
			i := vs[x]
			return (state>>uint(i))&1 == 1
		}

		return root.evaluate(fn)
	}

	N := 1 << uint(len(vs))

	for i := 0; i < N; i++ {
		if !check(i) {
			return false
		}
	}

	return true
}

// C (and), D(or), E (eq), I (implies, x => y, means if x is true, then y is true), N (not)
type Node struct {
	left, right *Node
	v           byte
}

func (node *Node) evaluate(ctx func(byte) bool) bool {
	if node.v == 'C' {
		return node.left.evaluate(ctx) && node.right.evaluate(ctx)
	}
	if node.v == 'D' {
		return node.left.evaluate(ctx) || node.right.evaluate(ctx)
	}
	if node.v == 'I' {
		return !node.left.evaluate(ctx) || node.right.evaluate(ctx)
	}
	if node.v == 'E' {
		return node.left.evaluate(ctx) == node.right.evaluate(ctx)
	}
	if node.v == 'N' {
		return !node.left.evaluate(ctx)
	}

	return ctx(node.v)
}

func ParseExpr(S string) *Node {

	var build func() *Node

	var pos int

	build = func() *Node {
		node := new(Node)
		node.v = S[pos]
		pos++
		if node.v == 'C' || node.v == 'D' || node.v == 'I' || node.v == 'E' {
			node.left = build()
			node.right = build()
		} else if node.v == 'N' {
			node.left = build()
		}
		return node
	}
	return build()
}
