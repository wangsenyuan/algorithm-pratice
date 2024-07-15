You are given a positive integer 𝑛
. Find 𝑘
 positive integers 𝑎1,𝑎2,…,𝑎𝑘
, such that:

𝑎1+𝑎2+…+𝑎𝑘=𝑛
𝐿𝐶𝑀(𝑎1,𝑎2,…,𝑎𝑘)≤𝑛2
Here 𝐿𝐶𝑀
 is the least common multiple of numbers 𝑎1,𝑎2,…,𝑎𝑘
.

We can show that for given constraints the answer always exists.


### ideas. 
1. 如果 k = n (n = k >= 3), 那么就是 [1...]
2. 如果 k <= n / 2
3. n/2, 1, 1, 1 (k - 1 + n / 2 = n)
4. 假设n是偶数, k = n / 2 + 1
5. 如果 k > n / 2, 那么就用k-1个1，
6. k - 1 >= n / 2, 剩下那个 = n - (k - 1) <= n / 2
7. 如果 k < n / 2 => 2 * k < n
8. let x = n / k
9. 如果 x * k = n, good
10. x * k < n
11. x * (k + 1) > n
12. 2 * x, x, x, x, 1, 1, 1
13. 然后用 2 * x, 尽量的填充，剩下的都用1?
14. 2 * x * a + x * b + c = n
15. (2 * a + b) * x + c = n
16. a + b + c = k
17. x = n / k
18. 比如 n = 9, k = 4,   2 + 2 + 2 + 3 = 9
19. 2 + 2 + 4 + 1 = 9
20. r中肯定可以从中找出一个x, r个 2 * x, r个1， 身下的都是x