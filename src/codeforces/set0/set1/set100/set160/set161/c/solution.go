package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line := readNNums(reader, 4)

	res := solve(line[0], line[1], line[2], line[3])

	fmt.Println(res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(l1, r1, l2, r2 int) int {
	var ans int
	var dfs func(a, b, c, d, i int)

	dfs = func(a, b, c, d, i int) {
		if a > b || c > d || i == 0 {
			return
		}
		ans = max(ans, min(b-c, b-a, d-a, d-c)+1)
		if b-a < ans || d-c < ans {
			return
		}
		mid := 1 << (i - 1)
		dfs(a, min(b, mid-1), c, min(d, mid-1), i-1)
		dfs(a, min(b, mid-1), max(1, c-mid), d-mid, i-1)
		dfs(max(1, a-mid), b-mid, c, min(d, mid-1), i-1)
		dfs(max(1, a-mid), b-mid, max(c-mid, 1), d-mid, i-1)
	}

	dfs(l1, r1, l2, r2, 30)
	return ans
}

func min(nums ...int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < res {
			res = nums[i]
		}
	}
	return res
}

func max(nums ...int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
