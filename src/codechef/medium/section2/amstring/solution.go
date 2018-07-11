package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		N, K := readTwoNums(scanner)
		scanner.Scan()
		A := scanner.Text()
		scanner.Scan()
		B := scanner.Text()
		res := solve(N, K, A, B)
		fmt.Printf("%d\n", res)
		tc--
	}
}

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

func solve(N, K int, A, B string) int {

	find := func(L1, L2 int) int {
		R1, R2 := L1, L2
		var ans int
		var diff int
		if A[L1] != B[L2] {
			diff = 1
		}

		for R1 < N && R2 < N {
			if diff > K {
				ans += R1 - L1
				if A[L1] != B[L2] {
					diff--
				}
				L1++
				L2++
			} else {
				R1++
				R2++
				if R1 < N && R2 < N && A[R1] != B[R2] {
					diff++
				}
			}
		}
		ll := min(N-L1, N-L2)

		return ans + ll*(ll+1)/2
	}
	var ans int
	for L1 := 0; L1 < N; L1++ {
		ans += find(L1, 0)
	}

	for L2 := 1; L2 < N; L2++ {
		ans += find(0, L2)
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
