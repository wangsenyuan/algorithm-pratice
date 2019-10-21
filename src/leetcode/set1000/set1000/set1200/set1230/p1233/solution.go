package p1233

import (
	"sort"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)

	var res []string

	for i := 0; i < len(folder); i++ {
		if len(res) == 0 {
			res = append(res, folder[i])
		} else {
			last := res[len(res)-1]
			if !isSubfolder(last, folder[i]) {
				res = append(res, folder[i])
			}
		}
	}

	return res
}

func isSubfolder(a, b string) bool {
	var i int
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	if i < len(a) {
		return false
	}
	// i == len(a)
	return i == len(b) || b[i] == '/'
}
