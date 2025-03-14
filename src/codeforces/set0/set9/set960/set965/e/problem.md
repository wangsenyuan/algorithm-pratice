Arkady's code contains 𝑛
 variables. Each variable has a unique name consisting of lowercase English letters only. One day Arkady decided to shorten his code.

He wants to replace each variable name with its non-empty prefix so that these new names are still unique (however, a new name of some variable can coincide with some old name of another or same variable). Among such possibilities he wants to find the way with the smallest possible total length of the new names.

A string 𝑎
 is a prefix of a string 𝑏
 if you can delete some (possibly none) characters from the end of 𝑏
 and obtain 𝑎
.

Please find this minimum possible total length of new names.

### ideas
1. 考虑 abc, abd 答案 = [a, ab] = 3?
2. dfs on trees？