package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	S, _ := reader.ReadString('\n')
	solver := solve(n, k, S)
	q := readNum(reader)
	var buf bytes.Buffer
	for q > 0 {
		q--
		tmp, _ := reader.ReadString('\n')
		ans := solver.Query(len(tmp)-1, tmp)
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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

type Solver struct {
	dp [][]int64
	k  int
}

func solve(n, k int, S string) *Solver {
	segs := make([][]Pair, k)
	for i := 0; i < k; i++ {
		segs[i] = make([]Pair, 0, 2)
	}

	for i, p := 1, 0; i <= n; i++ {
		if i == n || S[i] != S[i-1] {
			j := int(S[i-1] - 'a')
			segs[j] = append(segs[j], Pair{p, i})
			p = i
		}
	}

	K := 1 << uint(k)
	dp := make([][]int64, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int64, K)
	}
	last := make([]int, k)
	prev := make([]int64, n+1)
	for i := 0; i < k; i++ {
		if len(segs[i]) == 0 {
			continue
		}
		sz := len(segs[i])
		for j := 0; j < k; j++ {
			last[j] = -1
		}
		for j := 0; j < sz; j++ {
			prev[j] = int64(segs[i][j].second - segs[i][j].first)
			dp[i][0] += prev[j] * prev[j]
			if j > 0 {
				prev[j] += prev[j-1]
			}
		}

		for j := 1; j < sz; j++ {
			var mask int
			for u := segs[i][j-1].second; u < segs[i][j].first; u++ {
				mask |= 1 << uint(S[u]-'a')
			}
			for b := 0; b < k; b++ {
				if (mask>>uint(b))&1 == 1 {
					last[b] = j - 1
				}
			}
			cur := j - 1
			l := int64(segs[i][j].second - segs[i][j].first)
			for cur != -1 {
				tmp := cur
				cur = -1
				for b := 0; b < k; b++ {
					if last[b] < tmp {
						cur = max(cur, last[b])
					}
				}
				if cur >= 0 {
					dp[i][mask] += 2 * l * (prev[tmp] - prev[cur])
				} else {
					dp[i][mask] += 2 * l * prev[tmp]
				}
				for b := 0; b < k; b++ {
					if last[b] == cur {
						mask |= 1 << uint(b)
					}
				}
			}
		}

		for j := 0; j < k; j++ {
			for mask := 0; mask < K; mask++ {
				if (mask>>uint(j))&1 == 1 {
					dp[i][mask] += dp[i][mask^(1<<uint(j))]
				}
			}
		}
	}
	return &Solver{dp, k}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

func (solver *Solver) Query(n int, s string) int64 {
	var mask int
	for i := 0; i < n; i++ {
		mask |= 1 << uint(s[i]-'a')
	}
	var ans int64
	for j := 0; j < solver.k; j++ {
		if (mask>>uint(j))&1 == 0 {
			ans += solver.dp[j][mask]
		}
	}

	return ans
}
