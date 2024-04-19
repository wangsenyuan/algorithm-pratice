package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a, b := readTwoNums(reader)

	fmt.Println(solve(a, b))
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(a, b int) int {
	ds := getDivisors(a + b)
	da := getDivisors(a)
	db := getDivisors(b)

	check2 := func(w int, h int, s int, x []int) bool {
		// x[i]越小，s/x[i] 越大
		i := sort.Search(len(x), func(i int) bool { return x[i] > w })
		if i == 0 {
			return false
		}
		u := x[i-1]
		return s/u <= h
	}

	check := func(w int, h int) bool {
		return check2(w, h, a, da) || check2(h, w, a, da) || check2(w, h, b, db) || check2(h, w, b, db)
	}

	var best = 2 + 2*(a+b)
	for _, x := range ds {
		y := (a + b) / x
		if check(x, y) {
			best = min(best, 2*(x+y))
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func getDivisors(num int) []int {
	var res []int
	for i := 1; i <= num/i; i++ {
		if num%i == 0 {
			res = append(res, i)
			if num != i*i {
				res = append(res, num/i)
			}
		}
	}
	sort.Ints(res)
	return res
}
