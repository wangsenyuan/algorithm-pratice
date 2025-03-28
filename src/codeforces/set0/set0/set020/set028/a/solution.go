package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	nails := make([][]int, n)
	for i := range n {
		nails[i] = readNNums(reader, 2)
	}
	rods := readNNums(reader, m)
	res := solve(nails, rods)
	if res == nil {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
	fmt.Println(buf.String())
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

func solve(nails [][]int, rods []int) []int {
	has := make(map[int][]int)
	for i, r := range rods {
		has[r] = append(has[r], i+1)
	}
	n := len(nails)

	check := func(s int) []int {
		// 从s开始, 它是第一个中点
		pos := make(map[int]int)
		ans := make([]int, n)
		for i := range n {
			ans[i] = -1
		}
		for i := range n / 2 {
			id := 2*i + s
			a := nails[(id-1+n)%n]
			b := nails[id]
			c := nails[(id+1)%n]
			expect := hamiltonian_distance(a, b) + hamiltonian_distance(b, c)
			if pos[expect] == len(has[expect]) {
				return nil
			}
			ans[id] = has[expect][pos[expect]]
			pos[expect]++
		}
		return ans
	}
	res := check(0)
	if len(res) == 0 {
		res = check(1)
	}
	return res
}

func hamiltonian_distance(a []int, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(num int) int {
	return max(num, -num)
}
