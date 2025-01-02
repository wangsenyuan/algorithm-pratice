package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

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
	n, m := readTwoNums(reader)
	a := make([][]int, n)
	for i := range n {
		a[i] = readNNums(reader, n)
	}
	q := make([][]int, m)
	for i := range m {
		q[i] = readNNums(reader, 4)
	}
	return solve(a, q)
}

func solve(a [][]int, q [][]int) []int {
	n := len(a)

	s := make([][][3]int, n+1)
	for i := range n + 1 {
		s[i] = make([][3]int, n+1)
	}

	for i := range n {
		for j := range n {
			v := a[i][j]
			s[i+1][j+1][0] = s[i+1][j][0] + s[i][j+1][0] - s[i][j][0] + v
			s[i+1][j+1][1] = s[i+1][j][1] + s[i][j+1][1] - s[i][j][1] + v*i
			s[i+1][j+1][2] = s[i+1][j][2] + s[i][j+1][2] - s[i][j][2] + v*(j+1)
		}
	}

	query := func(r1, c1, r2, c2 int, t int) int {
		return s[r2][c2][t] - s[r2][c1][t] - s[r1][c2][t] + s[r1][c1][t]
	}

	ans := make([]int, len(q))
	for i, cur := range q {
		r1, c1, r2, c2 := cur[0]-1, cur[1]-1, cur[2], cur[3]
		s0 := query(r1, c1, r2, c2, 0)
		s1 := query(r1, c1, r2, c2, 1)
		s2 := query(r1, c1, r2, c2, 2)
		ans[i] = (c2-c1)*(s1-r1*s0) + s2 - c1*s0
	}
	return ans
}
