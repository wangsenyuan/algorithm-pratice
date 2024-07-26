After Fox Ciel got off a bus, she found that the bus she was on was a wrong bus and she lost her way in a strange town. However, she fortunately met her friend Beaver Taro and asked which way to go to her castle. Taro's response to her was a string s, and she tried to remember the string s correctly.

However, Ciel feels n strings b1, b2, ... , bn are really boring, and unfortunately she dislikes to remember a string that contains a boring substring. To make the thing worse, what she can remember is only the contiguous substring of s.

Determine the longest contiguous substring of s that does not contain any boring string, so that she can remember the longest part of Taro's response.

Input
In the first line there is a string s. The length of s will be between 1 and 105, inclusive.

In the second line there is a single integer n (1 ≤ n ≤ 10). Next n lines, there is a string bi (1 ≤ i ≤ n). Each length of bi will be between 1 and 10, inclusive.

Each character of the given strings will be either a English alphabet (both lowercase and uppercase) or a underscore ('_') or a digit. Assume that these strings are case-sensitive.

Output
Output in the first line two space-separated integers len and pos: the length of the longest contiguous substring of s that does not contain any bi, and the first position of the substring (0-indexed). The position pos must be between 0 and |s| - len inclusive, where |s| is the length of string s.

If there are several solutions, output any.


### ideas
1. len(s) <= 1e5, len(b[i]) <= 10, n <= 10
2. dp[i] 表示以i开始的最长的子串, dp[i] = min(j) - 1, where j 是i后面，任何一个子串b的结束位置
3. 使用suffix array