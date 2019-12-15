package p1286

type CombinationIterator struct {
	str     string
	comb    int
	n       int
	pos     []int
	cur     int
	hasMore bool
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	n := len(characters)
	// abc, 2 => ab, ac, bc
	// abcd, 2 => ab, ac, ad, bc, bd, cd
	pos := make([]int, combinationLength)

	for i := 0; i < combinationLength; i++ {
		pos[i] = i
	}

	// init pos [0, 1]

	return CombinationIterator{characters, combinationLength, n, pos, combinationLength - 1, true}
}

func (this *CombinationIterator) Next() string {
	comp := this.comb
	pos := this.pos
	str := this.str

	res := make([]byte, comp)

	for i := 0; i < comp; i++ {
		res[i] = str[pos[i]]
	}
	cur := this.cur
	n := this.n - 1
	j := cur

	for j >= 0 && pos[j] == n {
		j--
		n--
	}

	if j < 0 {
		this.hasMore = false
	} else {
		pos[j]++
		for j < this.comb-1 {
			pos[j+1] = pos[j] + 1
			j++
		}
	}

	return string(res)
}

func (this *CombinationIterator) HasNext() bool {
	return this.hasMore
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
