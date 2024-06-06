You are given an array 𝑎
 consisting of 𝑛
 nonnegative integers.

You can swap the elements at positions 𝑖
 and 𝑗
 if 𝑎𝑖 𝖷𝖮𝖱 𝑎𝑗<4
, where 𝖷𝖮𝖱
 is the bitwise XOR operation.

Find the lexicographically smallest array that can be made with any number of swaps.

An array 𝑥
 is lexicographically smaller than an array 𝑦
 if in the first position where 𝑥
 and 𝑦
 differ, 𝑥𝑖<𝑦𝑖
.


### ideas
1. a ^ b < 4
2. b ^ c < 4
3. a ^ c ?    a ^ c = a ^ b ^ b ^ c => (0, 1, 2, 3) ^ (0, 1, 2, 3) < 4
4. 也就是说它这个是可以传递的
5. 也就是说这里要进行分组（那些 a ^ b < 4 的部分要分在一组内）
6. 