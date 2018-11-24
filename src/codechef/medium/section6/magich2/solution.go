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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)
	for tc > 0 {
		tc--
		var n, k uint64
		scanner.Scan()
		pos := readUint64(scanner.Bytes(), 0, &n)
		readUint64(scanner.Bytes(), pos+1, &k)
		res := solve(n, k)
		fmt.Printf("%.7f\n", res)
	}
}

func solve(n, k uint64) float64 {
	if n <= 2 || k == 0 {
		return 1.0 / float64(n)
	}

	if n&3 == 2 {
		n = n >> 1
		k--
	}

	for n > 1 && k > 0 {
		k--
		n = (n + 1) >> 1
	}
	return 1.0 / float64(n)

	// var cal func(n, k uint64) float64

	// cal = func(n, k uint64) float64 {
	// 	if n == 0 {
	// 		return 1.0
	// 	}
	// 	if k == 0 {
	// 		return 1.0 / float64(n)
	// 	}

	// 	if n%4 == 0 {
	// 		return cal(n/2, k-1)
	// 	}
	// 	if n%4 == 1 {
	// 		return cal(n/2+1, k-1)
	// 	}
	// 	if n%4 == 2 {
	// 		return cal(n/2, k-1)
	// 	}
	// 	return cal(n/2+1, k-1)
	// }

	// if n%4 == 2 {
	// 	return cal(n/2+1, k-1)
	// }

	// return cal(n, k)
}
