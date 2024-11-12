You are given a string 𝑠
 consisting only of first 20
 lowercase Latin letters ('a', 'b', ..., 't').

Recall that the substring 𝑠[𝑙;𝑟]
 of the string 𝑠
 is the string 𝑠𝑙𝑠𝑙+1…𝑠𝑟
. For example, the substrings of "codeforces" are "code", "force", "f", "for", but not "coder" and "top".

You can perform the following operation no more than once: choose some substring 𝑠[𝑙;𝑟]
 and reverse it (i.e. the string 𝑠𝑙𝑠𝑙+1…𝑠𝑟
 becomes 𝑠𝑟𝑠𝑟−1…𝑠𝑙
).

Your goal is to maximize the length of the maximum substring of 𝑠
 consisting of distinct (i.e. unique) characters.

The string consists of distinct characters if no character in this string appears more than once. For example, strings "abcde", "arctg" and "minecraft" consist of distinct characters but strings "codeforces", "abacaba" do not consist of distinct characters.

### ideas
1. 如果l...r是distinct字符组成的（这个长度不会超过20）
2. 那么，希望在右边，或者左边，找到[l...r]的补集
3. 这个补集的大小理论上可以超过1e6, (pow(2, 10))
4. got了
5. 不用迭代位置，直接迭代集合就可以了