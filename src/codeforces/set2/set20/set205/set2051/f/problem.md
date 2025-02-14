Consider a deck of 𝑛
 cards. The positions in the deck are numbered from 1
 to 𝑛
 from top to bottom. A joker is located at position 𝑚
.

𝑞
 operations are applied sequentially to the deck. During the 𝑖
-th operation, you need to take the card at position 𝑎𝑖
 and move it either to the beginning or to the end of the deck. For example, if the deck is [2,1,3,5,4]
, and 𝑎𝑖=2
, then after the operation the deck will be either [1,2,3,5,4]
 (the card from the second position moved to the beginning) or [2,3,5,4,1]
 (the card from the second position moved to the end).

Your task is to calculate the number of distinct positions where the joker can be after each operation.


### ideas
1. 一开始joker在的位置是i
2. 如果取走的是它上面的一张牌（比如i-1), 并放回到顶部 => 位置不变，还是i
3. 如果放回到底部，那么joker现在的位置就变成了i-1
4. 同理如果取走的是下面的一张牌，新的位置，有可能变成i+1
5. 如果取走的就是i，新的位置就变成了 1或者n
6. 所以把这个牌堆看作一个环是更合理的，如果取走的位置是可能位置的前面，那么这些位置往前移动一次
7. 如果取走的是后面的位置，往后移动
8. 如果是取走的是可能位置的一个，那么久需要产生新的位置（如果已经包含了1/n， 似乎是没有效果的？)
9. 假设有个数据结构，它保存的是某个区间的开始位置，以及它的大小
10. 那么在这个区间的后面取走一张牌，所有的这些区间都变大1
11. 感觉最多有3个区间
12. 一开始的区间，在某个时刻出现的两头的区间，然后通过位置判断，这三个区间的变化情况（有可能会merge在一起）