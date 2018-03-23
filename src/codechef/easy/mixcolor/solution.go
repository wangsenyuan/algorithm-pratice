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
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
		t--
	}
}

func solve1(n int, A []int) int {
	nums := make(map[int]bool)

	for _, a := range A {
		nums[a] = true
	}

	return len(A) - len(nums)
}

func solve(n int, A []int) int {
	mergeSort(n, A)

	var j int
	var cnt int
	for i := 1; i <= n; i++ {
		if i < n && A[i] == A[j] {
			continue
		}
		j = i
		cnt++
	}
	return n - cnt
}

func mergeSort(n int, A []int) {

	var rec func(left, right int)
	help := make([]int, n)

	rec = func(left, right int) {
		if left+1 >= right {
			return
		}

		mid := left + (right-left)/2
		rec(left, mid)
		rec(mid, right)

		i, j, k := left, mid, left
		for i < mid && j < right {
			if A[i] <= A[j] {
				help[k] = A[i]
				i++
				k++
			} else {
				help[k] = A[j]
				j++
				k++
			}
		}

		for i < mid {
			help[k] = A[i]
			i++
			k++
		}
		for j < right {
			help[k] = A[j]
			j++
			k++
		}
		copy(A[left:right], help[left:right])
	}
	rec(0, n)
}
