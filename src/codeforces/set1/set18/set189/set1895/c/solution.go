package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readNum(reader)
	s := readString(reader)
	ss := strings.Split(s, " ")
	res := solve(ss)
	fmt.Println(res)
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

func solve(s []string) int {
	ans := 0
	cnt := make([][]int, 6)
	for i := range cnt {
		cnt[i] = make([]int, 46)
	}

	for _, y := range s {
		l := len(y)
		sumDigits := 0
		for i := 0; i < l; i++ {
			digit := int(y[i] - '0')
			sumDigits += digit
		}
		cnt[l][sumDigits]++
	}

	for _, L := range s {
		for lenr := len(L) % 2; lenr <= len(L); lenr += 2 {
			l := len(L) + lenr
			suml := 0
			for _, c := range L[:l/2] {
				digit, _ := strconv.Atoi(string(c))
				suml += digit
			}
			sumr := 0
			for _, c := range L[l/2:] {
				digit, _ := strconv.Atoi(string(c))
				sumr += digit
			}
			if suml-sumr >= 0 {
				ans += cnt[lenr][suml-sumr]
			}
		}
	}

	for _, R := range s {
		for lenl := len(R) % 2; lenl < len(R); lenl += 2 {
			l := len(R) + lenl
			suml := 0
			for _, c := range R[len(R)-l/2:] {
				digit, _ := strconv.Atoi(string(c))
				suml += digit
			}
			sumr := 0
			for _, c := range R[:len(R)-l/2] {
				digit, _ := strconv.Atoi(string(c))
				sumr += digit
			}
			if suml-sumr >= 0 {
				ans += cnt[lenl][suml-sumr]
			}
		}
	}

	return ans
}
