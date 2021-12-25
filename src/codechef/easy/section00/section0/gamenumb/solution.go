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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNBigNums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n, k := readTwoNums(scanner)

		A := readNBigNums(scanner, n)
		D := readNBigNums(scanner, n)
		B := readNBigNums(scanner, k)
		res := solve(n, A, D, k, B)
		fmt.Println(res)
	}
}

func solve(n int, A []int64, D []int64, k int, B []int64) int64 {
	var count int64

	cards := make([]Card, n)
	for i := 0; i < n; i++ {
		cards[i] = Card{A[i], D[i]}
		count += D[i]
	}

	sort.Sort(Cards(cards))

	left, right := int64(1), count
	for i := 1; i <= k; i++ {
		b := B[i-1]
		if i%2 == 1 {
			left += count - b
		} else {
			right -= count - b
		}
		count = b
	}

	var ans int64
	var a, b int64
	for i := 0; i < n; i++ {
		a = b + 1
		b = a + cards[i].d - 1
		x := max(a, left)
		y := min(b, right)
		if x <= y {
			ans += cards[i].a * (y - x + 1)
		}
	}

	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

type Card struct {
	a int64
	d int64
}

type Cards []Card

func (this Cards) Len() int {
	return len(this)
}

func (this Cards) Less(i, j int) bool {
	return this[i].a < this[j].a
}

func (this Cards) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
