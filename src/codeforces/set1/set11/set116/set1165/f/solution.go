package main

import (
	"bufio"
	"fmt"
	"os"
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
	n, m := readTwoNums(reader)
	k := readNNums(reader, n)
	special := make([][]int, m)
	for i := range m {
		special[i] = readNNums(reader, 2)
	}
	return solve(k, special)
}

func solve(k []int, special [][]int) int {
	n := len(k)
	arr := make([]int, n)

	var md int
	for _, sp := range special {
		md = max(md, sp[0])
	}
	at := make([][]int, md+1)
	for _, sp := range special {
		at[sp[0]] = append(at[sp[0]], sp[1]-1)
	}

	check := func(days int) bool {
		// cur 是当前剩余挣到的钱
		copy(arr, k)
		cur := days
		p := min(days, md)
		for d := p; d > 0; d-- {
			for _, j := range at[d] {
				w := min(arr[j], cur, p)
				arr[j] -= w
				cur -= w
				p -= w
			}
			if p == 0 {
				break
			}
			if p == d {
				p--
			}
		}

		return 2*sum(arr) <= cur
	}

	return search(2*sum(k), check)
}

func sum(arr []int) int {
	var res int
	for i := range arr {
		res += arr[i]
	}
	return res
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
