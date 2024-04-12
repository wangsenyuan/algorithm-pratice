Vasya has a sequence 𝑎
consisting of 𝑛
integers 𝑎1,𝑎2,…,𝑎𝑛
. Vasya may pefrom the following operation: choose some number from the sequence and swap any pair of bits in its binary
representation. For example, Vasya can transform number 6
(…000000001102)
into 3
(…000000000112)
, 12
(…0000000011002)
, 1026
(…100000000102)
and many others. Vasya can use this operation any (possibly zero) number of times on any number from the sequence.

Vasya names a sequence as good one, if, using operation mentioned above, he can obtain the sequence with bitwise
exclusive or of all elements equal to 0
.

For the given sequence 𝑎1,𝑎2,…,𝑎𝑛
Vasya'd like to calculate number of integer pairs (𝑙,𝑟)
such that 1≤𝑙≤𝑟≤𝑛
and sequence 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟
is good.

### ideas

1. transform不变的是1的个数
2. 考虑如果给定了l...r如何判断这个序列是good的？
3. 如果这个序列的长度是2，那么只需要判断，两个数的x的个数是否相等
4. 如果序列的长度是3呢？
5. 似乎就是一个个数 x的 xor 是不是0
6. (6, 7, 14) 表明不是这样的
7. 考虑长度为3的情况，假设它们1的个数分别为a,b,c (a <= b <= c)
8. 考虑b和c有一部分抵消了，剩下x个，如果 x = a, 那么就是ok的，b = x + w, c = x + v
9. 考虑抵消部分的数量是y， 那么有 b + c - 2 * y = a
10. 2 * y = (b + c) - a
11. 如果 b + c和a的parity是一致的，那么就肯定存在答案
12. 也就是说这里，只有当a, b, c 都是奇数时，肯定不行
13. 考虑有几堆石头，每次必须从偶数堆头部取掉一个石头，然后是否可以取完？
14. 。。。。没想法～
15. 首先sum必须是偶数，但是仅仅是偶数也是不行的
16. 比如一堆有2个，另外一堆有10个，那么两次后，就无法继续了
17. 所以sum 是偶数，sum - max(x) >= max(x)
18. sum >= 2 * max(x)
19. 这里x最大是61
20. 所以，区间超过长度超过61的时候，只需要考虑奇偶性

#### solution

Since we can swap any pair of bits in the number, so all we need to know is just the number of ones in its binary
representation. Let create array 𝑏1,𝑏2,…,𝑏𝑛
, where 𝑏𝑖
is number of ones in binary representation of 𝑎𝑖
. Forget about array 𝑎
, we will work with array 𝑏
.

Let's find a way to decide whether fixed segment is good or not. It can be proven that two conditions must be met. At
first, sum of 𝑏𝑖
at this segment should be even. At second, maximal element should be lower or equal to the sum of all other elements.

We will iterate over left borders of subsegments in descending order and for each left border 𝑙
calculate number of right borders 𝑟
such that [𝑙,𝑟]
is good.

Let's, as first, "forget" about condition on maximum and calculate 𝑐𝑛𝑡𝐴𝑙𝑙(𝑙)
— number of right borders 𝑟
, such that sum on segment [𝑙,𝑟]
is even and left border 𝑙
is fixed. We can calculate it by counting 𝑆0
and 𝑆1
— the number of suffixes of array with even sum of 𝑏𝑖
and number of suffixes with odd sum. If the current sum ∑𝑖=𝑙𝑛𝑏𝑖
is even then 𝑐𝑛𝑡𝐴𝑙𝑙(𝑙)=𝑆0
since ∑𝑖=𝑙𝑟𝑏𝑖=∑𝑖=𝑙𝑛𝑏𝑖−∑𝑖=𝑟+1𝑛𝑏𝑖
. If ∑𝑖=𝑙𝑛𝑏𝑖
is odd then 𝑐𝑛𝑡𝐴𝑙𝑙(𝑙)=𝑆1
.

Since we forgot about condition on maximum, some bad segments were counted. Since 𝑎𝑖≤1018
then 𝑏𝑖<61
. That's why if length of the segment ≥61
, condition on the maximum is always true. So, for a fixed 𝑙
we can iterate over all right borders in the [𝑙,𝑙+61]
and subtract number of segments with even sum and too large maximum (since these segments were counted in the answer).