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

const H = 9

const inf = 1 << 30

func solve(a []int) []int {
	// 最多是 512 - 1
	dp := make([]int, 1<<H)
	for i := 0; i < len(dp); i++ {
		dp[i] = inf
	}
	dp[0] = 0
	fp := make([]int, 1<<H)
	for _, num := range a {
		copy(fp, dp)
		for x := 0; x < 1<<H; x++ {
			if fp[x] < num {
				dp[x^num] = min(dp[x^num], num)
			}
		}
	}
	var ans []int
	for x := 0; x < 1<<H; x++ {
		if dp[x] < inf {
			ans = append(ans, x)
		}
	}
	return ans
}
