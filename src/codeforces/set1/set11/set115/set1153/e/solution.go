package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	ask := func(r1 int, c1 int, r2 int, c2 int) int {
		fmt.Printf("? %d %d %d %d\n", r1, c1, r2, c2)
		return readNum(reader)
	}

	res := solve(n, ask)

	fmt.Printf("! %d %d %d %d\n", res[0], res[1], res[2], res[3])
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

func solve(n int, ask func(int, int, int, int) int) []int {
	cols := make([]int, n+1)
	for i := 1; i < n; i++ {
		cols[i] = ask(1, i, n, i) & 1
		cols[n] ^= cols[i]
	}

	var odd []int
	for i := 1; i <= n; i++ {
		if cols[i] == 1 {
			odd = append(odd, i)
		}
	}
	if len(odd) == 2 {
		return solve1(n, ask, odd[0], odd[1])
	}
	// len(odd) == 0
	return solve2(n, ask)
}

func solve1(n int, ask func(int, int, int, int) int, c1 int, c2 int) []int {
	// c1 != c2
	// can use binary search
	get := func(c int) int {
		l, r := 1, n
		for l < r {
			mid := (l + r) / 2
			v := ask(l, c, mid, c)
			if v&1 == 1 {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return r
	}
	r1 := get(c1)
	r2 := get(c2)
	return []int{r1, c1, r2, c2}
}

func solve2(n int, ask func(int, int, int, int) int) []int {
	row := make([]int, n+1)
	for i := 1; i < n; i++ {
		row[i] = ask(i, 1, i, n) & 1
		row[n] ^= row[i]
	}
	// 肯定有两个奇数
	var odd []int
	for i := 1; i <= n; i++ {
		if row[i] == 1 {
			odd = append(odd, i)
		}
	}
	r1 := odd[0]
	r2 := odd[1]
	// 只需要查r1
	get := func(i int) int {
		l, r := 1, n
		for l < r {
			mid := (l + r) / 2
			v := ask(i, l, i, mid)
			if v&1 == 1 {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return r
	}
	c := get(r1)
	return []int{r1, c, r2, c}
}
