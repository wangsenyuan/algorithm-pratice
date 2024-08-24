You are given an array 𝑎1,𝑎2,…,𝑎𝑛
, consisting of 𝑛
 positive integers.

Initially you are standing at index 1
 and have a score equal to 𝑎1
. You can perform two kinds of moves:

move right — go from your current index 𝑥
 to 𝑥+1
 and add 𝑎𝑥+1
 to your score. This move can only be performed if 𝑥<𝑛
.
move left — go from your current index 𝑥
 to 𝑥−1
 and add 𝑎𝑥−1
 to your score. This move can only be performed if 𝑥>1
. Also, you can't perform two or more moves to the left in a row.
You want to perform exactly 𝑘
 moves. Also, there should be no more than 𝑧
 moves to the left among them.

What is the maximum score you can achieve?

#### ideas
1. 这里有两个点，一个是向左的总次数，不能超过z, 还有一个是，不能有连续两次向左
2. 所以，通过第二点，可以知道，向左后，只能马上回来
3. dp[i][j] 表示目前在位置i，且使用了j次向左（且返回了位置i）的最大值
4. dp[i][j] = dp[i-1][j] + a[i] (不使用向左的)
5.            dp[i][j-1] + a[i] + a[i-1] (多使用一次向左，但是马上要回来)