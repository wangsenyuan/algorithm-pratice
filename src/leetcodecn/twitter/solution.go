package twitter

import "sort"

type Twitter struct {
	users map[int]*Person
	time  int
}

type Item struct {
	tweetId int
	timeAt  int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].timeAt > items[j].timeAt
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

type Person struct {
	tweets  []Item
	front   int
	end     int
	size    int
	follows map[int]int
	id      int
}

func NewUser(id int) *Person {
	tweets := make([]Item, 11)
	follows := make(map[int]int)
	return &Person{tweets, 0, 0, 0, follows, id}
}

func (user *Person) PostTweet(tweet Item) {
	user.tweets[user.end] = tweet
	user.size++
	if user.size == 11 {
		user.front++
		user.front %= len(user.tweets)
		user.size--
	}
	user.end++
	user.end %= len(user.tweets)
}

func (user *Person) Follow(followeeId int) {
	if user == nil || user.id == followeeId {
		return
	}
	user.follows[followeeId] = 1
}

func (user *Person) UnFollow(followeeId int) {
	if user == nil {
		return
	}
	delete(user.follows, followeeId)
}

func (user *Person) GetTweets() []Item {
	if user == nil {
		return nil
	}
	res := make([]Item, 0, 10)

	n := len(user.tweets)

	for i := user.front; i != user.end; i = (i + 1) % n {
		res = append(res, user.tweets[i])
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	users := make(map[int]*Person)
	var time int
	return Twitter{users, time}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, found := this.users[userId]; !found {
		this.users[userId] = NewUser(userId)
	}
	tweet := Item{tweetId, this.time}
	this.time++
	this.users[userId].PostTweet(tweet)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	res := make([]Item, 0, 100)

	if _, found := this.users[userId]; !found {
		return nil
	}

	follows := this.users[userId].follows

	for k := range follows {
		res = append(res, this.users[k].GetTweets()...)
	}

	res = append(res, this.users[userId].GetTweets()...)

	sort.Sort(Items(res))

	ans := make([]int, min(10, len(res)))

	for i := 0; i < 10 && i < len(res); i++ {
		ans[i] = res[i].tweetId
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, found := this.users[followerId]; !found {
		this.users[followerId] = NewUser(followerId)
	}

	this.users[followerId].Follow(followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.users[followerId].UnFollow(followeeId)
}

/**
* Your Twitter object will be instantiated and called as such:
* obj := Constructor();
* obj.PostTweet(userId,tweetId);
* param_2 := obj.GetNewsFeed(userId);
* obj.Follow(followerId,followeeId);
* obj.Unfollow(followerId,followeeId);
 */
