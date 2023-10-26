package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		nums := make([][]int, n)
		for i := 0; i < n; i++ {
			s, _ := reader.ReadBytes('\n')
			var m int
			pos := readInt(s, 0, &m)
			nums[i] = make([]int, m)
			for j := 0; j < m; j++ {
				pos = readInt(s, pos+1, &nums[i][j])
			}
		}
		res := solve(nums)
		if res {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const N = 200010

var cnt [N]int

func solve(nums [][]int) bool {
	// find negative result
	for _, num := range nums {
		for _, x := range num {
			cnt[x-1]++
		}
	}

	res := false

	for _, num := range nums {
		bad := false
		for _, x := range num {
			x--
			if cnt[x] == 1 {
				// 无法使用这个数，x只出现了一次， 使用它，放入任何一组，都会使一组为1， 一组为0
				bad = true
				break
			}
		}
		if !bad {
			res = true
			break
		}
	}

	for _, num := range nums {
		for _, x := range num {
			cnt[x-1] = 0
		}
	}

	return res
}
