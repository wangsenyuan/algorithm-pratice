package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)

	for range tc {
		_, res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, op := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", op[0], op[1]))
		}
	}

	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (a []int, res [][]int) {
	n := readNum(reader)
	a = readNNums(reader, n)
	b := make([]int, n)
	copy(b, a)
	res = solve(b)
	return
}

func solve(a []int) [][]int {
	// n := len(a)
	cnt := make([]int, 3)
	for _, x := range a {
		cnt[x]++
	}

	pos := func(x int) int {
		if x < cnt[0] {
			return 0
		}
		if x < cnt[0]+cnt[1] {
			return 1
		}
		return 2
	}

	vip := make([][][]int, 3)
	for i := range 3 {
		vip[i] = make([][]int, 3)
	}

	add := func(i int) {
		x := a[i]
		j := pos(i)
		vip[x][j] = append(vip[x][j], i)
	}
	n := len(a)

	for i := range n {
		add(i)
	}

	remove := func(i int) {
		x := a[i]
		j := pos(i)
		vip[x][j] = vip[x][j][:len(vip[x][j])-1]
	}

	var res [][]int
	swap := func(i int, j int) {
		res = append(res, []int{i + 1, j + 1})
		remove(i)
		remove(j)
		a[i], a[j] = a[j], a[i]
		add(i)
		add(j)
	}

	for {
		ok := false
		for len(vip[1][0]) > 0 && len(vip[0][1]) > 0 {
			// 1 在0的位置，交换0在1的位置
			swap(last(vip[1][0]), last(vip[0][1]))
			ok = true
		}

		for len(vip[1][0]) > 0 && len(vip[0][2]) > 0 {
			// 1在0的位置，交换0在2的位置
			swap(last(vip[1][0]), last(vip[0][2]))
			ok = true
		}

		for len(vip[1][2]) > 0 && len(vip[2][1]) > 0 {
			swap(last(vip[1][2]), last(vip[2][1]))
			ok = true
		}

		for len(vip[1][2]) > 0 && len(vip[2][0]) > 0 {
			swap(last(vip[1][2]), last(vip[2][0]))
			ok = true
		}

		if !ok {
			break
		}
	}

	if len(vip[0][2]) > 0 {
		// 2.... 1.....0
		swap(last(vip[1][1]), last(vip[0][2]))
		for {
			swap(last(vip[1][2]), last(vip[2][0]))
			if len(vip[0][2]) == 0 {
				break
			}
			swap(last(vip[1][0]), last(vip[0][2]))
		}
		swap(last(vip[1][0]), last(vip[0][1]))
	}

	return res
}

func last(arr []int) int {
	return arr[len(arr)-1]
}
