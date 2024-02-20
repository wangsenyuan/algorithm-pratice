Let's define the anti-beauty of a multiset {𝑏1,𝑏2,…,𝑏𝑙𝑒𝑛}
as the number of occurrences of the number 𝑙𝑒𝑛
in the multiset.

You are given 𝑚
multisets, where the 𝑖
-th multiset contains 𝑛𝑖
distinct elements, specifically: 𝑐𝑖,1
copies of the number 𝑎𝑖,1
, 𝑐𝑖,2
copies of the number 𝑎𝑖,2,…,𝑐𝑖,𝑛𝑖
copies of the number 𝑎𝑖,𝑛𝑖
. It is guaranteed that 𝑎𝑖,1<𝑎𝑖,2<…<𝑎𝑖,𝑛𝑖
. You are also given numbers 𝑙1,𝑙2,…,𝑙𝑚
and 𝑟1,𝑟2,…,𝑟𝑚
such that 1≤𝑙𝑖≤𝑟𝑖≤𝑐𝑖,1+…+𝑐𝑖,𝑛𝑖
.

Let's create a multiset 𝑋
, initially empty. Then, for each 𝑖
from 1
to 𝑚
, you must perform the following action exactly once:

Choose some 𝑣𝑖
such that 𝑙𝑖≤𝑣𝑖≤𝑟𝑖
Choose any 𝑣𝑖
numbers from the 𝑖
-th multiset and add them to the multiset 𝑋
.
You need to choose 𝑣1,…,𝑣𝑚
and the added numbers in such a way that the resulting multiset 𝑋
has the minimum possible anti-beauty.

### thoughts

1. 好绕口的一个题目，应该举一个例子说明anti-beauty是什么
2. 假设X = {10,11,11,11,12,12,13,13,13,13}, 共有10个元素，所以len = 10， 然后10出现了1次，所以它的anti-beauty = 1
3. anti-beauty = 集合的大小（len)在集合中出现的次数
4. 有m个集合，每个集合的元素通过num和它的出现次数cnt给定，并指定了每个集合挑选数的下限和上限 (l, r)
5. 要计算结果集合的最小的anti-beauty
6. 假设最后的集合的大小是n，这里希望n出现的次数越少越好，n的下限和上限是知道的
7. 如果n比较小(= sum(l)), 那么在每个集合中，应该选择尽可能大的数
8. 但是如果n比较大，那又要选择尽可能小的数
9. 这个答案，不会超过 sum(l), 如果所有的数都一样，那就选择最小的size即可
10. 考虑最后的集合的sz > sum(l), 且它有最小的anti-beauty=x,
11. x <= sum(l)
12. 是否可以，在不增加x的情况下，删除多余的数到sum(l)呢？
13. x表示 sz 出现了x次,
14. sz减1后，需要sz-1出现x次
15. 如果 sz-1这个数，在原来的集合中，就有x次，那么只要删除掉一个sz即可（如果x = 0, 则最好删除掉sz-1)
16. 如果 sz-1没有x次 sz-1 > x, 那这个时候，就要删除掉sz - 1这个数
17. 这个过程不大对，貌似没法达成目的
18. 知道下限和上限 L, R
19. 假设是n，那么应该尽可能的选择两头的数，且两头的数肯定是连续的
20. 从L开始，逐渐的增大sz，这时候要选择sz个数，那么必然是选择所有大于sz的数，假设有a个，选择所有小于sz的数，假设有b个
21. 如果 a + b >= sz 答案就是0
22. 如果 a + b < sz, 那么答案就是 sz - (a + b), 剩下的数，必须用sz来代替
23. 如果 R - L 大于所有不同的数的个数， 答案就是0，所以迭代的数量 <= n
24. 那剩下的就是如何更新a和b了
25. 不管那个数在那个集合内，我们希望的仍然是两头的数，但是每个集合能够取的数是有限制的
26. 且一开始的状态也是有限制的