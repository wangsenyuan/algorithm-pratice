Recently, Dima met with Sasha in a philatelic store, and since then they are collecting coins together. Their favorite
occupation is to sort collections of coins. Sasha likes having things in order, that is why he wants his coins to be
arranged in a row in such a way that firstly come coins out of circulation, and then come coins still in circulation.

For arranging coins Dima uses the following algorithm. One step of his algorithm looks like the following:

He looks through all the coins from left to right;
If he sees that the i-th coin is still in circulation, and (i + 1)-th coin is already out of circulation, he exchanges
these two coins and continues watching coins from (i + 1)-th.
Dima repeats the procedure above until it happens that no two coins were exchanged during this procedure. Dima calls
hardness of ordering the number of steps required for him according to the algorithm above to sort the sequence, e.g.
the number of times he looks through the coins from the very beginning. For example, for the ordered sequence hardness
of ordering equals one.

Today Sasha invited Dima and proposed him a game. First he puts n coins in a row, all of them are out of circulation.
Then Sasha chooses one of the coins out of circulation and replaces it with a coin in circulation for n times. During
this process Sasha constantly asks Dima what is the hardness of ordering of the sequence.

The task is more complicated because Dima should not touch the coins and he should determine hardness of ordering in his
mind. Help Dima with this task.

### ideas

1. Let's denote as O coin out of circulation, and as X — coin is circulation.
2. 考虑在给定的一个状态下，如何将O移动到前面， X移动到后面
3. 最左边的X，会移动到下一个X的前面；然后下一个X移动到再下一个X的前面，直到最后一个位置
4. 处理的次数 = 不在expect位置的X的个数
5. 一次操作，能够将一个不在位置的X移动到它应在的位置
6. 假设总数是cnt个X到目前位置，那么操作的次数就是1...n - cnt 区间内的X的数量