After struggling and failing many times, Ujan decided to try to clean up his house again. He decided to get his strings in order first.

Ujan has two distinct strings 𝑠
 and 𝑡
 of length 𝑛
 consisting of only of lowercase English characters. He wants to make them equal. Since Ujan is lazy, he will perform the following operation at most 2𝑛
 times: he takes two positions 𝑖
 and 𝑗
 (1≤𝑖,𝑗≤𝑛
, the values 𝑖
 and 𝑗
 can be equal or different), and swaps the characters 𝑠𝑖
 and 𝑡𝑗
.

Ujan's goal is to make the strings 𝑠
 and 𝑡
 equal. He does not need to minimize the number of performed operations: any sequence of operations of length 2𝑛
 or shorter is suitable.


### ideas
1. i = j 没有用处。
2. 所以i != j, 交换后至少 s[i] = t[i] 或者 s[j] = t[j]
3. 首先可以尽量的使用使的, s[i] = t[i], 且 s[j] = t[j]的交换
4. 然后剩下的就是 s[i] = t[i] or s[j] = t[j]的那些位置 （这样的位置，应该会形成环，比如3位的环，只需要3次）
5. 所以，最多只需要n次，就可以