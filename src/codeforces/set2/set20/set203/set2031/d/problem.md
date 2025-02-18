Dedicated to pushing himself to his limits, Penchick challenged himself to survive the midday sun in the Arabian Desert!

While trekking along a linear oasis, Penchick spots a desert rabbit preparing to jump along a line of palm trees. There are 𝑛
 trees, each with a height denoted by 𝑎𝑖
.

The rabbit can jump from the 𝑖
-th tree to the 𝑗
-th tree if exactly one of the following conditions is true:

𝑗<𝑖
 and 𝑎𝑗>𝑎𝑖
: the rabbit can jump backward to a taller tree.
𝑗>𝑖
 and 𝑎𝑗<𝑎𝑖
: the rabbit can jump forward to a shorter tree.
For each 𝑖
 from 1
 to 𝑛
, determine the maximum height among all trees that the rabbit can reach if it starts from the 𝑖
-th tree.

### ideas
1. 如果 a[i] > a[j], i < j, 那么兔子可以在这两棵树间相互跳（也就是说这是一条双向的边）
2. i1, i2, 如果 a[i1], a[i2] 都大于j，且都在j的前面
3. 如果 a[i1] > a[i2], 且 i1 < i2, 那么j只要要i1（或者i2)连接就可以了
4. 如果 a[i1] <= a[i2], 那么j也应该只和i2连接？如果i1,和i2是最后的跳入点，显然i2更优
5. 如果不是，比如存在某个i，可以从i1跳过去， 必然满足a[i] > a[i1], 且i < i1
6. 那么此时 i < i2成立，如果 a[i] > a[i2]成立，那么从i2也是可以跳过去的
7. 所以结论是，j连接到左边，比a[j]大的，最近的点
8. dp[j]如果只考虑左边 = 左边最大的a[i], 但是如果它右边存在某个点k, j < k, a[j] > a[k] => 那么k后面最大的点
9. 那么此时，这个k越大越好（能够选择的范围就越多）