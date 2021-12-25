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

	fn := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	for tc > 0 {
		tc--
		scanner.Scan()
		var N uint64
		readUint64(scanner.Bytes(), 0, &N)
		fail := solve(N, fn)
		if fail {
			break
		}
	}
}

func check(N uint64) bool {
	if N&1 == 1 {
		return true
	}
	if N%4 == 2 {
		// losing position
		return false
	}

	return check(N / 4)
}

func solve(N uint64, fn func() string) bool {

	win := check(N)

	if !win {
		fmt.Println("Lose")
		res := fn()

		return res == "WA"
	}

	fmt.Println("Win")

	var res string
	for N != 0 {
		if N%2 == 0 {
			fmt.Println("/2")
			N /= 2
		} else if N > 1 && check(N-1) {
			fmt.Println("+1")
			N++
		} else {
			fmt.Println("-1")
			N--
		}
		res = fn()
		if res == "WA" {
			return true
		}
		if res == "GG" {
			return false
		}
		if res == "/2" {
			N /= 2
		} else if res == "+1" {
			N++
		} else {
			N--
		}
		if N == 0 {
			// other won, how can it be
			return true
		}
	}
	return false
}
