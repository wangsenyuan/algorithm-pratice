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
		n, k, l := readThreeNums(reader)
		d := readNNums(reader, n)
		res := solve(n, k, l, d)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(n int, k int, l int, d []int) bool {
	// n 有 3 * 1e5
	type pair struct {
		first  int
		second int
	}

	swim := func(cur []pair, i int) []pair {
		// 从现在的区间游到下一个位置去
		if d[i] > l {
			// always fail
			return nil
		}
		x := l - d[i]
		if x == 0 {
			// 只有一个位置可以, 0
			for _, p := range cur {
				if p.second == 2*k-1 {
					// p.second = 2 * k - 1, 下个时刻，p[i] = 0, 所以，可以等等
					return []pair{{0, 0}}
				}
			}
			return nil
		}
		if x > k {
			x = k
		}
		// 安全的区域是 {0, x}, {2 * k - x, 2 * k - 1}
		var res []pair

		for _, p := range cur {
			if p.second == 2*k-1 {
				// 从当前时刻开始游泳，可以在时刻 2 * k 到达下个位置，且这个位置很安全
				res = append(res, pair{0, x})
				y := max(p.first+1, 2*k-x)
				if y <= 2*k-1 {
					res = append(res, pair{y, 2*k - 1})
				}
			} else if p.second+1 >= 2*k-x && p.second+1 <= 2*k-1 {
				res = append(res, pair{p.second + 1, 2*k - 1})
			}
			if p.first+1 <= x {
				res = append(res, pair{p.first + 1, x})
			}

		}
		sort.Slice(res, func(i, j int) bool {
			return res[i].first < res[j].first
		})

		var arr []pair

		// 连接重叠区域
		for i := 0; i < len(res); {
			a, b := res[i].first, res[i].second
			for i < len(res) && res[i].first <= b {
				b = max(b, res[i].second)
				if a <= k && b >= k {
					// 如果能够到达k，就能到达最后一个位置
					b = 2*k - 1
				}
				i++
			}
			arr = append(arr, pair{a, b})
		}

		return arr
	}

	dp := []pair{{0, 2*k - 1}}

	for i := 0; i < n; i++ {
		ndp := swim(dp, i)
		if len(ndp) == 0 {
			return false
		}
		dp = ndp
	}

	return true
}
