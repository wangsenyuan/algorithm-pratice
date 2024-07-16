You are given an integer 𝑛
. You have to apply 𝑚
 operations to it.

In a single operation, you must replace every digit 𝑑
 of the number with the decimal representation of integer 𝑑+1
. For example, 1912
 becomes 21023
 after applying the operation once.

You have to find the length of 𝑛
 after applying 𝑚
 operations. Since the answer can be very large, print it modulo 109+7
.

### ideas
1. 1912 => 21023 => 32134 => ... => 87689 => 987910
2. 考虑0 经过 10次变成了 10, 特别的 x 经过 10 - x次后变成了10
3. 然后10 => 10次变成 910 (+1) 1021
4. f(1, m) = 1经过m次后的长度
5. f(0, m) = 0经过m次后的长度
6. f(0, 1) = 1, f(0, 10) = 2
7. f(10, m) = f(1, m) + f(0, m) = f(10, m - 9) + f(10, m - 10)
8. 把10去掉， f(m) = f(m - 9) + f(m - 10) 这个可以提前算好
9. f(1912, m) = f(10, m - 9) + f(10, m - 1) + f(10, m - 9) + f(10, m - 8)
10. 