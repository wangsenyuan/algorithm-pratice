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

	a := readString(reader)
	b := readString(reader)
	res := solve(a, b)

	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

func solve(a string, b string) bool {
	first := parse(a)
	second := parse(b)

	buf := []int{0, 0, 0}

	var dfs func(i int, flag int) bool

	dfs = func(i int, flag int) bool {
		if i == len(buf) {
			return check(first, buf)
		}
		for j := 0; j < len(second); j++ {
			if (flag>>j)&1 == 0 {
				buf[i] = second[j]
				if dfs(i+1, flag|(1<<j)) {
					return true
				}
			}
		}
		return false
	}

	return dfs(0, 0)
}

func check(a []int, b []int) bool {
	// in DD:MM:YY format now
	// a is good
	if !valid(b) {
		return false
	}
	if a[2]-b[2] > 18 {
		return true
	}
	if a[2]-b[2] < 18 {
		return false
	}
	// a[0] - b[0] = 18
	if a[1] > b[1] {
		return true
	}
	if a[1] < b[1] {
		return false
	}
	// same month
	return a[0] >= b[0]
}

var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func valid(arr []int) bool {

	if arr[1] == 0 || arr[1] > 12 {
		return false
	}

	d := days[arr[1]-1]

	if arr[2]%4 == 0 && arr[1] == 2 {
		d++
	}

	if arr[0] == 0 || arr[0] > d {
		return false
	}

	return true
}

func parse(s string) []int {
	x := strings.Split(s, ".")
	res := make([]int, len(x))
	for i, cur := range x {
		res[i], _ = strconv.Atoi(cur)
	}
	return res
}
