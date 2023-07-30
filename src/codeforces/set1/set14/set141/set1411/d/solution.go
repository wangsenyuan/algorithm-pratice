package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	x, y := readTwoNums(reader)
	res := solve(s, x, y)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const inf = math.MaxInt64 >> 1

func solve(s string, x int, y int) int64 {
	if x > y {
		return solve(flip(s), y, x)
	}
	// x <= y
	// so prefer 01 always
	cnt := make([]int64, 2)
	var one int64
	var sum int64
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			cnt[1]++
			one++
		} else {
			cnt[0]++
			sum += one
		}
	}
	ans := cnt[0]*cnt[1]*int64(x) + int64(y-x)*sum
	var zero int64
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			zero++
		} else if s[i] == '1' {
			one--
		} else {
			cnt[0]--
			cnt[1]++
			sum -= one
			sum += zero
			ans = min(ans, cnt[0]*cnt[1]*int64(x)+int64(y-x)*sum)
		}
	}
	return ans
}

func flip(s string) string {
	buf := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			buf[i] = '1'
		} else if s[i] == '1' {
			buf[i] = '0'
		}
	}
	return string(buf)
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
