There is a shop that sells action figures near Monocarp's house. A new set of action figures will be released shortly; this set contains 𝑛
 figures, the 𝑖
-th figure costs 𝑖
 coins and is available for purchase from day 𝑖
 to day 𝑛
.

For each of the 𝑛
 days, Monocarp knows whether he can visit the shop.

Every time Monocarp visits the shop, he can buy any number of action figures which are sold in the shop (of course, he cannot buy an action figure that is not yet available for purchase). If Monocarp buys at least two figures during the same day, he gets a discount equal to the cost of the most expensive figure he buys (in other words, he gets the most expensive of the figures he buys for free).

Monocarp wants to buy exactly one 1
-st figure, one 2
-nd figure, ..., one 𝑛
-th figure from the set. He cannot buy the same figure twice. What is the minimum amount of money he has to spend?

### ideas
1. 是不是倒过来考虑肯合适？
2. 最后的情况下，如果只剩一个，那么把它和前面的合并，从而省下最后的，交换前一个是更优的方案
3. 因为一个的时候无法省钱
4. 