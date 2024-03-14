We will call a non-empty string balanced if it contains the same number of plus and minus signs. For example:
strings "+--+" and "++-+--" are balanced, and strings "+--", "--" and "" are not balanced.

We will call a string promising if the string can be made balanced by several (possibly zero) uses of the following
operation:

replace two adjacent minus signs with one plus sign.
In particular, every balanced string is promising. However, the converse is not true: not every promising string is
balanced.

For example, the string "-+---" is promising, because you can replace two adjacent minuses with plus and get a balanced
string "-++-", or get another balanced string "-+-+".

How many non-empty substrings of the given string 𝑠
are promising? Each non-empty promising substring must be counted in the answer as many times as it occurs in string 𝑠
.

Recall that a substring is a sequence of consecutive characters of the string. For example, for string "+-+" its
substring are: "+-", "-+", "+", "+-+" (the string is a substring of itself) and some others. But the following strings
are not its substring: "--", "++", "-++".

### ideas

1. 如果 [l...r] 是 promising的，那么表示，通过某种处理它可以变成balanced
2. 且两个promising的子串连接后，仍然是promising的假设 dp[i]表示以i结尾的promising的个数
3. dp[i] = sum(dp[j-1], if [i...j] is promising)
4. 那怎么判断某个区间是 promising 的呢？
5. 如果这段的sum = x, x <= 0, 且 abs(x) % 3 = 0