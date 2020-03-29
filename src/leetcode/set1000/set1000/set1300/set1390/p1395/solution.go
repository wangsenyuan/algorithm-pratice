package p1395

type Event struct {
	id      int
	station string
	time    int
}

type Pair struct {
	first  int64
	second int64
}

type UndergroundSystem struct {
	checkIn  map[int]Event
	duration map[string]map[string]*Pair
}

func Constructor() UndergroundSystem {
	checkIn := make(map[int]Event)
	duration := make(map[string]map[string]*Pair)

	return UndergroundSystem{checkIn, duration}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	event := Event{id, stationName, t}
	this.checkIn[id] = event
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	event := this.checkIn[id]

	from := event.station

	if _, found := this.duration[from]; !found {
		this.duration[from] = make(map[string]*Pair)
	}

	prev := this.duration[from]

	if _, found := prev[stationName]; !found {
		prev[stationName] = &Pair{int64(t - event.time), 1}
	} else {
		prev[stationName].first += int64(t - event.time)
		prev[stationName].second++
	}

	delete(this.checkIn, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	entry := this.duration[startStation][endStation]
	x := entry.first
	y := entry.second
	return float64(x) / float64(y)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
