You are given a string 𝑠
. You can build new string 𝑝
 from 𝑠
 using the following operation no more than two times:

choose any subsequence 𝑠𝑖1,𝑠𝑖2,…,𝑠𝑖𝑘
 where 1≤𝑖1<𝑖2<⋯<𝑖𝑘≤|𝑠|
;
erase the chosen subsequence from 𝑠
 (𝑠
 can become empty);
concatenate chosen subsequence to the right of the string 𝑝
 (in other words, 𝑝=𝑝+𝑠𝑖1𝑠𝑖2…𝑠𝑖𝑘
).
Of course, initially the string 𝑝
 is empty.

For example, let 𝑠=ababcd
. At first, let's choose subsequence 𝑠1𝑠4𝑠5=abc
 — we will get 𝑠=bad
 and 𝑝=abc
. At second, let's choose 𝑠1𝑠2=ba
 — we will get 𝑠=d
 and 𝑝=abcba
. So we can build abcba
 from ababcd
.

Can you build a given string 𝑡
 using the algorithm above?

Input
The first line contains the single integer 𝑇
 (1≤𝑇≤100
) — the number of test cases.

Next 2𝑇
 lines contain test cases — two per test case. The first line contains string 𝑠
 consisting of lowercase Latin letters (1≤|𝑠|≤400
) — the initial string.

The second line contains string 𝑡
 consisting of lowercase Latin letters (1≤|𝑡|≤|𝑠|
) — the string you'd like to build.

It's guaranteed that the total length of strings 𝑠
 doesn't exceed 400
.

### ideas
1. 最多两次操作得到t
2. 那么先考虑一次的情, 如果 t是s的子序列，那么就可以在一次内完成
3. 然后考虑2次操作的情况
4. 假设两次的分界线在位置t[i]处，那么t[1...i], t[i+1....m]都必须是s的子序列
5. 特别的t[i+1...m]必须是移除t[1...i]后s的子序列
6. 前一个比较好处理，可以先通过 n * m 的时间预处理出来；但是后一个呢？
7. 还有一个问题，子序列移除的方式，似乎也是有关系的
8. 比如 s = cabad, t = a+cab, 如果先移除掉的是第一个a，那么没法得到cab
9. 所以还是要一起考虑
10. 如果分界点在i, 那么就得到了两个子串 t1, t2
11. dp[i][j] = 在满足t1[:i]是 t2[:j] 都是 s1的子串的情况下，最小的下标是多少
12. dp[i+1][j] = 如果 s[dp[i][j]:]后面第一个和t1[i+1]相等的下标
13. dp[i][j+1] = .... 的下标