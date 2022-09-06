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
		res := solve1(nums)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(nums []int) int64 {
	a, b, c, d := nums[0], nums[1], nums[2], nums[3]
	if a > b || c > d || d < a {
		return 0
	}

	if b > d {
		b = d
	}

	if c < a {
		c = a
	}

	if b < c {
		return int64(b-a+1) * int64(d-c+1)
	}

	ans := int64(b-a+1) * int64(d-c+1)
	ans -= int64(b-c+1) * int64(b-c+2) / 2

	return ans
}

func solve1(nums []int) int64 {
	a, b, c, d := nums[0], nums[1], nums[2], nums[3]

	if a > b || c > d || d < a {
		return 0
	}

	// need x < y
	// b >= c
	// a < b, c < d & b >= c
	// formula not that
	var sum int64
	for x := a; x <= b; x++ {
		y := max(x+1, c)
		if y <= d {
			sum += int64(d - y + 1)
		}
	}

	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
