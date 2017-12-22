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
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = sign * tmp
	return i
}

func readNum(scanner *bufio.Scanner) (x int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &x)
	return
}

func readFourNums(scanner *bufio.Scanner) (a int, b int, c int, d int) {
	scanner.Scan()
	bs := scanner.Bytes()
	x := readInt(bs, 0, &a)
	x = readInt(bs, x+1, &b)
	x = readInt(bs, x+1, &c)
	readInt(bs, x+1, &d)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for i := 0; i < t; i++ {
		n := readNum(scanner)
		flights := make([][]int, n)
		for j := 0; j < n; j++ {
			a, b, c, d := readFourNums(scanner)
			flights[j] = []int{a, b, c, d}
		}
		from, begin, end, birthday := readFourNums(scanner)
		can, cnt := solve(n, flights, from, begin, end, birthday)
		if can {
			fmt.Printf("Yes %d", cnt)
		} else {
			fmt.Println("No")
		}
	}
}

func solve(n int, flights [][]int, from int, begin int, end int, birthday int) (bool, int) {
	// fmt.Printf("[debug] %v\n", flights)
	sort.Sort(Flights(flights))

	city, time, cnt := from, begin, 0
	mark := make([]bool, n)
	for {
		if city == end && time <= birthday {
			return true, cnt
		}
		i := sort.Search(n, func(j int) bool {
			x := flights[j]
			return x[0] >= city || (x[0] == city && x[1] >= time)
		})

		if i == n || mark[i] {
			return false, -1
		}
		tmp := flights[i]
		if tmp[0] != city {
			return false, -1
		}
		city, time = tmp[2], tmp[3]
		cnt++
		mark[i] = true
	}
}

type Flights [][]int

func (this Flights) Len() int {
	return len(this)
}

func (this Flights) Less(i, j int) bool {
	a, b := this[i], this[j]
	return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
}

func (this Flights) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
