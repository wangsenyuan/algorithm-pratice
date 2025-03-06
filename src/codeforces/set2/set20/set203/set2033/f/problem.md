Kosuke is too lazy. He will not give you any legend, just the task:

Fibonacci numbers are defined as follows:

𝑓(1)=𝑓(2)=1
.
𝑓(𝑛)=𝑓(𝑛−1)+𝑓(𝑛−2)
 (3≤𝑛)
We denote 𝐺(𝑛,𝑘)
 as an index of the 𝑛
-th Fibonacci number that is divisible by 𝑘
. For given 𝑛
 and 𝑘
, compute 𝐺(𝑛,𝑘)
.
As this number can be too big, output it by modulo 109+7
.

For example: 𝐺(3,2)=9
 because the 3
-rd Fibonacci number that is divisible by 2
 is 34
. [1,1,2,3,5,8,13,21,34]
.

### ideas
1. n is huge, k <= 1e5
2. 如果能知道前i个fib 数，里面有多少个数 % k = 0, 
3. 那么就可以变成2分？ 好像也不大行。 因为n太大了，这个i估计也很大。大到无法用二分来表示
4. h(i) = (h(i-1) + h(i-2)) % k
5. 有没有可能 h(i) = 0, h(i+1) 也是0？
6. 还有就是它会不会是一个循环？
7. h(i) = 0, h(i+1) = h(i-1), h(i+2) = h(i+1) + h(i) = h(i+1)
8. h(i+3) = 2 * h(i+1) 
9. h(i+4) = 3 * h(i+1)
10. h(i+k) = (k-1) * h(i+1) 变成了0， 
11. h(i+k+1) = (k - 2) * h(i+1) 感觉循环起来了
12. k= 2时， [1,1,2,3,5,8,13,21,34,55,89]
13.  [1, 1, 0, 1, 1, 0, 1, 1, 0]
14.  k = 3时
15.  [1, 1, 2, 0, 2, 2, 1, 0, 1, 1, 2, 0, 2, 2, 1, 0]
16.  那就把前2 * k个找出来