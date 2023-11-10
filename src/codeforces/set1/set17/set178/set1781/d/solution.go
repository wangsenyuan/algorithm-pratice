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
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func solve(a []int) int {
	n := len(a)

	// n <= 50
	//  a[i] + x
	// a[i] + x is square,
	// a[j] + x is square too
	// a[j] - a[i] = u * u - v * v = (u + v) * (u - v)
	// 就看它们的差值是否满足这样的关系
	// 找到了 u, and v
	// 就找到了对应的x

	check := func(x int) int {
		var res int
		for i := 0; i < n; i++ {
			if checkSquare(x + a[i]) {
				res++
			}
		}
		return res
	}

	best := check(0)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := a[j] - a[i]
			// diff = (u + v) * (u - v)
			for k := 1; k*k <= diff; k++ {
				if diff%k == 0 {
					q := k + diff/k
					if q%2 == 0 {
						q /= 2
						if q*q > a[j] {
							best = max(best, check(q*q-a[j]))
						}
					}
				}
			}
		}
	}

	return max(best, 1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func checkSquare(x int) bool {
	y := int(math.Sqrt(float64(x)))
	return y*y == x
}
