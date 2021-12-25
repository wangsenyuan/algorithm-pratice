package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		res := solve(n, A)
		if res {
			fmt.Println("Yes")
			for i := 0; i < n; i++ {
				if i < n - 1 {
					fmt.Printf("%d ", A[i])
				} else {
					fmt.Printf("%d\n", A[i])
				}
			}
		} else {
			fmt.Println("No")
		}
	}
}

func solve(n int, A []int) bool {
	if n == 1 {
		return false
	}

	pivot := A[0]
	cnt := 1
	for i := 1; i < n; i++ {
		if A[i] == pivot {
			cnt++
		} else {
			cnt--
		}

		if cnt < 0 {
			pivot = A[i]
			cnt = 1
		}
	}

	cnt = 0
	for i := 0; i < n; i++ {
		if A[i] == pivot {
			cnt++
		}
	}

	if cnt * 2 > n {
		return false
	}

	B := make([]int, n)
	copy(B, A)

	sort.Ints(B)
	var xx []int
	cnt = 0
	for i := 0; i < n; i++ {
		if A[i] == B[i] {
			xx = append(xx, A[i])
		}
	}

	half := len(xx) / 2

	swap(xx, 0, len(xx) - 1)
	swap(xx, 0, half - 1)
	swap(xx, half, len(xx) - 1)


	for i, j := 0, 0; i < n; i++ {
		if A[i] == B[i] {
			B[i] = xx[j]
			j++
		}
	}
	copy(A, B)
	return true
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func swap(nums []int, i int, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
