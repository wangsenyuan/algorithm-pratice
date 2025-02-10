package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, ok, res := process(reader)
	if !ok {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (s string, t string, ok bool, res [][]int) {
	readString(reader)
	s = readString(reader)
	t = readString(reader)
	ok, res = solve(s, t)
	return
}

func solve(s string, t string) (ok bool, res [][]int) {
	if s == t {
		return true, nil
	}
	cnt := make([]int, 2)
	// 01 for 1, 10 for 0
	arr := make([][]int, 4)
	for i := range s {
		x := int(s[i] - 'a')
		y := int(t[i] - 'a')
		cnt[x]++
		cnt[y]++
		arr[(x<<1)|y] = append(arr[(x<<1)|y], i)
	}
	if cnt[0]%2 == 1 {
		return false, nil
	}

	// 01 是 a/b结构
	for len(arr[1]) > 1 {
		u := arr[1][0]
		v := arr[1][1]
		res = append(res, []int{u + 1, v + 1})
		// 变成了 b/b, 和 a/a
		arr[3] = append(arr[1], u)
		arr[0] = append(arr[0], v)
		arr[1] = arr[1][2:]
	}

	for len(arr[2]) > 1 {
		// b/a
		u := arr[2][0]
		v := arr[2][1]
		res = append(res, []int{u + 1, v + 1})
		arr[0] = append(arr[0], u)
		arr[3] = append(arr[3], v)
		arr[2] = arr[2][2:]
	}

	if len(arr[1]) > 0 {
		// a/b 和 b/a
		u := arr[1][0]
		v := arr[2][0]
		res = append(res, []int{u + 1, u + 1})
		// a/b => b/a
		res = append(res, []int{u + 1, v + 1})
	}

	return true, res
}
