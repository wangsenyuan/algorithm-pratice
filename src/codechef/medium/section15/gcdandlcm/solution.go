package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		var n uint64
		s, _ := reader.ReadBytes('\n')
		readUint64(s, 0, &n)

		res := solve(int64(n))

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

var perfect_sqaures map[int64]int64

func init() {
	perfect_sqaures = make(map[int64]int64)
	for x := int64(1); x <= 100000; x++ {
		perfect_sqaures[x*x] = x
	}
}

func solve(N int64) int {
	var res int
	/*we have got root(n) values of m
	and for every such m we need to find integers j , k such that
	(1 + j * j)  * (1 + k * k) = m
	if we simply write it as j * j * k * k = m
	we get bound of j * k as sqrt(m)
	*/
	for i := int64(1); i <= N/i; i++ {
		if N%(i*i) != 0 {
			continue
		}
		m := N / (i * i)

		for j := int64(1); j <= (m-1)/j; j++ {
			if m%(1+j*j) != 0 {
				continue
			}
			p := m / (1 + j*j)
			k := int64(math.Sqrt(float64(p - 1)))
			if k*k == p-1 && p >= 2 && gcd(j, k) == 1 {
				res++
			}
		}
	}

	return res
}

func solve1(N int64) int {
	res := make(map[Pair]int)

	// a * a + b * b + g * g + l * l = n
	for a := int64(1); a*a <= N; a++ {

		for l := a; l*l <= N; l += a {
			x := N - a*a - l*l
			if x <= 0 {
				break
			}
			x *= l * l
			y := l*l + a*a
			if x%y == 0 {
				z := x / y
				// z is b * b
				b := perfect_sqaures[z]
				g := gcd(a, b)
				if z < N && b > 0 && a*b/g == l && b*b+g*g == N-a*a-l*l {
					res[Pair{a, b}]++
				}
			}
		}
	}

	return len(res)
}

type Pair struct {
	first, second int64
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
