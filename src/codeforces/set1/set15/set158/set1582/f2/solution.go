package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	s := fmt.Sprintf("%v", res)
	buf.WriteString(fmt.Sprintf("%s\n", s[1:len(s)-1]))
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

const H = 13

const inf = 1 << 30

func solve(a []int) []int {
	dp := make([]int, 1<<H)
	for i := 0; i < len(dp); i++ {
		dp[i] = inf
	}
	dp[0] = -1

	v := make([][]int, 1<<H)
	for i, num := range a {
		v[num] = append(v[num], i)
	}

	for i := 1; i < 1<<H; i++ {
		for j := 0; j < 1<<H; j++ {
			// 这里迭代i是递增的，所以找到的 dp[j](位置)其数值肯定是 < i 的
			p := sort.SearchInts(v[i], dp[j])
			if p == len(v[i]) {
				continue
			}
			dp[j^i] = min(dp[j^i], v[i][p])
		}
	}

	var ans []int

	for i := 0; i < 1<<H; i++ {
		if dp[i] < inf {
			ans = append(ans, i)
		}
	}

	return ans
}
func solve1(a []int) []int {
	ok := make([]bool, 1<<H)
	ok[0] = true
	g := make([][]int16, 1<<H)
	r := make([]int16, 1<<H)
	for i := 0; i < len(g); i++ {
		g[i] = append(g[i], 0)
		r[i] = 1 << H
	}

	for _, x := range a {
		for _, k := range g[x] {
			to := k ^ int16(x)
			ok[to] = true
			for r[to] > int16(x) {
				r[to]--
				if r[to] > int16(x) {
					g[r[to]] = append(g[r[to]], to)
				}
			}
		}
		g[x] = make([]int16, 0, 1)
	}

	var ans []int
	for i := 0; i < 1<<H; i++ {
		if ok[i] {
			ans = append(ans, i)
		}
	}
	return ans
}
