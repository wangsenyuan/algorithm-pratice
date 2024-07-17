Now you get Baby Ehab's first words: "Given an integer 𝑛
, find the longest subsequence of [1,2,…,𝑛−1]
 whose product is 1
 modulo 𝑛
." Please solve the problem.

A sequence 𝑏
 is a subsequence of an array 𝑎
 if 𝑏
 can be obtained from 𝑎
 by deleting some (possibly all) elements. The product of an empty subsequence is equal to 1
.

### ideas
1. a * b % n = 1
2. 如果在这个过程中，出现任何一个数，整除了n,那么就是不行的
3. 比如8, 2 * 4 % 8 = 0
4. a % n = 1 => a = x * n + 1
5. (a - 1) % n = 0
6. 等价于 这些数相乘的product = n + 1 的倍数

### solution
So first observe that the subsequence can't contain any element that isn't coprime with 𝑛
. Why? Because then its product won't be coprime with 𝑛
, so when you take it modulo 𝑛
, it can't be 1
. In mathier words, 𝑔𝑐𝑑(𝑝𝑟𝑜𝑑 𝑚𝑜𝑑 𝑛,𝑛)=𝑔𝑐𝑑(𝑝𝑟𝑜𝑑,𝑛)≠1
. Now, let's take all elements less than 𝑛
 and coprime with it, and let's look at their product modulo 𝑛
; call it 𝑝
. If 𝑝
 is 1
, you can take all these elements. Otherwise, you should take them all except for 𝑝
. It belongs to them because 𝑝
 is coprime with 𝑛
, since 𝑔𝑐𝑑(𝑝 𝑚𝑜𝑑 𝑛,𝑛)=𝑔𝑐𝑑(𝑝,𝑛)=1
 since all the elements in 𝑝
 are coprime with 𝑛
.

