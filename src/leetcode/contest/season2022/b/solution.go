package b

type DiscountSystem struct {
	acts      map[int]*Activity
	userCount map[int]map[int]int
}

type Activity struct {
	id         int
	priceLimit int
	number     int
	userLimit  int
	discount   int
	index      int
}

func NewActivity(id int, priceLimit int, discount int, number int, userLimit int) *Activity {
	act := new(Activity)
	act.id = id
	act.priceLimit = priceLimit
	act.discount = discount
	act.number = number
	act.userLimit = userLimit
	return act
}

func Constructor() DiscountSystem {
	acts := make(map[int]*Activity)
	userCount := make(map[int]map[int]int)
	return DiscountSystem{acts, userCount}
}

func (this *DiscountSystem) AddActivity(actId int, priceLimit int, discount int, number int, userLimit int) {

	act := NewActivity(actId, priceLimit, discount, number, userLimit)

	this.acts[actId] = act
	this.userCount[actId] = make(map[int]int)
}

const INF = 1 << 28

func (this *DiscountSystem) RemoveActivity(actId int) {
	delete(this.acts, actId)
	delete(this.userCount, actId)
}

func (this *DiscountSystem) Consume(userId int, cost int) int {
	var picked *Activity
	for id, act := range this.acts {
		if act.number > 0 && act.priceLimit <= cost && this.userCount[id][userId] < act.userLimit {
			if picked == nil || picked.discount < act.discount || picked.discount == act.discount && picked.id > id {
				picked = act
			}
		}
	}
	ret := cost
	if picked != nil {
		picked.number--
		this.userCount[picked.id][userId]++
		ret = cost - picked.discount
	}
	return ret
}

/**
 * Your DiscountSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddActivity(actId,priceLimit,discount,number,userLimit);
 * obj.RemoveActivity(actId);
 * param_3 := obj.Consume(userId,cost);
 */
