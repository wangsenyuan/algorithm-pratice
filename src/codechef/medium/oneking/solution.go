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
		n := readNum(scanner)
		kingdoms := make([][]int, n)
		for i := 0; i < n; i++ {
			kingdoms[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(n, kingdoms))
		t--
	}
}

func solve(n int, kingdoms [][]int) int {
	sort.Sort(Kingdoms(kingdoms))

	var ans int
	prev := -1

	for _, kingdom := range kingdoms {
		left, right := kingdom[0], kingdom[1]
		if left > prev {
			ans++
			prev = right
		}
	}
	return ans
}

type Kingdoms [][]int

func (this Kingdoms) Len() int {
	return len(this)
}

func (this Kingdoms) Less(i, j int) bool {
	return this[i][1] < this[j][1]
}

func (this Kingdoms) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
