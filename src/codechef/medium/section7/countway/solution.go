package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, K := readTwoNums(reader)
	A := readNNums(reader, n)
	res := solve(A, K)
	fmt.Println(res)
}

const MOD = 13313

func solve(t []int, k int) int {
	sort.Ints(t)
	var freq []int
	for i, j := 1, 0; i <= len(t); i++ {
		if i == len(t) || t[i] > t[i-1] {
			freq = append(freq, i-j)
			j = i
		}
	}

	var loop func(freq []int) []int64

	loop = func(freq []int) []int64 {
		if len(freq) == 1 {
			res := make([]int64, freq[0]+1)
			for i := 0; i < len(res); i++ {
				res[i] = 1
			}
			return res
		}
		var a, b []int
		for i := 0; i < len(freq); i++ {
			if i < len(freq)/2 {
				a = append(a, freq[i])
			} else {
				b = append(b, freq[i])
			}
		}
		return karatsuba(loop(a), loop(b))
	}

	res := loop(freq)

	return int(res[k])
}

func brute(a, b []int64) []int64 {
	if len(a) == 0 || len(b) == 0 {
		return []int64{}
	}

	r := make([]int64, len(a)+len(b)-1)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			r[i+j] += a[i] * b[j]
		}
	}
	for i := 0; i < len(r); i++ {
		r[i] %= MOD
	}

	return r
}

func karatsuba(a, b []int64) []int64 {
	if max(len(a), len(b)) < 100 {
		return brute(a, b)
	}
	half := (max(len(a), len(b)) + 1) / 2
	a2 := make([]int64, max(0, len(a)-half))
	for i := half; i < len(a); i++ {
		a2[i-half] = a[i]
	}
	if half < len(a) {
		a = a[:half]
	}

	b2 := make([]int64, max(0, len(b)-half))
	for i := half; i < len(b); i++ {
		b2[i-half] = b[i]
	}
	if half < len(b) {
		b = b[:half]
	}
	z0 := karatsuba(a, b)
	z2 := karatsuba(a2, b2)

	for i := 0; i < len(a2); i++ {
		a[i] += a2[i]
	}
	for i := 0; i < len(b2); i++ {
		b[i] += b2[i]
	}

	z1 := karatsuba(a, b)

	r := make([]int64, max(len(z2)+2*half, len(z1)+half))

	for i := 0; i < len(z2); i++ {
		r[i+half] -= z2[i]
		r[i+2*half] += z2[i]
	}

	for i := 0; i < len(z1); i++ {
		r[i+half] += z1[i]
	}

	for i := 0; i < len(z0); i++ {
		r[i] += z0[i]
		r[i+half] -= z0[i]
	}

	for i := 0; i < len(r); i++ {
		r[i] %= MOD
		if r[i] < 0 {
			r[i] += MOD
		}
	}
	n := len(r)
	for n > 0 && r[n-1] == 0 {
		n--
	}
	return r[:n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
