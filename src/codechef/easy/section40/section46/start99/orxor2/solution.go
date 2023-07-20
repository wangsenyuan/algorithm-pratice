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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
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
	if len(bs) == 0 || bs[0] == '\n' {
		return readNum(reader)
	}
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

func solve(A []int) int {
	// max f(s1) ^ f(s2)
	// 考虑到s1中不断添加元素进去的后，会发生什么？
	sort.Ints(A)

	n := len(A)

	if A[0] == A[n-1] {
		return 0
	}

	if n == 2 {
		return A[0] ^ A[1]
	}

	var hi int

	for i := 30; i >= 0; i-- {
		var cnt int
		for _, num := range A {
			cnt += (num >> i) & 1
		}
		if cnt > 0 && cnt < n {
			hi = i
			break
		}
	}
	which := make([]int, n)
	for i := 0; i < n; i++ {
		which[i] = (A[i] >> hi) & 1
		which[i]++
	}

	set := NewDSU(n)
	// restore when fail
	tmp_parent := make([]int, n)
	tmp_size := make([]int, n)

	for d := hi - 1; d >= 0; d-- {
		cnt := make([]int, 2)
		for i := 0; i < n; i++ {
			if (A[i]>>d)&1 == 1 {
				if which[i] == 1 {
					cnt[0]++
				} else {
					cnt[1]++
				}
			}
		}
		if cnt[0] == 0 {
			continue
		}

		if cnt[1] == 0 {
			prev := -1
			for i := 0; i < n; i++ {
				if (A[i]>>d)&1 == 1 {
					if prev >= 0 {
						set.Union(prev, i)
					}
					prev = i
				}
			}
		} else {
			// both have
			for i := 0; i < n; i++ {
				tmp_parent[i] = set.Find(i)
				tmp_size[i] = set.cnt[tmp_parent[i]]
			}
			prev := -1
			var cnt_of_set1 int
			for i := 0; i < n; i++ {
				if (A[i]>>d)&1 == 1 && which[i] == 1 {
					if prev >= 0 {
						set.Union(prev, i)
					}
					prev = i
				}
				if which[i] == 1 {
					cnt_of_set1++
				}
			}
			if set.cnt[set.Find(prev)] == cnt_of_set1 {
				// can't move, need to restore
				for i := 0; i < n; i++ {
					set.arr[i] = tmp_parent[i]
					set.cnt[i] = tmp_size[i]
				}
			} else {
				for i := 0; i < n; i++ {
					if set.Find(i) == set.Find(prev) {
						which[i]++
					}
				}
			}
		}
	}

	ans := make([]int, 2)

	for i := 0; i < n; i++ {
		if which[i] == 1 {
			ans[0] |= A[i]
		} else {
			ans[1] |= A[i]
		}
	}

	return ans[0] ^ ans[1]
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}
