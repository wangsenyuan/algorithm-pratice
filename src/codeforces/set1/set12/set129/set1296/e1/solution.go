package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	s := readString(reader)
	res := solve(s[:n])
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	fmt.Println(res)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(s string) string {
	s1, s2 := int(s[0]-'a'), 0
	n := len(s)
	ans := make([]byte, n)
	ans[0] = '0'
	for i := 1; i < n; i++ {
		x := int(s[i] - 'a')
		if x >= s1 {
			s1 = x
			ans[i] = '0'
		} else if x >= s2 {
			s2 = x
			ans[i] = '1'
		} else {
			return ""
		}
	}
	return string(ans)
}

func solve1(s string) string {
	cnt, ans := solve2(s)
	if cnt > 2 {
		return ""
	}
	n := len(s)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = '0'
		if ans[i] == 2 {
			buf[i] = '1'
		}
	}

	return string(buf)
}

func solve2(s string) (int, []int) {
	n := len(s)
	fp := make([]int, 27)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
		x := int(s[i] - 'a')
		v := slices.Max(fp[x+1:]) + 1
		dp[i] = v
		fp[x] = max(fp[x], v)
	}

	return slices.Max(fp), dp
}
