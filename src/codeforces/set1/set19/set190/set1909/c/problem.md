https://codeforces.com/problemset/problem/1909/C

### problem

You have 𝑛
intervals [𝑙1,𝑟1],[𝑙2,𝑟2],…,[𝑙𝑛,𝑟𝑛]
, such that 𝑙𝑖<𝑟𝑖
for each 𝑖
, and all the endpoints of the intervals are distinct.

The 𝑖
-th interval has weight 𝑐𝑖
per unit length. Therefore, the weight of the 𝑖
-th interval is 𝑐𝑖⋅(𝑟𝑖−𝑙𝑖)
.

You don't like large weights, so you want to make the sum of weights of the intervals as small as possible. It turns out
you can perform all the following three operations:

rearrange the elements in the array 𝑙
in any order;
rearrange the elements in the array 𝑟
in any order;
rearrange the elements in the array 𝑐
in any order.
However, after performing all of the operations, the intervals must still be valid (i.e., for each 𝑖
, 𝑙𝑖<𝑟𝑖
must hold).

What's the minimum possible sum of weights of the intervals after performing the operations?

### ideas

1. 先考虑几个例子
2. [l1, r1, c1], [l2, r2, c2]
3. 如果r1 <= l2(或者 r2 <= l1) 也就是没有交集的情况下，
4. 如果进行交换，就会产生重叠的部分，那么结果肯定会更大
5. 所以，第一个想法应该是尽量的减少重叠的部分
6. 但是如果必须要重叠时，怎么处理 l1 < l2 < r1 < r2
7. 假设c一样的情况下
8. (r1 - l1) * c + (r2 - l2) * c = (r1 + r2 - -l1 - l2) * c
8. (r1 - l2) * c + (r2 - l1) * c = ....
9. c不一样的情况下假设c1 < c2 (r1 - l1) * c1 + (r2 - l2) * c2
10. = (r1 - l2 + l2 - l1) * c1 + (r2 - r1 + r1 - l2) * c2
11. = (l2 - l1) * c1 + (r1 - l2) * c1 + (r2 - r1) * c2 + (r1 - l2) * c2
12. = (l2 - l1) * c1 + (r1 - l2) * (c1 + c2) + (r2 - r1) * c2
13. 可以看到重叠的部分，怎么选择c1和c2，都会被计算到，但是两头，如果让l1和r2pair,那么就可以让两头使用较小的数
14. 尽量减小重叠区域，无法减少时，使用大区间包裹小区间的方式