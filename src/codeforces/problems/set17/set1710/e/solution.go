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
	A := readNNums(reader, n)
	B := readNNums(reader, m)
	res := solve(A, B)
	fmt.Println(res)
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
func solve(a []int, b []int) int {
	m := len(a)
	n := len(b)
	A := make([]int, m+1)
	B := make([]int, n+1)
	copy(A[1:], a)
	copy(B[1:], b)

	sort.Ints(A)
	sort.Ints(B)

	da := make([]int, m+1)
	db := make([]int, n+1)

	va := sort.SearchInts(A, a[0])
	vb := sort.SearchInts(B, b[0])

	check := func(x int) bool {
		var sum int64

		for j, i := n, 0; j > 0; j-- {
			for i < m && A[i+1]+B[j] <= x {
				i++
			}
			db[j] = i
			sum += int64(i)
		}

		for i, j := m, 0; i > 0; i-- {
			for j < n && A[i]+B[j+1] <= x {
				j++
			}
			da[i] = j
		}

		f1 := sum
		f2 := sum

		for i, j := m, n; i >= 1; i-- {
			sum -= int64(min(j, da[i]))
			sum += int64(n - max(j, da[i]))
			f1 = max(f1, sum)
			f2 = max(f2, sum-boolValue(i <= va)*boolValue(j < vb))
			for j > 0 && min(db[j], i-1) <= m-max(db[j], i-1) {
				sum -= int64(min(db[j], i-1))
				sum += int64(m - max(db[j], i-1))
				j--
				f1 = max(f1, sum)
				f2 = max(f2, sum-boolValue(i <= va)*boolValue(j < vb))
			}
		}

		return f1 == f2
	}

	l := A[1] + B[1]
	r := a[0] + b[0]

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

type Number interface {
	int64 | int
}

func min[T Number](a T, b T) T {
	if a <= b {
		return a
	}
	return b
}

func max[T Number](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func boolValue(a bool) int64 {
	if a {
		return 1
	}
	return 0
}
