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

	var tot int

	for tc > 0 {
		tc--
		n := readNum(scanner)
		requests := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			requests[i] = scanner.Text()
		}
		res := solve(n, requests)
		fmt.Println(res)
		tot += res
	}

	fmt.Println(tot)
}

func solve(n int, requests []string) int {
	table := make([][]int, 4)

	for i := 0; i < 4; i++ {
		table[i] = make([]int, 4)
	}

	for _, req := range requests {
		x := req[0] - 'A'
		var y int
		if req[2] == '3' {
			y = 1
		} else if req[2] == '6' {
			y = 2
		} else if req[2] == '9' {
			y = 3
		}
		table[x][y]++
	}

	items := make([]Item, 16)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			items[i*4+j] = Item{i, j, table[i][j]}
		}
	}

	sort.Sort(Items(items))

	prices := []int{100, 75, 50, 25}

	var row, col int

	var res int

	var p int

	for i := 0; i < 16 && p < 4; i++ {
		cur := items[i]
		x, y, v := cur.row, cur.col, cur.value

		if v == 0 {
			break
		}

		if row&(1<<uint(x)) > 0 || col&(1<<uint(y)) > 0 {
			continue
		}

		row |= 1 << uint(x)
		col |= 1 << uint(y)

		res += prices[p] * v
		p++
	}

	for i := 0; i < 4; i++ {
		if row&(1<<uint(i)) == 0 {
			res -= 100
		}
	}

	return res
}

type Item struct {
	row, col int
	value    int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].value > items[j].value
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
