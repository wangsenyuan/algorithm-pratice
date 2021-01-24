package p1737

func minCharacters(a string, b string) int {
	x := solve1(a, b)
	y := solve2(a, b)
	z := solve2(b, a)

	return min(x, min(y, z))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(a, b string) int {
	// to make a == b with same charaters
	cnt := make([]int, 26)
	for i := 0; i < len(a); i++ {
		cnt[int(a[i]-'a')]++
	}

	for i := 0; i < len(b); i++ {
		cnt[int(b[i]-'a')]++
	}

	var mx int
	for i := 0; i < 26; i++ {
		if cnt[i] > mx {
			mx = cnt[i]
		}
	}

	return len(a) + len(b) - mx
}

func solve2(a, b string) int {
	var best = len(a) + len(b)
	// a < b
	for x := 'b'; x <= 'z'; x++ {
		// all b >= x && all a < x
		tmp := changeLess(a, byte(x)) + changeGe(b, byte(x))
		if tmp < best {
			best = tmp
		}
	}

	return best
}

func changeLess(s string, x byte) int {
	var cnt int
	for i := 0; i < len(s); i++ {
		if s[i] >= x {
			cnt++
		}
	}
	return cnt
}

func changeGe(s string, x byte) int {
	var cnt int
	for i := 0; i < len(s); i++ {
		if s[i] < x {
			cnt++
		}
	}
	return cnt
}
