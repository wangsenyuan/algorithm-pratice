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
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const MAX_X = 1010

var factors [MAX_X][]int

func init() {
	factors[1] = append(factors[1], 1)
	for i := 2; i < MAX_X; i++ {
		for j := 1; j <= i/j; j++ {
			if i%j == 0 {
				factors[i] = append(factors[i], j)
				if i/j != j {
					factors[i] = append(factors[i], i/j)
				}
			}
		}
		sort.Ints(factors[i])
	}
}

func solve(A []int) []int {
	// gcd(B[0]...B[i]) max
	// B[0] = max_of(A)
	// A[i] <= 1000
	// 最多20个， 剩下的随便
	cnt := make([]int, 1001)
	var mx int
	n := len(A)

	for i := 0; i < n; i++ {
		cnt[A[i]]++
		if mx < A[i] {
			mx = A[i]
		}
	}

	B := make([]int, n)
	B[0] = mx
	cnt[mx]--
	g := mx
	for i := 1; i < n; i++ {
		if g == 1 {
			x := 1
			for i < n {
				for cnt[x] == 0 {
					x++
				}
				B[i] = x
				cnt[x]--
				i++
			}

			break
		}
		for j := len(factors[g]) - 1; j >= 0 && B[i] == 0; j-- {
			x := factors[g][j]
			// 以它为GCD的，取越小的越好
			for y := x; y <= mx; y += x {
				if cnt[y] > 0 {
					B[i] = y
					cnt[y]--
					break
				}
			}
		}
		g = gcd(g, B[i])
	}

	return B
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
