Frog Gorf is traveling through Swamp kingdom. Unfortunately, after a poor jump, he fell into a well of 𝑛
 meters depth. Now Gorf is on the bottom of the well and has a long way up.

The surface of the well's walls vary in quality: somewhere they are slippery, but somewhere have convenient ledges. In other words, if Gorf is on 𝑥
 meters below ground level, then in one jump he can go up on any integer distance from 0
 to 𝑎𝑥
 meters inclusive. (Note that Gorf can't jump down, only up).

Unfortunately, Gorf has to take a break after each jump (including jump on 0
 meters). And after jumping up to position 𝑥
 meters below ground level, he'll slip exactly 𝑏𝑥
 meters down while resting.

Calculate the minimum number of jumps Gorf needs to reach ground level.


### ideas
1. 如果跳跃到i出后，i - a[i] < 0 那么frog就跳出来了
2. dp[i] 表示跳跃到i处所需的最少次数
3. dp[i] = dp[j] + 1 如果 j - x + b[j-x] = i, x <= a[i]
4.   let j - x = k, k + b[k] = i
5.   k >= j - a[j]
6. 所以对于每个位置k，可以计算它落脚的位置i (k + b[k])
7. 通过这层关系，假设知道了d[j], 就没法批量的更新d[i]了，因为它们不是连续的
8. 假设fp[k] = dp[j] + 1 这个是可以批量更新的
9. 即使能批量更新了，似乎还是不大行
10. 好像是可以的，从n到1迭代， dp[i] = fp[k] + 1,
11. 然后使用dp[i]去更新fp[.]
12. 但是有个问题，就是同一个i，可以有多个k
13. 就是有个优化点，是，如果i已经到达了，那么从k过去的就不需要再到i了，（因为i已经有一个更快的路径了）
14. 这是一个图，就是 j 到 j - a[i] ... j - 1范围中的节点，有1的边
15. k到 k + b[k]有0的边
16. 然后计算最短路径
17. 倒是有一个办法，就是范围查询 + 范围更新