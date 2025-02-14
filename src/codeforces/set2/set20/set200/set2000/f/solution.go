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
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	rects := make([][]int, n)
	for i := range n {
		rects[i] = readNNums(reader, 2)
	}
	return solve(k, rects)
}

const H = 101

const inf = 1 << 60

func solve(k int, rects [][]int) int {
	dp := make([]int, k+1)
	fp := make([]int, k+1)
	for i := range k + 1 {
		dp[i] = inf
		fp[i] = inf
	}
	dp[0] = 0

	for _, rect := range rects {
		w, h := rect[0], rect[1]

		for x, v := range dp {
			if v == inf {
				break
			}
			a, b := w, h
			var sum int
			for c := 1; c <= w+h && x+c <= k; c++ {
				if a <= b {
					sum += a
					b--
				} else {
					sum += b
					a--
				}

				fp[x+c] = min(fp[x+c], v+sum)
			}
		}

		for x := 0; x <= k; x++ {
			dp[x] = min(dp[x], fp[x])
			fp[x] = inf
		}
	}

	if dp[k] == inf {
		return -1
	}

	return dp[k]
}
