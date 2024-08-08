Let's call the string beautiful if it does not contain a substring of length at least 2
, which is a palindrome. Recall that a palindrome is a string that reads the same way from the first character to the last and from the last character to the first. For example, the strings a, bab, acca, bcabcbacb are palindromes, but the strings ab, abbbaa, cccb are not.

Let's define cost of a string as the minimum number of operations so that the string becomes beautiful, if in one operation it is allowed to change any character of the string to one of the first 3
 letters of the Latin alphabet (in lowercase).

You are given a string 𝑠
 of length 𝑛
, each character of the string is one of the first 3
 letters of the Latin alphabet (in lowercase).

You have to answer 𝑚
 queries — calculate the cost of the substring of the string 𝑠
 from 𝑙𝑖
-th to 𝑟𝑖
-th position, inclusive.

Input
The first line contains two integers 𝑛
 and 𝑚
 (1≤𝑛,𝑚≤2⋅105
) — the length of the string 𝑠
 and the number of queries.

The second line contains the string 𝑠
, it consists of 𝑛
 characters, each character one of the first 3
 Latin letters.

The following 𝑚
 lines contain two integers 𝑙𝑖
 and 𝑟𝑖
 (1≤𝑙𝑖≤𝑟𝑖≤𝑛
) — parameters of the 𝑖
-th query.

Output
For each query, print a single integer — the cost of the substring of the string 𝑠
 from 𝑙𝑖
-th to 𝑟𝑖
-th position, inclusive.

### ideas
1. aabc 这个字符串可以在一次操作后变成 cabc
2. abbc 这个字符长需要两次abca类似这样
3. 但是给定一个区间，如何计算它的最小修改数呢？
4. abcabc这种pattern