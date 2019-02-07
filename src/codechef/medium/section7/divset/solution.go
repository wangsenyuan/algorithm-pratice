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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)
	for tc > 0 {
		tc--
		N, K, C := readThreeNums(scanner)
		A := readNNums(scanner, N)
		fmt.Println(solve(N, K, C, A))
	}
}

const INF = 1<<29 + 1

func solve(N, K, C int, A []int) int {
	sort.Ints(A)

	nums := make([]int, N)
	sz := make([]int, N)
	check := func(md int) bool {
		for i := 0; i < md; i++ {
			sz[i] = 0
		}

		var nw int
		for i := 0; i < N; i++ {
			if sz[nw] == 0 || (nums[nw] < INF/C && nums[nw]*C <= A[i]) {
				sz[nw]++
				nums[nw] = A[i]
				nw = (nw + 1) % md
			}
		}

		return sz[md-1] >= K
	}

	left, right := 0, N+1
	for left < right-1 {
		mid := (left + right) / 2
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
