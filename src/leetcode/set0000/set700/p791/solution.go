package p791

import "sort"

func customSortString(S string, T string) string {
	ord := make([]int, 26)
	for i := 0; i < 26; i++ {
		ord[i] = 26
	}
	for i := 0; i < len(S); i++ {
		ord[S[i]-'a'] = i
	}

	co := CO{[]byte(T), ord}
	sort.Sort(co)

	return string(co.bs)
}

type CO struct {
	bs  []byte
	ord []int
}

func (this CO) Len() int {
	return len(this.bs)
}

func (this CO) Less(i, j int) bool {
	a, b := this.bs[i]-'a', this.bs[j]-'a'
	return this.ord[a] < this.ord[b]
}

func (this CO) Swap(i, j int) {
	this.bs[i], this.bs[j] = this.bs[j], this.bs[i]
}
