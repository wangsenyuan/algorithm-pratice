Princess Heidi decided to give orders to all her K Rebel ship commanders in person. Unfortunately, she is currently travelling through hyperspace, and will leave it only at N specific moments t1, t2, ..., tN. The meetings with commanders must therefore start and stop at those times. Namely, each commander will board her ship at some time ti and disembark at some later time tj. Of course, Heidi needs to meet with all commanders, and no two meetings can be held during the same time. Two commanders cannot even meet at the beginnings/endings of the hyperspace jumps, because too many ships in one position could give out their coordinates to the enemy.

Your task is to find minimum time that Princess Heidi has to spend on meetings, with her schedule satisfying the conditions above.

### ideas
1. sort t first
2. dp[i][j] = 表示在t[i]处结束会见j个指挥官时的最小花费时间
3. dp[i][j] = min of dp[ii][j-1] +  t[i] - t[ii]（这个可以记录dp[i][j-1] - t[i]的最小值，达到 n * k的复杂性)
4. 但还是太慢了
5. 如果把相邻位置减一减，得到x1, x2, x3....那么如果选择了x1, 那么就不能选择x2,如果选择了x2, 就不能选择x1，x3
6. 