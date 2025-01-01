package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	r := readNum(reader)
	res := solve(r)
	if len(res) == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println(res[0], res[1])
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

func solve(r int) []int {
	// x**2 + 2 * xy + x + 1 = r

	for x := 1; x*x+x+1 < r; x++ {
		w := r - x*x - x - 1
		if w%(2*x) == 0 {
			return []int{x, w / (2 * x)}
		}
	}
	return nil
}
