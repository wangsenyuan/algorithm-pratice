package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	ask := func(a, b, c int) (int, int) {
		fmt.Printf("1 %d %d %d\n", a, b, c)
		return readTwoNums(reader)
	}

	sanitize := func(x, y int) int {
		fmt.Printf("2 %d %d\n", x, y)
		return readNum(reader)
	}

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		solve(n, m, ask, sanitize)
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

func solve(n int, m int, ask func(a, b, c int) (int, int), sanitize func(x, y int) int) {
	a := make(map[int]int)
	b := make(map[int]int)
	d := make(map[int]int)

	for i := 1; i <= 128; i++ {
		for j := 1; j <= 128; j++ {
			if gcd(i, j) == 1 {
				u, dist := ask(i, j, 0)
				if _, found := a[u]; found {
					a1 := int64(a[u])
					b1 := int64(b[u])
					d1 := int64(-d[u])
					a2 := int64(i)
					b2 := int64(j)
					d2 := int64(-dist)
					den := a1*b2 - a2*b1
					x := (b1*d2 - b2*d1) / den
					y := (a2*d1 - a1*d2) / den
					v := sanitize(int(x), int(y))
					if v == 1 {
						return
					}
				}
				a[u] = i
				b[u] = j
				d[u] = dist
			}
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
