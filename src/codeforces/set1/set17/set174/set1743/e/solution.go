package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := readNInt64s(reader, 2)
	b := readNInt64s(reader, 2)
	h, s := readTwoNums(reader)
	res := solve(a, b, h, s)

	fmt.Println(res)
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
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

const inf = 1 << 60

func solve(a []int64, b []int64, h int, s int) int64 {
	// dp[x] 表示摧毁x单位，所需的最少时间
	ps := []int64{a[0], b[0]}
	ts := []int64{a[1], b[1]}

	dp := make([]int64, h+1)
	for i := 0; i <= h; i++ {
		dp[i] = inf
	}
	dp[0] = 0

	var ans int64 = inf

	x := int64(s)

	for i := 0; i < h; i++ {
		// dp[i]表示两个laser同时可用
		// j是shot laser k的次数
		for j := 1; j <= h-i; j++ {
			for k := 0; k < 2; k++ {
				ni := min(int64(h), int64(i)+int64(j)*(ps[k]-x)+int64(j)*ts[k]/ts[k^1]*(ps[k^1]-x))
				ii := int(ni)
				if ii == h {
					ans = min(ans, dp[i]+int64(j)*ts[k])
				}
				// 最后一次让两个同时发射
				if int64(j)*ts[k] >= ts[k^1] {
					ni = min(int64(h), int64(i)+int64(j-1)*(ps[k]-x)+(int64(j)*ts[k]/ts[k^1]-1)*(ps[k^1]-x)+(ps[0]+ps[1]-x))
					ii = int(ni)
					dp[ii] = min(dp[ii], dp[i]+int64(j)*ts[k])
				}
			}
		}
	}

	return min(ans, dp[h])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(arr ...int64) int64 {
	res := arr[0]
	for i := 0; i < len(arr); i++ {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}
