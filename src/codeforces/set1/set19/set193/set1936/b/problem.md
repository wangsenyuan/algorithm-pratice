There is a one-dimensional grid of length 𝑛
. The 𝑖
-th cell of the grid contains a character 𝑠𝑖
, which is either '<' or '>'.

When a pinball is placed on one of the cells, it moves according to the following rules:

If the pinball is on the 𝑖
-th cell and 𝑠𝑖
 is '<', the pinball moves one cell to the left in the next second. If 𝑠𝑖
 is '>', it moves one cell to the right.
After the pinball has moved, the character 𝑠𝑖
 is inverted (i. e. if 𝑠𝑖
 used to be '<', it becomes '>', and vice versa).
The pinball stops moving when it leaves the grid: either from the left border or from the right one.
You need to answer 𝑛
 independent queries. In the 𝑖
-th query, a pinball will be placed on the 𝑖
-th cell. Note that we always place a pinball on the initial grid.

For each query, calculate how many seconds it takes the pinball to leave the grid. It can be shown that the pinball will always leave the grid within a finite number of steps.


### ideas
1. 假设放在位置i, 且s[i] = R （👉）， 然后它一直移动，要么移出去了，要么移动到某个位置j, s[j] = L (👈)
2. 在移动到j的过程中， i....j-1 都变成了向左，如果它移动了j，那么下一个位置，是i前面的某个位置k, s[k] = R
3. 依次震荡，每次都至少消耗掉一个位置，所以最多在n次后，停止
4. 但如此，显然不行，时间至少是n*n, TLE
5. 考虑位置i和k，从位置i到k后，到达位置k后，j的符号发生了变化，其他的都回复到了原来的位置(还有k也是如此)
6. 而且是k和j中间所有的i都是这样次 >>>>>>><
7. 这里有个潜在的点，就是像连续的R，似乎可以同时考虑，
8. 好像也不是
9. 还是没有找到题眼
10. 如果从i开始，它是向右的，如果右边假设有a个向左的字符，然后考虑i后边有b个向右的字符，那么可以有一个结论是
11. 如果 a <= b， 那么这个球肯定是从右边离开，
12. 如果 a > b 那么这个球肯定是从左边离开
13. 假设它从右边离开，在所有的a对 >< 中间
14. 这个时间是这样累加的， j1 - i + j1 - k1 + j2 - k1 + j2 - k2 + ... + ja - ka + n - ka
15. -i + 2 * j1 - 2 * k1 + 2 * j2 - 2 * k2 + ... + 
16. -i + 2 * (j1 + ... ja) + n - 2 * (k1 + ... + ka)
17. 如果是从左边出去的话，那么就是 -i + 2 * (j1 + ... + ja) + j(a+1) - 2 * (k1 + ... + ka) - 0
18. 这个思路应该是对的，只是细节可能会出错