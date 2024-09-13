Today, as a friendship gift, Bakry gave Badawy 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 and challenged him to choose an integer 𝑋
 such that the value max1≤𝑖≤𝑛(𝑎𝑖⊕𝑋)
 is minimum possible, where ⊕
 denotes the bitwise XOR operation.

As always, Badawy is too lazy, so you decided to help him and find the minimum possible value of max1≤𝑖≤𝑛(𝑎𝑖⊕𝑋)
.

### ideas
1. 从高位开始考虑, 如果当前位都是0, 那么x这一位也应该是0; 如果都是1，那么x[i] = 1
2. 如果这一位有0，有1，那么似乎有点复杂
3. x[i] = 0, 肯定有一部分得到1，如果x[i] = 1, 还是有1的结果
4. 但是总不能随便放吧。
5. [1, 2, 3] [01, 10, 11] when x = 3 (11)
6. 确实不能随便放，假设这一位放置了x[i] = 1，那么结果看，res[i] >= 1 << i (肯定存在某个数字a[?] ^ x >= 1 << i )
7. 比如上面的i = 1, 和 i = 0, 这两位都存在0/1的情况， 所以最后的结果肯定是大于等于 1<< 1, 1 << 0
8. 所以，这里肯定确定的是，最后的结果，是最高位的i, 既有0/1的情况
9. res |= 1 << h
10. 在结果确定是包含1 << h的情况下，然后看下一位，就是看下一位能否值为0
11. 