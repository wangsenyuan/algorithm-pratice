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
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

var primes []int
var lpf []int

const MAX_A = 1000010

func init() {

	lpf = make([]int, MAX_A)

	for i := 2; i < MAX_A; i++ {
		if lpf[i] == 0 {
			primes = append(primes, i)
			for j := i; j < MAX_A; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
}

func pow(a, b int) int64 {
	r := int64(1)
	A := int64(a)

	for b > 0 {
		if b&1 == 1 {
			r *= A
		}

		A *= A
		b >>= 1
	}

	return r
}

func solve(A []int) []int64 {
	var g int
	for _, a := range A {
		g = gcd(g, a)
	}

	res := make([]int64, len(A))

	for i, a := range A {
		res[i] = int64(a) / int64(g)
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func solve1(A []int) []int64 {
	// x | y * A[i]
	// 如果 x 整除 A[i], 它的贡献是1, 否则，
	// let g = gcd(A[i], x), 那么 y = x / g
	// 假设x  = a * b * c, y = 1
	// 如果 A[.] 不包括a (任意A[i]), y *= a
	n := len(A)
	cnts := make([]map[int]int, n)

	min_cnt := make(map[int]int)
	tot_cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		cnts[i] = make(map[int]int)
		x := A[i]
		for x > 1 {
			cnts[i][lpf[x]]++
			x /= lpf[x]
		}

		for k, v := range cnts[i] {
			if w, ok := min_cnt[k]; !ok {
				min_cnt[k] = v
			} else {
				min_cnt[k] = min(v, w)
			}
			tot_cnt[k]++
		}
	}

	res := make([]int64, n)

	for i := 0; i < n; i++ {
		var y int64 = 1
		for k, v := range cnts[i] {
			w := min_cnt[k]

			if tot_cnt[k] != n {
				w = 0
			}

			if w < v {
				y *= pow(k, v-w)
			}
		}
		res[i] = y
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
