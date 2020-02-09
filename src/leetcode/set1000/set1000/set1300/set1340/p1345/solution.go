package p1345

import "sort"

type TweetCounts struct {
	tweets map[string][]int
}

func Constructor() TweetCounts {
	tweets := make(map[string][]int)

	return TweetCounts{tweets}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	if twts, found := this.tweets[tweetName]; found {
		twts = append(twts, time)
		// sort.Ints(twts)
		this.tweets[tweetName] = twts
	} else {
		twts = make([]int, 0, 100)
		twts = append(twts, time)
		this.tweets[tweetName] = twts
	}
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	interval := 60
	if freq == "hour" {
		interval = 60 * 60
	} else if freq == "day" {
		interval = 24 * 60 * 60
	}

	tot := endTime - startTime + 1

	n := (tot + interval - 1) / interval

	res := make([]int, n)

	twts := this.tweets[tweetName]

	sort.Ints(twts)

	for p := 0; p < n; p++ {
		i := sort.Search(len(twts), func(i int) bool {
			return twts[i] >= min(startTime+p*interval, endTime)
		})

		j := sort.Search(len(twts), func(j int) bool {
			return twts[j] >= min(startTime+(p+1)*interval, endTime+1)
		})

		res[p] = j - i
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */
