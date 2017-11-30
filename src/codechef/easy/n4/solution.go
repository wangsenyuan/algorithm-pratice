package main

import (
	"bufio"
	"fmt"
	"os"
)

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func main() {
	ds, path, parent := preCompute()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := toInt(scanner.Bytes())

	for i := 0; i < t; i++ {
		scanner.Scan()
		s := scanner.Text()
		d := ds[s]
		fmt.Println(d)

		for i := 0; i < d; i++ {
			fmt.Println(path[s])
			s = parent[s]
		}
	}
}

func preCompute() (map[string]int, map[string]string, map[string]string) {

	m := [][]int{
		{0, 3, 6, 8, 5, 2},
		{1, 4, 7, 9, 6, 3},
		{6, 9, 11, 12, 10, 8},
	}

	rotate := func(src string, x int, r int) string {
		res := []byte(src)

		if r == 0 {
			tmp := res[m[x][0]]

			for i := 0; i < 5; i++ {
				res[m[x][i]] = res[m[x][i+1]]
			}
			res[m[x][5]] = tmp
		} else {
			tmp := res[m[x][5]]
			for i := 5; i > 0; i-- {
				res[m[x][i]] = res[m[x][i-1]]
			}
			res[m[x][0]] = tmp
		}

		return string(res)
	}
	n := 715
	ds := make(map[string]int)
	path := make(map[string]string)
	parent := make(map[string]string)
	start := "0001001011000"
	ds[start] = 0

	que := make([]string, n)
	head, tail := 0, 0
	que[tail] = start
	tail++

	for head < tail {
		tmp := tail
		for head < tmp {
			x := que[head]
			// fmt.Printf("[debug]%d: %s\n", head, x)
			head++
			a := ds[x]
			for i := 0; i < 3; i++ {
				for j := 0; j < 2; j++ {
					z := fmt.Sprintf("%d %d", i, 1-j)
					y := rotate(x, i, j)
					// fmt.Printf("[debug] %s - (%s) -> %s\n", x, z, y)
					if y != x {
						if _, found := ds[y]; !found {
							path[y] = z
							parent[y] = x
							ds[y] = a + 1
							que[tail] = y
							//fmt.Printf("[debug]%d: %s\n", tail, y)
							tail++
						}
					}
				}
			}
		}
	}

	return ds, path, parent
}
