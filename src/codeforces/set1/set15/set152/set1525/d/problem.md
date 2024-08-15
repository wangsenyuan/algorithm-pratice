There are 𝑛
 armchairs, numbered from 1
 to 𝑛
 from left to right. Some armchairs are occupied by people (at most one person per armchair), others are not. The number of occupied armchairs is not greater than 𝑛2
.

For some reason, you would like to tell people to move from their armchairs to some other ones. If the 𝑖
-th armchair is occupied by someone and the 𝑗
-th armchair is not, you can tell the person sitting in the 𝑖
-th armchair to move to the 𝑗
-th armchair. The time it takes a person to move from the 𝑖
-th armchair to the 𝑗
-th one is |𝑖−𝑗|
 minutes. You may perform this operation any number of times, but these operations must be done sequentially, i. e. you cannot tell a person to move until the person you asked to move in the last operation has finished moving to their destination armchair.

You want to achieve the following situation: every seat that was initially occupied must be free. What is the minimum time you need to do it?

### ideas
1. dp[i][j] = 到i为止，有多少个人，还没有安排好
2. 结果 = dp[n][0]
3. 如果 s[i] = 1, 被占了， dp[i][j] => dp[i][j+1] - i (当前这个人，需要被移动到后面)
4.     s[i] = 0, d[i][j] => dp[i][j-1] + i 有一个人被安排了