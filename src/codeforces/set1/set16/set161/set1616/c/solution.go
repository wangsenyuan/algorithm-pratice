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
	}
	fmt.Print(buf.String())
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
	// 等差数列
	n := len(A)

	if n <= 2 {
		return 0
	}

	ps := make([][]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = make([]Pair, n)
		for j := i + 1; j < n; j++ {
			ps[i][j] = gcd_pair(A[j]-A[i], j-i)
		}
	}

	check := func(i int, j int) int {
		// y = a * x + b
		// A[i] = a * i + b
		// A[j] = a * j + b
		// a = (A[j] - A[i]) / (j - i)
		// b = A[i] - a * i
		var res int
		for k := 0; k < n; k++ {
			// a * k + b = A[k]
			if i == k || j == k {
				continue
			}
			var ai, aj Pair
			if k < i {
				ai = ps[k][i]
			} else {
				ai = ps[i][k]
			}

			if k < j {
				aj = ps[k][j]
			} else {
				aj = ps[j][k]
			}

			if ai != aj {
				res++
			}
			// else 斜率一致，且通过同一个点，肯定在一条线上
		}
		return res
	}

	// 任意两点 i, j 做为基准
	res := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp := check(i, j)
			if tmp < res {
				res = tmp
			}
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}

func gcd_pair(a, b int) Pair {
	g := gcd(a, b)
	a /= g
	b /= g
	return Pair{a, b}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
