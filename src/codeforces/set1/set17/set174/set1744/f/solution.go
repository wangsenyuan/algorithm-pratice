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
		n := readNum(reader)
		p := readNNums(reader, n)
		res := solve(p)
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

func solve(p []int) int64 {
	n := len(p)
	pos := make([]int, n)

	for i, x := range p {
		pos[x] = i
	}

	ans := 1
	l, r := pos[0], pos[0]

	for m := 2; m <= n; m++ {
		i := pos[(m-1)/2]
		l = min(l, min(i, n-m))
		r = max(r, max(i, m-1))
		// 只要包含了区间[l...r], 且长度为m的区间，即满足 mex > median (m-1) / 2
		ans += max(m-(r-l), 0)
	}

	return int64(ans)
}

func solve1(p []int) int64 {
	n := len(p)
	if n == 1 {
		return 1
	}
	// mex(arr) > med(arr)
	//  => 假设arr的长为sz,
	//      sz 是奇数  0, 1, .. sz/2 ... 必须存在arr中
	//      sz 偶数    0, 1, ...sz/2 - 1.... 必须在arr中
	//   （0..... (sz - 1) / 2) 必须在arr中
	// 那么从0开始
	// 假设[l...r] 的 mex是 x
	// 那么下一个数，应该是 x的位置通过扩展后，x会增加，然后判断是否满足条件

	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[p[i]] = i
	}
	// [1, 0, 2]
	var ans int64

	// 当 mex = 1 时，包含0的，且长度 <= 2 的都是candidate
	// 假设 [l....r] 是包含了[0...mex-1]的最小的集合
	// sz = r - l + 1 这个集合在考虑mex时，必须是存在的
	// 然后它可以向右，或者向左延伸
	l, r := pos[0], pos[0]
	mex := 1
	vis := make([]bool, n+1)
	for mex < n {
		for vis[mex] {
			mex++
		}
		if mex == n {
			break
		}
		// vis[mex] = false
		cur := pos[mex]
		if cur < l {
			// expand left to cur position
			for l > cur {
				// 如果当前位置是l, r - l + 1 = sz => (sz - 1) / 2 <= mex
				// => sz <= 2 * mex
				// 假设 mex = 2, 那么这里 sz 最大到4， 是因为如果sz = 5, 那么 med = 2
				//     mex = 3, 那么这里sz最大到6？
				r0 := min(n, l+2*mex) - 1
				if r0 >= r {
					ans += int64(r0 - r + 1)
				}
				l--
				vis[p[l]] = true
			}
		} else {
			// expand right to cur position
			for r < cur {
				l0 := max(-1, r-2*mex) + 1
				if l0 <= l {
					ans += int64(l - l0 + 1)
				}
				r++
				vis[p[r]] = true
			}
		}
		mex++
	}

	return ans + 1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
