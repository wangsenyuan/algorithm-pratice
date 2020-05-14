package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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

	tc := readNum(reader)
	for tc > 0 {
		tc--
		L, _ := reader.ReadString(' ')
		R, _ := reader.ReadString('\n')
		fmt.Println(solve(L[:len(L)-1], R[:len(R)-1]))
	}
}

func solve(L, R string) int {
	L = leftPadZero(L, len(R)-len(L))

	n := len(R)

	var n4, n7 int

	var i int

	for i < n && L[i] == R[i] {
		if L[i] == '4' {
			n4++
		} else if L[i] == '7' {
			n7++
		}
		i++
	}

	var best int

	if i < n {
		// L[i] < R[i]
		x := int(L[i] - '0')
		y := int(R[i] - '0')

		if x < 4 && 4 < y {
			best = max(best, count(n4+1, n7, n-i-1))
		}
		if x < 7 && 7 < y {
			best = max(best, count(n4, n7+1, n-i-1))
		}

		if x+1 < y {
			best = max(best, count(n4, n7, n-i-1))
		}
		a, b := n4, n7
		if x == 4 {
			a++
		}
		// try stay same as L
		for k := i + 1; k < n; k++ {
			// same as L before k
			// it is already less than R, and at position k, we make it larger than L
			if L[k] < '7' {
				best = max(best, count(a, b+1, n-k-1))
			}
			if L[k] < '4' {
				best = max(best, count(a+1, b, n-k-1))
			}
			if L[k] == '4' {
				a++
			} else if L[k] == '7' {
				b++
			}
		}

		best = max(best, count(a, b, 0))

		a, b = n4, n7
		if y == 7 {
			b++
		}

		for k := i + 1; k < n; k++ {
			if R[k] > '7' {
				best = max(best, count(a, b+1, n-k-1))
			}
			if R[k] > '4' {
				best = max(best, count(a+1, b, n-k-1))
			}
			if R[k] == '4' {
				a++
			} else if R[k] == '7' {
				b++
			}
		}

		best = max(best, count(a, b, 0))

	} else {
		best = n4 * n7
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func leftPadZero(s string, n int) string {
	var buf bytes.Buffer

	for n > 0 {
		buf.WriteByte('0')
		n--
	}
	buf.WriteString(s)
	return buf.String()
}

func count(cnt4, cnt7 int, n int) int {
	best := cnt4 * cnt7

	for a := 0; a <= n; a++ {
		x := (cnt4 + a) * (cnt7 + n - a)
		if x > best {
			best = x
		}
	}
	return best
}
