You are given a string 𝑠
 of length 𝑛
 consisting of lowercase Latin characters. Find the length of the shortest string 𝑘
 such that several (possibly one) copies of 𝑘
 can be concatenated together to form a string with the same length as 𝑠
 and, at most, one different character.

More formally, find the length of the shortest string 𝑘
 such that 𝑐=𝑘+⋯+𝑘𝑥 times
 for some positive integer 𝑥
, strings 𝑠
 and 𝑐
 has the same length and 𝑐𝑖≠𝑠𝑖
 for at most one 𝑖
 (i.e. there exist 0
 or 1
 such positions).

### ideas
1. k = n时，是成立的
2. k肯定是n个一个因子
3. 然后在给定k的情况下，可以在O(n)的时间内检查？
4. 是可以的
5. 