package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)

		buf.WriteString(fmt.Sprintf("%d\n", res))

		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", A[i]))
		}
		buf.WriteByte('\n')
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

func solve(A []int) int {
	var ans int

	n := len(A)

	for i := 0; i < n; i++ {
		if A[i] == 0 {
			A[i]++
			ans++
		}
	}

	if connected(A) {
		return ans
	}

	for i := 0; i < n; i++ {
		A[i]++
		if connected(A) {
			return ans + 1
		}
		A[i]--
	}

	for i := 0; i < n; i++ {
		if A[i] > 1 {
			A[i]--
			if connected(A) {
				return ans + 1
			}
			A[i]++
		}
	}

	var m int

	for i := 0; i < n; i++ {
		m = max(m, A[i]&(-A[i]))
	}

	for i := 0; i < n; i++ {
		if A[i]&-A[i] == m {
			A[i]--
			break
		}
	}

	for i := 0; i < n; i++ {
		if A[i]&-A[i] == m {
			A[i]++
			break
		}
	}

	return ans + 2
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func connected(arr []int) bool {
	set := NewSet(32)
	var all int
	for _, a := range arr {
		all |= a
		last := -1
		for i := 0; i < 32; i++ {
			if (a>>i)&1 == 1 {
				if last >= 0 {
					set.Union(last, i)
				}
				last = i
			}
		}
	}
	var cnt int
	var root int
	for i := 0; i < 32; i++ {
		if (all>>i)&1 == 1 {
			cnt++
			root = set.Find(i)
		}
	}

	return set.cnt[root] == cnt
}

type Set struct {
	arr []int
	cnt []int
}

func NewSet(n int) *Set {
	set := new(Set)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	return set
}

func (set *Set) Find(x int) int {
	if set.arr[x] != x {
		set.arr[x] = set.Find(set.arr[x])
	}
	return set.arr[x]
}

func (set *Set) Union(a, b int) bool {
	a = set.Find(a)
	b = set.Find(b)
	if a == b {
		return false
	}
	if set.cnt[a] < set.cnt[b] {
		a, b = b, a
	}
	set.cnt[a] += set.cnt[b]
	set.arr[b] = a
	return true
}
