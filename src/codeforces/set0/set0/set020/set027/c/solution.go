package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := process(reader)
	fmt.Println(len(res))
	if len(res) > 0 {
		fmt.Println(res[0], res[1], res[2])
	}
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

func process(reader *bufio.Reader) (a []int, res []int) {
	n := readNum(reader)
	a = readNNums(reader, n)
	res = solve(a)
	return
}

func solve(a []int) []int {
	if sort.IntsAreSorted(a) {
		return nil
	}
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] *= -1
	}
	if sort.IntsAreSorted(a) {
		return nil
	}
	for i := 0; i < n; i++ {
		a[i] *= -1
	}
	// 前面的最大值，最小值
	prev := []int{0, 0}
	dp := make([][]int, n)
	for i := 1; i < n; i++ {
		dp[i] = make([]int, 2)
		copy(dp[i], prev)
		if a[i] > a[prev[0]] {
			prev[0] = i
		}
		if a[i] < a[prev[1]] {
			prev[1] = i
		}
	}
	suf := []int{n - 1, n - 1}
	for i := n - 2; i > 0; i-- {
		x := a[i]
		if a[dp[i][1]] < x && a[suf[1]] < x {
			return []int{dp[i][1] + 1, i + 1, suf[1] + 1}
		}
		if a[dp[i][0]] > x && a[suf[0]] > x {
			return []int{dp[i][0] + 1, i + 1, suf[0] + 1}
		}
		if x > a[suf[0]] {
			suf[0] = i
		}
		if x < a[suf[1]] {
			suf[1] = i
		}
	}

	return nil
}
