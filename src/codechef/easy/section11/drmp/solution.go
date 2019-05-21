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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
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
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		var M int64
		scanner.Scan()
		readInt64(scanner.Bytes(), 0, &M)
		res := solve(M)
		fmt.Println(len(res))
		for _, num := range res {
			fmt.Println(num)
		}
	}
}

func solve(M int64) []int64 {
	pr := make([][]int64, 30)


	x := M
	var cnt int

	for i := int64(2); i * i <= M; i++ {
		if x % i != 0 {
			continue
		}
		var c int64
		for x % i == 0 {
			x /= i
			c++
		}
		pr[cnt] = []int64{i, c}
		cnt++
	}

	if x > 1 {
		pr[cnt] = []int64{x, 1}
		cnt++
	}

	var check func(ind int, x int64)

	var ans []int64

	check = func(ind int, x int64) {
		if ind == -1 {
			ans = append(ans, x)
			return
		}
		p := int64(1)
		for i := 0; i <= int(2 * pr[ind][1]); i++ {
			if x > M / p {
				break
			}
			check(ind - 1, x * p)
			if p > M / pr[ind][0] {
				break
			}
			p *= pr[ind][0]
		}
	}

	check(cnt - 1, 1)

	sort.Sort(Int64Slice(ans))

	for i := 0; i < len(ans); i++ {
		ans[i] += M
	}

	return ans
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