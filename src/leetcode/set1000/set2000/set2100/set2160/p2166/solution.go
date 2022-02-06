package p2166

import "bytes"

type Bitset struct {
	size int
	bits map[int]bool
	flip bool
}

func Constructor(size int) Bitset {
	bits := make(map[int]bool)
	return Bitset{size, bits, false}
}

func (this *Bitset) Fix(idx int) {
	if !this.flip {
		this.bits[idx] = true
	} else {
		// fliped
		if this.bits[idx] {
			delete(this.bits, idx)
		}
	}
}

func (this *Bitset) Unfix(idx int) {
	if !this.flip {
		delete(this.bits, idx)
	} else {
		if !this.bits[idx] {
			this.bits[idx] = true
		}
	}
}

func (this *Bitset) Flip() {
	this.flip = !this.flip
}

func (this *Bitset) All() bool {
	if !this.flip {
		return len(this.bits) == this.size
	}
	return len(this.bits) == 0
}

func (this *Bitset) One() bool {
	if !this.flip {
		return len(this.bits) > 0
	}
	return len(this.bits) < this.size
}

func (this *Bitset) Count() int {
	if !this.flip {
		return len(this.bits)
	}
	return this.size - len(this.bits)
}

func (this *Bitset) ToString() string {
	var buf bytes.Buffer
	if !this.flip {
		for i := 0; i < this.size; i++ {
			if this.bits[i] {
				buf.WriteByte('1')
			} else {
				buf.WriteByte('0')
			}
		}
	} else {
		for i := 0; i < this.size; i++ {
			if this.bits[i] {
				buf.WriteByte('0')
			} else {
				buf.WriteByte('1')
			}
		}
	}
	return buf.String()
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */
