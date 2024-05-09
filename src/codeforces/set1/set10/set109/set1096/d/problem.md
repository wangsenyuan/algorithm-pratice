Vasya is preparing a contest, and now he has written a statement for an easy problem. The statement is a string of
length 𝑛
consisting of lowercase Latin latters. Vasya thinks that the statement can be considered hard if it contains a
subsequence hard; otherwise the statement is easy. For example, hard, hzazrzd, haaaaard can be considered hard
statements, while har, hart and drah are easy statements.

Vasya doesn't want the statement to be hard. He may remove some characters from the statement in order to make it easy.
But, of course, some parts of the statement can be crucial to understanding. Initially the ambiguity of the statement is
0
, and removing 𝑖
-th character increases the ambiguity by 𝑎𝑖
(the index of each character is considered as it was in the original statement, so, for example, if you delete character
r from hard, and then character d, the index of d is still 4
even though you delete it from the string had).

Vasya wants to calculate the minimum ambiguity of the statement, if he removes some characters (possibly zero) so that
the statement is easy. Help him to do it!

Recall that subsequence is a sequence that can be derived from another sequence by deleting some elements without
changing the order of the remaining elements.

### ideas

1. 显然的，除了 hard, 外，其他的字符都不用删除
2. 总的cost = sum of a[i]
3. 为了让删除的cost最小，应该使得保留的cost最大
4. dp[0] 表示没有 hard
5. dp[1] 表示前面有h时的最大cost
6. dp[2] 表示有ha的最大cost
7. dp[3] 表示有har的最大cost
8. 但是比如遇到a，如果前面没有h，那么把它保留下来就没有问题，有h时，它就变成了dp[1]