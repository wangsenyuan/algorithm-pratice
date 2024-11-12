You are given a string 𝑠
 consisting of lowercase Latin letters and 𝑞
 queries for this string.

Recall that the substring 𝑠[𝑙;𝑟]
 of the string 𝑠
 is the string 𝑠𝑙𝑠𝑙+1…𝑠𝑟
. For example, the substrings of "codeforces" are "code", "force", "f", "for", but not "coder" and "top".

There are two types of queries:

1 𝑝𝑜𝑠 𝑐
 (1≤𝑝𝑜𝑠≤|𝑠|
, 𝑐
 is lowercase Latin letter): replace 𝑠𝑝𝑜𝑠
 with 𝑐
 (set 𝑠𝑝𝑜𝑠:=𝑐
);
2 𝑙 𝑟
 (1≤𝑙≤𝑟≤|𝑠|
): calculate the number of distinct characters in the substring 𝑠[𝑙;𝑟]
.

### ideas
1. 每个字符建立一个序列，然后计算是否在l...r中即可