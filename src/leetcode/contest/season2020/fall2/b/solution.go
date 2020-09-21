package b

func isMagic(target []int) bool {
	n := len(target)

	if n == 1 {
		return true
	}
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i + 1
	}
	back := make([]int, n)

	shuffle := func(pos int) {
		var j = pos
		for i := pos + 1; i < n; i += 2 {
			back[j] = perm[i]
			j++
		}
		for i := pos; i < n; i += 2 {
			back[j] = perm[i]
			j++
		}
		copy(perm, back)
	}

	shuffle(0)

	var k int

	for k < n && target[k] == perm[k] {
		k++
	}
	if k == 0 {
		return false
	}
	if k == n {
		return true
	}

	pos := k

	for pos < n {
		shuffle(pos)
		j := pos
		for j < pos+k && j < n && target[j] == perm[j] {
			j++
		}
		if j == n {
			break
		}
		if j < pos+k {
			return false
		}
		pos += k
	}
	return true
}
