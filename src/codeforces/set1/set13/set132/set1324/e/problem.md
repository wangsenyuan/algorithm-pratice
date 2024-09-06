Vova had a pretty weird sleeping schedule. There are ℎ
 hours in a day. Vova will sleep exactly 𝑛
 times. The 𝑖
-th time he will sleep exactly after 𝑎𝑖
 hours from the time he woke up. You can assume that Vova woke up exactly at the beginning of this story (the initial time is 0
). Each time Vova sleeps exactly one day (in other words, ℎ
 hours).

Vova thinks that the 𝑖
-th sleeping time is good if he starts to sleep between hours 𝑙
 and 𝑟
 inclusive.

Vova can control himself and before the 𝑖
-th time can choose between two options: go to sleep after 𝑎𝑖
 hours or after 𝑎𝑖−1
 hours.

Your task is to say the maximum number of good sleeping times Vova can obtain if he acts optimally.

#### ideas
1. dp[i][j] 表示到i结束后，是否能够到达时间j
2. dp[i][j] = dp[i-1][j - a[i]] or dp[i-1][j-a[i] + 1]
3. 但是有个问题在于，某个时刻它是good的，那么它就应该被计算，但是它应该只被计算一次
4. 