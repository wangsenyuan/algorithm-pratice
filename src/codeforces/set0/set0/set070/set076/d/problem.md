Bitwise exclusive OR (or bitwise addition modulo two) is a binary operation which is equivalent to applying logical exclusive OR to every pair of bits located on the same positions in binary notation of operands. In other words, a binary digit of the result is equal to 1 if and only if bits on the respective positions in the operands are different.

For example, if X = 10910 = 11011012, Y = 4110 = 1010012, then:

X xor Y  =  6810  =  10001002.
Write a program, which takes two non-negative integers A and B as an input and finds two non-negative integers X and Y, which satisfy the following conditions:

A = X + Y
B  =  X xor Y, where xor is bitwise exclusive or.
X is the smallest number among all numbers for which the first two conditions are true.


### ideas
1. x + y = x | y + x & y
2. b[i] = 1, 表明这一位要么x[i]= 1, 或者y[i] = 1, 为了使得x最小（可以先考虑y[i] = 1, x[i] = 0)
3. 其他位，要么x[i], y[i] = 1, 要么都为0
4. 如果a[i] = 1，a[i-1....j-1] = 0, b[i] = 0, b[i-1....j]都是1, b[j-1] = 0，那么可以有 x[j-1] = 1, y[i-1...j-1] = 1组成
5. 如果a[i] = 1, b[i] = 1, 那么只有一种 x[i] = 0, y[i] = 1组成
6. 如果a[i] = 0, b[i] = 0, x[i] = 0, y[i] = 0(排除掉其他情况)
7. a[i] = 0, b[i] = 1, 排除掉其他情况