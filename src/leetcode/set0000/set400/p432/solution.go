package main

import "fmt"

func main() {
	allOne := Constructor()

	allOne.Inc("hello")
	fmt.Println(allOne.GetMaxKey())
	fmt.Println(allOne.GetMinKey())
	allOne.Inc("hello")

	fmt.Println(allOne.GetMinKey())
	fmt.Println(allOne.GetMaxKey())

	allOne.Inc("leet")

	fmt.Println(allOne.GetMaxKey())
	fmt.Println(allOne.GetMinKey())
}

type AllOne struct {
	kv map[string]int
	vk map[int]map[string]bool
	mx int
	mn int
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{kv: make(map[string]int), vk: make(map[int]map[string]bool), mn: 1 << 30}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if v, ok := this.kv[key]; ok {
		this.kv[key] = v + 1
		delete(this.vk[v], key)
		if len(this.vk[v]) == 0 {
			delete(this.vk, v)
		}
		if this.vk[v+1] == nil {
			this.vk[v+1] = make(map[string]bool)
		}
		this.vk[v+1][key] = true
		if v+1 > this.mx {
			this.mx = v + 1
		}

		if v == this.mn {
			if len(this.vk[v]) == 0 {
				this.mn = v + 1
			}
		}

		return
	}
	this.kv[key] = 1
	if this.vk[1] == nil {
		this.vk[1] = make(map[string]bool)
	}
	this.vk[1][key] = true
	if 1 > this.mx {
		this.mx = 1
	}

	if 1 < this.mn {
		this.mn = 1
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if v, ok := this.kv[key]; ok {
		if v > 1 {
			this.kv[key] = v - 1
			this.vk[v-1][key] = true
		} else {
			delete(this.kv, key)
		}
		delete(this.vk[v], key)
		if v-1 < this.mn {
			this.mn = v - 1
		}
		if v == this.mx {
			if len(this.vk[v]) == 0 {
				this.mx = v - 1
				delete(this.vk, v)
			}
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	vs := this.vk[this.mx]
	for v := range vs {
		return v
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	vs := this.vk[this.mn]
	for v := range vs {
		return v
	}
	return ""
}
