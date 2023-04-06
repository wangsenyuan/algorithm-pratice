package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		S := readString(reader)
		res := solve(n, S)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, S string) int64 {
	// D的数量也有关系，还有R的数量
	// dp[i][j][k] = 使用前S[...k]能够到达(i, j)
	// 如果 S[k+1] = R => dp[i][j+1][k+1] = true
	// 如果 S[k] = R => dp[i][j+1][k] = true
	// 如果  = D => dp[i+1][j][k(+1)] = true
	// 显然没法所有的格子都去扫一遍
	// 首先S[m]（最后一个)如果是D，那么只能达到最右边一列后，然后不断向下
	// 要到达最右边，需要有n-1个R，假设现在有x个R，也就是需要添加 （n - 1 - x)个R
	// 在x个R的后面，连续的R是不是要当做一个来看待的？
	// 假设有y个位置， y ** (n - 1 - x)
	// 然后考虑添加D的数量，最后的D和其他的D不一样
	m := len(S)
	// always go down at last step
	// cnt[0] for R, cnt[1] for D
	cnt := make([]int, 2)

	for i := 0; i < m; i++ {
		if S[i] == 'R' {
			cnt[0]++
		} else {
			cnt[1]++
		}
	}
	if cnt[0] == 0 || cnt[0] == m {
		return int64(n)
	}

	count := func(S string) int64 {
		ld := find(S, 'R')
		sum := int64(ld) * int64(n-1)
		var y int64
		for i := m - 1; i >= ld; i-- {
			if S[i] == 'D' {
				sum += y
			} else {
				y++
			}
		}
		return sum
	}

	res := int64(n) * int64(n)

	// normal one
	res -= count(S)
	res -= count(flip(S))
	return res
}

func flip(s string) string {
	buf := []byte(s)
	for i := 0; i < len(s); i++ {
		if buf[i] == 'R' {
			buf[i] = 'D'
		} else {
			buf[i] = 'R'
		}
	}
	return string(buf)
}

func find(s string, x byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == x {
			return i
		}
	}
	return -1
}
