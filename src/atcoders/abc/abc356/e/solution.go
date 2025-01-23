package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

func solve(a []int) int {
	var mx int
	for _, x := range a {
		if x > mx {
			mx = x
		}
	}
	// 如果x的倍数贡献时w，那么对于某个y, x % y = 0, 那么对y的贡献 = y / x * w
	cnt := make([]int, mx+1)
	for _, x := range a {
		cnt[x]++
	}

	sum := make([]int, mx+2)

	var res int
	for x := mx; x > 0; x-- {
		res += cnt[x] * (cnt[x] - 1) / 2
		res += cnt[x] * (sum[x+1] - sum[min(mx+1, 2*x)])

		for y := 2 * x; y <= mx; y += x {
			res += y / x * (sum[y] - sum[min(y+x, mx+1)]) * cnt[x]
		}
		sum[x] = cnt[x] + sum[x+1]
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b 
}

func solve1(a []int) int {
	sort.Ints(a)
	var res int
	n := len(a)
	mx := a[n-1]
	for i := 0; i < n; {
		j := i
		for i < n && a[j] == a[i] {
			i++
		}
		cnt := i - j

		res += (cnt - 1) * cnt / 2

		if i == n {
			break
		}
		x := a[j]
		for y := a[i] / x * x; y <= mx; {
			l := search(i, n, func(j int) bool {
				return a[j] >= y
			})
			r := search(l, n, func(j int) bool {
				return a[j] >= y+x
			})
			res += cnt * y / x * (r - l)
			if r == n {
				break
			}
			y = a[r] / x * x
		}
	}

	return res
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
