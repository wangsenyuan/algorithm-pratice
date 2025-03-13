There is a string s, consisting of capital Latin letters. Let's denote its current length as |s|. During one move it is allowed to apply one of the following operations to it:

INSERT pos ch — insert a letter ch in the string s in the position pos (1 ≤ pos ≤ |s| + 1, A ≤ ch ≤ Z). The letter ch becomes the pos-th symbol of the string s, at that the letters shift aside and the length of the string increases by 1.
DELETE pos — delete a character number pos (1 ≤ pos ≤ |s|) from the string s. At that the letters shift together and the length of the string decreases by 1.
REPLACE pos ch — the letter in the position pos of the line s is replaced by ch (1 ≤ pos ≤ |s|, A ≤ ch ≤ Z). At that the length of the string does not change.
Your task is to find in which minimal number of moves one can get a t string from an s string. You should also find the sequence of actions leading to the required results.

### ideas
1. dp[i][j] 表示通过 s[i]得到t[j]的最小操作数
2. dp[i][j] = dp[i-1][j-1] 如果 s[i] = t[j]
3. s[i] != t[j], 怎么操作呢？
4. dp[i][j] = dp[i][j-1] + 1, 把t[j]Insert到i
5.  如果在i处insert了一个字符t[j], 这里相当于把t[j]给移除掉
6.          = dp[i-1][j] + 1 把s[i]删除掉，
7.          = dp[i-1][j-1] + 1, 把s[i]替换成t[j]