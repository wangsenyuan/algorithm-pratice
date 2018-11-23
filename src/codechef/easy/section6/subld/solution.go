package main

import (
	"bufio"
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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const MAX_K = 1e12 + 3

func main() {
	// var num, base int64

	// num = 9
	// base = 1
	// for {
	// 	k := compute(num)
	// 	if k > MAX_K {
	// 		break
	// 	}
	// 	fmt.Printf("%d: %d\n", num, k)
	// 	num = num*10 + 9
	// 	base *= 10
	// }

	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		var k uint64
		scanner.Scan()
		readUint64(scanner.Bytes(), 0, &k)
		fmt.Println(solve(int64(k)))
	}
}

// var records []int64

// func init() {
// 	records = []int64{
// 		1,
// 		28,
// 		310,
// 		3138,
// 		31427,
// 		314323,
// 		3143290,
// 		31432972,
// 		314329797,
// 		3143298050,
// 		31432980589,
// 	}
// }

func solve(k int64) int64 {
	var left, right int64 = 0, 1e14

	for left < right {
		mid := left + (right-left)/2
		cnt := compute(mid)
		if cnt > k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right - 1
}

func compute(n int64) int64 {
	var power int64 = 1
	for n >= power*10 {
		power *= 10
	}

	var steps int64
	for n > 0 {
		first := n / power
		rem := n % power
		tmp := rem/first + 1
		n = n - tmp*first
		steps += tmp
		if n < power {
			power /= 10
		}
	}
	return steps + 1
}
