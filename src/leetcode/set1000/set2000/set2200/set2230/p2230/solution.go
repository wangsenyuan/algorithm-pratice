package p2230

import "sort"

func largestInteger(num int) int {
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	n := len(arr)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	var even []*Pair
	var odd []*Pair
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			even = append(even, &Pair{arr[i], i})
		} else {
			odd = append(odd, &Pair{arr[i], i})
		}
	}

	sort.Sort(Pairs(even))
	sort.Sort(Pairs(odd))

	var res int

	for i, j := 0, 0; i < len(even) || j < len(odd); {
		if i == len(even) || (j < len(odd) && even[i].second > odd[j].second) {
			res = res*10 + odd[j].first
			j++
		} else {
			res = res*10 + even[i].first
			i++
		}
	}
	return res
}

type Pair struct {
	first  int
	second int
}

type Pairs []*Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first > this[j].first
}

func (this Pairs) Swap(i, j int) {
	a, b := this[i].second, this[j].second
	this[i], this[j] = this[j], this[i]
	this[i].second = a
	this[j].second = b
}
