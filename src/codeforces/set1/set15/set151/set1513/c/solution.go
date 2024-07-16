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
		s, _ := reader.ReadBytes('\n')
		var i int
		for i < len(s) && s[i] != ' ' {
			i++
		}
		num := string(s[:i])
		var m int
		readInt(s, i+1, &m)
		res := solve(num, m)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const M = 2_000_10

var F [M]int

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func init() {
	// F[i] = 10变化i次后的长度
	for i := 0; i < 9; i++ {
		// 0， 1， 2， 。。。。 9
		F[i] = 2
	}
	// 109
	F[9] = 3

	for i := 10; i < M; i++ {
		// 2110 = 4
		F[i] = add(F[i-9], F[i-10])
	}
}

func solve(n string, m int) int {
	get := func(x int, m int) int {
		m -= 10 - x
		if m < 0 {
			return 1
		}

		return F[m]
	}

	var res int

	for i := 0; i < len(n); i++ {
		d := int(n[i] - '0')
		res = add(res, get(d, m))
	}

	return res
}
