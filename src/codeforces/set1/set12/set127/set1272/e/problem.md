You are given an array 𝑎
 consisting of 𝑛
 integers. In one move, you can jump from the position 𝑖
 to the position 𝑖−𝑎𝑖
 (if 1≤𝑖−𝑎𝑖
) or to the position 𝑖+𝑎𝑖
 (if 𝑖+𝑎𝑖≤𝑛
).

For each position 𝑖
 from 1
 to 𝑛
 you want to know the minimum the number of moves required to reach any position 𝑗
 such that 𝑎𝑗
 has the opposite parity from 𝑎𝑖
 (i.e. if 𝑎𝑖
 is odd then 𝑎𝑗
 has to be even and vice versa).

 ### ideas
 1. 找到最近的那几组先