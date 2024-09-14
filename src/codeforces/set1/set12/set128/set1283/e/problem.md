Oh, New Year. The time to gather all your friends and reflect on the heartwarming events of the past year...

𝑛
 friends live in a city which can be represented as a number line. The 𝑖
-th friend lives in a house with an integer coordinate 𝑥𝑖
. The 𝑖
-th friend can come celebrate the New Year to the house with coordinate 𝑥𝑖−1
, 𝑥𝑖+1
 or stay at 𝑥𝑖
. Each friend is allowed to move no more than once.

For all friends 1≤𝑥𝑖≤𝑛
 holds, however, they can come to houses with coordinates 0
 and 𝑛+1
 (if their houses are at 1
 or 𝑛
, respectively).

For example, let the initial positions be 𝑥=[1,2,4,4]
. The final ones then can be [1,3,3,4]
, [0,2,3,3]
, [2,2,5,5]
, [2,1,3,5]
 and so on. The number of occupied houses is the number of distinct positions among the final ones.

So all friends choose the moves they want to perform. After that the number of occupied houses is calculated. What is the minimum and the maximum number of occupied houses can there be?

### ideas
1. 在算最小值时，同一个位置的应该一起移动；
2. dp[i][0] 表示前i个人处理完后，且第i个人留在原地时的最小值
3. dp[i][1] 往前移动， dp[i][2]往后移动
4. 现在处理i+1的,他往前移动，当前仅当和前面的人能重叠时，才应该移动；否则他都应该往后移动 
5. 所以可以不用dp，直接用贪心就可以了
6. 计算最大值。
7. 似乎还是可以贪心
8. 现在处理位置0， 如果cnt[1] > 0, 就移动一个到这里来
9. 否则就到位置1
10. 在处理位置i时，如果cnt[i] > 0 且, pos[i-1] is free => 优先填满位置i-1