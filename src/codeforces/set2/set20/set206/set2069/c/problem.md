Let's call an integer sequence beautiful if the following conditions hold:

its length is at least 3
;
for every element except the first one, there is an element to the left less than it;
for every element except the last one, there is an element to the right larger than it;
For example, [1,4,2,4,7]
 and [1,2,4,8]
 are beautiful, but [1,2]
, [2,2,4]
, and [1,3,5,3]
 are not.

Recall that a subsequence is a sequence that can be obtained from another sequence by removing some elements without changing the order of the remaining elements.

You are given an integer array 𝑎
 of size 𝑛
, where every element is from 1
 to 3
. Your task is to calculate the number of beautiful subsequences of the array 𝑎
. Since the answer might be large, print it modulo 998244353
.