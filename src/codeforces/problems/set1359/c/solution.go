package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		H, C, T := readThreeNums(reader)

		fmt.Println(solve(H, C, T))
	}
}

func solve(h, c, t int) int64 {
	if h <= t {
		return 1
	}

	if 2*t <= h+c {
		return 2
	}

	calc := func(n int64) float64 {
		y := float64(2*n - 1)
		x := float64(n*int64(h)) / y
		z := float64((n-1)*int64(c)) / y
		return x + z
	}

	var left, right int64 = 2, math.MaxInt32

	for left < right {
		mid := (left + right) / 2
		x := calc(mid)
		if x <= float64(t) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	x := calc(right - 1)
	y := calc(right)
	if math.Abs(x-float64(t)) <= math.Abs(y-float64(t)) {
		right--
	}
	return 2*right - 1
}

func solve1(h, c, t int) int64 {

	if h <= t {
		return 1
	}

	if 2*t <= h+c {
		return 2
	}

	a := int64(h - t)
	b := int64(2*t - c - h)
	k := int64(2*(a/b)) + 1
	val1 := abs(k/2*int64(c) + (k+1)/2*int64(h) - int64(t)*k)
	val2 := abs((k+2)/2*int64(c) + (k+3)/2*int64(h) - int64(t)*(k+2))
	if val1*(k+2) <= val2*k {
		return k
	}
	return k + 2
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}
