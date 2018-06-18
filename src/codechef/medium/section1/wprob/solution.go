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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	x = readInt(scanner.Bytes(), x+1, &b)
	readInt(scanner.Bytes(), x+1, &c)
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
		scanner.Scan()
		s := scanner.Bytes()
		fmt.Println(solve(len(s), s))
		t--
	}
}

func solve(n int, s []byte) int64 {
	// cnt := make([]int, n)

	// check := func(r, g, b byte) int64 {
	// 	var rc int
	// 	var ans int64
	// 	var pos int
	// 	for i := 0; i < n; i++ {
	// 		cnt[i] = rc
	// 		if s[i] == r {
	// 			// need to put s[i] into pos position
	// 			ans += int64(i - pos)
	// 			rc++
	// 			pos++
	// 		}
	// 	}

	// 	for i := 0; i < n; i++ {
	// 		if s[i] == g {
	// 			//only need to take care of g
	// 			// j is the new position after replace red
	// 			j := i + rc - cnt[i]
	// 			ans += int64(j - pos)
	// 			pos++
	// 		}
	// 	}
	// 	return ans
	// }

	check := func(r, g, b byte) int64 {
		var ans int64
		var cg, cb int64

		for i := 0; i < n; i++ {
			if s[i] == r {
				ans += cg + cb
			} else if s[i] == g {
				ans += cb
				cg++
			} else {
				cb++
			}
		}
		return ans
	}

	cand := make([]int64, 6)
	cand[0] = check('r', 'g', 'b')
	cand[1] = check('r', 'b', 'g')
	cand[2] = check('g', 'r', 'b')
	cand[3] = check('g', 'b', 'r')
	cand[4] = check('b', 'r', 'g')
	cand[5] = check('b', 'g', 'r')

	sort.Sort(Int64Slice(cand))

	return cand[0]
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
