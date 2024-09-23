package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	A, B, n := readThreeNums(reader)
	queries := make([][]int, n)
	for i := 0; i < n; i++ {
		queries[i] = readNNums(reader, 3)
	}
	res := solve(A, B, queries)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

func solve(A int, B int, queries [][]int) []int {
	n := len(queries)

	query := func(cur []int) int {
		l, t, m := cur[0], cur[1], cur[2]
		first := A + (l-1)*B
		if t < first {
			return -1
		}
		if t == first {
			return l
		}

		r := search(l, 1<<30, func(j int) bool {
			last := A + (j-1)*B
			sum := (first + last) * (j - l + 1) / 2
			return last > t || sum > t*m
		})

		return r - 1
	}

	ans := make([]int, n)

	for i, cur := range queries {
		ans[i] = query(cur)
	}

	return ans
}

func solve1(A int, B int, queries [][]int) []int {
	n := len(queries)

	query := func(cur []int) int {
		l, t, m := cur[0], cur[1], cur[2]
		first := A + (l-1)*B
		if t < first {
			return -1
		}
		if t == first {
			return l
		}
		// 先考虑2 * m 长度的是否能够在t秒内
		r := l + 2*m - 1
		last := A + (r-1)*B
		sum := (first + last) * (r - l + 1) / 2
		cost := (sum + m - 1) / m
		if cost <= t {
			res := search(r, 1<<30, func(j int) bool {
				tmp := (first + A + (j-1)*B) * (j - l + 1) / 2
				return (tmp+m-1)/m > t
			})
			return res - 1
		}
		res := search(l, r+1, func(j int) bool {
			last := A + (j-1)*B
			if last > t {
				return true
			}
			if j-l+1 <= m {
				return false
			}
			// 在区间(m,2*m)之间
			// 要算出，在进行不浪费的进行了x次操作后，剩余m个（非0）的最后一个的高度是多少
			// 在最后一次不浪费操作后，肯定将(j - m)个变成了0（否则的话，就可以继续操作
			// 然后剩下 m个非0 的元素，假设其中最高的是i
			// 那么就有 sum(s[i+1] .... s[j]) / m <= t
			// sum = (tmp + last) * (j - i) / 2
			i := search(l, l+m+1, func(i int) bool {
				tmp := A + (i-1)*B
				sum := (tmp + last) * (j - i + 1) / 2
				return (sum+m-1)/m <= t
			})
			tmp := A + (i-1)*B
			sum := (tmp + last) * (j - i + 1) / 2
			x := (sum + m - 1) / m
			if i == l {
				return x > t
			}
			// 它前面那个要花全部的时间
			return x+tmp-B > t
		})

		return res - 1
	}

	ans := make([]int, n)

	for i, cur := range queries {
		ans[i] = query(cur)
	}

	return ans
}

func search(l int, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
