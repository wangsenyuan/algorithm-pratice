You are given two strings a and b. You have to remove the minimum possible number of consecutive (standing one after another) characters from string b in such a way that it becomes a subsequence of string a. It can happen that you will not need to remove any characters at all, or maybe you will have to remove all of the characters from b and make it empty.

Subsequence of string s is any such string that can be obtained by erasing zero or more characters (not necessarily consecutive) from string s.

### ideas
1. 删除t中连续的一段，使得剩余部分事s的一个子序列
2. 假设删除的是t[l...r]
3. 那么剩余的t[..l] 和 t[r...] 都是s的子序列
4. 假设 dp[r](r是s的下标)表示在s的后缀s[r..]中，能够匹配的最长的t的后缀（子序列匹配）
5. 然后计算t[...l]需要的最短前缀x，那么这个l对应的r = max(l+1, dp[x+1])
6. 