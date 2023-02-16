package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	Q := make([][]int, n)

	for i := 0; i < n; i++ {
		Q[i] = readNNums(reader, 3)
	}
	res := solve(n, m, Q)
	var buf bytes.Buffer

	for i := 0; i < m; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}

	buf.WriteByte('\n')

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

func solve(n, m int, Q [][]int) []int {
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		ans[i] = -1
	}

	M := int64(m)

	seen := make([]bool, M+1)
	seen[0] = true
	vis := make([]bool, M+1)
	for i := 1; i <= n; i++ {
		copy(vis, seen)

		t, x, y := Q[i-1][0], Q[i-1][1], Q[i-1][2]
		for b := 0; b <= m; b++ {
			cur := int64(b)
			if seen[cur] {
				operate(&cur, t, int64(x))
				for a := 1; a <= y; a++ {
					if cur > M || cur <= 0 || seen[cur] {
						break
					}
					vis[cur] = true
					ans[cur-1] = i
					operate(&cur, t, int64(x))
				}
			}
		}

		copy(seen, vis)
	}

	return ans
}

const DIV = 1e5

func ceil(x, y int64) int64 {
	return (x + y - 1) / y
}

func add(cur *int64, x int64) {
	*cur += ceil(x, DIV)
}

func mul(cur *int64, x int64) {
	*cur = ceil(*cur*x, DIV)
}

func operate(cur *int64, t int, x int64) {
	if t == 1 {
		add(cur, x)
	} else {
		mul(cur, x)
	}
}
