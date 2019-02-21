package main

import (
	"bufio"
	"fmt"
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

func query(x, y int) int {
	fmt.Printf("Q %d %d", x, y)
	var res int
	fmt.Scanf("%d", &res)
	return res
}

func main() {
	var tc int
	fmt.Scanf("%d", &tc)
	for tc > 0 {
		tc--
		a, b, c, d := solve(query)
		fmt.Printf("A %d %d %d %d\n", a, b, c, d)
		fmt.Scanln()
	}
}

const INF = 1e9

func solve(q func(int, int) int) (x1 int, y1 int, x2 int, y2 int) {
	a := q(0, 0)
	b := q(INF, 0)
	//x1 + x2 = a + INF - b
	if a == 0 {
		//x1, y1 has to be zero
		x1 = 0
		y1 = 0
		x2 = INF - b
		c := q(x2, INF)
		y2 = INF - c
		return
	}

	x := (a - b + INF) / 2
	y1 = q(x, 0)
	x1 = a - y1
	x2 = INF + y1 - b
	y2 = INF - q(x2, INF)
	return
}
