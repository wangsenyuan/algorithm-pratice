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
		h, k := readTwoNums(reader)
		X := readNNums(reader, 1<<uint(h))
		res := solve(h, k, X)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const INF = 1 << 30

func solve(H int, K int, X []int) int {
	n := len(X)
	dp := make([]int, n)
	A := make([]int, n)
	copy(A, X)
	who := make([]int, n)
	for i := 0; i < n; i++ {
		who[i] = i
	}

	tree := NewTree(n)

	for i := 0; i < n; i++ {
		tree.Set(i, 0)
	}

	fp := make([]int, n)
	hu := make([]int, n)
	B := make([]int, n)

	for h := 0; h < H; h++ {
		sz := 1 << uint(h)
		for i := 0; i < n; i += sz {
			left := i ^ sz
			right := left + sz
			less := left
			more := left
			for j := i; j < i+sz; j++ {
				if dp[j] == INF {
					fp[j] = INF
					continue
				}
				for less < right && A[less] < A[j] {
					less++
				}
				for more < right && A[more] <= A[j]+K {
					more++
				}
				v1 := tree.GetMin(left, less)
				v2 := tree.GetMin(less, more)
				if v2 < INF {
					v2++
				}
				fp[j] = min(v1, v2)
				if fp[j] < INF {
					fp[j] += dp[j]
				}
			}
		}
		for i := 0; i < n; i += sz * 2 {
			left := i
			right := i + sz*2
			mid := i + sz
			p1, p2, p3 := left, mid, left
			for p1 < mid || p2 < right {
				if p2 == right || p1 < mid && A[p1] < A[p2] {
					hu[p3] = who[p1]
					B[p3] = A[p1]
					dp[p3] = fp[p1]
					p1++
				} else {
					hu[p3] = who[p2]
					B[p3] = A[p2]
					dp[p3] = fp[p2]
					p2++
				}

				p3++
			}
		}

		copy(A, B)
		copy(who, hu)

		for i := 0; i < n; i++ {
			tree.Set(i, dp[i])
		}
	}

	for i := 0; i < n; i++ {
		if who[i] == 0 {
			if dp[i] < INF {
				return dp[i]
			}
			return -1
		}
	}
	return -1
}

type Tree struct {
	arr []int
	sz  int
}

func NewTree(n int) *Tree {
	t := new(Tree)
	t.arr = make([]int, 2*n)
	for i := 0; i < len(t.arr); i++ {
		t.arr[i] = INF
	}
	t.sz = n
	return t
}

func (t *Tree) Set(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = min(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *Tree) GetMin(l, r int) int {
	l += t.sz
	r += t.sz
	res := INF
	for l < r {
		if l&1 == 1 {
			res = min(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func solve1(h int, K int, X []int) int {
	n := 1 << uint(h)

	cur := make([][]Pair, n)
	for i := 0; i < n; i++ {
		cur[i] = make([]Pair, 1)
		cur[i][0] = Pair{X[i], 0}
	}

	process := func(a, b []Pair) []Pair {
		var j, k int
		c := make([]Pair, len(a))
		var p int
		for i := 0; i < len(a); i++ {
			var cnt = INF
			// need to find a range in b, that b[j].first - K <= a[i].first
			for j < len(b) && b[j].first-K > a[i].first {
				j++
			}
			if j < len(b) {
				cnt = a[i].second + b[j].second + 1
			}

			for k < len(b) && b[k].first >= a[i].first {
				k++
			}

			if k < len(b) {
				cnt = min(cnt, a[i].second+b[k].second)
			}
			if cnt < INF {
				tmp := Pair{a[i].first, cnt}
				for p > 0 && c[p-1].second >= tmp.second {
					p--
				}
				c[p] = tmp
				p++
			}

		}
		return c[:p]
	}

	merge := func(idx int) []Pair {
		a, b := cur[idx], cur[idx+1]
		// merge a & b
		A := process(a, b)
		if idx == 0 {
			return A
		}
		B := process(b, a)
		// let's merge A & B
		C := make([]Pair, len(A)+len(B))

		var i, j, p int
		for i < len(A) && j < len(B) {
			var tmp Pair
			if A[i].first > B[j].first {
				tmp = A[i]
				i++
			} else if A[i].first < B[j].first {
				tmp = B[j]
				j++
			} else {
				tmp = Pair{A[i].first, min(A[i].second, B[j].second)}
				i++
				j++
			}

			for p > 0 && C[p-1].second >= tmp.second {
				p--
			}
			C[p] = tmp
			p++
		}

		for i < len(A) {
			for p > 0 && C[p-1].second >= A[i].second {
				p--
			}
			C[p] = A[i]
			p++
			i++
		}

		for j < len(B) {
			for p > 0 && C[p-1].second >= B[j].second {
				p--
			}
			C[p] = B[j]
			p++
			j++
		}

		return C[:p]
	}

	for u := 1; u <= h; u++ {
		next := make([][]Pair, len(cur)/2)
		for i := 0; i < len(cur); i += 2 {
			next[i/2] = merge(i)
			if len(next[i/2]) == 0 {
				return -1
			}
		}
		cur = next
	}
	return cur[0][0].second
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}
