Alice and Bob are playing a game in the shop. There are 𝑛
 items in the shop; each item has two parameters: 𝑎𝑖
 (item price for Alice) and 𝑏𝑖
 (item price for Bob).

Alice wants to choose a subset (possibly empty) of items and buy them. After that, Bob does the following:

if Alice bought less than 𝑘
 items, Bob can take all of them for free;
otherwise, he will take 𝑘
 items for free that Alice bought (Bob chooses which 𝑘
 items it will be), and for the rest of the chosen items, Bob will buy them from Alice and pay 𝑏𝑖
 for the 𝑖
-th item.
Alice's profit is equal to ∑𝑖∈𝑆𝑏𝑖−∑𝑗∈𝑇𝑎𝑗
, where 𝑆
 is the set of items Bob buys from Alice, and 𝑇
 is the set of items Alice buys from the shop. In other words, Alice's profit is the difference between the amount Bob pays her and the amount she spends buying the items.

Alice wants to maximize her profit, Bob wants to minimize Alice's profit. Your task is to calculate Alice's profit if both Alice and Bob act optimally.

### ideas
1. alice选择少于k个没啥意义。因为这时候他的收益是负的。不如啥都不选，收益等于0
2. 所以，alice肯定选了 > k 个
3. 在选择了m个以后，成本等于sum(a)， 收益等于 sum(b least (m - k)) 个
4. 但是从n个里面选择m个，这个复杂性太高。
5. 还需要观察
6. 如果按照b降序排，确实可以很快的计算出收益，但却不一定是成本最低的
7. 选择最贵b个（它们肯定会被free掉）但是它们的成本却不一定是最低的
8. 保证成本最低的时候，（按照a升序），却不一定能保证收益
9. 如果一个item b[i]极大，但是 a[i]极小，那么它肯定应该被选中
10. 难道就是把所有b[i] - a[i] >= 0选中就可以了吗？
11. 然后那些 b[i] - a[i] < 0的部分，假设选中了，且b[i] > 所有其他的，那么它的贡献就是-a[i]
12. 显然是不划算的，如果 b[i]不是最大的那部分，它同样不会有贡献

### solution
Let's sort the array in descending order based on the array 𝑏
. For a fixed set of Alice's items, Bob will take the first 𝑘
 of them for free (because they are the most expensive) and pay for the rest.

Now we can iterate over the first item that Bob will pay (denote it as 𝑖
). Alice has to buy the cheapest 𝑘
 items among 1,2,…,𝑖−1
 (denote the sum of these values as 𝑓
), because Bob can take them for free. Bob has to pay for each of the items among 𝑖,𝑖+1,…,𝑛
 that Alice will buy. So Alice will buy all the items with 𝑏𝑖−𝑎𝑖>0
 (denote the sum of these values as 𝑝
). Then the Alice's profit is 𝑝−𝑓
.

Thus, we got a solution that works in 𝑂(𝑛2)
. In order to speed up this solution, we have to calculate the values 𝑓
 and 𝑝
 faster than 𝑂(𝑛)
. We can do it as follows: while iterating over the value of 𝑖
, let's store "free" items in the ordered set, and when the size of this set becomes larger than 𝑘
, remove the most expensive element from it; and the value of 𝑝
 can be calculated using prefix sums (over the values max(0,𝑏𝑖−𝑎𝑖)
) or maintaining a variable (and update it when moving to the next value of 𝑖
).