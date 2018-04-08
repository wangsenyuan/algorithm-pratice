package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)

	A := make([]int, n)

	for i := 0; i < n; i++ {
		A[i] = readNum(scanner)
	}

	fmt.Println(solve(A))
}

func solve(A []int) int {
	n := len(A)
	bits := make([][]int, 31)
	for i := 0; i < 31; i++ {
		bits[i] = make([]int, 0, n)
	}

	for i, a := range A {
		var j int
		for a > 0 {
			if a&1 == 1 {
				bits[j] = append(bits[j], i)
			}
			a >>= 1
			j++
		}
	}

	from := 30
	for from >= 0 && len(bits[from]) < 2 {
		from--
	}

	if from < 0 {
		return 0
	}

	xs := bits[from]

	res := 1 << uint(from)

	for i := from - 1; i >= 0; i-- {
		ys := bits[i]
		if len(ys) >= 2 {
			tmp := merge(xs, ys)
			if len(tmp) >= 2 {
				res |= 1 << uint(i)
				xs = tmp
			}
		}
	}

	return res
}

func merge(x, y []int) []int {
	z := make([]int, 0, len(x))

	for i, j := 0, 0; i < len(x) && j < len(y); {
		if x[i] == y[j] {
			z = append(z, x[i])
			i++
			j++
		} else if x[i] < y[j] {
			i++
		} else {
			j++
		}
	}

	return z
}
