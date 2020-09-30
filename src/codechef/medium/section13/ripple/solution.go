package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a, _ := reader.ReadString('\n')
		b, _ := reader.ReadString('\n')
		res := solve(a[:n], b[:n])
		fmt.Println(res)
	}
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

func solve(a, b string) string {
	// len(y) >= len(x)
	x := []byte( calc(a))
	y := []byte( calc(b))
	res := make([]byte, 0, len(y))
	var borrow int
	for i := 0; i < len(y); i++ {
		yy := y[i]
		var xx byte
		if i >= len(x) {
			xx = '0'
		} else {
			xx = x[i]
		}
		if int(yy-xx) < borrow {
			res = append(res, byte('0'+int(yy-xx)+2-borrow))
			borrow = 1
		} else {
			res = append(res, byte('0'+int(yy-xx)-borrow))
			borrow = 0
		}
	}
	var p = len(res)
	for p > 0 && res[p-1] == '0' {
		p--
	}
	if p == 0 {
		return "0"
	}
	res = res[:p]
	reverse(res)
	return string(res)
}

func calc(s string) string {
	n := len(s)

	var ca int
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			ca++
		}
	}
	var cra int
	res := make([]byte, 0, n+1)
	for i := n - 1; i >= 0; i-- {
		res = append(res, byte('0'+(cra+ca)%2))
		cra = (ca + cra) / 2
		if s[i] == '1' {
			ca--
		}
	}
	for cra > 0 {
		res = append(res, byte('0'+cra%2))
		cra /= 2
	}

	//reverse(res)

	return string(res)
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
