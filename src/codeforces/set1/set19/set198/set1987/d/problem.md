Alice and Bob are playing a game. Initially, there are 𝑛
 cakes, with the 𝑖
-th cake having a tastiness value of 𝑎𝑖
.

Alice and Bob take turns eating them, with Alice starting first:

In her turn, Alice chooses and eats any remaining cake whose tastiness is strictly greater than the maximum tastiness of any of the cakes she's eaten before that. Note that on the first turn, she can choose any cake.
In his turn, Bob chooses any remaining cake and eats it.
The game ends when the current player can't eat a suitable cake. Let 𝑥
 be the number of cakes that Alice ate. Then, Alice wants to maximize 𝑥
, while Bob wants to minimize 𝑥
.

Find out how many cakes Alice will eat if both players play optimally.

### ideas
1. alice的策略是尽量选择当前最低的那个，这样子她后续的限制更小
2. bob的策略似乎有点复杂，
3. 如果把最高的删除掉，似乎不一定是最优的，假设最高的有2个，但是次高的有一个，那么先取掉次高的似乎是更优的选择
4. alice同一个数，只能取一次。但是bob却可以取多次。且bob一旦开始取这个数的时候，必须保证在alice reach前能够全部取完
5. 所以bob除了第一个数外，他应该取尽量个数少的那部分
6. 假设alice第一个取的数是x，那么bob应该取（x+1)后面的个数最少的那些。
7. 如果 x+2有两个，但是 y > x + ?, 只有一个，那么先取x+2似乎是一个更优的策略
8. 因为，如果先取y，虽然可以保证y不被取到，但是x+2却会被取掉
9. 所以不大对
10. alice的策略其实还是固定的，就是尽量的取当前最小的那个
11. 在这个策略的前提下，dp[i] = dp[i+1] 如果 a[i] != a[i+1], 但是这样子，没法体现处bob的操作
12. bob处理后，alice的策略似乎没有变化，这就很麻烦了。但是显然不是这样子的
13. bob的策略肯定会影响到alice才对
14. dp[i][j] 表示前i个处理后，alice已经获得了j个cake
15. 但是这里关键的地方在于，甚至都不能确定alice、或者bob现在在操作的是哪个？
16. dp[i] = 从i开始的最大值
17. dp[i] = min(i...j 里面 alice 能够获得的cakes + dp[j])
18. 现在的问题就变成了，如何计算[i...j]这段内，进行偶数次操作，且在最优的情况下，alice能获得的最大值？
19. 这个是不是就变成了一个子问题？
20. 如果 cnt[j] <= distinct(i...j-1), 那么 dp[i] = distinct(i...j-1) + dp[j+1] 
21. 假设在i...j段内，alice得分为w，那么
22. 感觉还得用优先级队列
23. 就是如果当前的j可以被bob全部取完时，可以考虑将其全部取掉（以阻止alice获取）
24. 但是如果来不及取完时，可以反悔前面的，也就是操作次数最多的那个，把当前的给取掉
25. dp[i][k] = dp[i][k-1] + arr[i] if dp[i][k-1] + arr[i] <= i - k
26. dp[k] = dp[k-1] + arr[i] if dp[k-1] + arr[i] <= i - k
27. dp[k-1] + k <= i - arr[i]
28. 也就是说对于那些 dp[k-1] + k <= i - arr[i]的部分，dp[k] = min(dp[k], dp[k-1] + arr[i])
29. 

### solution
Let's consider Alice's strategy. Notice that if both players play optimally, Alice eating a cake with a tastiness value of 𝑡
 is equivalent to her eating all remaining cakes with 𝑎𝑖≤𝑡
. Since it is better for her to have more cakes to choose from later on, she will choose the minimum possible 𝑡
 on each turn.

Now let's consider Bob's strategy. Suppose Bob ate at least one cake with a tastiness value of 𝑡
. Then he has to have eaten all of them, because eating only some of them does not affect Alice's possible moves, resulting in wasted turns (he might be forced to waste moves at the end of the game when there are some leftover cakes).

Let 𝐴1,…,𝐴𝑚
 be the sorted unique values of 𝑎1,…,𝑎𝑛
, and let 𝑐𝑖
 be the number of occurrences of 𝐴𝑖
 in 𝑎
. Then the game is reduced to Alice finding the first 𝑐𝑖>0
 and assigning 𝑐𝑖:=0
, and Bob choosing any 𝑐𝑖>0
 and decreasing it by 1
.

Since Bob's optimal strategy is for each 𝑐𝑖
 either not to touch it or make 𝑐𝑖=0
, we can model it as him selecting some set of indices 1≤𝑖1<…<𝑖𝑘≤𝑚
 and zeroing out 𝑐𝑖1,𝑐𝑖2,…,𝑐𝑖𝑘
 in this order. To be able to zero out 𝑐𝑖𝑝
 (1≤𝑝≤𝑘
), Alice must not have gotten to the value 𝑐𝑖𝑝
. In 𝑐𝑖1+…+𝑐𝑖𝑝
 turns, Alice will zero out exactly that many values. Additionally, Bob would have zeroed out 𝑝−1
 values before index 𝑖𝑝
. So, 𝑐𝑖1+…+𝑐𝑖𝑝+𝑝−1
 must be less than 𝑖𝑝
. Transforming this a bit, we get that the condition ∑𝑝𝑗=1𝑐𝑖𝑗≤𝑖𝑝−𝑝
 must hold for all 1≤𝑝≤𝑘
.

Bob's objective to maximize the size of the set of incides 1≤𝑖1<…<𝑖𝑘≤𝑚
. Let 𝑑𝑝[𝑖][𝑘]=
 the minimum possible ∑𝑘𝑗=1𝑐𝑖𝑗
 over all valid sets of indices 1≤𝑖1<…<𝑖𝑘≤𝑖
. Initialize 𝑑𝑝[0][0]=0
, and everything else to +∞
. The main idea is to grow valid sets of indices one index at a time.

We will iterate over all 𝑖
 from 1
 to 𝑚
, and then for each 𝑘
 from 0
 to 𝑚
. The transitions are:

𝑑𝑝[𝑖][𝑘]=min(𝑑𝑝[𝑖][𝑘],𝑑[𝑖−1][𝑘])
, which corresponds to not using the current index in the set.
Let 𝑠:=𝑑𝑝[𝑖−1][𝑘−1]+𝑐𝑖
. If 𝑠≤𝑖𝑘−𝑘
, update 𝑑𝑝[𝑖][𝑘]=min(𝑑𝑝[𝑖][𝑘],𝑠)
, which is equivalent to adding the current index to the set. We only need to check the condition on the last index because it is already satisfied for all the previous indices.
The answer to the problem is 𝑚−𝑦
, where 𝑦
 is the largest value where 𝑑𝑝[𝑚][𝑦]≤+∞
.

Complexity: (𝑛2)
Note: it is possible to solve in (𝑛log𝑛)
 by making an additional observation.