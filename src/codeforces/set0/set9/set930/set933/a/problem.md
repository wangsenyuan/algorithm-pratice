A dragon symbolizes wisdom, power and wealth. On Lunar New Year's Day, people model a dragon with bamboo strips and
clothes, raise them with rods, and hold the rods high and low to resemble a flying dragon.

A performer holding the rod low is represented by a 1, while one holding it high is represented by a 2. Thus, the line
of performers can be represented by a sequence a1, a2, ..., an.

Little Tommy is among them. He would like to choose an interval [l, r] (1 ≤ l ≤ r ≤ n), then reverse al, al + 1, ..., ar
so that the length of the longest non-decreasing subsequence of the new sequence is maximum.

A non-decreasing subsequence is a sequence of indices p1, p2, ..., pk, such that p1<p2<...<pk and ap1 ≤ ap2 ≤ ... ≤ apk.
The length of the subsequence is k.

### ideas

1. a[i] = 1, 2
2. 如果 l...r进行 reverse， 那么 对l前没有影响，对r后面没有影响
3. 那么最好的情况是在l前面考虑1，在r的后面考虑2，然后在中间是 1.。。2
4. 所以，这里仅需要考虑dp[l][r]表示，最长的递减序列
5. dp[l][r] = max(dp[l+1][r] + 1 if a[l] = 2, dp[l][r-1] + 1 if a[r] = 1)
6. 如果中间都是1/2，不用处理，貌似也不会错