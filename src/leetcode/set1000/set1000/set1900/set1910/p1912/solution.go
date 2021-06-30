package p1912

import "container/heap"

type MovieRentingSystem struct {
	mem   map[int]*Available
	shops map[int]map[int]*Movie
	rents *Rents
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	mem := make(map[int]*Available)

	shops := make(map[int]map[int]*Movie)

	for _, entry := range entries {
		movie := new(Movie)
		movie.shop = entry[0]
		movie.movie = entry[1]
		movie.price = entry[2]
		if _, found := mem[entry[1]]; !found {
			tmp := make(Available, 0, 2)
			mem[entry[1]] = &tmp
		}

		heap.Push(mem[entry[1]], movie)

		if _, found := shops[entry[0]]; !found {
			shops[entry[0]] = make(map[int]*Movie)
		}

		shops[entry[0]][entry[1]] = movie
	}

	rents := make(Rents, 0, 10)

	return MovieRentingSystem{mem, shops, &rents}
}

func (this *MovieRentingSystem) Search(movie int) []int {
	res := make([]int, 0, 5)

	arr := make([]*Movie, 0, 5)
	avi := this.mem[movie]

	cnt := 5
	for avi != nil && avi.Len() > 0 && cnt > 0 {
		cur := heap.Pop(avi).(*Movie)
		res = append(res, cur.shop)
		arr = append(arr, cur)
		cnt--
	}

	for i := 0; i < len(arr); i++ {
		heap.Push(avi, arr[i])
	}

	return res
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	avi := this.mem[movie]
	cur := this.shops[shop][movie]

	price := cur.price
	// make cur.prive as 0, to pop it from avi
	avi.updatePrice(cur, 0)
	heap.Pop(avi)

	tmp := new(Movie)
	tmp.shop = shop
	tmp.movie = movie
	tmp.price = price

	heap.Push(this.rents, tmp)

	// update the reference
	this.shops[shop][movie] = tmp
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	tmp := this.shops[shop][movie]
	price := tmp.price

	this.rents.updatePrice(tmp, 0)
	heap.Pop(this.rents)

	cur := new(Movie)
	cur.shop = shop
	cur.movie = movie
	cur.price = price
	heap.Push(this.mem[movie], cur)
	this.shops[shop][movie] = cur
}

func (this *MovieRentingSystem) Report() [][]int {
	res := make([][]int, 0, 5)
	arr := make([]*Movie, 0, 5)

	cnt := 5

	for cnt > 0 && this.rents.Len() > 0 {
		cnt--
		tmp := heap.Pop(this.rents).(*Movie)
		res = append(res, []int{tmp.shop, tmp.movie})
		arr = append(arr, tmp)
	}

	for i := 0; i < len(arr); i++ {
		heap.Push(this.rents, arr[i])
	}

	return res
}

type Movie struct {
	shop  int
	movie int
	price int
	index int
}

type Available []*Movie

func (this Available) Len() int {
	return len(this)
}

func (this Available) Less(i, j int) bool {
	return this[i].price < this[j].price || this[i].price == this[j].price && this[i].shop < this[j].shop
}

func (this Available) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *Available) Push(x interface{}) {
	i := x.(*Movie)
	i.index = len(*this)
	*this = append(*this, i)
}

func (this *Available) Pop() interface{} {
	old := *this
	n := len(old)
	res := old[n-1]
	*this = old[:n-1]
	res.index = -1
	return res
}

func (this *Available) updatePrice(movie *Movie, price int) {
	movie.price = price
	heap.Fix(this, movie.index)
}

type Rents []*Movie

func (this Rents) Len() int {
	return len(this)
}

func (this Rents) Less(i, j int) bool {
	return this[i].price < this[j].price ||
		this[i].price == this[j].price && this[i].shop < this[j].shop ||
		this[i].price == this[j].price && this[i].shop == this[j].shop && this[i].movie < this[j].movie
}

func (this Rents) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *Rents) Push(x interface{}) {
	i := x.(*Movie)
	i.index = len(*this)
	*this = append(*this, i)
}

func (this *Rents) Pop() interface{} {
	old := *this
	n := len(old)
	res := old[n-1]
	*this = old[:n-1]
	res.index = -1
	return res
}

func (this *Rents) updatePrice(movie *Movie, price int) {
	movie.price = price
	heap.Fix(this, movie.index)
}
