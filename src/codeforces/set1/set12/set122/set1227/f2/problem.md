The problem is about a test containing 𝑛
one-choice-questions. Each of the questions contains 𝑘
options, and only one of them is correct. The answer to the 𝑖
-th question is ℎ𝑖
, and if your answer of the question 𝑖
is ℎ𝑖
, you earn 1
point, otherwise, you earn 0
points for this question. The values ℎ1,ℎ2,…,ℎ𝑛
are known to you in this problem.

However, you have a mistake in your program. It moves the answer clockwise! Consider all the 𝑛
answers are written in a circle. Due to the mistake in your program, they are shifted by one cyclically.

Formally, the mistake moves the answer for the question 𝑖
to the question 𝑖mod𝑛+1
. So it moves the answer for the question 1
to question 2
, the answer for the question 2
to the question 3
, ..., the answer for the question 𝑛
to the question 1
.

We call all the 𝑛
answers together an answer suit. There are 𝑘𝑛
possible answer suits in total.

You're wondering, how many answer suits satisfy the following condition: after moving clockwise by 1
, the total number of points of the new answer suit is strictly larger than the number of points of the old one. You
need to find the answer modulo 998244353
.

For example, if 𝑛=5
, and your answer suit is 𝑎=[1,2,3,4,5]
, it will submitted as 𝑎′=[5,1,2,3,4]
because of a mistake. If the correct answer suit is ℎ=[5,2,2,3,4]
, the answer suit 𝑎
earns 1
point and the answer suite 𝑎′
earns 4
points. Since 4>1
, the answer suit 𝑎=[1,2,3,4,5]
should be counted.

### ideas

1. 这个感觉要用到矩阵了
2. dp[i][j] = dp[i+1][j] （如果 h[i] = h[i+1])
3.          = dp[i+1][j-1] + dp[i+1][j+1] + (k - 2) * dp[i+1][j]
4. 那么中间那些相等的，其实可以省略掉，只剩下相邻的数字，不相同的序列
5. 那么所有的都是 dp[j] = dp[j-1] + dp[j+1] + (k-2) * dp[j]

### solution

First of all, special judge for 𝑘=1
, where the answer is zero.

Let 𝑑
be the difference between the points for latest answer suit and the previous one. An valid answer suit means 𝑑>0
.

For positions satisfying ℎ𝑖=ℎ𝑖 mod 𝑛+1
, the answer for this position will not affect 𝑑
. Assume there's 𝑡
positions which ℎ𝑖≠ℎ𝑖 mod 𝑛+1
.

For a fixed position 𝑖
which ℎ𝑖≠ℎ𝑖 mod 𝑛+1
, let your answer for this position is 𝑎𝑖
. If 𝑎𝑖=ℎ𝑖
, then the 𝑑
value will decrease by 1. We call this kind of position as a decreasing position. If 𝑎𝑖=ℎ𝑖 mod 𝑛+1
, then the 𝑑
value increase by 1. We call this kind of position as a increasing position. Otherwise 𝑑
value will not be affected, we call this kind of position zero position.

Obviously the number of increasing position should be exact larger then the decreasing position. So let's enumerate the
number of zero positions. We can find the answer is equal to 𝑘𝑛−𝑡×∑0≤𝑖≤𝑡−1[(𝑘−2)𝑖×(𝑡𝑖)×∑[𝑡−𝑖2]+1≤𝑗≤𝑡−𝑖(𝑡−𝑖𝑗)]
. 𝑖
represent the number of zero positions and 𝑗
represent the number of increasing positions.

The only problem is how to calculate ∑[𝑡−𝑖2]+1≤𝑗≤𝑡−𝑖(𝑡−𝑖𝑗)
quickly. Due to (𝑛𝑥)=(𝑛𝑛−𝑥)
, we can tell that when 𝑡−𝑖
is odd, ∑[𝑡−𝑖2]+1≤𝑗≤𝑡−𝑖(𝑡−𝑖𝑗)=2𝑡−𝑖−1
. Otherwise it is equal to 2𝑡−𝑖−(𝑡−𝑖𝑡−𝑖2)2
.