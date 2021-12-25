package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	tmp := int64(0)
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * int64(sign)
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)
	buf := new(bytes.Buffer)
	for t > 0 {
		scanner.Scan()
		bs := scanner.Bytes()
		var n int
		x := readInt(bs, 0, &n)
		var k int64
		readInt64(bs, x+1, &k)
		res := solve(n, k)

		if len(res) == 0 {
			buf.WriteString("-1")
		} else {
			//fmt.Printf("%v\n", res)
			buf.WriteString(strconv.Itoa(res[0]))
			for i := 1; i < n; i++ {
				buf.WriteString(" ")
				buf.WriteString(strconv.Itoa(res[i]))
			}
		}
		buf.WriteString("\n")

		t--
	}
	fmt.Println(buf.String())
}

func solve(n int, k int64) []int {
	m := int64(n)
	minimum := m*(m+1)/2 - 1
	half := m / 2
	x := half + 1
	y := x + 1
	maximum := int64(0)

	if n%2 == 0 {
		maximum = (half-1)*(y+m) + x
	} else {
		maximum = half * (y + m)
	}

	if k < minimum || k > maximum {
		return nil
	}

	res := make([]int, n)
	left, right, x, c := 1, n, 0, minimum

	for c < maximum && c+int64(right-left-1) <= k {
		c += int64(right - left - 1)
		res[x] = left
		x++
		left++
		res[x] = right
		x++
		right--
	}

	for i := left; i <= right; i++ {
		res[x] = i
		x++
	}

	if c < k {
		s := int(k - c)
		for i := 0; i < s; i++ {
			res[n-i-1] = res[n-i-2]
		}
		res[n-1-s] = right
	}

	return res
}
