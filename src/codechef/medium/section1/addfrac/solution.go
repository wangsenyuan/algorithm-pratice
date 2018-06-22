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

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		A := make([][]int, n)

		for i := 0; i < n; i++ {
			scanner.Scan()
			var x, y int
			var j int
			bs := scanner.Bytes()

			for bs[j] != '/' {
				x = x*10 + int(bs[j]-'0')
				j++
			}

			j++
			for j < len(bs) {
				y = y*10 + int(bs[j]-'0')
				j++
			}

			A[i] = []int{x, y}
		}

		B := solve(n, A)

		for i := 0; i < n; i++ {
			x, y := B[i][0], B[i][1]
			fmt.Printf("%d/%d\n", x, y)
		}
		fmt.Println()
		t--
		if t > 0 {
			scanner.Scan()
		}
	}
}

func solve(n int, A [][]int) [][]int64 {
	AA := make([][]int64, n)
	for i := 0; i < n; i++ {
		AA[i] = copyIntSlice(A[i])
	}

	upto := make([]int, n)
	upto[n-1] = n - 1

	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for j < n {
			x, y := AA[j][0], AA[j][1]
			tmp1 := float64(AA[i][0]) / float64(AA[i][1])
			tmp2 := float64(AA[i][0]+x) / float64(AA[i][1]+y)
			if tmp1 < tmp2 {
				AA[i][0] += x
				AA[i][1] += y
				j = upto[j] + 1
			} else {
				upto[i] = j - 1
				break
			}
		}
		if j >= n {
			upto[i] = n - 1
		}
	}

	for i := 0; i < n; i++ {
		g := gcd(AA[i][0], AA[i][1])
		AA[i][0] /= g
		AA[i][1] /= g
	}
	return AA
}

func copyIntSlice(X []int) []int64 {
	x, y := int64(X[0]), int64(X[1])
	return []int64{x, y}
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
