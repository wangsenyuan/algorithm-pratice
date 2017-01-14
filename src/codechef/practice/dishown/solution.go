package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n, q int
		fmt.Scanf("%d", &n)

		uf := &UF{make([]int, n)}

		s := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &s[i])
			uf.elems[i] = i
		}

		fmt.Scanf("%d", &q)

		for q > 0 {
			var k, x, y int
			fmt.Scanf("%d", &k)

			if k == 0 {
				fmt.Scanf("%d %d", &x, &y)
				x--
				y--
				if uf.sameSet(x, y) {
					fmt.Println("Invalid query!")
					continue
				}

				a := uf.find(x)
				b := uf.find(y)
				if s[a] > s[b] {
					uf.union(a, b)
				} else if s[a] < s[b] {
					uf.union(b, a)
				}

			} else {
				fmt.Scanf("%d", &x)
				fmt.Println(uf.find(x - 1) + 1)
			}

			q--
		}

		t--
	}
}

type UF struct {
	elems []int
}

func (uf *UF) find(a int) int {
	if uf.elems[a] != a {
		uf.elems[a] = uf.find(uf.elems[a])
	}

	return uf.elems[a]
}

func (uf *UF) sameSet(a, b int) bool {
	return uf.find(a) == uf.find(b)
}

func (uf *UF) union(a, b int) {
	uf.elems[b] = a
}
