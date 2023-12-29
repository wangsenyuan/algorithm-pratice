package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tc int
	fmt.Fscan(in, &tc)

	a := make([]int, N)
	b := make([]int, N)

	for tc > 0 {
		tc--
		var n int
		fmt.Fscan(in, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &b[i])
		}

		res := solve(a[:n], b[:n])

		fmt.Fprintln(out, res)
	}

}

const N = 200010

var freq [][]int

var ps [N]Pair

func init() {
	freq = make([][]int, 635)
	for i := 0; i < 635; i++ {
		freq[i] = make([]int, N)
	}
}

func solve(a []int, b []int) int {
	// n
	// 1 <= a[i] <= n
	// 1 <= b[i] <= n
	// a[i] * a[j] = b[i] + b[j]
	// when a[i] & b[i] is given
	// for a[i] * x <= n
	// square(n)
	n := len(a)

	for i := 0; i < n; i++ {
		ps[i] = Pair{a[i], b[i]}
	}

	sort.Slice(ps[:n], func(i, j int) bool {
		return ps[i].first < ps[j].first
	})

	var res int

	lim := int(math.Sqrt(2*float64(n))) + 2
	// 计算 a[i] = a[j] <= lim
	// a[i] * a[j] = b[i] + b[j]

	for i := 0; i < n && ps[i].first < lim; {
		j := i
		for i < n && ps[i].first == ps[j].first {
			//x, y := p.first, p.second
			u, v := ps[i].first, ps[i].first*ps[i].first-ps[i].second
			if v > 0 && v <= n {
				res += freq[u][v]
			}
			freq[u][ps[i].second]++
			i++
		}
		for j < i {
			freq[ps[j].first][ps[j].second]--
			j++
		}
	}
	// now only a[i] < a[j]
	for i := 0; i < n; {
		j := i
		// y = x * a[i] - b[i]
		for i < n && ps[i].first == ps[j].first {
			for x := 1; x*ps[i].first <= 2*n && x < ps[i].first && x < lim; x++ {
				y := x*ps[i].first - ps[i].second
				if y > 0 && y <= n {
					res += freq[x][y]
				}
			}

			i++
		}

		for j < i {
			if ps[j].first < lim {
				freq[ps[j].first][ps[j].second]++
			}
			j++
		}
	}

	for i := 0; i < n; i++ {
		if ps[i].first < lim {
			freq[ps[i].first][ps[i].second]--
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}
