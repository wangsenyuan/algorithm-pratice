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
		n, k := readTwoNums(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, k, A))
	}
}

func solve(n int, k int, A []int) int64 {
	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		cnt[A[i]]++
	}

	arr := make([]int, 0, len(cnt))
	right := 0
	var total int64

	for _, v := range cnt {
		if v > 1 {
			arr = append(arr, v)
			total += int64(v)
			right = max(right, v)
		}
	}

	K := int64(k)
	sort.Ints(arr)

	count := func(x int) int64 {
		var sum int64

		for i := 0; i < len(arr); i++ {
			if arr[i] > x {
				sum += int64(arr[i] - x)
			}
		}

		return sum
	}

	x := sort.Search(right, func(x int) bool {
		return count(x) <= K
	})

	y := count(x)
	// y <= K

	for i := 0; i < len(arr); i++ {
		if arr[i] > x {
			arr[i] = x
		}
	}

	for i := len(arr) - 1; i >= 0 && y < K; i-- {
		arr[i]--
		y++
	}

	y = 0

	for i := 0; i < len(arr); i++ {
		c := int64(arr[i])
		y += c * (c - 1) / 2
	}

	N := int64(n)

	return N*(N-1)/2 - y
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
