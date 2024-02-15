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

	n, m, k := readThreeNums(reader)
	s := readString(reader)
	res := solve(n, m, k, s)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(res[i])
		buf.WriteByte('\n')
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

func solve(n int, m int, k int, s string) []string {
	buf := []byte(s)
	sort.Slice(buf, func(i, j int) bool {
		return buf[i] < buf[j]
	})

	mat := make([][]byte, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			mat[i][j] = ' '
		}
	}
	var it int
	var begin int
	for j := 0; j < m; j++ {
		for i := begin; i < k; i++ {
			mat[i][j] = buf[it]
			it++
			if i > begin && mat[i-1][j] != mat[i][j] {
				begin = i
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == ' ' {
				mat[i][j] = buf[it]
				it++
			}
		}
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = string(mat[i])
	}
	return res
}
