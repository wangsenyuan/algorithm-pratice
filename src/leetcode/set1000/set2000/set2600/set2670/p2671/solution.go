package p2671

type FrequencyTracker struct {
	freq map[int]int
	cnt  map[int]int
}

func Constructor() FrequencyTracker {
	freq := make(map[int]int)
	cnt := make(map[int]int)
	return FrequencyTracker{freq, cnt}
}

func (this *FrequencyTracker) Add(number int) {
	if f, ok := this.cnt[number]; !ok {
		// first time
		this.cnt[number]++
		this.freq[1]++
		return
	} else {
		if v, ok2 := this.freq[f]; ok2 {
			if v > 1 {
				this.freq[f] = v - 1
			} else {
				delete(this.freq, f)
			}
		}
		this.cnt[number]++
		this.freq[this.cnt[number]]++
	}
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if f, ok := this.cnt[number]; !ok {
		return
	} else {
		if f == 1 {
			delete(this.cnt, number)
		} else {
			this.cnt[number] = f - 1
		}
		v := this.freq[f]
		// v >= 1
		if v > 1 {
			this.freq[f]--
		} else {
			delete(this.freq, f)
		}
		f--
		if f > 0 {
			this.freq[f]++
		}
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.freq[frequency] > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
