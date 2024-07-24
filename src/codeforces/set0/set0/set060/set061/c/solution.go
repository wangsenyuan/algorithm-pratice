package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	first := readString(reader)
	num := readString(reader)
	var i int
	for first[i] != ' ' {
		i++
	}
	a := first[:i]
	b := first[i+1:]
	res := solve(a, b, num)
	fmt.Println(res)
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

func solve(a, b string, num string) string {
	var x int
	readInt([]byte(a), 0, &x)
	x = convert10(num, x)

	if b == "R" {
		return convertToRoman(x)
	}
	var y int
	readInt([]byte(b), 0, &y)
	return convert(x, y)
}

// 26 in base 7 = 2 * 7 + 6 = 20 in base 10
// A9 in base 11 = 10 * 11 + 9 = 119 in base 10

func convert10(num string, base int) int {
	n := len(num)
	var res int
	for i := 0; i < n; i++ {
		var x int
		if num[i] >= '0' && num[i] <= '9' {
			// base >= 10
			x = int(num[i] - '0')
		} else {
			x = int(num[i]-'A') + 10
		}

		res *= base
		res += x
	}
	return res
}

const SMALL = "IXCM"
const BIG = "VLDM"

func convertToRoman(num int) string {
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	// reverse(arr)
	var buf []byte
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] <= 3 {
			for j := 0; j < arr[i]; j++ {
				buf = append(buf, SMALL[i])
			}
		} else if arr[i] == 4 {
			buf = append(buf, SMALL[i], BIG[i])
		} else if arr[i] == 5 {
			buf = append(buf, BIG[i])
		} else if arr[i] < 9 {
			buf = append(buf, BIG[i])
			for j := 0; j < arr[i]-5; j++ {
				buf = append(buf, SMALL[i])
			}
		} else {
			buf = append(buf, SMALL[i], SMALL[i+1])
		}
	}

	return string(buf)
}

func convert(num int, base int) string {
	var buf []byte
	for num > 0 {
		r := num % base
		num /= base
		if r < 10 {
			buf = append(buf, byte(r+'0'))
		} else {
			r -= 10
			buf = append(buf, byte(r+'A'))
		}
	}
	if len(buf) == 0 {
		return "0"
	}
	reverse(buf)
	return string(buf)
}

func reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
