package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	r := readNNums(reader, n)
	b := readNNums(reader, m)
	g := readNNums(reader, k)

	res := solve(r, b, g)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const inf = 1 << 50

func solve(r []int, b []int, g []int) int {
	sort.Ints(r)
	reverse(r)
	sort.Ints(b)
	reverse(b)
	sort.Ints(g)
	reverse(g)

	nr := len(r)
	nb := len(b)
	ng := len(g)
	dp := make([][]int, nr+1)
	fp := make([][]int, nr+1)
	for i := 0; i <= nr; i++ {
		dp[i] = make([]int, nb+1)
		fp[i] = make([]int, nb+1)
		for j := 0; j <= nb; j++ {
			dp[i][j] = -inf
			fp[i][j] = -inf
		}
	}

	dp[0][0] = 0

	for i := 2; i <= nr+nb+ng; i += 2 {
		for j := 0; j <= nr; j++ {
			// 红色的用了i个
			for k := 0; k <= nb && j+k <= i; k++ {
				//绿色的用了k个
				// 蓝色的用l个
				l := i - j - k
				if l > ng {
					continue
				}
				if l < 0 {
					break
				}
				if j > 0 && k > 0 {
					fp[j][k] = max(fp[j][k], dp[j-1][k-1]+r[j-1]*b[k-1])
				}
				if j > 0 && l > 0 {
					fp[j][k] = max(fp[j][k], dp[j-1][k]+r[j-1]*g[l-1])
				}
				if k > 0 && l > 0 {
					fp[j][k] = max(fp[j][k], dp[j][k-1]+b[k-1]*g[l-1])
				}
			}
		}

		for j := 0; j <= nr; j++ {
			for k := 0; k <= nb; k++ {
				dp[j][k] = max(dp[j][k], fp[j][k])
				fp[j][k] = -inf
			}
		}
	}

	var best int
	for i := 0; i <= nr; i++ {
		for j := 0; j <= nb; j++ {
			best = max(best, dp[i][j])
		}
	}
	return best
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
