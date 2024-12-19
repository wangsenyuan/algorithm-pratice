Suppose you are working in some audio streaming service. The service has 𝑛
 active users and 109
 tracks users can listen to. Users can like tracks and, based on likes, the service should recommend them new tracks.

Tracks are numbered from 1
 to 109
. It turned out that tracks the 𝑖
-th user likes form a segment [𝑙𝑖,𝑟𝑖]
.

Let's say that the user 𝑗
 is a predictor for user 𝑖
 (𝑗≠𝑖
) if user 𝑗
 likes all tracks the 𝑖
-th user likes (and, possibly, some other tracks too).

Also, let's say that a track is strongly recommended for user 𝑖
 if the track is not liked by the 𝑖
-th user yet, but it is liked by every predictor for the 𝑖
-th user.

Calculate the number of strongly recommended tracks for each user 𝑖
. If a user doesn't have any predictors, then print 0
 for that user.

Input
The first line contains one integer 𝑡
 (1≤𝑡≤104
) — the number of test cases. Next, 𝑡
 cases follow.

The first line of each test case contains one integer 𝑛
 (1≤𝑛≤2⋅105
) — the number of users.

The next 𝑛
 lines contain two integers 𝑙𝑖
 and 𝑟𝑖
 per line (1≤𝑙𝑖≤𝑟𝑖≤109
) — the segment of tracks the 𝑖
-th user likes.

Additional constraint on the input: the sum of 𝑛
 over all test cases doesn't exceed 2⋅105
.

Output
For each test case, print 𝑛
 integers, where the 𝑖
-th integer is the number of strongly recommended tracks for the 𝑖
-th user (or 0
, if that user doesn't have any predictors).

### ideas
1. 假设用户a，他喜欢的歌单是从10...20, 然后有用户B喜欢的歌单是从5...21
2. 那么用户b是a的前导
3. 同时c也是a的前导，且c喜欢的歌单是 7.....22
4. 那么这里7...9,21就是可以推荐给a的曲目
5. a的所有前导，就是那些在la前面开始，且在lb后面结束的部分
6. 在la的前面的最大值，和lb后面的最小值
7. 是不是也可以从大区间到小区间的处理呢？