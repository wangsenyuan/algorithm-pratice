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
		n, k := readTwoNums(reader)
		nums := readNNums(reader, n)
		res, segs := solve(nums, k)
		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
		for _, seg := range segs {
			buf.WriteString(fmt.Sprintf("%d %d\n", seg[0], seg[1]))
		}
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

func solve(A []int, k int) ([]int, [][]int) {
	n := len(A)
	// n <= 1e5
	// if d works, then d + 1 works
	// 可不可以迭代y，从高往低，
	B := make([]int, n)
	copy(B, A)
	sort.Ints(B)

	res := []int{-1, -1}
	var hi int
	var cnt int
	for y, x, i, j := n, n+1, n-1, n-1; y >= 1; y-- {
		// 在[x..y] 区间内的数字 >= hi + lo + k
		for i >= 0 && B[i] > y {
			cnt--
			hi++
			i--
		}
		// A[i] <= y
		// n - (hi - lo) > (hi - lo)
		for x > 1 && hi+(j+1)+k > cnt {
			x--
			for j >= 0 && B[j] >= x {
				cnt++
				j--
			}
		}
		if hi+(j+1)+k <= cnt {
			if res[0] < 0 || res[1]-res[0] > y-x {
				res = []int{x, y}
			}
		}
	}
	var segs [][]int
	expect := 1
	prev := 0
	for i := 0; i < n; i++ {
		if A[i] >= res[0] && A[i] <= res[1] {
			B[i] = 1
		} else {
			B[i] = -1
		}
		if i > 0 {
			B[i] += B[i-1]
		}
		if B[i] == expect {
			if expect == k {
				segs = append(segs, []int{prev + 1, n})
				break
			} else {
				segs = append(segs, []int{prev + 1, i + 1})
			}
			prev = i + 1
			expect++
		}
	}

	return res, segs
}
