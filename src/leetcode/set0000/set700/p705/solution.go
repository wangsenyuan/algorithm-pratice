package p705

type MyHashSet struct {
	table []bool
}

const H = 1000010

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	table := make([]bool, H)

	return MyHashSet{table}
}

func (this *MyHashSet) Add(key int) {
	this.table[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.table[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.table[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
