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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)
		A := readNNums(scanner, n)
		res := solve(n, k, A)
		if res {
			fmt.Println("YES")
			for i := 0; i < n-1; i++ {
				fmt.Printf("%d ", A[i])
			}
			fmt.Printf("%d\n", A[n-1])
		} else {
			fmt.Println("NO")
		}
	}
}

func solve(n int, k int, A []int) bool {
	if k >= 3 {
		if A[0] == -1 {
			if n == 1 || A[1] == -1 {
				A[0] = 1
			} else if A[1] > 1 {
				A[0] = A[1] - 1
			} else {
				// A[1] == 1
				A[0] = 2
			}
		}

		for i := 1; i < n-1; i++ {
			prev := A[i-1]
			after := A[i+1]
			if A[i] == -1 {
				if after > 0 {
					if prev != 1 && after != 1 {
						A[i] = 1
					} else if prev != 2 && after != 2 {
						A[i] = 2
					} else {
						A[i] = 3
					}
				} else {
					if prev != 1 {
						A[i] = 1
					} else {
						A[i] = 2
					}
				}
			} else if A[i] == A[i-1] {
				// original one conflict
				return false
			}
		}

		if A[n-1] < 0 {
			if A[n-2] != 1 {
				A[n-1] = 1
			} else {
				A[n-1] = 2
			}
		}

		if A[n-1] == A[n-2] {
			return false
		}
		return true
	}

	if k == 1 {
		if n > 1 {
			return false
		}
		A[0] = 1
		return true
	}

	// k == 2
	var i int

	for i < n && A[i] < 0 {
		i++
	}
	// i == n || A[i] > 0
	for j := i - 1; j >= 0; j-- {
		if j+1 == n || A[j+1] == 2 {
			A[j] = 1
		} else {
			A[j] = 2
		}
	}

	for {
		for i < n && A[i] > 0 {
			if i > 0 && A[i] == A[i-1] {
				return false
			}
			i++
		}
		// i == n || A[i] < 0
		if i == n {
			return true
		}
		// A[i] < 0 
		for i < n && A[i] < 0 {
			if A[i-1] == 2 {
				A[i] = 1
			} else {
				A[i] = 2
			}
			i++
		}
		if i == n {
			return true
		}
	}

}
