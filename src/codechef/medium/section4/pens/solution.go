package main

import (
	"bufio"
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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)

		scanner.Scan()
		S := scanner.Text()

		P := make([]string, k)

		for i := 0; i < k; i++ {
			scanner.Scan()
			P[i] = scanner.Text()
		}
		res := solve(n, k, S, P)

		for i := 0; i < n; i++ {
			if i < n-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
	}
}

const MAX_L = 20

var can []int

func init() {
	can = make([]int, 1<<MAX_L+1)
}

func solve(n, k int, S string, P []string) []int {
	for i := 0; i < len(can); i++ {
		can[i] = 0
	}

	for j, p := range P {
		var f int
		for i := 0; i < len(p); i++ {
			x := uint(p[i] - 'a')
			f |= 1 << x
		}
		can[f] = j + 1
	}

	for mask := 1<<MAX_L - 1; mask > 0; mask-- {
		if can[mask] == 0 {
			continue
		}

		cur := mask

		for cur > 0 {
			can[mask^(cur&(-cur))] = can[mask]
			cur ^= cur & -cur
		}
	}

	res := make([]int, n)

	for i := 0; i < n; {
		var j int = i
		var f int
		for j < n {
			x := uint(S[j] - 'a')
			if can[f|(1<<x)] == 0 {
				break
			}
			f |= 1 << x
			j++
		}

		for k := i; k < j; k++ {
			res[k] = can[f]
		}

		i = j
	}

	return res
}

func solve1(n, k int, S string, P []string) []int {

	root := NewNode()

	for j, p := range P {
		var f int
		for i := 0; i < len(p); i++ {
			x := uint(p[i] - 'a')
			f |= 1 << x
		}

		node := root

		for i := 20; i >= 0; i-- {
			x := (f >> uint(i)) & 1
			if node.children[x] == nil {
				node.children[x] = NewNode()
			}
			node = node.children[x]
		}

		node.index = j + 1
	}

	var find func(node *Node, f int, pos int) int

	find = func(node *Node, f int, pos int) int {
		if node == nil {
			return 0
		}

		if pos < 0 {
			return node.index
		}

		x := (f >> uint(pos)) & 1

		if x == 1 {
			return find(node.children[1], f, pos-1)
		}
		ans := find(node.children[0], f, pos-1)
		if ans > 0 {
			return ans
		}
		return find(node.children[1], f, pos-1)
	}

	res := make([]int, n)

	for i := 0; i < n; {
		var cand int
		var j = i
		var f int
		for j < n {
			x := uint(S[j] - 'a')
			f |= 1 << x
			tmp := find(root, f, 20)
			if tmp == 0 {
				break
			}
			cand = tmp
			j++
		}

		for k := i; k < j; k++ {
			res[k] = cand
		}

		i = j
	}

	return res
}

type Node struct {
	children []*Node
	index    int
}

func NewNode() *Node {
	node := new(Node)
	node.children = make([]*Node, 2)
	return node
}
