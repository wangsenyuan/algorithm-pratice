package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := make([]float64, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')
		readFloat64(s, 0, &a[i])
	}

	res := solve(a)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func readFloat64(bytes []byte, from int, val *float64) int {
	i := from
	var sign float64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var real int64

	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		real = real*10 + int64(bytes[i]-'0')
		i++
	}

	if i == len(bytes) || bytes[i] != '.' {
		*val = float64(real)
		return i
	}

	// bytes[i] == '.'
	i++

	var fraq float64
	var base float64 = 0.1
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		fraq += base * float64(bytes[i]-'0')
		base /= 10
		i++
	}

	*val = (float64(real) + fraq) * sign

	return i
}

const eps = 1e-6

func solve(a []float64) []int {

	var sum int
	n := len(a)

	ans := make([]int, n)
	for i, num := range a {
		ans[i] = floor(num)
		sum += ans[i]
	}
	var i int
	for sum != 0 {
		if math.Abs(float64(ans[i])-a[i]) < eps {
			// a[i] is an integer
			i++
			continue
		}
		ans[i]++
		sum++
		i++
	}

	return ans
}

func floor(num float64) int {
	if num >= 0 {
		return int(num)
	}
	// num < 0
	ans := int(-num)
	if math.Abs(float64(ans)+num) < eps {
		// -1.0, things
		return -ans
	}

	return -ans - 1
}
