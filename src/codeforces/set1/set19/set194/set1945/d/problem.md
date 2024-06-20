The guys lined up in a queue of 𝑛
 people, starting with person number 𝑖=1
, to ask Serafim the Owl about the meaning of life. Unfortunately, Kirill was very busy writing the legend for this problem, so he arrived a little later and stood at the end of the line after the 𝑛
-th person. Kirill is completely dissatisfied with this situation, so he decided to bribe some people ahead of him.

For the 𝑖
-th person in the queue, Kirill knows two values: 𝑎𝑖
 and 𝑏𝑖
. If at the moment Kirill is standing at position 𝑖
, then he can choose any position 𝑗
 such that 𝑗<𝑖
 and exchange places with the person at position 𝑗
. In this case, Kirill will have to pay him 𝑎𝑗
 coins. And for each 𝑘
 such that 𝑗<𝑘<𝑖
, Kirill will have to pay 𝑏𝑘
 coins to the person at position 𝑘
. Kirill can perform this action any number of times.

Kirill is thrifty, so he wants to spend as few coins as possible, but he doesn't want to wait too long, so Kirill believes he should be among the first 𝑚
 people in line.

Help Kirill determine the minimum number of coins he will have to spend in order to not wait too long.


### ideas
1. dp[i] 表示 Kirill在位置i时所花的最少的钱数
2. dp[j] = a[j] + sum(j+1....i-1) + dp[i]
3.       = a[j] + suf(j+1) - suf(i) + dp[i]
4.       = a[j] + suf(j+1) + dp[i] - suf[i]