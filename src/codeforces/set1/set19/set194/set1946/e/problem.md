Some permutation of length 𝑛
 is guessed.

You are given the indices of its prefix maximums and suffix maximums.

Recall that a permutation of length 𝑘
 is an array of size 𝑘
 such that each integer from 1
 to 𝑘
 occurs exactly once.

Prefix maximums are the elements that are the maximum on the prefix ending at that element. More formally, the element 𝑎𝑖
 is a prefix maximum if 𝑎𝑖>𝑎𝑗
 for every 𝑗<𝑖
.

Similarly, suffix maximums are defined, the element 𝑎𝑖
 is a suffix maximum if 𝑎𝑖>𝑎𝑗
 for every 𝑗>𝑖
.

You need to output the number of different permutations that could have been guessed.

As this number can be very large, output the answer modulo 109+7
.

#### ideas
1. 先考虑只有prefix的情况下，应该怎么处理？
2. 最后一个元素肯定是n，假设对应的下标是t[m] 那么它后面到n的都是小于n的数，这样的数 = C(n - 1, n - t[m])个
3. 所有小于n的数，都可以放在后面的位置
4. 然后前面一个t[m-1], 这个数应该是多少呢？好像它也不能随便的小（它太小了，前面的位置不够了）也不能随便大了，太大了，也不能是n
5. 假设prefix的长度是w, 那么应该选择 w - 1 个数(最后一个数肯定是n)，选择w-1个数，后作为prefix
6. C(n-1, w - 1), 然后剩余的n - w 个数的位置，似乎是比较固定的，但也不是那么固定
7. 比如n-1如果没有被选中，那它只能出现在n的后面,1没有被选中，那它可以在所有被选中的数后面
8. 咋搞？
9. 两个下标之间，可以放入w = t[i] - t[i-1] - 1 个数
10. 这些数必须比t[i-1]小， 如果t[i-1] = x, C(x - 1, w) * P[w] ?
11. 似乎也不对， 在i-1前面的已经被选中的 < x, C(x - (i - 1), w) * P[w]
12. dp[i][x]表示第i个放置x时的计数 dp[i][x] =  C(x - i, w) * P[w]
13. 但是问题时为啥它不和 dp[i-1][y]有关系呢？