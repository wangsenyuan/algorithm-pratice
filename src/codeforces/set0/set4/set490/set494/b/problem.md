Hamed has recently found a string t and suddenly became quite fond of it. He spent several days trying to find all occurrences of t in other strings he had. Finally he became tired and started thinking about the following problem. Given a string s how many ways are there to extract k ≥ 1 non-overlapping substrings from it such that each of them contains string t as a substring? More formally, you need to calculate the number of ways to choose two sequences a1, a2, ..., ak and b1, b2, ..., bk satisfying the following requirements:

k ≥ 1



  t is a substring of string saisai + 1... sbi (string s is considered as 1-indexed).
As the number of ways can be rather large print it modulo 109 + 7.

### ideas
1. 先理解题目，找出所有的方案，a[1]...a[k], b[1]....b[k] (a[i], b[i]都是下标)
2. 满足这样的条件 k >= 1
3. 且对所有的i(1...k) a[i] <= b[i] (这两个序列可以选择相同的下标)
4. b[i-1] < a[i]
5. 假设有 [l1, r1], [l2, r2], ... [lm, rm] （s中共m个t串）
6. 那么k最大就是到m?
7. 假设 a[1], b[1]抱在 l1, r1, 那么 a[2] > r1肯定需要成立
8. 因为 r1 <= b[1] < a[2]
9. dp[i]表示第i个匹配的结果
10. dp[i] = sum dp[j] * (j - r)  where j > r = dp[j] * j - dp[j] * r
11. 所以需要知道后面有多少个j > r (这个比较容易)
12. 