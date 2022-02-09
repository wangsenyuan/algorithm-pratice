package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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
	// eventully it will became equal
	n := len(A)
	// sum / avg
	prev := make([]int, n)
	var cnt0 int
	var sum int
	for i, num := range A {
		sum += num
		if sum == 0 {
			cnt0++
		}
		prev[i] = sum
	}

	bruteForce := func(div int) int {
		mul := 1
		for i := 0; i < n && mul < prev[n-1]/div; i++ {
			if prev[i] == mul*div {
				mul++
			}
		}
		if prev[n-1]/div != mul {
			mul = 1
		}
		return n - mul
	}

	best := n - 1
	if sum == 0 {
		best = n - cnt0
	}

	for i := 1; i <= abs(sum)/i; i++ {
		if sum%i == 0 {
			best = min(best, bruteForce(sum/i))

			best = min(best, bruteForce(i*(sum/abs(sum))))
		}
	}

	return best
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func search(n int, fn func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
