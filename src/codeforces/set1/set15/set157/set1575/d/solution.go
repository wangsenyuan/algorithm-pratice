package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(s string) int {
	if len(s) > 1 && s[0] == '0' {
		return 0
	}

	var x []int

	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			x = append(x, i)
		}
	}

	if len(x) == 0 {
		return solve_x(x, -1, s)
	}
	var res int
	for i := 0; i < 10; i++ {
		if x[0] == 0 && i == 0 && len(s) > 1 {
			continue
		}
		res += solve_x(x, i, s)
	}

	return res
}

func solve_x(pos []int, x int, s string) int {
	buf := []byte(s)
	for _, i := range pos {
		buf[i] = byte(x + '0')
	}
	// now buf only have digits and _

	n := len(buf)
	if n == 1 {
		if buf[0] == '0' || buf[0] == '_' {
			return 1
		}
		return 0
	}

	var suf int
	if can_get(buf[n-2:], "00") && len(buf) > 2 {
		suf++
	}
	if can_get(buf[n-2:], "25") {
		suf++
	}
	if can_get(buf[n-2:], "75") {
		suf++
	}
	if can_get(buf[n-2:], "50") {
		suf++
	}
	if suf == 0 {
		return 0
	}

	ways := 1

	for i := 0; i < n-2; i++ {
		if buf[i] == '_' {
			if i == 0 {
				ways *= 9
			} else {
				ways *= 10
			}
		}
	}

	return ways * suf
}

func can_get(s []byte, expect string) bool {
	if s[0] == '_' || s[0] == expect[0] {
		return s[1] == '_' || s[1] == expect[1]
	}
	return false
}
