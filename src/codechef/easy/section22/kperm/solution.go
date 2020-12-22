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
		s, _ := reader.ReadBytes('\n')
		var n, r int
		pos := readInt(s, 0, &n) + 1
		pos = readInt(s, pos, &r) + 1
		var k uint64
		readUint64(s, pos, &k)

		res := solve(n, r, int64(k))
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
	}
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const MAX_N = 505
const MAX_R = 200005
const INF int64 = 1e18

var dp [MAX_N][MAX_R]int64

func init() {
	dp[1][0] = 1

	for i := 2; i < MAX_N; i++ {
		i2 := i * (i - 1) / 2
		for j := 0; j < i2+2; j++ {
			var sum int64
			for k := 0; k <= j; k++ {
				if sum > INF {
					sum = -1
					break
				}
				if dp[i-1][j-k] == -1 {
					sum = -1
					break
				}
				sum += dp[i-1][j-k]
			}
			if sum > INF {
				sum = -1
			}
			dp[i][j] = sum
		}
		//dp[i][0] = -1
	}
}

func solve(n int, r int, k int64) []int {
	if dp[n][r] > -1 && k > dp[n][r] {
		return nil
	}

	ans := make([]int, 0, n)

	for i := n; i >= 1; i-- {
		if i == 1 {
			ans = append(ans, r)
			break
		}
		for j := 0; j < i; j++ {
			if dp[i-1][r] == -1 || dp[i-1][r] >= k {
				ans = append(ans, j)
				break
			}
			k -= dp[i-1][r]
			r--
		}
	}
	res := make([]int, 0, n)
	taken := make([]bool, n)
	for i := 0; i < n; i++ {
		var ctr int
		for j := 0; j < n; j++ {
			if !taken[j] {
				ctr++
			}
			if ctr > ans[i] {
				taken[j] = true
				res = append(res, j+1)
				break
			}
		}
	}
	return res
}
