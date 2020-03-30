package main

import "fmt"

func main() {
	var u, v uint64
	fmt.Scanf("%d %d", &u, &v)
	res, found := solve(u, v)
	if !found {
		fmt.Println(-1)
	} else {
		fmt.Println(len(res))
		for i := 0; i < len(res); i++ {
			fmt.Printf("%d ", res[i])
		}
	}
}

func solve(u uint64, v uint64) ([]uint64, bool) {
	if u > v || u&1 != v&1 {
		return nil, false
	}

	if u == 0 && v == 0 {
		return nil, true
	}

	if u == v {
		return []uint64{u}, true
	}

	x := (v - u) >> 1
	// u, x, x is an answer with length 3
	for i := 0; i < 63; i++ {
		if (x>>i)&1 == 1 {
			if (u>>i)&1 == 1 {
				// can' be true
				return []uint64{u, x, x}, true
			}
		}
	}

	return []uint64{u + x, x}, true
}
