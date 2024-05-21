In the new messenger for the students of the Master's Assistance Center, Keftemerum, an update is planned, in which developers want to optimize the set of messages shown to the user. There are a total of 𝑛
 messages. Each message is characterized by two integers 𝑎𝑖
 and 𝑏𝑖
. The time spent reading the set of messages with numbers 𝑝1,𝑝2,…,𝑝𝑘
 (1≤𝑝𝑖≤𝑛
, all 𝑝𝑖
 are distinct) is calculated by the formula:

∑𝑖=1𝑘𝑎𝑝𝑖+∑𝑖=1𝑘−1|𝑏𝑝𝑖−𝑏𝑝𝑖+1|
Note that the time to read a set of messages consisting of one message with number 𝑝1
 is equal to 𝑎𝑝1
. Also, the time to read an empty set of messages is considered to be 0
.

The user can determine the time 𝑙
 that he is willing to spend in the messenger. The messenger must inform the user of the maximum possible size of the set of messages, the reading time of which does not exceed 𝑙
. Note that the maximum size of the set of messages can be equal to 0
.

The developers of the popular messenger failed to implement this function, so they asked you to solve this problem.

### ideas
1. 不知如何入手
2. 看起来麻烦的就是绝对值那个，看看有没有办法去掉
3. 对于x, y = p[i], p[i+1], 
4. b[x] 和 b[y]的大小是定的
5. 如果x和y是相邻的两个，那么它们的贡献是确定的
6. 假设按照b对 (a, b)进行升序排列后，是不是对结果是没有影响的？
7. 确实是没有影响的，因为我们可以调整p的顺序，从而得到同样的结果
8. 那么现在就是结果就变成了 sum(a[i]) + sum(b[i+1] - b[i])
9. 这样子，岂不是k越大越好？
10. 忘记了l的限制
11. dp[i][j] = x 前i个，使用j个集合的情况下，所能得到的最小值是多少？
12. 这样子b[i]是相互抵消的，只有第一个，和最后一个是有影响的
13. 是否可以用m个，那这m个肯定是连续的