Baby Ehab has a piece of Cut and Stick with an array 𝑎
of length 𝑛
written on it. He plans to grab a pair of scissors and do the following to it:

pick a range (𝑙,𝑟)
and cut out every element 𝑎𝑙
, 𝑎𝑙+1
, ..., 𝑎𝑟
in this range;
stick some of the elements together in the same order they were in the array;
end up with multiple pieces, where every piece contains some of the elements and every element belongs to some piece.
More formally, he partitions the sequence 𝑎𝑙
, 𝑎𝑙+1
, ..., 𝑎𝑟
into subsequences. He thinks a partitioning is beautiful if for every piece (subsequence) it holds that, if it has
length 𝑥
, then no value occurs strictly more than ⌈𝑥/2⌉
times in it.

He didn't pick a range yet, so he's wondering: for 𝑞
ranges (𝑙,𝑟)
, what is the minimum number of pieces he needs to partition the elements 𝑎𝑙
, 𝑎𝑙+1
, ..., 𝑎𝑟
into so that the partitioning is beautiful.

A sequence 𝑏
is a subsequence of an array 𝑎
if 𝑏
can be obtained from 𝑎
by deleting some (possibly zero) elements. Note that it does not have to be contiguous.

### ideas

1. 每段的长度不需要相同
2. 假设其中有一段的长度是x, 且x是偶数
3. 那么将x+1没有坏处。这是因为，在增加前，没有数超过一半, 假设正好一半，x/2, 如果新添加的数是这个多数，它不会破坏这个性质
4. 如果不是，还可以继续增加。
5. 但是会破坏后面的那段吗？ [2, 2, 1, 1] [1, 2, 2]
6. 貌似确实会破坏后面这段的性质
7. 这么复杂，感觉答案要么是1，要么是2，否则无从下手呐
8. 也不对哦。如果全部一样 =》 只能是这个长度
9. 如果x能成立，那么 x + 1是否一定成立？
10. 比如上面的分成两段是成立的，那么分成整段，似乎也是成立的
11. 怎么证明？[l0...r0] [l1...r1] ... [li, ri] ...
12. 每段都是满足 最多的数不超过一半
13. 极端的情况，每个都是奇数，且正都是一半，且都相同 [1, 2, 2], [2, 2, 1]
14. 很简单，如果存在奇数的，那么从里面分出一个来，就可以了
15. 也就是说，符合单调性
16. 剩下的就是怎么快速的找出x段？
17. 还是不大行呐。x可以二分，但是判断不行呐
18. 对于位置i，记录j, [i...j] 是最大的beautiful subarray
19. 然后如果在[l...r]中包含w个这样的，那么答案最多是w+2
20. 然后考虑如果从左边开始，从i开始，进过x个partion能够到达的最远距离 dp[i][x]
21. dp[i][x] >= r
22. 还有就是 fp[i][x] = l 表示从右开始，能够达到的最左位置
23. fp[r][x] <= l
24. 又理解错了，

### Solution

Suppose the query-interval has length 𝑚
. Let's call an element super-frequent if it occurs more than ⌈𝑚2⌉
 times in it, with frequency 𝑓
. If there's no super-frequent element, then we can just put all the elements in 1
 subsequence. Otherwise, we need the partitioning. Let's call the rest of the elements (other than the super-frequent one) good elements. One way to partition is to put all the 𝑚−𝑓
 good elements with 𝑚−𝑓+1
 super-frequent elements; then, put every remaining occurrence of the super-frequent element in a subsequence on its own. The number of subsequences we need here is then 1+𝑓−(𝑚−𝑓+1)=2∗𝑓−𝑚
. There's no way to improve this, because: for every subsequence we add, the number of occurrences of the super-frequent element minus the number of good elements is at most 1
, so by making it exactly 1
 in each subsequence, we get an optimal construction. Now, the problem boils down to calculating 𝑓
. Note that calculating the most frequent element in general is a well-known slow problem. It's usually solved with MO's algorithm in 𝑂(𝑛𝑛√𝑙𝑜𝑔(𝑛))
, maybe with a tricky optimization to 𝑂(𝑛𝑛√)
. However, notice that we only need the most frequent element if it occurs more than ⌈𝑚2⌉
 times. How can we use this fact?

