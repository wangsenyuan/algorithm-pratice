package main

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	for {
		var s string
		n, err := fmt.Scan(&s)
		if err != nil || n == 0 {
			break
		}
		buf.WriteString(s)
	}
	res := solve(buf.String())

	buf.Reset()

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const table_start = "<table>"
const table_end = "</table>"
const tr_start = "<tr>"
const tr_end = "</tr>"
const td_start = "<td>"
const td_end = "</td>"

func solve(expr string) []int {

	expr = strings.ReplaceAll(expr, "\n", "")
	expr = strings.ReplaceAll(expr, "\t", "")
	expr = strings.ReplaceAll(expr, " ", "")

	var res []int

	var loop func(expr string)
	var loop_tr func(expr string) int
	var loop_td func(expr string)

	loop = func(expr string) {
		// <table><tr><td>
		if len(expr) == 0 {
			return
		}
		var cnt int
		stack := make([]int, len(expr))
		var p int
		start := len(table_start)
		end := len(expr) - len(table_end)
		var open int
		for i := start; i < end; {
			if i+len(tr_start) <= end && expr[i:i+len(tr_start)] == tr_start {
				open++
				if open == 1 {
					stack[p] = i
					p++
				}
				i += len(tr_start)
			} else if i+len(tr_end) <= end && expr[i:i+len(tr_end)] == tr_end {
				open--
				if open == 0 {
					j := stack[p-1]
					p--
					cnt += loop_tr(expr[j : i+len(tr_end)])
				}
				i += len(tr_end)
			} else {
				i++
			}
		}
		res = append(res, cnt)
	}

	loop_tr = func(expr string) int {
		// <tr><td></td></tr>
		start := len(tr_start)
		end := len(expr) - len(tr_end)
		var cnt int

		stack := make([]int, len(expr))
		var p int

		var open int

		for i := start; i < end; {
			if i+len(td_start) < end && expr[i:i+len(td_start)] == td_start {
				open++
				if open == 1 {
					cnt++
					stack[p] = i
					p++
				}
				i += len(td_start)
			} else if i+len(td_end) <= end && expr[i:i+len(td_end)] == td_end {
				open--
				if open == 0 {
					j := stack[p-1]
					p--
					loop_td(expr[j : i+len(td_end)])
				}
				i += len(td_end)
			} else {
				i++
			}
		}
		return cnt
	}

	loop_td = func(expr string) {
		if len(expr) == len(td_start)+len(td_end) {
			return
		}
		loop(expr[len(td_start) : len(expr)-len(td_end)])
	}

	loop(expr)

	// reverse(res)
	sort.Ints(res)

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
