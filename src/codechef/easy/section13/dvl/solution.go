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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)

	A := readNNums(scanner, n)

	fmt.Println(solve(n, A))
}

const MOD = 987654319

const MAX_X = 1000000001

func solve(n int, A []int) int64 {
	// dp[i][j] is for B < C, and count of A that A > B && A > C
	dp := make([][]int64, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int64, n)
	}

	root := NewNode(1, MAX_X)

	for i := 0; i < n; i++ {
		b := A[i]
		for j := i + 1; j < n; j++ {
			c := A[j]
			if b < c {
				dp[i][j] = int64(i - root.CountNumLt(c+1))
			}
			dp[i][j] += dp[i][j-1]
			if dp[i][j] >= MOD {
				dp[i][j] -= MOD
			}
		}
		root.SetValue(b)
	}

	//fp[i][j] is for B > D
	fp := make([]int64, n)

	for i := 0; i < n; i++ {
		b := A[i]
		for j := i + 2; j < n; j++ {
			d := A[j]
			if b > d {
				fp[j] += dp[i][j]
				if fp[j] >= MOD {
					fp[j] -= MOD
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = 0
		}
	}

	root = NewNode(0, MAX_X)

	for i := n - 1; i >= 0; i-- {
		f := A[i]
		for j := i - 1; j >= 0; j-- {
			e := A[j]
			if f > e {
				dp[i][j] = int64(root.CountNumLt(e))
			}
			dp[i][j] += dp[i][j+1]
			if dp[i][j] >= MOD {
				dp[i][j] -= MOD
			}
		}
		root.SetValue(f)
	}
	gp := make([]int64, n)

	for i := n - 1; i >= 0; i-- {
		f := A[i]
		for j := i - 1; j >= 0; j-- {
			d := A[j]
			if d > f {
				gp[j] += dp[i][j]
				if gp[j] >= MOD {
					gp[j] -= MOD
				}
			}
		}
	}

	var res int64

	for i := 2; i < n-1; i++ {
		res += (fp[i] * gp[i]) % MOD
		res %= MOD
	}

	return res
}

type Node struct {
	count       int
	start, end  int
	left, right *Node
}

func NewNode(start, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	return node
}

func (node *Node) pushDown() {
	if node.start+1 < node.end && node.left == nil {
		mid := (node.start + node.end) / 2
		node.left = NewNode(node.start, mid)
		node.right = NewNode(mid, node.end)
	}
}

func (node *Node) SetValue(v int) {
	node.pushDown()

	node.count++

	if node.start+1 == node.end {
		return
	}

	mid := (node.start + node.end) / 2

	if v < mid {
		node.left.SetValue(v)
	} else {
		node.right.SetValue(v)
	}
}

func (node *Node) CountNumLt(v int) int {
	// node is for range [start, end)
	if node.end <= v {
		return node.count
	}

	if node.start > v {
		return 0
	}

	if node.left == nil {
		return 0
	}

	// node.start <= v < node.end

	mid := (node.start + node.end) / 2

	if v < mid {
		return node.left.CountNumLt(v)
	}

	return node.left.count + node.right.CountNumLt(v)
}
