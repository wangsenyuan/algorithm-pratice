package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	ask := func(n int) int {
		fmt.Println(n)
		return readNum(reader)
	}

	for range tc {
		fn := solve(ask)
		fmt.Println("!")
		res := fn(readNum(reader))
		fmt.Println(res)
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

const H = 30

func solve(ask func(int) int) (response func(int) int) {
	var n1, n2 int
	for i := H - 1; i >= 0; i-- {
		if i&1 == 0 {
			n1 |= 1 << i
		} else {
			n2 |= 1 << i
		}
	}
	q1 := ask(n1)
	q2 := ask(n2)

	var x, y int
	// 特殊处理最低的两位
	if q2&1 == 1 {
		x |= 1
	} else if q2&2 == 2 {
		x |= 1
		y |= 1
	}
	if q1&2 == 0 {
		// 10
		x |= 2
	} else if q1&4 == 4 {
		// 有进位
		x |= 2
		y |= 2
	}

	for i := 2; i < H; i += 2 {
		if (q2>>i)&1 == 0 {
			x |= 1 << i
		} else if (q2>>i)&2 == 2 {
			x |= 1 << i
			y |= 1 << i
		}
		if (q1>>i)&2 == 0 {
			// 那么 x和y肯定是不同的
			x |= 2 << i
		} else if (q1>>i)&4 == 4 {
			// 看有没有进位

			x |= 2 << i
			y |= 2 << i
		}
	}

	return func(m int) int {
		return (m | x) + (m | y)
	}
}
