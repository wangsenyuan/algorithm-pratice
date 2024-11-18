package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, m, k := readThreeNums(reader)
	a := readNNums(reader, n)
	offers := make([][]int, m)
	for i := range m {
		offers[i] = readNNums(reader, 2)
	}
	return solve(a, k, offers)
}

func solve(a []int, k int, offers [][]int) int {
	sort.Ints(a)
	// 只需要考虑前k个就可以了
	save := make([]int, k+1)

	for _, cur := range offers {
		x, y := cur[0], cur[1]
		if x > k {
			continue
		}
		save[x] = max(save[x], y)
	}

	for i := 1; i <= k; i++ {
		save[i] = max(save[i], save[i-1])
	}
	// dp[i]等于节省的最大值
	dp := make([]int, k+1)
	sum := make([]int, k+1)
	for i := 0; i < k; i++ {
		dp[i+1] = dp[i]
		sum[i+1] = sum[i] + a[i]
		for j := 1; j <= i+1; j++ {
			w := save[j]
			if w == 0 {
				continue
			}
			// 比如j=2， w = 1
			// sum[i+1 - 1] - sum[i+1 - 2]
			dp[i+1] = max(dp[i+1], dp[i+1-j]+sum[i+1-(j-w)]-sum[i+1-j])
		}
	}

	return sum[k] - dp[k]
}
