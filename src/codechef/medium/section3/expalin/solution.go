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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		scanner.Scan()
		fmt.Printf("%d\n", solve(scanner.Bytes()))
		tc--
	}
}

func solve(S []byte) int64 {
	n := int64(len(S))

	ans := n

	for p := int64(2); p <= n; p++ {
		//center at i
		for i := int64(1); i*p <= n; i++ {
			if i%p == 0 {
				left, right := i/p, i*p
				for S[left-1] == S[right-1] {
					ans++
					if left%p != 0 || right*p > n {
						break
					}
					left /= p
					right *= p
				}
			}
		}

		// center between i & i * p
		for i := int64(1); i*p <= n; i++ {
			left, right := i, i*p
			for S[left-1] == S[right-1] {
				ans++
				if left%p != 0 || right*p > n {
					break
				}
				left /= p
				right *= p
			}
		}
	}

	return ans
}
