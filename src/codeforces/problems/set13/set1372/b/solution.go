package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	for tc > 0 {
		tc--
		n := readNum(reader)
		a, b := solve(n)
		fmt.Printf("%d %d\n", a, b)
	}
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

const N = 100010

var primes []int

func init() {
	primes = make([]int, 0, N)
	set := make([]bool, N)
	for x := 2; x < N; x++ {
		if !set[x] {
			primes = append(primes, x)
			for y := 2 * x; y < N; y += x {
				set[y] = true
			}
		}
	}
}

func solve(n int) (int, int) {
	if n%2 == 0 {
		return n / 2, n / 2
	}
	// n is odd

	for i := 0; i < len(primes) && n > primes[i]; i++ {
		if n%primes[i] == 0 {
			a := n / primes[i]
			b := n - a
			return a, b
		}
	}
	// n is prime
	return 1, n - 1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
