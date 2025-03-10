package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if len(res) == 0 {
		fmt.Println("No solution")
		return
	}
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}

	buf.WriteTo(os.Stdout)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func process(reader *bufio.Reader) []string {
	n := readNum(reader)
	dates := make([]string, n)
	for i := range n {
		dates[i] = readString(reader)
	}

	return solve(dates)
}

func solve(dates []string) []string {
	// 不能小于prev
	n := len(dates)
	res := make([]string, n)

	for i, cur := range dates {
		prev := "1000"
		if i > 0 {
			prev = res[i-1]
		}
		var j int
		for j < 4 && prev[j] == cur[j] {
			j++
		}
		// 相等的部分不能变小
		if j == 4 {
			res[i] = cur
			continue
		}
		// cur[j] != prev[j]
		buf := []byte(cur)
		if j == 3 || cur[j+1:] >= prev[j+1:] {
			// 这里可以使用和prev[j]一样的字符
			buf[j] = prev[j]
			res[i] = string(buf)
			continue
		}
		if cur[j] > prev[j]+1 {
			buf[j] = prev[j] + 1
		} else if cur[j] == prev[j]+1 {
			j++
			for j < 4 && buf[j] == '0' {
				j++
			}
			if j < 4 {
				buf[j] = '0'
			}
		} else {
			// cur[j] < prev[j]
			if prev[j] == '9' {
				// 这时候，只能把前面的变大了
				for j >= 0 && prev[j] == '9' {
					j--
				}
				if j < 0 {
					return nil
				}
				buf[j] = prev[j] + 1
			} else {
				// 变成prev[j]，结果比prev更小了
				buf[j] = prev[j] + 1
			}
		}
		res[i] = string(buf)
	}

	if res[n-1] > "2011" {
		return nil
	}

	return res
}
