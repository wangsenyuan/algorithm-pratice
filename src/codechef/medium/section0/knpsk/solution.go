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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := readNum(scanner)
	W := make([]int, n)
	C := make([]int, n)
	for i := 0; i < n; i++ {
		W[i], C[i] = readTwoNums(scanner)
	}
	res := solve(n, W, C)
	fmt.Printf("%d", res[0])
	for i := 1; i < len(res); i++ {
		fmt.Printf(" %d", res[i])
	}
	fmt.Println()
}

func solve(n int, W []int, C []int) []int64 {
	vc := make([]int, n)
	i, j := 0, n-1
	for k := 0; k < n; k++ {
		if W[k] == 1 {
			vc[i] = C[k]
			i++
		} else {
			vc[j] = C[k]
			j--
		}
	}

	sort.Ints(vc[:i])
	sort.Ints(vc[j+1:])

	var S int
	for a := 0; a < n; a++ {
		S += W[a]
	}
	res := make([]int64, S+1)

	var cur int64
	for w, u, v := 2, n-1, i-1; w <= S; w += 2 {
		if u == j && v < 0 {
			break
		}
		var x, y int64
		if u > j {
			x = int64(vc[u])
		}
		if v > 0 {
			y = int64(vc[v] + vc[v-1])
		} else if v == 0 {
			y = int64(vc[v])
		}
		if x > y {
			cur += x
			u--
		} else {
			cur += y
			v--
			v--
		}
		res[w] = cur
	}
	cur = 0
	if i > 0 {
		cur = int64(vc[i-1])
		i--
	}
	res[1] = cur
	for w, u, v := 3, n-1, i-1; w <= S; w += 2 {
		if u == j && v < 0 {
			break
		}
		var x, y int64
		if u > j {
			x = int64(vc[u])
		}
		if v > 0 {
			y = int64(vc[v] + vc[v-1])
		} else if v == 0 {
			y = int64(vc[v])
		}
		if x > y {
			cur += x
			u--
		} else {
			cur += y
			v--
			v--
		}
		res[w] = cur
	}

	return res[1:]
}
