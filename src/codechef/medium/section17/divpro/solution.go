package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		nums := readNInt64s(reader, 2)
		res := solve(int(nums[0]), nums[1])
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const MAX_L = 36
const MAX_L1 = (MAX_L + 1) / 2

var F = []int{2, 3, 5, 7}
var N = []int{3 * MAX_L1, 2 * MAX_L1, MAX_L1, MAX_L1}

var G = [][]int{
	{0, 0, 1, 0, 2, 0, 1, 0, 3, 0},
	{0, 0, 0, 1, 0, 0, 1, 0, 0, 2},
	{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
}

var pw [][]int64

var ans [MAX_L + 1][][]int64

func init() {
	pw = make([][]int64, 4)
	for i := 0; i < 4; i++ {
		pw[i] = make([]int64, N[i]+1)
		pw[i][0] = 1
		for j := 1; j <= N[i]; j++ {
			pw[i][j] = int64(F[i]) * pw[i][j-1]
		}
	}
	dp := make([][][][][]uint, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][][][]uint, 2*N[0]+1)
		for a := 0; a < len(dp[i]); a++ {
			dp[i][a] = make([][][]uint, 2*N[1]+1)
			for b := 0; b < len(dp[i][a]); b++ {
				dp[i][a][b] = make([][]uint, 2*N[2]+1)
				for c := 0; c < len(dp[i][a][b]); c++ {
					dp[i][a][b][c] = make([]uint, 2*N[3]+1)
				}
			}
		}
	}

	var p int
	dp[0][N[0]][N[1]][N[2]][N[3]] = 1

	for l := 1; l <= MAX_L; l++ {
		ans[l] = make([][]int64, 0, 1)
		p ^= 1
		l1 := l / 2
		l2 := (l - 1) - l1
		for a := 0; a < len(dp[p]); a++ {
			for b := 0; b < len(dp[p][a]); b++ {
				for c := 0; c < len(dp[p][a][b]); c++ {
					for d := 0; d < len(dp[p][a][b][c]); d++ {
						dp[p][a][b][c][d] = 0
					}
				}
			}
		}
		for n2 := -3 * l2; n2 <= 3*l1; n2++ {
			for n3 := -2 * l2; n3 <= 2*l1; n3++ {
				for n5 := -l2; n5 <= l1; n5++ {
					for n7 := -l2; n7 <= l1; n7++ {
						cnt := dp[1^p][N[0]+n2][N[1]+n3][N[2]+n5][N[3]+n7]
						if cnt == 0 {
							continue
						}
						for d := 1; d <= 9; d++ {
							dp[p][N[0]-n2+G[0][d]][N[1]-n3+G[1][d]][N[2]-n5+G[2][d]][N[3]-n7+G[3][d]] += cnt
						}
					}
				}
			}
		}
		l1 = (l + 1) / 2
		for n2 := 0; n2 <= 3*l1; n2++ {
			for n3 := 0; n3 <= 2*l1; n3++ {
				for n5 := 0; n5 <= l1; n5++ {
					for n7 := 0; n7 <= l1; n7++ {
						cnt := dp[p][N[0]+n2][N[1]+n3][N[2]+n5][N[3]+n7]
						if cnt > 0 {
							v := pw[0][n2] * pw[1][n3] * pw[2][n5] * pw[3][n7]
							ans[l] = append(ans[l], []int64{v, int64(cnt)})
						}
					}
				}
			}
		}
		sort.Slice(ans[l], func(i, j int) bool {
			return ans[l][i][0] < ans[l][j][0]
		})
	}
}

func solve(L int, V int64) uint {
	if V == 0 {
		return solve0(L)
	}

	i := sort.Search(len(ans[L]), func(i int) bool {
		return ans[L][i][0] >= V
	})
	if i < len(ans[L]) && ans[L][i][0] == V {
		return uint(ans[L][i][1])
	}

	return 0
}

func solve0(L int) uint {
	x := (L - 1) / 2
	a := pow(9, L-x)
	b := pow(10, x)
	c := pow(9, x)
	b -= c
	if b < 0 {
		b += MOD
	}
	a *= b
	a %= MOD
	return uint(a)
}

const MOD int64 = 1 << 32

func pow(a, b int) int64 {
	A := int64(a)
	var res int64 = 1
	for b != 0 {
		if b&1 == 1 {
			res *= A
			res %= MOD
		}
		A *= A
		A %= MOD
		b >>= 1
	}
	return res
}
