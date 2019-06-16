package p990

func equationsPossible(equations []string) bool {
	set := InitUF(26)

	for _, eq := range equations {
		a, b := int(eq[0]-'a'), int(eq[3]-'a')
		flag := eq[1] == '='

		if flag {
			set.Union(a, b)
		}
	}

	for _, eq := range equations {
		a, b := int(eq[0]-'a'), int(eq[3]-'a')
		flag := eq[1] == '='
		if !flag {
			pa := set.Find(a)
			pb := set.Find(b)
			if pa == pb {
				return false
			}
		}
	}

	return true
}

type UF struct {
	set  []int
	size int
}

func InitUF(n int) UF {
	size := n
	set := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
	}
	return UF{set, size}
}

func (uf *UF) Find(u int) int {
	set := uf.set
	var loop func(u int) int

	loop = func(u int) int {
		if set[u] != u {
			set[u] = loop(set[u])
		}
		return set[u]
	}
	return loop(u)
}

func (uf *UF) Union(a, b int) {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa != pb {
		uf.set[pa] = pb
	}
}
