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
		if x >= 0 {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		} else {
			buf.WriteString("MORTAL\n")
		}
	}
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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
	tc := readNum(reader)
	res := make([]int, tc)
	for i := range tc {
		n, m := readTwoNums(reader)
		land := make([]string, n)
		for j := range n {
			land[j] = readString(reader)[:m]
		}
		res[i] = solve(land)
	}
	return res
}

func solve(land []string) int {
	n := len(land)
	m := len(land[0])
	row := make([]int, n)
	col := make([]int, m)
	var cnt int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if land[i][j] == 'A' {
				cnt++
				row[i]++
				col[j]++
			}
		}
	}
	if cnt == 0 {
		return -1
	}

	if cnt == n*m {
		return 0
	}

	if row[0] == m || row[n-1] == m || col[0] == n || col[m-1] == n {
		// 一次就够了
		return 1
	}

	for i := 0; i < n; i++ {
		if row[i] == m {
			return 2
		}
	}
	for j := 0; j < m; j++ {
		if col[j] == n {
			return 2
		}
	}
	if land[0][0] == 'A' || land[0][m-1] == 'A' || land[n-1][0] == 'A' || land[n-1][m-1] == 'A' {
		return 2
	}

	if row[0] > 0 || row[n-1] > 0 || col[0] > 0 || col[m-1] > 0 {
		return 3
	}

	return 4
}
