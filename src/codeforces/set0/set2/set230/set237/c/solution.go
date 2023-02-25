package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a, b, k := readThreeNums(reader)
	res := solve(a, b, k)
	fmt.Println(res)
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

func solve(a int, b int, k int) int {
	// 1 <= l <= b - a + 1
	// a.... a + l - 1 至少有k个prime numbers
	// a + 1, .... a + l ....
	// a + 2, .... a + l + 1
	var primes []int
	for x := a; x <= b; x++ {
		if isPrime(x) {
			primes = append(primes, x)
		}
	}

	check := func(l int) bool {
		var front, end int
		for x := a; x <= b-l+1; x++ {
			y := x + l - 1
			for end < len(primes) && primes[end] <= y {
				end++
			}

			for front < len(primes) && primes[front] < x {
				front++
			}
			if end-front < k {
				return false
			}
		}
		return true
	}
	if len(primes) < k {
		return -1
	}
	l, r := 1, b-a+1

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	for x := 3; x <= num/x; x += 2 {
		if num%x == 0 {
			return false
		}
	}
	return true
}
