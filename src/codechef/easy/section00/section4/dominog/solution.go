package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(x, y int) (int, int) {
		fmt.Printf("%d %d\n", x, y)
		return readTwoNums(reader)
	}

	for {
		n, m := readTwoNums(reader)
		if n == 0 {
			break
		}
		solve(n, m, ask)
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

func solve(n int, m int, ask func(r, c int) (int, int)) {
	cache := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			cache[i][j] = -1
		}
	}

	query := func(a, b int) (int, int) {
		if cache[a][b] < 0 {
			c, d := ask(a, b)
			cache[a][b] = (c-1)*m + d - 1
			cache[c][d] = (a-1)*m + (b - 1)
		}
		tmp := cache[a][b]
		c, d := tmp/m, tmp%m
		return c + 1, d + 1
	}

	var loop func(a, b, c, d int)

	loop = func(x1, y1, x2, y2 int) {
		if x1 == x2 && y1 == y2 {
			query(x1, y1)
			return
		}
		// x is height, y is width
		if x2-x1 >= y2-y1 {
			h := (x2 - x1 + 1) / 2
			top := (y2 - y1 + 1) * h
			bot := (y2 - y1 + 1) * (x2 - x1 - h)
			x := x1 + h
			y := y1
			for y <= y2 {
				u, v := query(x, y)
				if u == 0 && v == 0 {
					return
				}
				if u != x {
					// different row
					if u < x {
						top--
					} else {
						bot--
					}
					y++
				} else if v > y {
					// same row
					y += 2
				} else {
					y++
				}
			}
			if top%2 == 1 {
				loop(x1, y1, x-1, y2)
			} else {
				loop(x+1, y1, x2, y2)
			}
		} else {
			w := (y2 - y1 + 1) / 2
			left := (x2 - x1 + 1) * w
			right := (x2 - x1 + 1) * (y2 - y1 - w)
			y := y1 + w
			x := x1
			for x <= x2 {
				u, v := query(x, y)
				if u == 0 && v == 0 {
					return
				}
				if v != y {
					// diff col
					if v < y {
						left--
					} else {
						right--
					}
					x++
				} else if u > x {
					x += 2
				} else {
					x++
				}
			}
			if left%2 == 1 {
				loop(x1, y1, x2, y-1)
			} else {
				loop(x1, y+1, x2, y2)
			}
		}
	}

	loop(1, 1, n, m)
}
