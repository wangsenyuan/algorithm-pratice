package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, x, k := readThreeNums(reader)
	a := readNNums(reader, n)

	res := solve(a, k, x)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(a []int, k int, x int) int {
	var res int
	sort.Ints(a)
	n := len(a)
	if k == 0 {
		prev := 0
		for i := 0; i < n; i++ {
			if i > 0 && a[i] > a[i-1] {
				prev = i
			}
			if a[i]%x != 0 {
				w := (a[i] - 1) / x
				r := search(prev, n, func(j int) bool {
					return a[j]/x > w
				})
				l := search(prev, n, func(j int) bool {
					return a[j]/x >= w
				})
				res += r - l
			}

		}
	} else {
		for i := 0; i < n; i++ {
			w := (a[i] - 1) / x
			r := search(0, n, func(j int) bool {
				return a[j]/x-w > k
			})
			l := search(0, n, func(j int) bool {
				return a[j]/x-w >= k
			})
			res += r - l
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
