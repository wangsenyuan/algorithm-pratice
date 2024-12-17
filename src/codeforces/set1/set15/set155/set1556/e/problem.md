William has two arrays 𝑎
 and 𝑏
, each consisting of 𝑛
 items.

For some segments 𝑙..𝑟
 of these arrays William wants to know if it is possible to equalize the values of items in these segments using a balancing operation. Formally, the values are equalized if for each 𝑖
 from 𝑙
 to 𝑟
 holds 𝑎𝑖=𝑏𝑖
.

To perform a balancing operation an even number of indices must be selected, such that 𝑙≤𝑝𝑜𝑠1<𝑝𝑜𝑠2<⋯<𝑝𝑜𝑠𝑘≤𝑟
. Next the items of array a at positions 𝑝𝑜𝑠1,𝑝𝑜𝑠3,𝑝𝑜𝑠5,…
 get incremented by one and the items of array b at positions 𝑝𝑜𝑠2,𝑝𝑜𝑠4,𝑝𝑜𝑠6,…
 get incremented by one.

William wants to find out if it is possible to equalize the values of elements in two arrays for each segment using some number of balancing operations, and what is the minimal number of operations required for that. Note that for each segment the operations are performed independently.

### ideas
1. 首先选择长度为2的序列就可以了（必须是偶数）
2. 所有长度超过2的序列，都可以被拆成长度为2的序列的组合
3. 且两个长度为2的序列，可以重叠
4. a[1], a[2], a[3], b[1], b[2], b[3]
5. 那么就可选择（1， 2）使的位置1处相等，假设增加了a[1] + 1， a[2] - x (b[2]不用变)
6. 然后使得 a[2] + y, 此时a[3] - y
7. x = b[1] - a[1]
8. -x + y = b[2] - a[2]
9. -y = b[3] - a[3]
10. 0 = b[1] + b[2] + b[3] - (...) 和相等？ 这个貌似是必要条件
11. 除了这个，还有个问题 
12. a[l...i]的前缀和 <= b[l...i]
13. 这样子，才能前面加，后面减
14. 这个才是难点，要用到segment tree才行 或者dp？
15. dp[i]表示，考虑从i开始的后缀，最近的位置 sum_a[i...j] > sum_b[i...j]的
16. 但是啊但是，这个操作次数，却是要按照越长越好，越长，次数就越少
17. 所以是这个区间的最大值
18. 😄，搞定