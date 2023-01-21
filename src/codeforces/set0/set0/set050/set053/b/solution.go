package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	h, w := readTwoNums(reader)

	a, b := solve(h, w)

	fmt.Printf("%d %d\n", a, b)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
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

func solve(h int, w int) (int, int) {
	H := int64(h)
	W := int64(w)
	// a, b, if a is power of 2
	a, b := check(H, W)
	x, y := check(W, H)

	if a*b > x*y || a*b == x*y && a > y {
		return int(a), int(b)
	}
	return int(y), int(x)
}

func check(x int64, y int64) (int64, int64) {
	// x need to be the power of 2
	a := int64(1)

	var area int64
	var h, w int64

	for a <= x {

		// 1.25
		// a * 5 / 4
		b := a * 5 / 4
		for (b+1)*4 <= a*5 {
			b++
		}
		if b <= y {
			if area < a*b {
				area = a * b
				h, w = a, b
			}
		} else {
			if y*5 >= a*4 {
				// y >= a * 0.8
				if area < a*y {
					area = a * y
					h, w = a, y
				}
			}
		}
		a *= 2
	}
	return h, w
}
