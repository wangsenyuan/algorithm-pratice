package main

import (
	"bufio"
	"fmt"
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

	var n int

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		x, y := solve(num)
		fmt.Printf("%d %d ", x, y)
	}

}

// var A [25]int

// func init() {
// 	A[0] = 1
// 	delta := -1
// 	for i := 1; i < 25; i++ {
// 		A[i] = 2*A[i-1] + delta
// 		delta = -delta
// 	}
// }

// func solve(n int) (int, int) {
// 	return A[n-1], 1 << uint(n)
// }

func solve(n int) (int, int) {
	x, z := 1, -1

	for i := 0; i < n-1; i++ {
		x = x*2 + z
		z = -z
	}

	return x, 1 << uint(n)
}
