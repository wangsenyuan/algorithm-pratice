You are given an integer array 𝑎1,𝑎2,…,𝑎𝑛
, where 𝑎𝑖
 represents the number of blocks at the 𝑖
-th position. It is guaranteed that 1≤𝑎𝑖≤𝑛
.

In one operation you can choose a subset of indices of the given array and remove one block in each of these indices. You can't remove a block from a position without blocks.

All subsets that you choose should be different (unique).

You need to remove all blocks in the array using at most 𝑛+1
 operations. It can be proved that the answer always exists.

Input
The first line contains a single integer 𝑛
 (1≤𝑛≤103
) — length of the given array.

The second line contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤𝑛
) — numbers of blocks at positions 1,2,…,𝑛
.

Output
In the first line print an integer 𝑜𝑝
 (0≤𝑜𝑝≤𝑛+1
).

In each of the following 𝑜𝑝
 lines, print a binary string 𝑠
 of length 𝑛
. If 𝑠𝑖=
'0', it means that the position 𝑖
 is not in the chosen subset. Otherwise, 𝑠𝑖
 should be equal to '1' and the position 𝑖
 is in the chosen subset.

All binary strings should be distinct (unique) and 𝑎𝑖
 should be equal to the sum of 𝑠𝑖
 among all chosen binary strings.

If there are multiple possible answers, you can print any.

It can be proved that an answer always exists.

### ideas
1. 排序后，显然a[i] = n 的，必须至少出现在n个集合中
2. a[i] = x 的，必须出现在x个集合中
3. 那么假设有y个a[?] = x，那么就必须至少有y+1个集合，把这些等于x的放进去
4. 因为至多可以有n+1个集合，所以肯定有答案
5. 怎么分配呢？
6. 这些a[?] = x相同的那些，要放在y+1个中