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
	var i int
	for i < n {
		x = readInt(scanner.Bytes(), x+1, &res[i])
		i++
		if x == len(scanner.Bytes()) {
			break
		}
	}
	return res[:i]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		N, K := readTwoNums(scanner)
		A := readNNums(scanner, N)
		fmt.Println(solve(N, K, A))
		t--
	}
}

func solve(N, K int, A []int) int {
	arr := make([]int, N)

	var dfs func(i int, flag int, cnt int)

	rule1 := func(pos int) int {

		if pos > 0 && arr[pos] > arr[pos-1] {
			return 1
		}
		return 0
	}

	var ans int

	dfs = func(i int, flag int, cnt int) {
		if i == N {
			if cnt == K {
				ans++
			}
			return
		}

		if A[i] > 0 {
			a := A[i] - 1
			if flag&(1<<uint(a)) == 0 {
				// take A[i]
				arr[i] = A[i]
				dfs(i+1, flag|(1<<uint(a)), cnt+rule1(i))
			}
			return
		}

		for j := 0; j < N; j++ {
			if flag&(1<<uint(j)) == 0 {
				// take j + 1
				arr[i] = j + 1
				dfs(i+1, flag|(1<<uint(j)), cnt+rule1(i))
			}
		}
	}

	dfs(0, 0, 0)

	return ans
}

func fact(n int) int {
	res := 1

	for n > 0 {
		res *= n
		n--
	}

	return res
}
