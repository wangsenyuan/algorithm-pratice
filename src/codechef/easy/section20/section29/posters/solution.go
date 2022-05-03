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
		reader.ReadBytes('\n')
		n := readNum(reader)
		posters := make([][]int, n)
		for i := 0; i < n; i++ {
			posters[i] = readNNums(reader, 4)
		}
		res := solve(n, posters)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(n int, rects [][]int) int {
	g := make([][]bool, n)

	overlap := func(i, j int) bool {
		a, b := rects[i], rects[j]
		return a[0] < b[0] && b[1] < a[1] && b[2] < a[2] && a[3] < b[3]
	}

	for i := 0; i < n; i++ {
		g[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			g[i][j] = overlap(i, j)
		}
	}

	var ans int

	v := make([]bool, n)

	pair := make([]int, n)
	for i := 0; i < n; i++ {
		pair[i] = -1
	}

	var aug func(x int) bool

	aug = func(x int) bool {
		if v[x] {
			return false
		}
		v[x] = true
		for y := 0; y < n; y++ {
			if g[x][y] && (pair[y] < 0 || aug(pair[y])) {
				pair[y] = x
				return true
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v[j] = false
		}
		if aug(i) {
			ans++
		}
	}

	return n - ans
}
