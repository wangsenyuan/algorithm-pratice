package p2074

type Robot struct {
	w     int
	h     int
	l     int
	x     int
	y     int
	dir   int
	moves int
}

var dirs = []string{"North", "East", "South", "West"}

// 0 for N, 1 for E, 2 for S, 3 for W

func Constructor(width int, height int) Robot {
	l := width - 1 + height - 1
	l *= 2
	return Robot{width, height, l, 0, 0, 1, 0}
}

func (this *Robot) Move(steps int) {
	this.moves += steps
	num := this.moves % this.l

	if num == 0 {
		this.x = 0
		this.y = 0
		this.dir = 2
		return
	}

	var x, y int
	d := 1
	if num < this.w {
		x = num
	} else {
		x = this.w - 1
		num -= x
		d = 0
		if num < this.h {
			y = num
		} else {
			y = this.h - 1
			num -= y
			d = 3
			if num < this.w {
				x = this.w - 1 - num
			} else {
				d = 2
				x = 0
				num -= this.w - 1
				y = this.h - 1 - num
			}
		}
	}

	this.x = x
	this.y = y
	this.dir = d
}

func (this *Robot) GetPos() []int {
	return []int{this.x, this.y}
}

func (this *Robot) GetDir() string {
	return dirs[this.dir]
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Move(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
