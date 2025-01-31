package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, I := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(a, I)
}

func solve(a []int, I int) int {
	n := len(a)

	sort.Ints(a)

	freq := make(map[int]int)

	var m int
	add := func(num int) {
		freq[num]++
		if freq[num] == 1 {
			m++
		}
	}

	rem := func(num int) {
		if freq[num] == 1 {
			m--
		}
		freq[num]--
	}

	I *= 8

	ans := n

	x := make([]int, n+1)
	x[1] = 0

	for i := 2; i <= n; i++ {
		x[i] = x[i-1]
		if 1<<x[i] < i {
			x[i]++
		}
	}

	for l, r := 0, 0; r < n; r++ {
		add(a[r])
		for x[m]*n > I {
			// distinct value 太多了
			rem(a[l])
			l++
		}
		ans = min(ans, n-(r-l+1))
	}

	return ans
}
