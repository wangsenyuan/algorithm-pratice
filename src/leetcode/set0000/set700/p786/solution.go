package p786

func kthSmallestPrimeFraction(A []int, K int) []int {

	compare := func(i, j, x, y int) int {
		return A[i]*A[y] - A[j]*A[x]
	}

	rank := func(p, q int) int {
		var res int

		for nu, de := 0, 1; nu < len(A)-1 && de < len(A); {
			if compare(nu, de, p, q) < 0 {
				res += len(A) - de
				nu++
				if nu > de {
					de++
				}
			} else {
				de++
			}
		}
		return res + 1
	}

	nu, de := 0, 1
	for {
		rk := rank(nu, de)
		if rk == K {
			return []int{A[nu], A[de]}
		} else if K < rk {
			de++
		} else {
			nu++
			if nu == de {
				de++
			}
		}
	}
}
