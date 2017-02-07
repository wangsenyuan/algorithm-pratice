package main

import "fmt"

func main() {
	var n, m, x int
	fmt.Scanf("%d %d %d", &n, &m, &x)

	bit := Bit{n, make([]int64, n+1)}

	bit.update(0, x)

	var tp string
	var a, b, c int
	for m > 0 {
		fmt.Scanf("%s", &tp)

		if tp == "Q" {
			fmt.Scanf("%d", &a)
			fmt.Println(bit.query(a - 1))
		} else {
			fmt.Scanf("%d %d %d", &a, &b, &c)
			a--
			b--
			bit.update(a, c)
			bit.update(b+1, -c)
		}

		m--
	}
}

type Bit struct {
	n   int
	arr []int64
}

func (bit Bit) update(i, val int) {
	i++
	for i <= bit.n {
		bit.arr[i] += int64(val)
		i += i & (-i)
	}
}

func (bit Bit) query(i int) int64 {
	i++
	var res int64
	for i > 0 {
		res += bit.arr[i]
		i -= i & (-i)
	}
	return res
}
