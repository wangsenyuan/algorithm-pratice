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
		l, r := readTwoNums(reader)
		res := solve(l, r)
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func repeat(num int, times int) int {
	var res int
	for times > 0 {
		res = res*10 + num
		times--
	}
	return res
}

func solve(l int, r int) int {
	if l == r {
		return l
	}
	// 最不lucky的数是0
	// 最lucky的数是8
	// dp[i][j] 等于到目前为止是否能得到最大值是i,最小值是j的数
	// 还需要知道是否在l...r的范围内
	// dp[i][j][x][y] x = 1 表示目前的数 == r, = 0 => < r
	//                y = 1, 表示目前的数 == l, = 0 => > l
	a := toDigits(l)
	b := toDigits(r)

	// 当两个数的长度不一致的时候
	// 比如 99 100, 那么就取 num = 9999即可呐
	if len(a) < len(b) {
		return repeat(9, len(a))
	}
	// len(a) = len(b)

	h := len(a)

	dp := make([][][][]bool, 10)
	fp := make([][][][]bool, 10)
	for i := 0; i < 10; i++ {
		dp[i] = make([][][]bool, 10)
		fp[i] = make([][][]bool, 10)
		for j := 0; j < 10; j++ {
			dp[i][j] = make([][]bool, 2)
			fp[i][j] = make([][]bool, 2)
			for k := 0; k < 2; k++ {
				dp[i][j][k] = make([]bool, 2)
				fp[i][j][k] = make([]bool, 2)
			}
		}
	}
	dp[0][9][1][1] = true

	// 这里还有个问题没有解决
	for d := h - 1; d >= 0; d-- {
		for i := 9; i >= 0; i-- {
			for j := 9; j >= 0; j-- {
				// i >= j
				for u := 0; u <= 1; u++ {
					for v := 0; v <= 1; v++ {
						if !dp[i][j][u][v] {
							continue
						}
						// 现在要取新的值x
						start := 0
						if v == 1 {
							start = a[d]
						}
						end := 9
						if u == 1 {
							end = b[d]
						}

						for x := start; x <= end; x++ {
							ni := max(i, x)
							nj := min(j, x)
							nu := 0
							if u == 1 && x == b[d] {
								nu = 1
							}
							nv := 0
							if v == 1 && x == a[d] {
								nv = 1
							}
							fp[ni][nj][nu][nv] = true
						}
					}
				}
			}
		}

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				for u := 0; u < 2; u++ {
					for v := 0; v < 2; v++ {
						dp[i][j][u][v] = fp[i][j][u][v]
						fp[i][j][u][v] = false
					}
				}
			}
		}
	}
	// 还要恢复这个数呢～
	best := []int{9, 0}
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			for u := 0; u < 2; u++ {
				for v := 0; v < 2; v++ {
					if dp[i][j][u][v] {
						if i-j < best[0]-best[1] {
							best[0] = i
							best[1] = j
						}
					}
				}
			}
		}
	}

	if best[0] == best[1] {
		return repeat(best[0], len(a))
	}

	// 最大值，最小值确定了,从前往后
	// eq = r so for?
	var dfs func(d int, el bool, er bool, num int) int

	dfs = func(d int, el bool, er bool, num int) int {
		if d < 0 {
			return num
		}

		start, end := best[1], best[0]
		if el {
			start = max(best[1], a[d])
		}
		if er {
			end = min(best[0], b[d])
		}

		// dead end
		if start > end {
			return -1
		}

		if !el && !er {
			// 已经没有限制了
			return dfs(d-1, false, false, num*10+start)
		}

		if start+1 < end {
			//如果能在这里取中间的结果，最好
			return dfs(d-1, false, false, num*10+start+1)
		}

		if !er && start+1 <= end {
			// 否则如果能刚好取比start 大1的数
			return dfs(d-1, false, false, num*10+start+1)
		}

		if !el && end-1 >= start {
			return dfs(d-1, false, false, num*10+end-1)
		}

		ans := dfs(d-1, el && start == a[d], er && start == b[d], num*10+start)
		if ans >= 0 {
			return ans
		}

		return dfs(d-1, el && end == a[d], er && end == b[d], num*10+end)
	}

	return dfs(h-1, true, true, 0)
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

func toDigits(num int) []int {
	var ds []int
	for num > 0 {
		ds = append(ds, num%10)
		num /= 10
	}

	return ds
}
