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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(A []int, B []int) int {
	n := len(A)
	C := make([]Pair, 2*n)

	for i := 0; i < n; i++ {
		C[i] = Pair{A[i], i}
		C[i+n] = Pair{B[i], i}
	}

	sort.Slice(C, func(i, j int) bool {
		return C[i].first < C[j].first
	})
	var res int = INF
	cnt := make(map[int]int)
	for l, r := 0, 0; l < len(C); l++ {
		for r < len(C) && len(cnt) < n {
			cnt[C[r].second]++
			r++
		}
		if len(cnt) == n {
			res = min(res, C[r-1].first-C[l].first)
		}
		cnt[C[l].second]--
		if cnt[C[l].second] == 0 {
			delete(cnt, C[l].second)
		}
	}

	return res
}

func solve1(A []int, B []int) int {
	// min(Amax - Amin), after swaps
	// swap(i), 有可能增加最大值，或者减少最小值
	// 假如我们最小值x的情况下，所能获得的最大值？
	//    如果 A[i] < x, 需要翻转, (B[i] < x, false)
	//        如果 A[i] > x && B[i] > x, 保留 min(A[i], B[i])
	// 至少是可操作的
	// 但是迭代 x, 是 O(n)的操作
	// A * B, 组成 P, Pi = (Ai, B[i])
	// 按照A[i] 升序排列, 迭代到i时，假如以i, 为最小值, 那么 j < i, 必须都翻转, 且不能有B[j] < A[i]
	//  max(B[j]) = y, min(B[j]) < A[i], break
	// i后面的部分呢? 可以不翻转，对于k来说, 如果 B[k] < A[k] && B[k] > A[i], 翻转
	//  如何高效的进行这个操作呢?
	//  同时按照B[i] 升序迭代，把所有小于A[i]的下标k设置为对应的a,
	// if len(A) == 1 {
	// 	return 0
	// }
	a := process(A, B)
	b := process(B, A)
	return min(a, b)
}

const INF = 1e9

func process(A []int, B []int) int {
	n := len(A)
	P := make([]Pair, n)

	for i := 0; i < n; i++ {
		P[i] = Pair{A[i], B[i]}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].first < P[j].first
	})

	// P sorted by first
	X := make([]Pair, n)
	for i := 0; i < n; i++ {
		X[i] = Pair{P[i].second, i}
	}

	sort.Slice(X, func(i, j int) bool {
		return X[i].first < X[j].first
	})

	arr := make([]int, 2*n)

	for i := n; i < 2*n; i++ {
		arr[i] = min(P[i-n].first, P[i-n].second)
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = max(arr[i*2], arr[i*2+1])
	}

	set := func(p int, v int) {
		p += n
		arr[p] = v
		for p > 1 {
			arr[p>>1] = max(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		l += n
		r += n
		var res int
		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	var res int = INF

	// X only for B
	var mx int
	var mn int = INF
	for i, j := 0, 0; i < n; i++ {
		// 以x为最小值
		x := P[i].first
		if mn < x {
			// x and later A[i] can't be min value
			break
		}
		for j < n && X[j].first < x {
			k := X[j].second
			if k > i {
				set(k, max(P[k].first, P[k].second))
			}
			j++
		}
		my := max(get(i+1, n), mx)
		if my >= x {
			res = min(res, my-x)
		}

		mx = max(mx, P[i].second)
		mn = min(mn, P[i].second)
	}

	return res
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

type Pair struct {
	first  int
	second int
}
