The DNA sequence for every living creature in Berland can be represented as a non-empty line consisting of lowercase Latin letters. Berland scientists found out that all the creatures evolve by stages. During one stage exactly one symbol of the DNA line is replaced by exactly two other ones. At that overall there are n permissible substitutions. The substitution ai->bici means that any one symbol ai can be replaced with two symbols bici. Every substitution could happen an unlimited number of times.

They say that two creatures with DNA sequences s1 and s2 can have a common ancestor if there exists such a DNA sequence s3 that throughout evolution it can result in s1 and s2, perhaps after a different number of stages. Your task is to find out by the given s1 and s2 whether the creatures possessing such DNA sequences can have a common ancestor. If the answer is positive, you have to find the length of the shortest sequence of the common ancestor’s DNA.

Input
The first line contains a non-empty DNA sequence s1, the second line contains a non-empty DNA sequence s2. The lengths of these lines do not exceed 50, the lines contain only lowercase Latin letters. The third line contains an integer n (0 ≤ n ≤ 50) — the number of permissible substitutions. Then follow n lines each of which describes a substitution in the format ai->bici. The characters ai, bi, and ci are lowercase Latin letters. Lines s1 and s2 can coincide, the list of substitutions can contain similar substitutions.

Output
If s1 and s2 cannot have a common ancestor, print -1. Otherwise print the length of the shortest sequence s3, from which s1 and s2 could have evolved.

### ideas
1. 肯定是dp。but how？
2. 假设有一个s3，那么len(s1) - len(s3)就是s1的代数（每次增加一个字符）
3. 同样的len(s2) - len(s3) 也是s2的代数
4. 如果知道s3，有办法得到s1/s2吗？
5. 似乎也挺麻烦的。
6. 从字符x开始，经过多少步以后，可以得到的字符串是可以知道的
7. 但是这个是指数级增长的，即使限制了高度，也是不行的
8. 但是似乎有个办法, 比如 s1中有ab, 那么肯定存在某个字符变成ab（最后一步），如果不存在，ab就是一开始就要存在的
9. dp[x][i][j] 表示从x进过i个阶段后，能够得到s1的一个连续区间， 从j开始
10. dp[x][i][j] = dp[y][i-1][j] && dp[z][i-1][j + ?] ? 不对
11.             = dp[y][i1][j] && dp[z][i2][j+?] 
12.  第一个字符经过i1步，第二个字符经过i2步，且max(i1, i2) = i - 1
13. 