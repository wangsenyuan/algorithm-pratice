package p859

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	n := len(A)

	if n <= 1 {
		return false
	}

	var diff int

	for i := 0; i < n; i++ {
		if A[i] != B[i] {
			diff++
		}
	}

	if diff > 2 {
		return false
	}
	if diff == 1 {
		return false
	}

	if diff == 2 {
		x, y := -1, -1
		for i := 0; i < n; i++ {
			if A[i] != B[i] {
				if x == -1 {
					x = i
				} else {
					y = i
				}
			}
		}

		if A[x] != B[y] || A[y] != B[x] {
			return false
		}
	}

	if diff == 0 {
		cnt := make([]int, 26)

		for i := 0; i < n; i++ {
			j := int(A[i] - 'a')
			cnt[j]++
			if cnt[j] == 2 {
				return true
			}
		}
		return false
	}

	return true
}
