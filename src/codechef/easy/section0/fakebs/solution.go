package main

import (
	"sort"
	"bufio"
	"os"
	"fmt"
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

func readNum(scanner *bufio.Scanner) int {
	var x int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &x)
	return x
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	bs := scanner.Bytes()
	pos := readInt(bs, 0, &a)
	pos = readInt(bs, pos+1, &b)
	return
}

func readTripNums(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	bs := scanner.Bytes()
	pos := readInt(bs, 0, &a)
	pos = readInt(bs, pos+1, &b)
	readInt(bs, pos+1, &c)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	scanner.Scan()
	bs := scanner.Bytes()
	x := -1
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for t > 0 {
		n, q := readTwoNums(scanner)
		A := readNNums(scanner, n)
		X := make([]int, q)
		for i := 0; i < q; i++ {
			X[i] = readNum(scanner)
		}
		res := solve(n, A, X)
		for i := 0; i < q; i++ {
			fmt.Println(res[i])
		}
		t--
	}
}

func solve(n int, A []int, X []int) []int {
	B := make([]int, n)
	copy(B, A)
	sort.Ints(B)

	index := make(map[int]int)
	sorted := make(map[int]int)

	for i := 0; i < n; i++ {
		index[A[i]] = i
		sorted[B[i]] = i
	}

	ans := make([]int, len(X))
	for i := 0; i < len(X); i++ {
		x := X[i]
		idx := index[x]
		nl := sorted[x]
		ng := n - nl - 1

		var ctl, ctg, totctl, totctg int

		low, high := 0, n-1
		for low < high {
			mid := (low + high) / 2
			if idx == mid {
				break
			}
			if idx < mid {
				high = mid - 1
				totctg++
				if A[mid] < x {
					ctg++
				}
			} else {
				low = mid + 1
				totctl++
				if A[mid] > x {
					ctl++
				}
			}
		}

		if totctl > nl || totctg > ng {
			ans[i] = -1
		} else {
			ans[i] = min(ctl, ctg) + abs(ctl-ctg)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
