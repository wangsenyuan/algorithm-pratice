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
		n, k := readTwoNums(scanner)
		arr := readNNums(scanner, n)

		res := solve(n, k, arr)

		fmt.Println(res)
	}
}

func solve(n, k int, arr []int) int64 {
	items := make([]Pair, n)

	for i := 0; i < n; i++ {
		items[i] = Pair{arr[i], i}
	}

	sort.Sort(PS(items))

	bit := make([]int, n+1)

	update := func(pos int) {
		pos++
		for pos <= n {
			bit[pos]++
			pos += pos & -pos
		}
	}

	get := func(pos int) int {
		var res int
		pos++
		for pos > 0 {
			res += bit[pos]
			pos -= pos & -pos
		}
		return res
	}

	K := int64(k)

	y := (K - 1) * (K) / 2

	var res int64

	var prev int64

	for i := 0; i < n; i++ {
		if i > 0 && items[i].first == items[i-1].first {
			res += prev
			continue
		}

		item := items[i]

		x := i - get(item.second)
		prev = int64(x) * K

		// then prev += i * ((K - 1) + K - 2 + .... + 1)
		prev += int64(i) * y

		res += prev

		update(item.second)
	}

	return res
}

type Pair struct {
	first  int
	second int
}

type PS []Pair

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
