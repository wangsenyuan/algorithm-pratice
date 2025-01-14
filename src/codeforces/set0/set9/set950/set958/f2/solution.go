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
	a := readNNums(reader, n)
	k := readNNums(reader, m)
	return solve(a, k)
}

func solve(a []int, k []int) int {
	m := len(k)
	n := len(a)
	freq := make([]int, m)
	for i := range n {
		a[i]--
	}

	var sum int
	for _, x := range k {
		sum += x
		if x == 0 {
			m--
		}
	}

	var cnt int

	add := func(num int) {
		freq[num]++
		if freq[num] == k[num] {
			cnt++
		}
	}
	rem := func(num int) {
		if freq[num] == k[num] {
			cnt--
		}
		freq[num]--
	}

	res := -1

	for l, r := 0, 0; r < n; r++ {
		add(a[r])
		for cnt == m {
			rem(a[l])
			if cnt < m {
				add(a[l])
				break
			}
			l++
		}
		if cnt == m {
			if res == -1 || res > r-l+1-sum {
				res = r - l + 1 - sum
			}
		}
	}
	return res
}
