This is the easy version of the problem. The only difference is that here 𝑘=2
. You can make hacks only if both the versions of the problem are solved.

This is an interactive problem.

Every decimal number has a base 𝑘
 equivalent. The individual digits of a base 𝑘
 number are called 𝑘
-its. Let's define the 𝑘
-itwise XOR of two 𝑘
-its 𝑎
 and 𝑏
 as (𝑎+𝑏)mod𝑘
.

The 𝑘
-itwise XOR of two base 𝑘
 numbers is equal to the new number formed by taking the 𝑘
-itwise XOR of their corresponding 𝑘
-its. The 𝑘
-itwise XOR of two decimal numbers 𝑎
 and 𝑏
 is denoted by 𝑎⊕𝑘𝑏
 and is equal to the decimal representation of the 𝑘
-itwise XOR of the base 𝑘
 representations of 𝑎
 and 𝑏
. All further numbers used in the statement below are in decimal unless specified. When 𝑘=2
 (it is always true in this version), the 𝑘
-itwise XOR is the same as the bitwise XOR.

You have hacked the criminal database of Rockport Police Department (RPD), also known as the Rap Sheet. But in order to access it, you require a password. You don't know it, but you are quite sure that it lies between 0
 and 𝑛−1
 inclusive. So, you have decided to guess it. Luckily, you can try at most 𝑛
 times without being blocked by the system. But the system is adaptive. Each time you make an incorrect guess, it changes the password. Specifically, if the password before the guess was 𝑥
, and you guess a different number 𝑦
, then the system changes the password to a number 𝑧
 such that 𝑥⊕𝑘𝑧=𝑦
. Guess the password and break into the system.


### ideas
1. 一开始的password 是 x,
2. 输入y, 新的password x = x ^ y
3. 有点凌乱了，如果n = 1, 那岂不是只能猜一次？如果猜一次，要怎么可能猜得到？
4. x >= 0...n-1, ok
5. 假设询问的序列是y[0], y[1], ... y[n-1]
6. 那么 x[0] = x
7. x[1] = x ^ y[0], x[2] = x[1] ^ y[1] = x ^ y[0] ^ y[1]
8.  x[i] = x ^ y[0] ^ y[1] ^ .... y[i-1]
9.  x[n] = x ^ y[0] ^ y[1] ^ ..... y[n-1]
10. x 属于 0...n-1
11. 假设某个 y[i]是正确答案，
12. y[i] = x ^ y[0]... ^ y[i-1] 代表的是啥？
13. 这里唯一能确定的是 y[0] ... y[i-1]不是正确答案
14. 显然序列y必须有一个模式，这个模式，可以提供一些信息
15. 假设 n = 4
16. x 属于(0, 1, 2, 3)
17. 假设 x = 2
18. y[0] = 1, x = 1 ^ 2 = 3
19. y[0] = n - 1, y[1] = (n - 1) ^ (n-2)...
20. y[i] = (n - 1) ^ (n - 2) ^ ... (n - i - 1)
21. x[1] = x ^ y[0] = x ^ (n-1)
22. x[2] = x ^ (n-1) ^ (n - 1) ^ (n - 2) = x ^ (n-2)
23. x[i] = x ^ (n - i)
24. 