You are given 𝑛
 integer numbers 𝑎1,𝑎2,…,𝑎𝑛
. Consider graph on 𝑛
 nodes, in which nodes 𝑖
, 𝑗
 (𝑖≠𝑗
) are connected if and only if, 𝑎𝑖
 AND 𝑎𝑗≠0
, where AND denotes the bitwise AND operation.

Find the length of the shortest cycle in this graph or determine that it doesn't have cycles at all.

### ideas
1. 考虑一个特殊的case, 11, 101, 110 (3, 5, 6)
2. 就可以出现一个长度3的cycle
3. 且如果存在一个cycle的话，貌似肯定不会超过长度60
4. 且假设存在7（111）的话，就没有必要选择(3) 能和3连接的，必然可以和7连接（反之则不成立）
5. 且7, 3, 1 可以形成一个cycle
6. 所以，感觉可以brute force？