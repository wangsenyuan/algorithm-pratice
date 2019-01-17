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
		line := readNNums(scanner, 3)
		N := line[0]
		a := line[1]
		b := line[2]
		nums := readNNums(scanner, N)
		res := solve(N, a, b, nums)
		if res {
			fmt.Println("BOB")
		} else {
			fmt.Println("ALICE")
		}
	}
}

func solve(N, a, b int, nums []int) bool {
	c := lcm(a, b)

	var cc int
	var ac int
	var bc int

	for i := 0; i < N; i++ {
		if nums[i]%a == 0 {
			ac++
		}
		if nums[i]%b == 0 {
			bc++
		}

		if nums[i]%c == 0 {
			cc++
		}
	}

	if cc == 0 {
		return ac > bc
	}
	ac -= cc
	bc -= cc

	return ac+1 > bc
}

func lcm(a, b int) int {
	g := gcd(a, b)
	x := int64(a) * int64(b)
	y := x / int64(g)
	return int(y)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
