Kevin discovered a binary string 𝑠
 that starts with 1 in the river at Moonlit River Park and handed it over to you. Your task is to select two non-empty substrings∗
 of 𝑠
 (which can be overlapped) to maximize the XOR value of these two substrings.

The XOR of two binary strings 𝑎
 and 𝑏
 is defined as the result of the ⊕
 operation applied to the two numbers obtained by interpreting 𝑎
 and 𝑏
 as binary numbers, with the leftmost bit representing the highest value. Here, ⊕
 denotes the bitwise XOR operation.

The strings you choose may have leading zeros.

### ideas
1. 一个肯定是s本身
2. brute force second
3. 但是第二个进行比较似乎也是 o(n)， 所以不大行
4. 应该是尽量的让左边成为1
5. 假设第一个0出现在i位，那么这里应该使用1
6. 