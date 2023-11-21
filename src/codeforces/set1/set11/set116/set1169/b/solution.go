package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	//tc := readNum(scanner)
	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		pairs := make([][]int, m)
		for i := 0; i < m; i++ {
			pairs[i] = readNNums(reader, 2)
		}
		res := solve(n, pairs)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Println(buf.String())
}

func solve(n int, pairs [][]int) bool {
	// x, y
	// pair[0] 里面的 (first, second) 肯定要么是x，要么是y

	find2 := func(x int, y int) bool {
		for _, p := range pairs {
			if p[0] == x || p[0] == y {
				continue
			}
			if p[1] == x || p[1] == y {
				continue
			}
			return false
		}

		return true
	}

	find := func(x int) bool {
		// given x, to find y
		var cnt int
		for _, p := range pairs {
			if p[0] != x && p[1] != x {
				return find2(x, p[0]) || find2(x, p[1])
			}
			cnt++
		}
		return cnt == len(pairs)
	}

	return find(pairs[0][0]) || find(pairs[0][1])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
