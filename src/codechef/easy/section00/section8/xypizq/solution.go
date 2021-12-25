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
		line := readNNums(scanner, 5)
		a, b := solve(line[0], line[1], line[2], line[3], line[4])
		fmt.Printf("%d %d\n", a, b)
	}
}

func solve(n int, t int, x, y, z int) (int64, int64) {
	N := int64(n)
	M := 2*N + 1

	if t == 3 {
		x, z = z, x
		t = 1
	} else if t == 4 {
		t = 2
	}
	if y == 0 {
		return 1, M
	}

	if t == 1 {
		if x < y {
			if z < y {
				return answer(int64(y-1), M)
			}
			return answer(M-int64(y+1), M)
		} else {
			if z < y {
				return answer(M-int64(y-1), M)
			}
			return answer(int64(y+1), M)
		}
	}
	return answer(M-2*int64(y), M)
}

func answer(x, y int64) (int64, int64) {
	z := gcd(x, y)
	return x / z, y / z
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
