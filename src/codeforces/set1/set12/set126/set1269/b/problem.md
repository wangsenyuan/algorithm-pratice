You are given a positive integer 𝑚
 and two integer sequence: 𝑎=[𝑎1,𝑎2,…,𝑎𝑛]
 and 𝑏=[𝑏1,𝑏2,…,𝑏𝑛]
. Both of these sequence have a length 𝑛
.

Permutation is a sequence of 𝑛
 different positive integers from 1
 to 𝑛
. For example, these sequences are permutations: [1]
, [1,2]
, [2,1]
, [6,7,3,4,1,2,5]
. These are not: [0]
, [1,1]
, [2,3]
.

You need to find the non-negative integer 𝑥
, and increase all elements of 𝑎𝑖
 by 𝑥
, modulo 𝑚
 (i.e. you want to change 𝑎𝑖
 to (𝑎𝑖+𝑥)mod𝑚
), so it would be possible to rearrange elements of 𝑎
 to make it equal 𝑏
, among them you need to find the smallest possible 𝑥
.

In other words, you need to find the smallest non-negative integer 𝑥
, for which it is possible to find some permutation 𝑝=[𝑝1,𝑝2,…,𝑝𝑛]
, such that for all 1≤𝑖≤𝑛
, (𝑎𝑖+𝑥)mod𝑚=𝑏𝑝𝑖
, where 𝑦mod𝑚
 — remainder of division of 𝑦
 by 𝑚
.

For example, if 𝑚=3
, 𝑎=[0,0,2,1],𝑏=[2,0,1,1]
, you can choose 𝑥=1
, and 𝑎
 will be equal to [1,1,0,2]
 and you can rearrange it to make it equal [2,0,1,1]
, which is equal to 𝑏
.

### ideas
1. (a[i] + x) % m = b[p[i]]
2. 假设有3个1，那么a[i]中，有同样个数的某个数w， (w + x) % m = 1
3. 