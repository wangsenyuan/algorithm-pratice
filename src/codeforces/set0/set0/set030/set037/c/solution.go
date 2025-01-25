package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	l := readNNums(reader, n)
	res := solve(l)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
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

type pair struct {
	first  int
	second int
}

func solve(l []int) []string {
	n := len(l)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{l[i], i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})

	m := arr[n-1].first
	buf := make([]byte, m+1)
	for i := 0; i <= m; i++ {
		buf[i] = '0'
	}
	res := make([]string, n)

	cur := []string{"0", "1"}

	for d, i := 1, 0; d <= m; d++ {
		var j int
		for i < n && arr[i].first == d {
			if j == len(cur) {
				return nil
			}
			res[arr[i].second] = cur[j]
			j++
			i++
		}

		if i == n {
			break
		}

		var next []string

		for j < len(cur) && len(next) < n-i {
			next = append(next, cur[j]+"0")
			next = append(next, cur[j]+"1")
			j++
		}
		cur = next
	}

	return res
}
