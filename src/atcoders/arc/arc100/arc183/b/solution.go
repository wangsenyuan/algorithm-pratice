package main

import (
	"bufio"
	"bytes"
	"os"
	"reflect"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for tc > 0 {
		tc--
		res := process(reader)
		if res {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
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

func process(reader *bufio.Reader) bool {
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(a, b, k)
}

func solve(a []int, b []int, k int) bool {
	if reflect.DeepEqual(a, b) {
		return true
	}

	n := len(b)

	if k == 1 {
		var arr []int
		for i := 0; i < n; i++ {
			if len(arr) == 0 || b[i] != arr[len(arr)-1] {
				arr = append(arr, b[i])
			}
		}
		var j int
		for i := 0; i < n; i++ {
			if j < len(arr) && a[i] == arr[j] {
				j++
			}
		}
		return j == len(arr)
	}
	// k >= 2
	vis := make(map[int]bool)
	for _, x := range a {
		vis[x] = true
	}
	found := false
	prev := make(map[int]int)
	for i, x := range b {
		if !vis[x] {
			// x在a中不存在
			return false
		}
		if j, ok := prev[x]; ok && i-j <= k {
			found = true
		}
		prev[x] = i
	}

	return found
}
