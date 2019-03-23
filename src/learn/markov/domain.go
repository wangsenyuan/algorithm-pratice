package main

import "hash/fnv"

const SIZE_TABLE = 32

type Suffix struct {
	word string
	next *Suffix
}

type State struct {
	prefix []string
	suffix *Suffix
	next   *State
}

type Map struct {
	table []*State
	size  int
}

func NewMap(hint int) Map {
	size := nextPow2(hint)
	table := make([]*State, size)
	return Map{table, size}
}

func nextPow2(n int) int {
	sz := 1
	for sz*2 <= n {
		sz *= 2
	}
	return sz
}

func (this *Map) AddWord(prefix []string, word string) {
	state := loopup(this, prefix, true)

	suffix := new(Suffix)
	suffix.word = word
	suffix.next = state.suffix
	state.suffix = suffix
}

func (this *Map) Lookup(s []string) *State {
	return loopup(this, s, false)
}

func loopup(mp *Map, s []string, create bool) *State {
	h := hash(s) % uint32(mp.size)
	state := mp.table[h]

	for p := state; p != nil; p = p.next {
		if sameStrs(p.prefix, s) {
			return p
		}
	}

	if create {
		state = new(State)
		state.prefix = make([]string, len(s))
		copy(state.prefix, s)
		state.next = mp.table[h]
		mp.table[h] = state
		return state
	}
	return nil
}

func sameStrs(a, b []string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func hash(s []string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s[0]))
	h.Write([]byte(s[1]))
	return h.Sum32()
}
