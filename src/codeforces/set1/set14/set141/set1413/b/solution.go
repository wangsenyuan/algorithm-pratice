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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		rows := make([][]int, n)
		for i := 0; i < n; i++ {
			rows[i] = readNNums(reader, m)
		}
		cols := make([][]int, m)
		for i := 0; i < m; i++ {
			cols[i] = readNNums(reader, n)
		}
		res := solve(n, m, rows, cols)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
			}
			buf.WriteByte('\n')
		}
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, m int, rows [][]int, cols [][]int) [][]int {
	// 只要按照第一列排就可以了
	first := make(map[int]int)
	for i := 0; i < n; i++ {
		first[rows[i][0]] = i
	}

	//mat := make([][]int, n)

	process := func(col []int) {
		for i := 0; i < n; i++ {
			first[col[i]] = i
		}
		sort.Slice(rows, func(i, j int) bool {
			a := rows[i]
			b := rows[j]
			return first[a[0]] < first[b[0]]
		})
	}

	for _, col := range cols {
		if _, ok := first[col[0]]; ok {
			// find  the first col
			process(col)
			break
		}
	}

	return rows
}
