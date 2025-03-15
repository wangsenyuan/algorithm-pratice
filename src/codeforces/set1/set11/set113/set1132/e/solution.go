package main

import (
	"bufio"
	"fmt"
	"os"
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
	w := readNum(reader)
	cnt := readNNums(reader, 8)
	return solve(w, cnt)
}

const N = 9

// lcm of (1...8)
const L = 840

func solve(W int, cnt []int) int {
	var dp [N][L * N]int
	for i := range dp {
		for j := range L * N {
			dp[i][j] = -1
		}
	}
	// dp[i][j] 表示前i个，选择的总重量为j, 能够选择的最大的L的个数
	dp[0][0] = 0
	for i := range 8 {
		for j := 0; j < L*N; j++ {
			if dp[i][j] < 0 {
				continue
			}
			// 最多可以选这么重的
			b := min(cnt[i], L/(i+1))
			for k := 0; k <= b; k++ {
				dp[i+1][j+k*(i+1)] = max(dp[i+1][j+k*(i+1)], dp[i][j]+(cnt[i]-k)/(L/(i+1)))
			}
		}
	}
	var ans int
	for j := 0; j < L*N; j++ {
		if j > W || dp[8][j] < 0 {
			continue
		}
		ans = max(ans, j+L*min(dp[8][j], (W-j)/L))
	}
	return ans
}
