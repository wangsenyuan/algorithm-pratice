Patrick calls a substring†
 of a binary string‡
 good if this substring contains exactly one 1.

Help Patrick count the number of binary strings 𝑠
 such that 𝑠
 contains exactly 𝑛
 good substrings and has no good substring of length strictly greater than 𝑘
. Note that substrings are differentiated by their location in the string, so if 𝑠=
 1010 you should count both occurrences of 10.

†
 A string 𝑎
 is a substring of a string 𝑏
 if 𝑎
 can be obtained from 𝑏
 by the deletion of several (possibly, zero or all) characters from the beginning and several (possibly, zero or all) characters from the end.

‡
 A binary string is a string that only contains the characters 0 and 1.

Input
Each test consists of multiple test cases. The first line contains a single integer 𝑡
 (1≤𝑡≤2500
) — the number of test cases. The description of the test cases follows.

The only line of each test case contains two integers 𝑛
 and 𝑘
 (1≤𝑛≤2500
, 1≤𝑘≤𝑛
) — the number of required good substrings and the maximum allowed length of a good substring.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 2500
.

Output
For each test case, output a single integer — the number of binary strings 𝑠
 such that 𝑠
 contains exactly 𝑛
 good substrings and has no good substring of length strictly greater than 𝑘
. Since this integer can be too large, output it modulo 998244353
.

### ideas
1. 先理解这个题目
2. good substring = 这个substring里面只包含一个1
3. 比如string 010101 => 010(3个) 01（3个）, 10(2个), 1（3）,都是good
4. 比如上面的string的， 不考虑k的影响（比如 k >= 3），f(s, k) = 11
5. 如果 k = 2, f(s, k) = 8
6. f(s, 1) = 3
7. 在一个good的substring中，只能有一个1，
8. 假设k = 3, 那么 000不是good的，因为它没有1
9. 0010 不是good的，因为它的长度超过了3
10. 但是 10101是good的，因为它里面最长的good substring 是 010
11. 脑子要炸了
12. 1的个数 <= n
13. 如果没有0， 字符串全部为0时，刚好是n
14. 如果有一个0， 如果它在第一位/或者最后一位，那么1的个数要减去1
15. 如果它在其他位置，那么1的个数要减去2 （n - 1) good + 1
16. 101 的 good = 1 + 10 + 01 + 1 = 4
17. 还没有找到题眼～
18. 假设没有k的限制，只有n的要求，是否能够计算出来？
19. 没有k的限制，这个数量是无穷的，因为即使 = 1， 它也是有无数个字符串满足条件
20. 还是要有k的限制，这里假设正好满足限制j （1到k)
21. 假设j = 1, dp[n] = 1
22. 如果j = 2, dp[1] = 2 (01/10)
23.     dp[2] = 3  (101)
24.     110 算什么？它有两个good substring (1, 10)
25.  dp[i][j] 表示到目前为止有i个good substring了， 且最后一个good的长度是j时的计数
26.  如果现在再放一个1, dp[i+?][?] 这里依赖前面有几个0
27.  如果有3个0，那么这里就是4， 
28.  dp[i][x] 到目前为止有i个good substring，且前面有x个0
29.  如果放置1， dp[i+x + 1][0] += dp[i][x]
30.  如果放置0 dp[i+1][x+1] += dp[i][x]
31.  但是k没法被处理
32.  先不考虑复杂行 dp[i][j][x]表示前面有i个good substring， 且最后一个substring的长度是j，其中0的个数为x
33.  dp[i+x+1][x+1][0] += dp[i][j][x] 当放置1的时候
34.  dp[i+j-x][j+1][x+1] += dp[i][j][x] 当放置0的时候
35.   ？等于 后面有x+1个0, ? = j + 1 - (x + 1)
36.   但是这个是 n * k * k 复杂性的
37.   还需要优化掉一个维度
38.   增加的delta + x = 第二个维度
39. 还差点～～
40. 


### solution

Let's first solve the problem where we are given some string 𝑠
 and must count the number of good substrings. To do this we use the technique of counting contributions. For every 1
 in 𝑠
, we find the number of good substrings containing that 1
. Consider the following example:

00001⏟𝑎10001⏟𝑎200000001𝑎30001⏟𝑎4000⏟𝑎5
The number of good substrings in this example is 𝑎1𝑎2+𝑎2𝑎3+𝑎3𝑎4+𝑎4𝑎5
. We can create such array for any string 𝑠
 and the number of good substrings of 𝑠
 is the sum of the products of adjacent elements of the array.

This motivates us to reformulate the problem. Instead, we count the number of arrays 𝑎1,𝑎2,...,𝑎𝑚
 such that every element is positive and the sum of the products of adjacent elements is exactly equal to 𝑛
. Furthermore, every pair of adjacent elements should have sum minus 1
 be less than or equal to 𝑘
. We can solve this with dynamic programming.

𝑑𝑝𝑖,𝑗=number of arrays with sum 𝑖 and last element 𝑗
𝑑𝑝𝑖,𝑗=∑𝑝=1min(⌊𝑖𝑗⌋,𝑘−𝑗+1)𝑑𝑝𝑖−𝑗⋅𝑝,𝑝
The key observation is that we only have to iterate 𝑝
 up to ⌊𝑖𝑗⌋
 (since if 𝑝
 is any greater, 𝑗⋅𝑝
 will exceed 𝑖
). At 𝑗=1
, we will iterate over at most ⌊𝑖1⌋
 values of 𝑝
. At 𝑗=2
, we will iterate over at most ⌊𝑖2⌋
 values of 𝑝
. In total, at each 𝑖
, we will iterate over at most ⌊𝑖1⌋+⌊𝑖2⌋+⋯+⌊𝑖𝑖⌋≈𝑖log𝑖
 values of 𝑝
. Thus, the time complexity of our solution is 𝑂(𝑛𝑘log𝑛)
.