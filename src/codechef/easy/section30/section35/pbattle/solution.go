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

	ps := make([]Pokenman, n)

	for i := 0; i < n; i++ {
		ps[i] = Pokenman{A[i], B[i]}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].ground_power < ps[j].ground_power
	})

	B = append(B, 0)

	sort.Ints(B)

	win := make([]int, n)

	m := len(B)
	arr := make([]int, 2*m)

	set := func(p int, v int) {
		p += m
		arr[p] = v
		for p > 0 {
			arr[p>>1] = arr[p] + arr[p^1]
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		l += m
		r += m
		var res int
		for l < r {
			if l&1 == 1 {
				res += arr[l]
				l++
			}
			if r&1 == 1 {
				r--
				res += arr[r]
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	index := func(wp int) int {
		return sort.SearchInts(B, wp)
	}

	for i := n - 1; i >= 0; i-- {
		// can win i from the ground arena
		j := index(ps[i].water_power)
		win[i] = i + get(0, j)
		set(j, 1)
	}

	var best int

	for i := 1; i < n; i++ {
		if win[i] > win[best] {
			best = i
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		if win[i] == win[best] {
			ans++
		}
	}
	return ans
}

type Pokenman struct {
	ground_power int
	water_power  int
}
