package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(arr []int) int {
		var buf bytes.Buffer
		buf.WriteString("?")
		for _, x := range arr {
			buf.WriteString(fmt.Sprintf(" %d", x))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
		return readNum(reader)
	}

	res := solve(ask)

	fmt.Printf("! %d\n", res)
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

func solve(ask func(arr []int) int) int {
	var a, b []int
	for i := 1; i <= 100; i++ {
		a = append(a, i)
		b = append(b, i<<7)
	}
	x := ask(a)
	y := ask(b)
	x = x >> 7 << 7
	y = y & (1<<7 - 1)
	return x | y
}
