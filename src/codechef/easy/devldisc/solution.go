package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)
	for i := 0; i < t; i++ {
		n := readNum(scanner)
		res := solve(n)
		if len(res) == 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(n)
			for _, x := range res {
				fmt.Printf("%d %d\n", x[0], x[1])
			}
			fmt.Println(6)
		}
	}
}

func solve(n int) [][]int {
	if n < 7 {
		return nil
	}

	res := make([][]int, n)
	var j int
	for i := 2; i <= 5; i++ {
		res[j] = []int{i - 1, i}
		j++
	}
	res[j] = []int{6, 2}
	j++
	res[j] = []int{6, 4}
	j++
	res[j] = []int{3, 7}
	j++
	for j < n {
		res[j] = []int{2, j + 1}
		j++
	}
	return res
}
