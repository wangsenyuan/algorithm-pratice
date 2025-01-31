package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	res, play := solve(n)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d %d\n", res[i].first, res[i].second))
	}

	fmt.Print(buf.String())

	q := readNum(reader)

	for q > 0 {
		q--
		l, r := readTwoNums(reader)
		ids := play(l, r)
		a, b := ids[0], ids[1]
		fmt.Printf("%d %d\n", a, b)
	}
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

type pair struct {
	first  int
	second int
}

func solve(n int) ([]pair, func(int, int) []int) {
	var res []pair

	h := bits.Len(uint(n))
	pos := make([][]int, n+1)

	for l := 1; l <= n; l++ {
		pos[l] = make([]int, h)
		for j := 0; l+(1<<j) <= n+1; j++ {
			pos[l][j] = len(res) + 1
			r := (l + 1<<j) - 1
			res = append(res, pair{l, r})
		}
	}

	play := func(l int, r int) []int {
		d := r - l + 1
		if d == 1 {
			return []int{pos[l][0], pos[l][0]}
		}
		var j int
		for 1<<(j+1) <= d {
			j++
		}
		a := pos[l][j]
		b := pos[r-(1<<j)+1][j]
		return []int{a, b}
	}

	return res, play
}
