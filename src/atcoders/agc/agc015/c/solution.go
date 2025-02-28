package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	buf.WriteTo(os.Stdout)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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
	n, m, k := readThreeNums(reader)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = readString(reader)[:m]
	}
	qs := make([][]int, k)
	for i := 0; i < k; i++ {
		qs[i] = readNNums(reader, 4)
	}
	return solve(a, qs)
}

func solve(a []string, qs [][]int) []int {
	n := len(a)
	m := len(a[0])

	sumV := make([][]int, n+1)
	sumE := make([][]int, n+1)
	rowE := make([][]int, n+1)
	colE := make([][]int, n+1)
	for i := range sumV {
		sumV[i] = make([]int, m+1)
		sumE[i] = make([]int, m)
		rowE[i] = make([]int, m+1)
		colE[i] = make([]int, m)
	}
	for i, row := range a {
		for j, v := range row {
			sumV[i+1][j+1] = sumV[i+1][j] + sumV[i][j+1] - sumV[i][j] + int(v-'0')
			if i < n-1 && j < m-1 {
				sumE[i+1][j+1] = sumE[i+1][j] + sumE[i][j+1] - sumE[i][j]
				if v == '1' {
					sumE[i+1][j+1] += int(row[j+1]-'0') + int(a[i+1][j]-'0')
				}
			}
			rowE[i][j+1] = rowE[i][j]
			if j < m-1 && v == '1' {
				rowE[i][j+1] += int(row[j+1] - '0')
			}
			colE[i+1][j] = colE[i][j]
			if i < n-1 && v == '1' {
				colE[i+1][j] += int(a[i+1][j] - '0')
			}
		}
	}

	get := func(r1 int, c1 int, r2 int, c2 int) int {
		r1--
		c1--
		v := sumV[r2][c2] - sumV[r2][c1] - sumV[r1][c2] + sumV[r1][c1]
		r2--
		c2--
		e := sumE[r2][c2] - sumE[r2][c1] - sumE[r1][c2] + sumE[r1][c1] +
			rowE[r2][c2] - rowE[r2][c1] +
			colE[r2][c2] - colE[r1][c2]
		return v - e
	}

	ans := make([]int, len(qs))

	for i, cur := range qs {
		ans[i] = get(cur[0], cur[1], cur[2], cur[3])
	}

	return ans
}
