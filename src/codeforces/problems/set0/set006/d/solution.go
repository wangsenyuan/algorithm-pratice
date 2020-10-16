package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, a, b := readThreeNums(reader)

	h := readNNums(reader, n)

	res := solve(n, a, b, h)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	fmt.Println(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

const INF = 10000

func solve(n int, a int, b int, h []int) []int {
	for i := 0; i < n; i++ {
		h[i]++
	}
	dp := make([][][]int, n)
	num := make([][][]int, n)
	hlh := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, h[i]+1)
		num[i] = make([][]int, h[i]+1)
		hlh[i] = make([][]int, h[i]+1)
		for j := 0; j <= h[i]; j++ {
			num[i][j] = make([]int, 20)
			dp[i][j] = make([]int, 20)
			hlh[i][j] = make([]int, 20)
			for k := 0; k < len(dp[i][j]); k++ {
				dp[i][j][k] = INF
			}
		}
	}
	dp[0][h[0]][0] = 0
	num[0][h[0]][0] = 0
	hlh[0][h[0]][0] = h[0]

	for i := 1; i < n-1; i++ {
		for j := 0; j <= h[i-1]; j++ {
			for k := 0; k < len(dp[i-1][j]); k++ {
				if dp[i-1][j][k] >= INF {
					continue
				}
				//first: need to make sure attacks at this archer can make h[i-1] to be zero
				for kk := 0; kk < 20; kk++ {
					jj := max(0, h[i]-b*k-a*kk)
					if kk*b >= j {
						if dp[i][jj][kk] > dp[i-1][j][k]+kk {
							dp[i][jj][kk] = dp[i-1][j][k] + kk
							num[i][jj][kk] = k
							hlh[i][jj][kk] = j
						}
					}
				}
			}
		}
	}

	for j := 0; j <= h[n-2]; j++ {
		for k := 0; k < len(dp[n-2][j]); k++ {
			if dp[n-2][j][k] >= INF {
				continue
			}
			if k*b >= h[n-1] {
				if dp[n-1][0][0] > dp[n-2][j][k] {
					dp[n-1][0][0] = dp[n-2][j][k]
					num[n-1][0][0] = k
					hlh[n-1][0][0] = j
				}
			}
		}
	}
	best := dp[n-1][0][0]
	res := make([]int, 0, best)
	ii := n - 1
	jj := 0
	kk := 0
	for best > 0 {
		k := num[ii][jj][kk]
		j := hlh[ii][jj][kk]
		//x := dp[ii][jj][kk] - dp[ii-1][j][k]
		for x := 0; x < k; x++ {
			res = append(res, ii)
		}
		best -= k
		ii--
		jj, kk = j, k
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
