package p2211

func countCollisions(directions string) int {

	var res int

	n := len(directions)

	var i int

	for i < n && directions[i] == 'L' {
		i++
	}

	for i < n {
		var r_cnt int
		for i < n && directions[i] == 'R' {
			r_cnt++
			i++
		}
		if i == n {
			break
		}
		// dir[i] == 'L'
		res += r_cnt

		if directions[i] == 'S' {
			for i < n && directions[i] == 'S' {
				i++
			}
			if i == n || directions[i] == 'R' {
				continue
			}
		}

		for i < n && directions[i] == 'L' {
			res++
			i++
		}
	}

	return res
}
