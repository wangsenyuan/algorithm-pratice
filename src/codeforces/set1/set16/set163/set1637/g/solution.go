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
		n := readNum(reader)
		ok, res := solve(n)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, op := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", op[0], op[1]))
			}
		}
	}

	fmt.Print(buf.String())
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(n int) (bool, [][]int) {
	// 1, 2 => 3, 1,
	if n == 2 {
		return false, nil
	}

	var rec func(n int, coeff int)
	var arr []int
	var ops [][]int

	rec = func(n int, coeff int) {
		if n <= 2 {
			for i := 1; i <= n; i++ {
				arr = append(arr, i*coeff)
			}
			return
		}
		p := 1
		for p*2 <= n {
			p *= 2
		}
		if p == n {
			arr = append(arr, n*coeff)
			p /= 2
			n--
		}
		arr = append(arr, p*coeff)
		for i := p + 1; i <= n; i++ {
			arr = append(arr, p*2*coeff)
			ops = append(ops, []int{i * coeff, (2*p - i) * coeff})
		}
		rec(2*p-n-1, coeff)
		rec(n-p, coeff*2)
	}

	rec(n, 1)

	sort.Ints(arr)
	ans := 1
	for ans < n {
		ans *= 2
	}

	for i := 0; ; i++ {
		if arr[i] == arr[i+1] {
			ops = append(ops, []int{arr[i], arr[i+1]})
			arr[i+1] *= 2
			copy(arr[i:], arr[i+1:])
			arr = arr[:len(arr)-1]
			break
		}
	}

	for _, x := range arr {
		for x != ans {
			ops = append(ops, []int{0, x})
			ops = append(ops, []int{x, x})
			x *= 2
		}
	}
	ops = append(ops, []int{0, ans})

	return true, ops
}
