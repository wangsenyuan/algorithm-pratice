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
		left, right := readTwoNums(scanner)
		fmt.Println(solve(left, right))
	}
}

const MAX_NUM = 1e9 + 10

var all []int64

func init() {
	all = make([]int64, 0, 111)

	generate := func(cnt int, seg int64, start int64) {
		num := start
		for num <= MAX_NUM && num > 0 {
			all = append(all, num)
			num = num<<uint(cnt*2) | seg
		}
	}

	for cnt := 1; cnt <= 30; cnt++ {
		ones := generateOnes(cnt)
		seg0 := ones << uint(cnt)
		seg1 := ones
		generate(cnt, seg0, seg0)
		generate(cnt, seg1, seg1)
		fmt.Fprintf(os.Stderr, "after %d, total %d got\n", cnt, len(all))
	}
	sort.Sort(Int64Slice(all))

	fmt.Fprintf(os.Stderr, "total %d period nums\n", len(all))
	fmt.Fprintf(os.Stderr, "%v\n", all)
}

func generateOnes(cnt int) int64 {
	var res int64

	for i := 0; i < cnt; i++ {
		res |= 1 << uint(i)
	}

	return res
}

func solve(left, right int) int {
	L := int64(left)
	R := int64(right)
	a := sort.Search(len(all), func(i int) bool {
		return all[i] > R
	})
	b := sort.Search(len(all), func(i int) bool {
		return all[i] >= L
	})
	return a - b
}

type Int64Slice []int64

func (this Int64Slice) Len() int {
	return len(this)
}

func (this Int64Slice) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64Slice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
