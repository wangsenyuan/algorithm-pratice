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
		b, c, d := readThreeNums(reader)
		res := solve(b, c, d)
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

func solve(b, c, d int) int {
	// a | b 肯定是 a & c 的超集，
	var a int

	for i := 0; i < 61; i++ {
		x := (b >> i) & 1
		y := (c >> i) & 1
		z := (d >> i) & 1
		if z == 0 {
			if y == 0 && x == 1 {
				return -1
			}
			if y == 1 {
				a |= 1 << i
			}
			continue
		}
		// z == 1
		if x == 0 && y == 1 {
			// a[i] = 0, 0 - 0
			// a[i] = 1, 1 - 1
			return -1
		}
		// x = 0,  y = 0, 那么a[i] = 1
		// x = 1, y = 0/1, a[i] = 0
		a |= (x ^ 1) << i
	}
	return a
}
