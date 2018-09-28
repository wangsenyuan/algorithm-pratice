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
	filleNNums(scanner, n, res)
	return res
}

func filleNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)
		filleNNums(scanner, n, nums)
		res := solve(n, k, nums)
		if res {
			fmt.Println("Nancy")
		} else {
			fmt.Println("Zeta")
		}
	}
}

const MAX = 2013

var grundy [MAX * MAX]int
var nums []int

func init() {
	var used [MAX]bool
	for stones := 1; stones < MAX; stones++ {
		for i := 0; i < MAX; i++ {
			used[i] = false
		}
		var next int
		for k := 1; k < MAX; k++ {
			if stones >= k {
				used[grundy[(stones-k)*MAX+k]] = true
				for next < MAX && used[next] {
					next++
				}
			}
			grundy[stones*MAX+k] = next
		}
	}
	nums = make([]int, 60)
}

func solve(n int, k int, stones []int) bool {
	var res int
	for i := 0; i < n; i++ {
		x := min(stones[i], k)
		res ^= grundy[stones[i]*MAX+x]
	}
	return res != 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
