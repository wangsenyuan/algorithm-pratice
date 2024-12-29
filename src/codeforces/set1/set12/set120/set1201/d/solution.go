package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

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

func process(reader *bufio.Reader) int {
	first := readNNums(reader, 4)
	n, m, k, q := first[0], first[1], first[2], first[3]
	treasures := make([][]int, k)
	for i := range k {
		treasures[i] = readNNums(reader, 2)
	}
	safe := readNNums(reader, q)
	return solve(n, m, treasures, safe)
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 60

func solve(n int, m int, treasures [][]int, safe []int) int {
	// prev[i] 表示第i列前面最近的safe col
	prev := make([]int, m+1)
	next := make([]int, m+1)
	sort.Ints(safe)

	for i, j := 1, 0; i <= m; i++ {
		if j < len(safe) && safe[j] == i {
			prev[i] = i
			j++
		} else if j > 0 {
			prev[i] = safe[j-1]
		} else {
			prev[i] = 0
		}
	}

	for i, j := m, len(safe)-1; i > 0; i-- {
		if j >= 0 && safe[j] == i {
			next[i] = i
			j--
		} else if j+1 < len(safe) {
			next[i] = safe[j+1]
		} else {
			next[i] = m + 1
		}
	}

	xs := make([][]int, n+1)
	at := make([][]int, n+1)

	for _, tr := range treasures {
		i, j := tr[0], tr[1]
		xs[i] = append(xs[i], j)
		x := prev[j]
		if x > 0 {
			at[i] = append(at[i], x)
		}
		if next[j] != x && next[j] <= m {
			at[i] = append(at[i], next[j])
		}
	}

	for i := 1; i <= n; i++ {
		sort.Ints(xs[i])
		sort.Ints(at[i])
		var p int
		for j := 1; j <= len(at[i]); j++ {
			if j == len(at[i]) || at[i][j] > at[i][j-1] {
				at[i][p] = at[i][j-1]
				p++
			}
		}
		at[i] = at[i][:p]
	}

	for len(at[n]) == 0 {
		n--
	}

	if n == 1 {
		return xs[1][len(xs[1])-1] - 1
	}

	sortAndMerge := func(arr []pair) []pair {
		slices.SortFunc(arr, func(a, b pair) int {
			return a.first - b.first
		})

		var p int
		for i := 0; i < len(arr); {
			j := arr[i].first
			k := arr[i].second
			for i < len(arr) && arr[i].first == j {
				k = min(k, arr[i].second)
				i++
			}
			arr[p] = pair{j, k}
			p++
		}
		return arr[:p]
	}

	var dp []pair

	if len(at[1]) > 0 {
		r := xs[1][len(xs[1])-1]
		for i := 0; i < len(at[1]); i++ {
			x := at[1][i]
			// 这里x不可能为0
			dp = append(dp, pair{x, r - 1 + abs(r-x)})
		}
		dp = sortAndMerge(dp)
	} else {
		dp = append(dp, pair{next[1], next[1] - 1})
	}
	// 每一行之需要处理，离treasure最近的那些点

	for r := 2; r <= n; r++ {
		// 先找到这行重要的点
		cur := at[r]
		if r == n {
			// 最后一行只需要停在某一个宝藏上即可
			cur = xs[r]
		}
		if len(cur) == 0 {
			// 全部dp+1
			for i := 0; i < len(dp); i++ {
				dp[i].second += 1
			}
			continue
		}

		var fp []pair
		// 如果是左边的safe column
		best := inf
		lx, rx := xs[r][0], xs[r][len(xs[r])-1]
		for i, j := 0, 0; i < len(cur); i++ {
			x := cur[i]
			for j < len(dp) && dp[j].first <= x {
				// 如果从
				// 先从下一行上来，然后移动到lx，再移动最远边
				// 再从最远边到达目前的位置
				tmp := dp[j].second + 1 + abs(lx-dp[j].first) + rx - lx
				best = min(best, tmp)
				j++
			}
			// rx < x 是有可能的
			fp = append(fp, pair{x, best + abs(rx-x)})
		}
		// 然后考虑右边的点
		best = inf
		for i, j := len(cur)-1, len(dp)-1; i >= 0; i-- {
			x := cur[i]
			for j >= 0 && dp[j].first >= x {
				tmp := dp[j].second + 1 + abs(rx-dp[j].first) + rx - lx
				best = min(best, tmp)
				j--
			}
			// x < lx 也是有可能的
			fp = append(fp, pair{x, best + abs(x-lx)})
		}

		dp = sortAndMerge(fp)
	}

	res := inf
	for _, v := range dp {
		res = min(res, v.second)
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
