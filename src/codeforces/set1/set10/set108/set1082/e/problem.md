You are given array 𝑎
of length 𝑛
. You can choose one segment [𝑙,𝑟]
(1≤𝑙≤𝑟≤𝑛
) and integer value 𝑘
(positive, negative or even zero) and change 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟
by 𝑘
each (i.e. 𝑎𝑖:=𝑎𝑖+𝑘
for each 𝑙≤𝑖≤𝑟
).

What is the maximum possible number of elements with value 𝑐
that can be obtained after one such operation?

Input
The first line contains two integers 𝑛
and 𝑐
(1≤𝑛≤5⋅105
, 1≤𝑐≤5⋅105
) — the length of array and the value 𝑐
to obtain.

The second line contains 𝑛
integers 𝑎1,𝑎2,…,𝑎𝑛
(1≤𝑎𝑖≤5⋅105
) — array 𝑎
.

Output
Print one integer — the maximum possible number of elements with value 𝑐
which can be obtained after performing operation described above.

### ideas

1. 问c的最大的次数是多少? 可以二分吗？好像是可以的
2. 如果 [l...r] 操作后能够得到x次， 同样肯定能得到x-1次
3. 如果选中了区间[l...r]那么就是把这个区间内的出现次数最多的给替换成c
4. 在给定r的时候，假设在l前面的c的个数是 cnt[l-1] + cnt_x[l...r] 最大
5. 考虑一个区间，如果这个区间内的最多的数是x，如果r+1也是x，那么扩展到r+1是更好的选择
6. 如果不是，那么最好是不要扩展
7. 所以，从这个角度看，只有到a[r] = x 是最多的那个数时，才需要处理
8. 如果在给定x的时候，那么l也肯定是x
9. 假设这时候有两个l1 < l2 都是x
10. 什么情况下选择l1， 什么情况下选择l2 呢？
11. 如果 l1 ...l2 中间没有c， 无疑选择l1更好， 如果有一个c，那么就选择l2
12. 但是如果 l1, l2, l3...c, l4 那么也是选择l1
13. 是不是这里可以二分l呢？

### solution

Let 𝑐𝑛𝑡(𝑙,𝑟,𝑥)
be a number of occurrences of number 𝑥
in subsegment [𝑙,𝑟]
.

The given task is equivalent to choosing [𝑙,𝑟]
and value 𝑑
such that 𝑎𝑛𝑠=𝑐𝑛𝑡(1,𝑙−1,𝑐)+𝑐𝑛𝑡(𝑙,𝑟,𝑑)+𝑐𝑛𝑡(𝑟+1,𝑛,𝑐)
is maximum possible. But with some transformations 𝑎𝑛𝑠=𝑐𝑛𝑡(1,𝑛,𝑐)+(𝑐𝑛𝑡(𝑙,𝑟,𝑑)−𝑐𝑛𝑡(𝑙,𝑟,𝑐))
so we need to maximize 𝑐𝑛𝑡(𝑙,𝑟,𝑑)−𝑐𝑛𝑡(𝑙,𝑟,𝑐)
.

Key observation is the next: if we fix some value 𝑑
then we can shrink each segment between consecutive occurrences of 𝑑
in one element with weight equal to −𝑐𝑛𝑡(𝑙𝑖,𝑟𝑖,𝑐)
. Then we need just to find subsegment with maximal sum — the standard task which can be solved in 𝑂(𝑐𝑛𝑡(1,𝑛,𝑑))
.

Finally, total complexity is ∑𝑑𝑂(𝑐𝑛𝑡(1,𝑛,𝑑))=𝑂(𝑛)
.