package main

import (
	"fmt"
	"bufio"
	"os"
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

	scanner.Scan()
	fmt.Println(solve(scanner.Bytes()))

}

func solve(S []byte) int64 {
	var a, b, c int
	var ans int64
	n := len(S)
	cnt := make(map[Pair]int64)
	cnt[Pair{0, 0}] = 1
	for i := 0; i < n; i++ {
		if S[i] == 'A' {
			a++
		} else if S[i] == 'B' {
			b++
		} else {
			c++
		}
		x := a - b
		y := b - c
		key := Pair{x, y}
		ans += cnt[key]
		cnt[key]++
	}

	return ans
}

type Pair struct {
	x, y int
}
