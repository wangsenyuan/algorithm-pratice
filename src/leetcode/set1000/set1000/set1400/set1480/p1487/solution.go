package p1487

import "fmt"

func getFolderNames(names []string) []string {
	seq := make(map[string]int)
	mem := make(map[string]bool)

	n := len(names)

	res := make([]string, n)

	for i := 0; i < n; i++ {
		name := names[i]

		if !mem[name] {
			seq[name] = 0
			mem[name] = true
			res[i] = name
			continue
		}

		j := seq[name] + 1

		for {
			tmp := fmt.Sprintf("%s(%d)", name, j)
			if !mem[tmp] {
				res[i] = tmp
				mem[tmp] = true
				seq[name] = j
				break
			}
			j++
		}
	}

	return res
}
