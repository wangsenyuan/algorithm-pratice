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
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(a, k)
}

func solve(a []int, k int) int {
	n := len(a)
	sort.Ints(a)
	dp := make([][]int, k+1)
	fp := make([][]int, k+1)
	for i := range k + 1 {
		dp[i] = make([]int, 6)
		fp[i] = make([]int, 6)
	}

	gp := make([]int, k+1)
	gp[1] = 1
	dp[1][0] = 1

	for i := 1; i < n; i++ {
		for j := 1; j <= k && j <= i+1; j++ {
			fp[j][0] = gp[j-1] + 1

			for d := 0; d <= 5; d++ {
				if a[i]-a[i-1]+d <= 5 {
					fp[j][a[i]-a[i-1]+d] = max(fp[j][a[i]-a[i-1]+d], dp[j][d]+1)
				}
			}
		}
		for j := 0; j <= k && j <= i+1; j++ {
			for d := 0; d <= 5; d++ {
				gp[j] = max(gp[j], fp[j][d])
				dp[j][d] = fp[j][d]
				fp[j][d] = 0
			}
		}
	}
	return gp[k]
}
