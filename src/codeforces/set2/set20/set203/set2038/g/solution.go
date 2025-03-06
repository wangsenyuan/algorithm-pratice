package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(t string) int {
		fmt.Printf("1 %s\n", t)
		return readNum(reader)
	}

	tc := readNum(reader)
	for range tc {
		n := readNum(reader)
		pos, x := solve(n, ask)
		fmt.Printf("0 %d %c\n", pos, x)
		res := readNum(reader)
		if res < 0 {
			break
		}
	}
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

func solve(n int, ask func(x string) int) (pos int, ans byte) {
	x := ask("1")
	if x == 0 {
		return 1, '0'
	}
	if x == n {
		return 1, '1'
	}
	// x > 0 & x < n
	y := ask("10")
	if y == 0 {
		// 000011
		return 1, '0'
	}
	// y > 0, 0?101010
	z := ask("11")
	if z+y == x {
		return n, '0'
	}
	return n, '1'
}
