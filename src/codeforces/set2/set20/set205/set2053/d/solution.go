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
		res := process(reader)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n, q := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, b, queries)
}

const mod = 998244353

func add(a int, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, mod-2)
}

func div(a, b int) int {
	return mul(a, inverse(b))
}

type pair struct {
	first  int
	second int
}

func solve(a []int, b []int, queries [][]int) []int {
	n := len(a)
	A := make([]int, n)
	B := make([]int, n)
	copy(A, a)
	copy(B, b)
	sort.Ints(A)
	sort.Ints(B)

	p := 1
	for i := range n {
		p = mul(p, min(A[i], B[i]))
	}

	f1 := func(i int) int {
		j := sort.Search(n, func(j int) bool {
			return A[j] > a[i]
		})
		a[i]++

		j--
		// a[j] == a[i]
		// 其实改变的是a[j], 使的a[j]+1, j 有可能等于i
		x := min(A[j], B[j])
		p = div(p, x)
		A[j]++
		x = min(A[j], B[j])
		p = mul(p, x)
		return p
	}

	f2 := func(i int) int {
		j := sort.Search(n, func(j int) bool {
			return B[j] > b[i]
		})
		b[i]++

		j--
		// a[j] == a[i]
		// 其实改变的是a[j], 使的a[j]+1, j 有可能等于i
		x := min(A[j], B[j])
		p = div(p, x)
		B[j]++
		x = min(A[j], B[j])
		p = mul(p, x)
		return p
	}

	var ans []int
	ans = append(ans, p)

	for _, cur := range queries {
		if cur[0] == 1 {
			ans = append(ans, f1(cur[1]-1))
		} else {
			ans = append(ans, f2(cur[1]-1))
		}
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
