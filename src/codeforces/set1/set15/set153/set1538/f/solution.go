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
		l, r := readTwoNums(reader)
		res := solve(l, r)
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

const D = 20

var F [D]int

func init() {
	// 0 -> 9 + (2) = 11
	F[0] = 0
	F[1] = 11

	for i := 2; i < D; i++ {
		// 10, 20, 30 .... 90 => 99 + 3
		// F[i-1] - i 是得到 9999 (i-1个9， 不要变成10000)
		F[i] = 9*F[i-1] + (F[i-1] - i) + i + 1
	}
}

func solve(l int, r int) int {
	var res int

	for r > 0 {
		res += r - l
		r /= 10
		l /= 10
	}
	return res
}

func solve1(l int, r int) int {

	calc := func(num int) int {
		res := num % 10
		d := 1
		num /= 10
		for num > 0 {
			res += (num % 10) * F[d]
			d++
			num /= 10
		}
		return res
	}

	return calc(r) - calc(l)
}
