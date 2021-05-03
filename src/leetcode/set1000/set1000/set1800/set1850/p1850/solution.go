package p1850

func getMinSwaps(num string, k int) int {
	buf := []byte(num)
	n := len(buf)
	for i := 0; i < k; i++ {
		j := n - 1
		for j > 0 && buf[j-1] >= buf[j] {
			j--
		}
		j--
		// buf[j] < buf[j+1]
		a := n - 1
		for {
			if buf[a] > buf[j] {
				break
			}
			a--
		}
		// buf[a] > buf[j]
		buf[j], buf[a] = buf[a], buf[j]
		j++
		a = n - 1
		for j < a {
			buf[j], buf[a] = buf[a], buf[j]
			j++
			a--
		}
	}
	var res int
	for i := 0; i < n; i++ {
		if buf[i] != num[i] {
			j := i + 1
			for j < n && buf[j] != num[i] {
				j++
			}
			// buf[j] = num[i]
			for j > i {
				res++
				buf[j] = buf[j-1]
				j--
			}
			buf[i] = num[i]
		}
	}

	return res
}
