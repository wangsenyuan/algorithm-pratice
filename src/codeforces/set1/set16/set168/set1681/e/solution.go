package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	D := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		D[i] = readNNums(reader, 4)
	}
	m := readNum(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 4)
	}

	res := solve(n, D, Q)
	var buf bytes.Buffer

	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const H = 20
const INF = 1e18

func solve(n int, D [][]int, Q [][]int) []int64 {
	// dp[i][j]
	// dp[i][0] 表示从(1, 1)到达第i层top行的最短距离
	// dp[i][1] ..                right行的  距离
	// 那这里有个问题，如果(x1, y1)靠近到(x2, y2) 所在层的最优路径上
	// 假设要到达max(x2, y2) - 1的top的门，必须经过max(x1, y1)top的门，且 (x1,y1)距离这个门更近
	// 那么结果就是 dp[j-1][0] - dp[i][0] + a + b
	// 但是如果dp[j-1][0]的最优路径是经过 dp[i][1], 就比较难处理了
	// 需要构成一棵树

	for i := 0; i < len(D); i++ {
		for j := 0; j < 4; j++ {
			D[i][j]--
		}
	}

	dp := make([][][][]int64, n-1)
	for i := 0; i < n-1; i++ {
		dp[i] = make([][][]int64, H)
		for d := 0; d < H; d++ {
			// 00 , 01, 10, 11
			dp[i][d] = make([][]int64, 2)
			for j := 0; j < 2; j++ {
				dp[i][d][j] = make([]int64, 2)
				for k := 0; k < 2; k++ {
					dp[i][d][j][k] = INF
				}
			}
		}
	}

	for i := 0; i < n-2; i++ {
		for k := 0; k < 2; k++ {
			dp[i][0][0][k] = distance(top(D[i][:2]), D[i+1][k*2:(k+1)*2]) + 1
			dp[i][0][1][k] = distance(right(D[i][2:]), D[i+1][k*2:(k+1)*2]) + 1
		}
	}

	for h := 1; h < H; h++ {
		for i := 0; i+(1<<(h-1)) < n-2; i++ {
			for j := 0; j < 2; j++ {
				for k := 0; k < 2; k++ {
					for t := 0; t < 2; t++ {
						dp[i][h][j][k] = min(dp[i][h][j][k], dp[i][h-1][j][t]+dp[i+(1<<(h-1))][h-1][t][k])
					}
				}
			}
		}
	}

	var query func(x1, y1, x2, y2 int) int64
	query = func(x1, y1, x2, y2 int) int64 {
		l1 := max(x1, y1)
		l2 := max(x2, y2)
		if l1 > l2 {
			return query(x2, y2, x1, y1)
		}
		if l1 == l2 {
			return distance([]int{x1, y1}, []int{x2, y2})
		}
		// l1 < l2
		ndp := []int64{distance(D[l1][:2], []int{x1, y1}), distance(D[l1][2:], []int{x1, y1})}
		for i := H - 1; i >= 0; i-- {
			if l1+(1<<i) < l2 {
				tmp := []int64{INF, INF}
				for j := 0; j < 2; j++ {
					for k := 0; k < 2; k++ {
						tmp[k] = min(tmp[k], ndp[j]+dp[l1][i][j][k])
					}
				}
				copy(ndp, tmp)
				l1 += (1 << i)
			}
		}
		ans := ndp[0] + distance(top(D[l1][:2]), []int{x2, y2})
		ans = min(ans, ndp[1]+distance(right(D[l1][2:]), []int{x2, y2}))
		return ans + 1
	}

	res := make([]int64, len(Q))
	for i, cur := range Q {
		for j := 0; j < 4; j++ {
			cur[j]--
		}
		res[i] = query(cur[0], cur[1], cur[2], cur[3])
	}

	return res
}

func top(coord []int) []int {
	res := []int{coord[0] + 1, coord[1]}
	return res
}

func right(coord []int) []int {
	return []int{coord[0], coord[1] + 1}
}

func distance(a []int, b []int) int64 {
	return int64(abs(b[1]-a[1]) + abs(a[0]-b[0]))
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Num interface {
	int | int64
}

func max[T Num](arr ...T) T {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}

func min[T Num](arr ...T) T {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < res {
			res = arr[i]
		}
	}
	return res
}
