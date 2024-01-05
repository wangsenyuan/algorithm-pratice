There are 𝑛
colorful balls arranged in a row. The balls are painted in 𝑛
distinct colors, denoted by numbers from 1
to 𝑛
. The 𝑖
-th ball from the left is painted in color 𝑐𝑖
. You want to reorder the balls so that the 𝑖
-th ball from the left has color 𝑖
. Additionally, you have 𝑘≥1
balls of color 0
that you can use in the reordering process.

Due to the strange properties of the balls, they can be reordered only by performing the following operations:

Place a ball of color 0
anywhere in the sequence (between any two consecutive balls, before the leftmost ball or after the rightmost ball) while
keeping the relative order of other balls. You can perform this operation no more than 𝑘
times, because you have only 𝑘
balls of color 0
.
Choose any ball of non-zero color such that at least one of the balls adjacent to him has color 0
, and move that ball (of non-zero color) anywhere in the sequence (between any two consecutive balls, before the
leftmost ball or after the rightmost ball) while keeping the relative order of other balls. You can perform this
operation as many times as you want, but for each operation you should pay 1
coin.
You can perform these operations in any order. After the last operation, all balls of color 0
magically disappear, leaving a sequence of 𝑛
balls of non-zero colors.

What is the minimum amount of coins you should spend on the operations of the second type, so that the 𝑖
-th ball from the left has color 𝑖
for all 𝑖
from 1
to 𝑛
after the disappearance of all balls of color zero? It can be shown that under the constraints of the problem, it is
always possible to reorder the balls in the required way.

### thoughts

1. 有一个color 0就可以完成重新排序，
2. 比如放在最右边，然后如果左边的球是期望的颜色，就把它移动到0 球的右边，否则移动到第i个位置
3. 这样子需要的是n的coin
4. 如果找到最右边第一个不在合理位置的球，从那里开始，可以减少cost
5. 但这样子还不足以解决这个问题；
    1. 如何在增加一个新的0色球时，减少cost？
    2. 如何快速的计算新的最优的cost
6. 最关键的问题是，这个操作到底作用什么？
7. dp[i][j][0/1] 表示前i个球，使用j个0色球排序后，且当0色球在最后一个位置（1）时的最优解
8. 如果 c[i+1]是目前的最大值 dp[i+1][j][0] = min(dp[i][j][0], dp[i][j][1]), 当前位置不需要更新
9. dp[i+1][j][1] = dp[i][j][1] + 1 (还需要把它从右边移动到左边)
9. 或者这里增加一个新的0色球，直接放在i+1球的右边
10. dp[i+1][j][1] = min(dp[i][j-1][0], dp[i][j-1][1]) + 1 (when c[i+1] 不是最大值 else 0 when it is a 最大值)
11. 这个策略不是最优的
12. 考虑在为止i，i+1中间放置一个0色球，
13. 可以一直往两边移动，知道遇到某个l balls[l] = l, balls[r] = r

### solution

Let's solve the problem for some fixed 𝑘
.

Consider the set 𝑆
of all balls that were never moved with operation of type 2
. Let's call balls from 𝑆
fixed and balls not from 𝑆
mobile.

The relative order of fixed balls never changes, so their colors must form an increasing sequence.

Let's define 𝑓(𝑆)
as the number of segments of mobile balls that the fixed balls divide sequence into. For example, if 𝑛=6
and 𝑆={3,4,6}
, then these segments are [1,2],[5,5]
and 𝑓(𝑆)
is equal to 2
.

As every mobile ball has to be moved at least once, there must be at least one zero-colored ball in each such segment,
whicn means that 𝑓(𝑆)≤𝑘
. Also, it means that we will need at least 𝑛−|𝑆|
operations of type 2
.

In fact, we can always place mobile balls at correct positions with exactly 𝑛−|𝑆|
operations. The proof is left as exercise for reader.

So the answer for 𝑘
is equal to minimum value of 𝑛−|𝑆|
over all sets 𝑆
of balls such that 𝑓(𝑆)≤𝑘
and the colors of balls in 𝑆
form an increasing sequence. This problem can be solved with dynamic programming: let 𝑑𝑝𝑖,𝑗
be maximum value of |𝑆|
if only balls from 1
to 𝑖
exist, ball 𝑖
belongs to 𝑆
and 𝑓(𝑆)
is equal to 𝑗
. To perform transitions from 𝑑𝑝𝑖,𝑗
you need to enumerate 𝑡
—the next ball from 𝑆
after 𝑖
. Then, if 𝑡=𝑖+1
, you transition to 𝑑𝑝𝑡,𝑗
, otherwise you transition to 𝑑𝑝𝑡,𝑗+1
. After each transition |𝑆|
increases by 1
, so you just update 𝑑𝑝𝑡,𝑗/𝑗+1
with value 𝑑𝑝𝑖,𝑗+1
.

There are 𝑂(𝑛2)
states and at most 𝑛
transitions from each state, so the complexity is 𝑂(𝑛3)
.

Solution can be optimized to 𝑂(𝑛2log𝑛)
with segment tree, but this was not required.