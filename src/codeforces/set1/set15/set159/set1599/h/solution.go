package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(x int, y int) int {
		fmt.Printf("? %d %d\n", x, y)
		return readNum(reader)
	}

	res := solve(ask)

	fmt.Printf("! %d %d %d %d\n", res[0], res[1], res[2], res[3])
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

const n = 1_000_000_000

func solve(ask func(int, int) int) []int {
	// x1 - 1 + y1 - 1 = a
	a := ask(1, 1)

	// x1 - 1 + n - y2 = b
	// b := ask(1, n)
	// n - x2 + y1 - 1 = c
	c := ask(n, 1)

	// n - x2 + n - y2 = d
	d := ask(n, n)

	// d - c = n - y2 - y1 + 1
	// y2 + y1 = n + 1 - (d - c)
	mid := ((n + 1) - (d - c)) / 2
	// x1 - 1 = e
	e := ask(1, mid)
	x1 := e + 1
	y1 := a + 2 - x1
	y2 := n + 1 - (d - c) - y1
	x2 := 2*n - d - y2

	return []int{x1, y1, x2, y2}
}
