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
	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, m, E)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
	}
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n, m int, E [][]int) []int {
	size := make([]int, n+1)

	for _, e := range E {
		size[e[0]]++
	}

	g := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		g[i] = make([]int, 0, size[i])
	}

	for _, e := range E {
		g[e[0]] = append(g[e[0]], e[1])
	}

	for i := 0; i <= n; i++ {
		sort.Ints(g[i])
	}

	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = -1
	}
	//K := len(g[0])
	var t int
	for _, v := range g[0] {
		t++
		arr[v] = t
	}

	for pos := 1; pos <= n; pos++ {
		if arr[pos] < 0 {
			return nil
		}
		if size[pos] != size[pos-1] &&
			size[pos]+1 != size[pos-1] {
			return nil
		}
		var dif1, dif2 int
		var which1, which2 int

		for _, key := range g[pos-1] {
			if !search(g[pos], key) {
				dif1++
				which1 = key
			}
		}

		for _, key := range g[pos] {
			if !search(g[pos-1], key) {
				dif2++
				which2 = key
			}
		}

		if dif1 > 1 || dif2 > 1 || which1 != pos {
			return nil
		}
		if which2 == 0 {
			continue
		}
		arr[which2] = arr[pos]
	}
	return arr[1:]
}

func search(arr []int, key int) bool {
	var left, right = 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] >= key {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if right < len(arr) && arr[right] == key {
		return true
	}
	return false
}
