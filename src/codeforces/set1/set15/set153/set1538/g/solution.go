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
		nums := readNNums(reader, 4)
		res := solve(nums[0], nums[1], nums[2], nums[3])
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

func solve(x int, y int, a int, b int) int {
	if a == b {
		return min(x/a, y/a)
	}
	// a != b
	// 当有m个类型1的gift set时，即 a个red, b个blue
	get := func(m int) float64 {
		w := float64(x - m*a)
		v := float64(y - m*b)

		n := min(w/float64(b), v/float64(a))

		return float64(m) + n
	}

	// a * m <= x
	// b * m <= y
	l, r := 0, min(x/a, y/b)

	for r-l > 3 {
		// 1, 2, 3, 4
		m1 := l + (r-l)/3
		m2 := r - (r-l)/3

		s1 := get(m1)
		s2 := get(m2)

		if s1 >= s2 {
			r = m2
		} else {
			l = m1
		}
	}

	// r - l < 3
	// ans := l
	var best int
	for i := l; i <= r; i++ {
		best = max(best, int(get(i)))
	}
	return best
}
