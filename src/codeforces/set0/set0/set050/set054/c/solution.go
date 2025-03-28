package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.11f\n", res)
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

func process(reader *bufio.Reader) float64 {
	n := readNum(reader)
	varirables := make([][]int, n)
	for i := range n {
		varirables[i] = readNNums(reader, 2)
	}
	k := readNum(reader)
	return solve(varirables, k)
}

func solve(variables [][]int, k int) float64 {
	n := len(variables)
	f := make([]float64, n)
	for i, cur := range variables {
		l, r := cur[0], cur[1]
		cnt := calc(r) - calc(l-1)
		f[i] = float64(cnt) / float64(r-l+1)
	}

	dp := make([]float64, n+1)
	dp[0] = 1

	for i := range n {
		for j := i + 1; j > 0; j-- {
			dp[j] = dp[j]*(1-f[i]) + dp[j-1]*f[i]
		}
		dp[0] *= 1 - f[i]
	}

	var res float64

	for j := n; j >= 0; j-- {
		if j*100 < n*k {
			break
		}
		res += dp[j]
	}

	return res
}

func calc(n int) int {
	var res int
	p10 := 1
	x := n
	for ; x >= 10; x /= 10 { // 正反算
		res += p10
		p10 *= 10
	}
	if x > 1 {
		res += p10
	} else if x == 1 {
		res += n - p10 + 1
	}
	return res
}
