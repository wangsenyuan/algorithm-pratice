package p1345

import (
	"reflect"
	"testing"
)

func TestSample1(t *testing.T) {
	tweetCounts := Constructor()

	tweetCounts.RecordTweet("tweet3", 0)
	tweetCounts.RecordTweet("tweet3", 60)
	tweetCounts.RecordTweet("tweet3", 10)                                    // All tweets correspond to "tweet3" with recorded times at 0, 10 and 60.
	res := tweetCounts.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59) // return [2]. The frequency is per minute (60 seconds), so there is one interval of time: 1) [0, 60> - > 2 tweets.
	if !reflect.DeepEqual(res, []int{2}) {
		t.Error("Fail")
	}
	res = tweetCounts.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60) // return [2, 1]. The frequency is per minute (60 seconds), so there are two intervals of time: 1) [0, 60> - > 2 tweets, and 2) [60,61> - > 1 tweet.

	if !reflect.DeepEqual(res, []int{2, 1}) {
		t.Error("Fail")
	}

	tweetCounts.RecordTweet("tweet3", 120) // All tweets correspond to "tweet3" with recorded times at 0, 10, 60 and 120.
	res = tweetCounts.GetTweetCountsPerFrequency("hour", "tweet3", 0, 210)

	if !reflect.DeepEqual(res, []int{4}) {
		t.Error("Fail")
	}
}
