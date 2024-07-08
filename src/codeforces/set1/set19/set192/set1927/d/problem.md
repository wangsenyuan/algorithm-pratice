You are given an array 𝑎
 of 𝑛
 integers, and 𝑞
 queries.

Each query is represented by two integers 𝑙
 and 𝑟
 (1≤𝑙≤𝑟≤𝑛
). Your task is to find, for each query, two indices 𝑖
 and 𝑗
 (or determine that they do not exist) such that:

𝑙≤𝑖≤𝑟
;
𝑙≤𝑗≤𝑟
;
𝑎𝑖≠𝑎𝑗
.
In other words, for each query, you need to find a pair of different elements among 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟
, or report that such a pair does not exist.