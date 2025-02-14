There are 
N slimes lined up in a row from left to right. For 
i=1,2,…,N, the 
i-th slime from the left has size 
A 
i
​
 .
For each 
K=1,2,…,N, solve the following problem.

Takahashi is the 
K-th slime from the left in the initial state. Find the maximum size that he can have after performing the following action any number of times, possibly zero:

Choose a slime adjacent to him that is strictly smaller than him, and absorb it. As a result, the absorbed slime disappears, and Takahashi's size increases by the size of the absorbed slime.
When a slime disappears due to absorption, the gap is immediately closed, and the slimes that were adjacent to the disappearing slime (if they exist) become adjacent (see the explanation in Sample Input 1).

### ideas
1. 如果是给定的k，使用贪心的方式，能吃到左边的就要吃左边的，能吃到右边就吃右边的
2. 但是对于所有的i要怎么计算呢？
3. 假设i能吃到区间 l....r区间内的，那么必须满足两个条件 sum[l...r] <= min(a[l-1], a[r+1])
4. 且 sum[i...j] > min(a[i-1], a[j+1])对于 i-1 >= l, j + 1 <= r 都要成立
5. 对于i,计算最小的j，满足 sum[i+1.j] <= a[i]
6. 对于j，计算最大的i，满足 sum[i.j-1] <= a[j]
7. 似乎还是没啥用呐，只要sum[i+1.j] > a[j+1], 那么a[i]
8. 如果从a[i]从低到高处理
9. 假设处理到了i,按照贪心的逻辑进行处理
10. 如果i处在之前的（特别是左边、右边最近的）的区间内，那么可以直接接上这个区间，然后去扩展
11. 如果i处在未确定的区间内，那么就正常的处理，直到和某个区间连接
12. 不大对。就算在某个区间内，或者链接到旁边的区间，它也不一定能完全吃掉
13. 比如 2, 3, 4, 3  最后一个3，在第一个3的区间内（它可以被第一个3吃掉，但是它却不能吃到其他的）
14. 构造一棵树，对于a[i]它的父节点是左边比它大的节点
15. 假设有这么一棵树，对于u，它的所有的子节点都可以被它吃掉。
16.  