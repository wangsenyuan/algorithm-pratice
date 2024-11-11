package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = readNNums(reader, 3)
	}

	res := solve(points)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
	}

	fmt.Print(buf.String())
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

func solve(points [][]int) [][]int {
	n := len(points)

	var res [][]int

	var loop func(d int, ids []int) int

	loop = func(d int, ids []int) int {
		if d == 3 {
			// 因为没有重叠的点， 到了第3维的时候，肯定只会有一个点
			return ids[0]
		}

		mp := make(map[int][]int)

		var pos []int

		for _, i := range ids {
			x := points[i][d]
			if mp[x] == nil {
				pos = append(pos, x)
			}
			mp[x] = append(mp[x], i)
		}

		var arr []int

		sort.Ints(pos)

		for _, v := range pos {
			vs := mp[v]
			tmp := loop(d+1, vs)
			if tmp != -1 {
				arr = append(arr, tmp)
			}
		}
		for i := 0; i+1 < len(arr); i += 2 {
			res = append(res, []int{arr[i] + 1, arr[i+1] + 1})
		}
		if len(arr)%2 == 0 {
			return -1
		}
		return arr[len(arr)-1]
	}

	ids := make([]int, n)
	for i := range n {
		ids[i] = i
	}

	loop(0, ids)

	return res
}
