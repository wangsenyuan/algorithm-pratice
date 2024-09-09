It is Professor R's last class of his teaching career. Every time Professor R taught a class, he gave a special problem for the students to solve. You being his favourite student, put your heart into solving it one last time.

You are given two polynomials 𝑓(𝑥)=𝑎0+𝑎1𝑥+⋯+𝑎𝑛−1𝑥𝑛−1
 and 𝑔(𝑥)=𝑏0+𝑏1𝑥+⋯+𝑏𝑚−1𝑥𝑚−1
, with positive integral coefficients. It is guaranteed that the cumulative GCD of the coefficients is equal to 1
 for both the given polynomials. In other words, 𝑔𝑐𝑑(𝑎0,𝑎1,…,𝑎𝑛−1)=𝑔𝑐𝑑(𝑏0,𝑏1,…,𝑏𝑚−1)=1
. Let ℎ(𝑥)=𝑓(𝑥)⋅𝑔(𝑥)
. Suppose that ℎ(𝑥)=𝑐0+𝑐1𝑥+⋯+𝑐𝑛+𝑚−2𝑥𝑛+𝑚−2
.

You are also given a prime number 𝑝
. Professor R challenges you to find any 𝑡
 such that 𝑐𝑡
 isn't divisible by 𝑝
. He guarantees you that under these conditions such 𝑡
 always exists. If there are several such 𝑡
, output any of them.

### ideas
1. c[0] = a[0] * b[0]
2. c[1] = a[0] * b[1] + a[1] * b[0]
3. c[2] = a[0] * b[2] + a[1] * b[1] + a[2] * b[0]
4. ...
5. c[i] = a[0] * b[i] + a[1] * b[i-1] + ... + a[i] * b[0]
6. c[t] 不能整除 p, 那是不是 a[0]...a[i]都不能有质因子p，还是只要有一个没有就可以呢？
7. 假设a[t], 和 b[0]没有质因子，那么他们a[t] * b[0] % p != 0
8. 假设 a[i]是第一个不包含因子p的数，如果b[0]也不包含p，那么似乎就找到了答案。因为其他所有的式子求余p都为0
9. 但是如果b[0]包含呢？假设b[j]是第一个不包含p因子的数
10. 那么 a[i] * b[j] 就不能整除p. t = i + j 吗？
11. c[i+j] = a[0] * b[i+j] + a[1] * b[i+j-1] ... + a[i] * b[j] + a[i+1] * b[j-1]
12. yes