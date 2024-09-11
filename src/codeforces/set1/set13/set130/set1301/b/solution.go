package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

const inf = 1 << 30

func solve(a []int) []int {
	mn, mx := inf, -inf

	for i, num := range a {
		if num == -1 {
			if i > 0 && a[i-1] != -1 {
				mn = min(mn, a[i-1])
				mx = max(mx, a[i-1])
			}
			if i+1 < len(a) && a[i+1] != -1 {
				mn = min(mn, a[i+1])
				mx = max(mx, a[i+1])
			}
		}
	}
	k := (mn + mx) / 2

	get := func(i int) int {
		if a[i] < 0 {
			return k
		}
		return a[i]
	}

	var m int
	for i := 1; i < len(a); i++ {
		v := get(i)
		w := get(i - 1)
		m = max(m, abs(v-w))
	}

	return []int{m, k}
}

func solve1(a []int) []int {
	n := len(a)
	var m0 int

	for i := 0; i+1 < n; i++ {
		if a[i] >= 0 && a[i+1] >= 0 {
			m0 = max(m0, abs(a[i]-a[i+1]))
		}
	}

	type pair struct {
		first  int
		second int
	}

	var arr []pair

	for i := 0; i < n; i++ {
		if a[i] >= 0 {
			continue
		}
		j := i
		for i < n && a[i] < 0 {
			i++
		}
		if j == 0 {
			if i == n {
				// 全部都是-1
				return []int{0, 0}
			}
			arr = append(arr, pair{-1, a[i]})
		} else {
			if i == n {
				arr = append(arr, pair{a[j-1], -1})
			} else {
				arr = append(arr, pair{a[j-1], a[i]})
			}
		}
	}

	// 也就是要找到一个k, 最小化 max(abs(k - first), abs(k - second))
	// 首先如果 k 很小，距离最大值会越远，如果k越大，就会离最小值越远
	// 所以k感觉应该取中间才对
	check := func(k int) int {
		var res int
		for _, cur := range arr {
			x, y := cur.first, cur.second
			x, y = min(x, y), max(x, y)
			if x < 0 {
				res = max(res, abs(y-k))
			} else {
				res = max(res, abs(x-k), abs(y-k))
			}
		}
		return res
	}

	l, r := 0, slices.Max(a)*2+1

	for r-l > 3 {
		m1 := l + (r-l)/3
		m2 := m1 + (r-l)/3
		ans1 := check(m1)
		ans2 := check(m2)
		if ans1 > ans2 {
			l = m1
		} else {
			r = m2
		}
	}
	m := inf
	k := l
	for i := l; i < r; i++ {
		tmp := check(i)
		if tmp < m {
			m = tmp
			k = i
		}
	}
	return []int{max(m0, m), k}
}

func abs(num int) int {
	return max(num, -num)
}
