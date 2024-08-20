You are given a sequence 𝑎
 consisting of 𝑛
 positive integers.

Let's define a three blocks palindrome as the sequence, consisting of at most two distinct elements (let these elements are 𝑎
 and 𝑏
, 𝑎
 can be equal 𝑏
) and is as follows: [𝑎,𝑎,…,𝑎𝑥,𝑏,𝑏,…,𝑏𝑦,𝑎,𝑎,…,𝑎𝑥]
. There 𝑥,𝑦
 are integers greater than or equal to 0
. For example, sequences []
, [2]
, [1,1]
, [1,2,1]
, [1,2,2,1]
 and [1,1,2,1,1]
 are three block palindromes but [1,2,3,2,1]
, [1,2,1,2,1]
 and [1,2]
 are not.

Your task is to choose the maximum by length subsequence of 𝑎
 that is a three blocks palindrome.

You have to answer 𝑡
 independent test cases.

Recall that the sequence 𝑡
 is a a subsequence of the sequence 𝑠
 if 𝑡
 can be derived from 𝑠
 by removing zero or more elements without changing the order of the remaining elements. For example, if 𝑠=[1,2,1,3,1,2,1]
, then possible subsequences are: [1,1,1,1]
, [3]
 and [1,2,1,3,1,2,1]
, but not [3,2,3]
 and [1,1,1,1,2]
.


### ideas
1. 现在a[i]可以到200， 所以不能用e1的做法了
2. 这里有个关注点在于，假如确定了l...r两个端点，那么中间是什么， 两端是什么并不重要，
3. 重要的是，中间最多的字符是什么，以及两端能够相等的最多的字符是什么
4. 当然不能迭代l...r
5. 固定r的时候， 迭代数字x，也就是希望存在这样一个l, 
6. 2 * min(pref[l][x], suf[r+1][x]) + pref[r][?] - pref[l][?] 最大
7. 这里不妨限制 pref[l][x] <= suf[r+1][x] （这个应该可以做到的吧？）
8. 2 * pref[l][x] + pref[r][?] - pref[l][?] 最大
9. 2 * (pref[l][x]- pref[l][?]) + pref[r][?] 最大
10. 好像还是不大行
11. 考虑 abbabbbaba,
12. 假设 存在一个最优的答案 aaabb...bbaaa 然后如果在b的第一个b的前面还有一个b，是否应该包括进去？似乎是不行的，包括一个b，必须至少舍弃两个a
13. 然后考虑如果 第一个b和最后一个b的中间还有两个a，似乎也不能设置b替换
14. 也就是说，最优的答案，既不是最长的a，也不是最长的b
15. a[l] = a[r]是确定的
16. 假设a[l] = a[r] = x, 且suf[r][x] <= pref[l][x], 那么这里l越小越好，否则就会浪费掉一些x，且还使得中间的部分也更小
17. 也就是l就是使得 suf[r][x] = pref[l][x] 的地方