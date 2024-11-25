You are given two strings 𝑠
 and 𝑡
 both of length 2
 and both consisting only of characters 'a', 'b' and 'c'.

Possible examples of strings 𝑠
 and 𝑡
: "ab", "ca", "bb".

You have to find a string 𝑟𝑒𝑠
 consisting of 3𝑛
 characters, 𝑛
 characters should be 'a', 𝑛
 characters should be 'b' and 𝑛
 characters should be 'c' and 𝑠
 and 𝑡
 should not occur in 𝑟𝑒𝑠
 as substrings.

A substring of a string is a contiguous subsequence of that string. So, the strings "ab", "ac" and "cc" are substrings of the string "abacc", but the strings "bc", "aa" and "cb" are not substrings of the string "abacc".

If there are multiple answers, you can print any of them.

### ideas
1. 假设 xy 是需要避免的
2. 是不是一定会有答案呢？
3. abc， 可以有6种组合， abc, acb, bac, bca, cab, cba
4. 如果s或者t是重复的aa,bb,cc 等，可以舍弃掉
5. 然后最多剩下两种组合 ab, bc, 然后就酸则 cab 这个组合
6. 如果是一种，那就更容易了