Given two positive integers 𝑛
 and 𝑘
, and another array 𝑎
 of 𝑛
 integers.

In one operation, you can select any subarray of size 𝑘
 of 𝑎
, then remove it from the array without changing the order of other elements. More formally, let (𝑙,𝑟)
 be an operation on subarray 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟
 such that 𝑟−𝑙+1=𝑘
, then performing this operation means replacing 𝑎
 with [𝑎1,…,𝑎𝑙−1,𝑎𝑟+1,…,𝑎𝑛]
.

For example, if 𝑎=[1,2,3,4,5]
 and we perform operation (3,5)
 on this array, it will become 𝑎=[1,2]
. Moreover, operation (2,4)
 results in 𝑎=[1,5]
, and operation (1,3)
 results in 𝑎=[4,5]
.

You have to repeat the operation while the length of 𝑎
 is greater than 𝑘
 (which means |𝑎|>𝑘
). What is the largest possible median†
 of all remaining elements of the array 𝑎
 after the process?

†
The median of an array of length 𝑛
 is the element whose index is ⌊(𝑛+1)/2⌋
 after we sort the elements in non-decreasing order. For example: 𝑚𝑒𝑑𝑖𝑎𝑛([2,1,5,4,3])=3
, 𝑚𝑒𝑑𝑖𝑎𝑛([5])=5
, and 𝑚𝑒𝑑𝑖𝑎𝑛([6,8,2,4])=4
.

### ideas
1. 能不能快速的检查a[i]是否能成为med?
2. 或者能不能快速的检查x，能够成为med, 也就是要有h个 < x， 且有 h个>= x
3. 因为每次删除都是k，所有最后剩余的m % k = n % k
4. 那是不是就剩余n % k 个元素，就可以了？
5. 假设对于x来说，是可以在某个 m长度（不是 n % k)时成立，也就是有一半小于等于x
6. 就让这个m = r + k 好了， 且它就在最中间，没法前面删除k个，后面删除k个，以保留x
7. 这样看，貌似确实没法把r 和 m对等起来
8. 比如 [1,2,3,4,5] x = 3， 且k = 3
9. 假设检查x，把所有的数字按照<= x(设置为1)， > x 设置为0
10. 那么其实就是希望1足够的多，通过删除连续的k个字符，使得1的个数至少占一半
11. dp[i] = 能够保留的最大的尽1的数量
12. 如果 a[i] = 0( > x), dp[i] = max(dp[i-1] - 1, dp[i-k]) (把这一段都删除掉)
13. 如果a[i] = 1, (<= x) dp[i] = max(dp[i-1] + 1, dp[i-k])
14. 如果 dp[n] >= 0, 那么就是可以得到x的
15. 不对～～～
16. 如果要检查a[i]是否可以作为median, 那么必须有h个<= a[i], h个>= a[i]的数存在
17. 怎么检查呢？这个变化也不是连续的。。。
18. 差一点点～～～
19. 对于给定的x，如果a[i] >= x， 设置为1，如果a[i] < x, 贡献为-1， 那么尽量的使得最后留下的array的和最大，那么表明，median至少 >= x
20. 假设数组是[1, 2, 3, 4, 5, 6]， x=2，那么数组的贡献为[-1, 1, 1, 1, 1, 1]那么至少表明2满足时，3，4，。。也有可能满足