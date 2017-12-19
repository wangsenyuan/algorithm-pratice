package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Bytes()
		if len(s) == 0 {
			break
		}
		winner, cnt := solve(s)
		if winner == 0 {
			fmt.Println("TIE")
		} else if winner == 1 {
			fmt.Printf("TEAM-A %d\n", cnt)
		} else {
			fmt.Printf("TEAM-B %d\n", cnt)
		}
	}

}

func solve(s []byte) (int, int) {
	x, y := 0, 0
	c, d := 5, 5
	for i := 0; i < 10; i++ {
		if s[i] == '1' {
			if i%2 == 0 {
				x++
			} else {
				y++
			}
		}
		if i%2 == 0 {
			c--
		} else {
			d--
		}

		if x+c < y {
			return -1, i + 1
		} else if y+d < x {
			return 1, i + 1
		}
	}

	if x < y {
		return -1, 10
	} else if x > y {
		return 1, 10
	}

	for i := 5; i < 10; i++ {
		a, b := s[2*i], s[2*i+1]
		if a > b {
			return 1, 2 * (i + 1)
		} else if a < b {
			return -1, 2 * (i + 1)
		}
	}

	return 0, 20
}
