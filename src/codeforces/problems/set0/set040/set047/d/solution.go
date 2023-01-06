package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	words := make([]string, m)

	for i := 0; i < m; i++ {
		words[i] = readString(reader)
	}

	res := solve(n, m, words)

	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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
func solve(n int, m int, attempts []string) int {

	first := attempts[0][:n]
	v := int(attempts[0][n+1] - '0')
	// v <= 5
	variants := make(map[string]bool)

	buf := []byte(first)

	var dfs func(i int, cnt int)

	dfs = func(i int, cnt int) {
		if cnt == 0 {
			variants[string(buf)] = true
			return
		}

		if i == len(first) {
			return
		}

		dfs(i+1, cnt)

		if first[i] == '0' {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}

		dfs(i+1, cnt-1)

		buf[i] = first[i]
	}

	dfs(0, v)

	check := func(x string, y string, v int) bool {
		var u int
		for i := 0; i < n; i++ {
			if x[i] != y[i] {
				u++
			}
		}
		return u == v
	}

	for i := 1; i < m; i++ {
		cur := attempts[i][:n]
		u := int(attempts[i][n+1] - '0')
		var not_valid []string
		for x := range variants {
			if !check(x, cur, u) {
				not_valid = append(not_valid, x)
			}
		}

		for _, x := range not_valid {
			delete(variants, x)
		}
	}

	return len(variants)
}
