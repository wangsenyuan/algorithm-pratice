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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n, m int
	scanner.Scan()
	bs := scanner.Bytes()
	nx := readInt(bs, 0, &n)
	readInt(bs, nx+1, &m)

	stations := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		bs = scanner.Bytes()
		readInt(bs, 0, &stations[i])
	}

	a, b := solve(n, stations, m)

	fmt.Printf("%d %d\n", a, b)
}

const MOD = 1000000007

func solve(n int, stations []int, m int) (int, int) {
	f := make([]int, n)
	g := make([]int, n)
	f[0] = 1
	g[0] = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && stations[j] <= stations[i]+m; j++ {
			if g[j] == 0 || g[i]+1 < g[j] {
				f[j] = f[i]
				g[j] = g[i] + 1
			} else if g[i]+1 == g[j] {
				f[j] += f[i]
				if f[j] > MOD {
					f[j] -= MOD
				}
			}
		}
	}
	stops := g[n-1] - 1
	ways := f[n-1]
	return stops, ways
}
