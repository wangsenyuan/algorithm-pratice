Ntarsis has been given a set 𝑆
, initially containing integers 1,2,3,…,101000
in sorted order. Every day, he will remove the 𝑎1
-th, 𝑎2
-th, …
, 𝑎𝑛
-th smallest numbers in 𝑆
simultaneously.

What is the smallest element in 𝑆
after 𝑘
days?

### thoughts

1. 如果没有1，那答案肯定是1
2. 如果要删除的数是连续的，就是 k * size + 1
3. 直接去模拟肯定是不行的
3. 考虑删除x次后，最小的数，如果是num
4. 假设 s = [1, 10]
5. 第一次删除后 [2,3... 9,11,...]
6. 第二次删除后 [3....9,11,13,...]
7. 第三次删除后 [4.........13,15,..]
8. 
7. 