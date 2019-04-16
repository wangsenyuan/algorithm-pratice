package main

import (
	"bufio"
	"fmt"
	"os"
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
	n, q := readTwoNums(scanner)
	ops := make([][]int, q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		bs := scanner.Bytes()
		var tp int
		pos := readInt(bs, 0, &tp)
		if tp == 1 {
			var x int
			readInt(bs, pos + 1, &x)
			ops[i] = []int{1, x}
		} else {
			ops[i] = []int{2}
		}
	}
	res := solve(n, q, ops)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

const MOD = 1000000007

func solve(n int, q int, ops [][]int) []int {
	// NO of edges
	m := 1 << (uint(n) + 1) - 2
	// left or right side
	side := n + 1
	top := 1
	bottom := 1 << uint(n)

	var ans []int

	for _, op := range ops {
		if op[0] == 2 {
			ans = append(ans, m)
		} else  {
			// edge always double
			m = (m * 2) % MOD
			tp := op[1]
			if tp == 1 || tp == 2 {
				//right or left

				m = (m + side) % MOD
				// side will no change
				top = (top * 2) % MOD
				bottom = (bottom * 2) % MOD
			} else if tp == 3 {
				//top
				m = (m + top) % MOD
				side = (side * 2) % MOD
				top = bottom
			} else {
				//bottom
				m = (m + bottom) % MOD
				side = (side * 2) % MOD
				bottom = top
			}
		}
	}

	return ans
}
