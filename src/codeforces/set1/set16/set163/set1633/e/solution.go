package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}
	line := readNNums(reader, 5)
	Q := readNNums(reader, line[0])
	res := solve(n, E, line[1], line[2], line[3], line[4], Q)
	fmt.Println(res)
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, E [][]int, k int, a, b, c int, Q []int) int64 {
	es := make([]Edge, len(E))
	for i, cur := range E {
		es[i] = Edge{cur[0] - 1, cur[1] - 1, 2 * cur[2]}
	}

	ev := []int{0}

	for i := 0; i < len(es); i++ {
		// from i, also add es[i].w into ev
		for j := i; j < len(es); j++ {
			ev = append(ev, (es[i].w+es[j].w)/2)
		}
	}

	ev = sortAndUnique(ev)
	var base []int64
	var cnt []int64

	set := make([]int, n)
	rk := make([]int, n)

	var find func(u int) int

	find = func(u int) int {
		if set[u] != u {
			set[u] = find(set[u])
		}
		return set[u]
	}

	unite := func(u, v int) bool {
		u = find(u)
		v = find(v)
		if u != v {
			if rk[u] < rk[v] {
				u, v = v, u
			}
			set[v] = u
			if rk[u] == rk[v] {
				rk[u]++
			}
			return true
		}
		return false
	}

	for _, x := range ev {
		sort.Slice(es, func(i, j int) bool {
			wa := abs(es[i].w - x)
			wb := abs(es[j].w - x)
			if wa != wb {
				return wa < wb
			}
			return es[i].w > es[j].w
		})

		for i := 0; i < n; i++ {
			set[i] = i
			rk[i] = 0
		}
		var curBase int64
		var curCnt int64
		for _, e := range es {
			if unite(e.u, e.v) {
				if x < e.w {
					curCnt++
					curBase += int64(e.w - x)
				} else {
					curBase += int64(x - e.w)
				}
			}
		}
		base = append(base, curBase)
		cnt = append(cnt, curCnt)
	}
	var x int
	var ans int64

	N := int64(n)

	for i := 0; i < k; i++ {
		if i < len(Q) {
			x = Q[i]
		} else {
			x = int((int64(x)*int64(a) + int64(b)) % int64(c))
		}
		j := search(len(ev), func(j int) bool {
			return ev[j] > 2*x
		})
		j--
		// ev[j] <= 2 * x
		// 那么需要知道在j的前面选了多少条边，在j的后面选了多少条边
		// cnt[j] 是在j后面选了多少条边
		tmp := (base[j] + (N-1-2*cnt[j])*int64(2*x-ev[j])) / 2
		ans ^= tmp
	}

	return ans
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
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

func sortAndUnique(arr []int) []int {
	sort.Ints(arr)
	var n int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] > arr[i-1] {
			arr[n] = arr[i-1]
			n++
		}
	}
	return arr[:n]
}

type Edge struct {
	u int
	v int
	w int
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
