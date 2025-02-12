You are given an integer array 𝑎
 of size 𝑛
. The elements of the array are numbered from 1
 to 𝑛
.

You can perform the following operation any number of times (possibly, zero): choose an index 𝑖
 from 1
 to 𝑛
; decrease 𝑎𝑖
 by 2
 and increase 𝑎(𝑖mod𝑛)+1
 by 1
.

After you perform the operations, all elements of the array should be non-negative equal integers.

Your task is to calculate the minimum number of operations you have to perform.

### ideas
1. 每次操作，将整个数组的sum减少1，那么至多在sum次后， 最后的结果变成了0
2. 但是有个前提是，在这个过程中，不产生负值
3. 如果最后的数是x，且x是个偶数，那么肯定可以得到 x-1 (所有位减去2，后面1位加上1)
4. 如果最后的数是x，且x是个奇数，且x > 1, 也同样可以得到 x - 1
5. 也就是说，如果能得到结果为x，肯定可以得到1
6. 所有，可以先检查所有的数，是否可以得到1。 如果不能 => -1 (不可能得到0)
7. 那怎么检查呢？或者更一般的，能否得到x？
8. 如果 a[i] > x, 且 a[i] - x 是偶数，那么只需要把 a[i] - x 的一半迁移到后面
9. 如果 a[i] > x， 且 a[i] - x 是奇数， 那么这时候肯定要迁移进来一个数
10. 这里主要有一个点，就是一个位置减少后，能够影响的距离最多是后面30个位置
11. 在给定x的情况下，看最后剩余的要么全部是偶数，全部是奇数
12. 