You are given 𝑛
 strings 𝑠1,𝑠2,…,𝑠𝑛
, each consisting of lowercase and uppercase English letters. In addition, it's guaranteed that each character occurs in each string at most twice. Find the longest common subsequence of these strings.

A string 𝑡
 is a subsequence of a string 𝑠
 if 𝑡
 can be obtained from 𝑠
 by deletion of several (possibly, zero or all) symbols.

Input
Each test consists of multiple test cases. The first line contains a single integer 𝑡
 (1≤𝑡≤5
) — the number of test cases. Description of the test cases follows.

The first line of each test case contains a single integer 𝑛
 (2≤𝑛≤10
) — the number of strings.

Each of the next 𝑛
 lines contains the corresponding string 𝑠𝑖
. Each 𝑠𝑖
 is non-empty, consists only of uppercase and lowercase English letters, and no character appears more than twice in each string.

Output
For each test case print the answer in two lines:

In the first line print the length of the longest common subsequence.

In the second line print the longest common subsequence. If there are multiple such subsequences, print any of them.

### ideas
1. 完全没idea～～～
2. 每个字符最多出现两次
3. 任意两个字符，在一个字符串中（出现的情况下）有三种关系
4. a在b的前面，b在a的前面，或者a，b相会包含 （abab)这样子
5. 那么统计ab的关系，如果它们在所有的字符串中，都符合a -> b, 
6. 那么如果最终的lcs，以a结尾，那么就可以扩展到以b结尾
7. 这样子，确实就变成一棵树了，找树里面最长的路径