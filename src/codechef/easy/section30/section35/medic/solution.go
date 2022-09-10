package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

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

const YYYYMMDD = "2006:01:02"

var month_31_days = []int{1, 3, 5, 7, 8, 10, 12}

func solve(s string) int {
	date, _ := time.Parse(YYYYMMDD, s)
	mod := date.Day() % 2
	for i := 0; ; i++ {
		day := date.Day()
		if day%2 != mod {
			return i
		}
		date = date.Add(time.Hour * 48)
	}
}

func nextDay(year int, month int, date int) (int, int, int) {
	date++
	if month == 2 {
		// feb
		if date <= 28 {
			return year, month, date
		}
		// date == 29
		if date == 29 && is_leap(year) {
			return year, month, date
		}
		date = 1
		month++
		return year, month, date
	}
	if date <= 30 {
		// good
		return year, month, date
	}
	// date = 31
	if date == 31 {
		for _, m := range month_31_days {
			if m == month {
				return year, month, date
			}
		}
	}
	date = 1
	month++
	if month == 13 {
		year++
		month = 1
	}
	return year, month, date
}

func is_leap(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}

		return true
	}

	return false
}
