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

1. 先考虑一个能work的方法
2. dp[i][j][k] 表示从i开始的后缀问题，其中第i个问题提交的答案是j，且得分是k时的得分
3. dp[i-1][h[i]][k+1] = dp[i][j][k] + 1
4. dp[i-1][u][k] = dp[i][j][k]
5. 也就是说，如果问题i-1提交的答案，正好时h[i]时，+1， 其他的不得分
6. 但是似乎有个问题时，第i个问题的answer，对前面的问题没有影响，只需要去匹配下一个问题的答案就可以了
7. dp[i][j]表示到i为止，得分为j的方案数
8. dp[i][j] = (k-1) * dp[i+1][j] + dp[i+1][j-1]
9. 第i个提交的答案正好时i+1正确答案，和提交其他的答案
10. dp[n][0] = k-1, dp[n][1] = 1
11. n = 3, k = 3
12. dp[3][0] = 2, dp[3][1] = 1
13. 那这样似乎和h[i]没有关系
14. 漏了一个关键点，就是现在的得分必须 > shift前的
15. 如果现在提交的答案是x, 且 x = h[i+1], 那么可以得一分
16. 但是如果 x = h[i], 那么要减一分