package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		n, F := readTwoNums(scanner)
		H := readNNums(scanner, n)

		fmt.Println(solve(n, F, H))
		t--
	}
}

func solve(n int, F int, H []int) int64 {
	sort.Ints(H)
	cnt := make(map[int]int64)

	for i := 0; i < n; i++ {
		h := H[i]
		for h%F == 0 {
			h /= F
		}
		cnt[h]++
	}

	var ans int64
	for _, v := range cnt {
		ans += v * (v - 1) / 2
	}
	return ans
}
