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
	a := readNNums(reader, n)
	res := solve(a)

	if len(res) == 0 {
		fmt.Println("Impossible")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("Possible\n")
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
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

func solve(a []int) []int {
	n := len(a)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] < a[id[j]] || a[id[i]] == a[id[j]] && id[i] < id[j]
	})

	res := make([]int, n)
	var p int
	for i := 0; i < n; {
		j := i
		p++
		for i < n && a[id[i]] == a[id[j]] && i-j+1 <= n-a[id[j]] {
			res[id[i]] = p
			i++
		}
		cnt := i - j
		if n-cnt != a[id[j]] {
			return nil
		}
	}

	return res
}
