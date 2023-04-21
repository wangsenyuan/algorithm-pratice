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
		n, x := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(A, x)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		if len(res) > 0 {
			buf.WriteByte('\n')
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

const MAX_NUM = 1000010

var lpf []int
var primes []int

func init() {
	lpf = make([]int, MAX_NUM)
	for i := 2; i < MAX_NUM; i++ {
		if lpf[i] == 0 {
			primes = append(primes, i)
			lpf[i] = i
			if MAX_NUM/i < i {
				continue
			}
			for j := i * i; j < MAX_NUM; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
}

func solve(A []int, X int) []int {
	var g int
	same := true
	for _, a := range A {
		g = gcd(g, a)
		if a != A[0] {
			same = false
		}
	}
	if same {
		return nil
	}
	if g > 1 {
		return []int{g}
	}
	// g == 1
	occ := make(map[int]int)
	var factors []int
	for _, a := range A {
		for a > 1 {
			x := lpf[a]
			occ[x]++
			if occ[x] == 1 {
				factors = append(factors, x)
			}
			a /= x
		}
	}

	sort.Ints(factors)

	for i := 0; i < len(primes) && primes[i] <= X; i++ {
		if i == len(factors) || factors[i] != primes[i] {
			return []int{primes[i]}
		}
	}
	return []int{2, 3}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
