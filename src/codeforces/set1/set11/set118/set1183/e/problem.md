A subsequence is a string that can be derived from another string by deleting some or no symbols without changing the order of the remaining symbols. Characters to be deleted are not required to go successively, there can be any gaps between them. For example, for the string "abaca" the following strings are subsequences: "abaca", "aba", "aaa", "a" and "" (empty string). But the following strings are not subsequences: "aabaca", "cb" and "bcaa".

You are given a string 𝑠
 consisting of 𝑛
 lowercase Latin letters.

In one move you can take any subsequence 𝑡
 of the given string and add it to the set 𝑆
. The set 𝑆
 can't contain duplicates. This move costs 𝑛−|𝑡|
, where |𝑡|
 is the length of the added subsequence (i.e. the price equals to the number of the deleted characters).

Your task is to find out the minimum possible total cost to obtain a set 𝑆
 of size 𝑘
 or report that it is impossible to do so.

 ### ideas
 1. 不知道咋入手～～～
 2. 如果k = 1, 显然把s放进去最优（cost = 0)
 3. 直观的感受应该是按照长度从长到短放进去
 4. 问题出在，如何去重？
 5. 比如字符串  abab ， 长度为2的子序列ab, 出现了3次 (但显然这3个，智能算一次)
 6. 所以问题是计算有多少个不同的字符串吗？