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
		n := readNum(scanner)
		A := readNNums(scanner, n)
		F := readNum(scanner)
		a, b := solve(n, A, F)
		if a < 0 {
			fmt.Println("impossible")
		} else {
			fmt.Println("possible")
			fmt.Printf("%d %d\n", a, b)
		}
	}
}

func solve(N int, A []int, F int) (int, int) {
	res := 1 << 30
	pos := -1
	leftPos := -1
	leftTot := 0

	for i := 1; i <= N; i++ {
		var now int
		if i%2 == 0 {
			leftPos = i - 1
			now = A[leftPos-1]
			leftTot++
		}
		nxt := i
		if i == N {
			nxt = 1
		}

		if A[nxt-1] > F {
			continue
		}

		right := N - i
		left := leftTot
		bit := uint(1)
		rightLast := N - 1
		leftLast := leftPos
		for left+right > 1 {
			if right&1 == 0 {
				rightLast -= 1 << (bit - 1)
				left = (left + 1) >> 1
				if leftLast >= 1 && leftLast&(1<<bit) > 0 {
					leftLast -= 1 << bit
				} else if leftLast >= 1 {
					now += A[leftLast-1]
				}
			} else {
				left >>= 1
				if leftLast < 1 {
					now += A[rightLast-1]
				}

				if leftLast >= 1 && leftLast&(1<<bit) == 0 {
					leftLast -= 1 << bit
				} else if leftLast >= 1 {
					now += A[leftLast-1]
				}
			}

			right >>= 1
			bit++
		}

		if now < res {
			res = now
			pos = i
		}
	}

	if pos > 0 {
		return pos, res + F
	}
	return -1, -1
}
