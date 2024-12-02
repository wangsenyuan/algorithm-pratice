You are given a permutation 𝑝
 of integers from 1
 to 𝑛
, where 𝑛
 is an even number.

Your goal is to sort the permutation. To do so, you can perform zero or more operations of the following type:

take two indices 𝑖
 and 𝑗
 such that 2⋅|𝑖−𝑗|≥𝑛
 and swap 𝑝𝑖
 and 𝑝𝑗
.
There is no need to minimize the number of operations, however you should use no more than 5⋅𝑛
 operations. One can show that it is always possible to do that.


### ideas
1. 假设1的位置是 p1, 如果p1在前半段，（假设p1 != 1)
2. p1 = i, 如果p[i]就在后半段，且 p[i] - p[1] >= n, 那么可以直接交换（i就回到自己的位置了）
3. 如果 p[i]无法被交换（在距离n内）
4. 那么可以选择 2 * n（或者1）用它i进行交换，然后再交换1
5. 假设在位置i处的数字，不是i
6. 那么如果p[i] - i >= n 直接交换
7. 否则的话，先将i交换到 2 * n位置处（如果可以的话），然后将i交换过来
8. i一定在p[i]的后面吗？如果是前半段，这个肯定是成立的（因为前面的都交换好了，是1，2，3，。。。。i-1)
9. 不可能有i
10. 但是问题出在i在后半段的时候，如果(i仍然在p[i]的后面)，但是无法交换到2 * n, 因为它们的距离它近
11. 这时只能和1交换
12. 但是和1交换后，1还的交换回去 (1, i) + (i, p[i]) + (1, p[i]) 3次交换
13. 