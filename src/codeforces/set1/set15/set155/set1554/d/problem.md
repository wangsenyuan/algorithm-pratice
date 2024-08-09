You are given an integer 𝑛
. Find any string 𝑠
 of length 𝑛
 consisting only of English lowercase letters such that each non-empty substring of 𝑠
 occurs in 𝑠
 an odd number of times. If there are multiple such strings, output any. It can be shown that such string always exists under the given constraints.

A string 𝑎
 is a substring of a string 𝑏
 if 𝑎
 can be obtained from 𝑏
 by deletion of several (possibly, zero or all) characters from the beginning and several (possibly, zero or all) characters from the end.

 ### ideas
 1. 考虑n = 5, 当然因为n < 26, 可以用abcde 来生成s，所有的都出现了1次
 2. 但是显然无法处理很长的数组
 3. abcba
 4. n / 26 如果是奇数
 5. abcbaccab