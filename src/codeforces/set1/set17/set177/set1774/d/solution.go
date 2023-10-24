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
		n, m := readTwoNums(reader)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = readNNums(reader, m)
		}
		cnt, res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", cnt))
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
		}
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

func solve(a [][]int) (int, [][]int) {
	// mim operation to make every row has same count of 1
	n := len(a)
	m := len(a[0])
	var cnt int
	// row[i] = cnt of 1 in this row
	row := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			row[i] += a[i][j]
		}
		cnt += row[i]
	}
	if cnt%n != 0 {
		return -1, nil
	}
	x := cnt / n
	// each row should have x cnt of 1
	// 如果, 然后按照1的个数进行排序吗？
	var res [][]int

	first := make([]int, n)
	second := make([]int, n)

	for c := 0; c < m; c++ {
		var u, v int
		for i := 0; i < n; i++ {
			if a[i][c] == 0 && row[i] < x {
				first[u] = i
				u++
			}
			if a[i][c] == 1 && row[i] > x {
				second[v] = i
				v++
			}
		}
		if u == 0 || u == n {
			// no swaps
			continue
		}

		for j := 0; j < u && j < v; j++ {
			// swap them is good
			res = append(res, []int{first[j] + 1, second[j] + 1, c + 1})
			row[first[j]]++
			row[second[j]]--
		}
	}

	return len(res), res
}
