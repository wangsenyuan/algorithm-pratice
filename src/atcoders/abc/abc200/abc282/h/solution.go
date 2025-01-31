package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, s := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)

	res := solve(s, a, b)

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

func solve(s int, a []int, b []int) int {
	n := len(a)

	tr := NewSegTree(n)
	for i := 0; i < n; i++ {
		tr.Update(i, a[i])
	}

	pref := make([]int, n+1)
	for i, num := range b {
		pref[i+1] = pref[i] + num
	}

	var res int

	var dfs func(l int, r int)

	dfs = func(l int, r int) {
		if l == r {
			if a[l]+b[l] <= s {
				res++
			}
			return
		}
		mid := tr.Get(l, r+1)
		i := mid.second
		if i-l < r-i {
			// iterate left
			for u := l; u <= i; u++ {
				v := search(i, r+1, func(v int) bool {
					return pref[v+1]-pref[u]+a[i] > s
				})

				res += v - i
			}
		} else {
			// iterate right
			for v := i; v <= r; v++ {
				u := search(l, i+1, func(u int) bool {
					return pref[v+1]-pref[u]+a[i] <= s
				})

				res += i - u + 1
			}
		}
		if l <= i-1 {
			dfs(l, i-1)
		}
		if i+1 <= r {
			dfs(i+1, r)
		}
	}

	dfs(0, n-1)

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

type pair struct {
	first  int
	second int
}

func min_pair(a, b pair) pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

const inf = 1 << 50

type SegTree []pair

func NewSegTree(n int) SegTree {
	arr := make([]pair, 2*n)
	for i := n; i < len(arr); i++ {
		arr[i] = pair{inf, i - n}
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = min_pair(arr[i*2], arr[i*2+1])
	}
	return arr
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = pair{v, p - n}

	for p > 1 {
		tr[p>>1] = min_pair(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) pair {
	n := len(tr) / 2
	l += n
	r += n
	res := pair{inf, -1}
	for l < r {
		if l&1 == 1 {
			res = min_pair(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min_pair(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
