During a normal walk in the forest, Katie has stumbled upon a mysterious code! However, the mysterious code had some characters unreadable. She has written down this code as a string 𝑐
 consisting of lowercase English characters and asterisks ("*"), where each of the asterisks denotes an unreadable character. Excited with her discovery, Katie has decided to recover the unreadable characters by replacing each asterisk with arbitrary lowercase English letter (different asterisks might be replaced with different letters).

Katie has a favorite string 𝑠
 and a not-so-favorite string 𝑡
 and she would love to recover the mysterious code so that it has as many occurrences of 𝑠
 as possible and as little occurrences of 𝑡
 as possible. Formally, let's denote 𝑓(𝑥,𝑦)
 as the number of occurrences of 𝑦
 in 𝑥
 (for example, 𝑓(𝑎𝑎𝑏𝑎𝑏𝑎,𝑎𝑏)=2
). Katie wants to recover the code 𝑐′
 conforming to the original 𝑐
, such that 𝑓(𝑐′,𝑠)−𝑓(𝑐′,𝑡)
 is largest possible. However, Katie is not very good at recovering codes in general, so she would like you to help her out.

 ### ideas
 1. dp[i][x][u][v] 表示最后c[i] = x时，最后u个字符和s相同，v个字符和t相同时的解
 2. dp[i+1][y][?][?] = 如果在c[i+1] = y 时的状态变化
 3. 这样是有问题的
 4. 在没有*的情况下，只要计算f(c, s) - f(c, t) 就可以了
 5. dp[i][u][v] 表示c[...i]的后缀和s的(u长度)前缀匹配，且和t的(v长度)的前缀匹配时的结果
 6. 从i-1到i时，如果c[i] = s[u], 那么u+1, 如果 c[i] = t[v], 那么v+1
 7. 如果 u+1 = len(s), +1, 如果 v+1 = len(t) -1
 8. 并重置到0