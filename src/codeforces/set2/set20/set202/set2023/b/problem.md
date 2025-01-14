It is already the year 3024
, ideas for problems have long run out, and the olympiad now takes place in a modified individual format. The olympiad consists of 𝑛
 problems, numbered from 1
 to 𝑛
. The 𝑖
-th problem has its own score 𝑎𝑖
 and a certain parameter 𝑏𝑖
 (1≤𝑏𝑖≤𝑛
).

Initially, the testing system gives the participant the first problem. When the participant is given the 𝑖
-th problem, they have two options:

They can submit the problem and receive 𝑎𝑖
 points;
They can skip the problem, in which case they will never be able to submit it.
Then, the testing system selects the next problem for the participant from problems with indices 𝑗
, such that:

If he submitted the 𝑖
-th problem, it looks at problems with indices 𝑗<𝑖
;
If he skipped the 𝑖
-th problem, it looks at problems with indices 𝑗≤𝑏𝑖
.
Among these problems, it selects the problem with the maximum index that it has not previously given to the participant (he has neither submitted nor skipped it before). If there is no such problem, then the competition for the participant ends, and their result is equal to the sum of points for all submitted problems. In particular, if the participant submits the first problem, then the competition for them ends. Note that the participant receives each problem at most once.

Prokhor has prepared thoroughly for the olympiad, and now he can submit any problem. Help him determine the maximum number of points he can achieve.

### ideas
1. 如果提交了1, 那么得分 = a[1] (结束)
2. 如果skip1， 那么就考虑b[1]
3. 假设希望能够submit i， 那么必须至少要通过某种方式到到 i以后的位置
4. 如果能够直接到到n，那么就一直submit回来
5. 到i为止的分数值和 sum[i] - 那些为了达到位置i，跳过的最小的分数之和
6. fp[i] = a[j] (b[j] = i) + fp[j] ? (这个是其中的一种情况， 就是到达位置j以后，直接跳到位置i)
7.       = a[j] + fp[j....i-1] 到达j的上方后，一直提交，直到到达位置j
8.       = fp[i+1...] (这里就出现问题了，要先计算出i+1，才能知道i的结果)
9. 如果fp[i] 最高到达i时候的最小的跳跃损失的话,
10. fp[i] = a[j] + min(fp[j....i-1])
11. 