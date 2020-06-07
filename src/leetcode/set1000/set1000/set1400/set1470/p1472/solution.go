package p1472

type BrowserHistory struct {
	history []string
	cur     int
}

func Constructor(homepage string) BrowserHistory {

	history := make([]string, 0, 10)
	history = append(history, homepage)
	return BrowserHistory{history, 0}
}

func (this *BrowserHistory) Visit(url string) {
	this.history = this.history[:this.cur+1]
	this.history = append(this.history, url)
	this.cur++
}

func (this *BrowserHistory) Back(steps int) string {
	if this.cur < 0 {
		return ""
	}
	if steps >= this.cur {
		steps = this.cur
	}
	this.cur -= steps
	return this.history[this.cur]
}

func (this *BrowserHistory) Forward(steps int) string {
	if len(this.history)-this.cur-1 < steps {
		steps = len(this.history) - this.cur - 1
	}
	this.cur += steps

	return this.history[this.cur]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
