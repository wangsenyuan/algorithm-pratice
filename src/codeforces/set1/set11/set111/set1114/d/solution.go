package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	fmt.Println(res)
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

func solve(a []int) int {
	n := len(a)
	var arr []int
	for i := 0; i < n; {
		arr = append(arr, a[i])
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
	}
	n = len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			dp[i][j] = n
		}
	}

	for r := 0; r < n; r++ {
		for l := r - 1; l >= 0; l-- {
			if arr[l] == arr[r] {
				dp[l][r] = dp[l+1][r-1] + 1
			} else {
				dp[l][r] = min(dp[l+1][r], dp[l][r-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}