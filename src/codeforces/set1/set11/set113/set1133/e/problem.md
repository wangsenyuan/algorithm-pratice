You are a coach at your local university. There are 𝑛
 students under your supervision, the programming skill of the 𝑖
-th student is 𝑎𝑖
.

You have to form 𝑘
 teams for yet another new programming competition. As you know, the more students are involved in competition the more probable the victory of your university is! So you have to form no more than 𝑘
 (and at least one) non-empty teams so that the total number of students in them is maximized. But you also know that each team should be balanced. It means that the programming skill of each pair of students in each team should differ by no more than 5
. Teams are independent from one another (it means that the difference between programming skills of two students from two different teams does not matter).

It is possible that some students not be included in any team at all.

Your task is to report the maximum possible total number of students in no more than 𝑘
 (and at least one) non-empty balanced teams.

 ### ideas
 1. sort a first
 2. dp[i][j][d] 表示到i为止，有j个team，且第一位和最后一位的差，不超过d，最多人数
 3. dp[i][j][0] = dp[i-1][j-1][?] + 1, dp[i-1][j][0] 如果 a[i] == a[i-1]
 4. 