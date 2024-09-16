package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line := readNNums(reader, 4)
	n, a, b, k := line[0], line[1], line[2], line[3]
	h := readNNums(reader, n)
	res := solve(a, b, k, h)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(a int, b int, k int, h []int) int {
	n := len(h)
	cnt := make([]int, n)

	for i, v := range h {
		r := v % (a + b)
		if r == 0 {
			// 这个将被bob杀死, 必须在最后的剩余b血量的时候，阻止bob出手
			cnt[i] = (b + a - 1) / a
			continue
		}
		if r <= a {
			cnt[i] = 0
			continue
		}
		r -= a
		// 阻止bob出手
		cnt[i] = (r + a - 1) / a
	}

	sort.Ints(cnt)
	var res int
	for i, c := range cnt {
		if k < c {
			break
		}
		res = i + 1
		k -= c
	}

	return res
}
