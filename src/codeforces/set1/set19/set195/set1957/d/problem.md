You are given an array 𝑎1,𝑎2,…,𝑎𝑛
. Find the number of tuples (𝑥,𝑦,𝑧
) such that:

1≤𝑥≤𝑦≤𝑧≤𝑛
, and
𝑓(𝑥,𝑦)⊕𝑓(𝑦,𝑧)>𝑓(𝑥,𝑧)
.
We define 𝑓(𝑙,𝑟)=𝑎𝑙⊕𝑎𝑙+1⊕…⊕𝑎𝑟
, where ⊕
 denotes the bitwise XOR operation.


 ### ideas
 1. f(x, y) = xor(a[x], a[x+1], ... a[y])
 2. f(x, y) ^ f(y, z) = a[y] ^ f(x, z) > f(x, z) 
 3. 在(x, z)  中间存在某个y, a[y] ^ f(x, z) 变大了
 4. 假设 v = f(x, z) 从高到低
 5. v[i] = 1，那么y[i] 必须等于 0， (否则变小了)
 6. v[i] = 0, 如果y[i] = 1 => good
 7. 也就是说，y需要和v[i]的高位不一样，且在某位v[i]= 0时， y[i] = 1
 8. 这是在已知(x, z)的情况下，找y， 这个复杂性时n * n（还不包括y)，TLE
 9. 在给定z的情况下 
 10. g(z) = a[1] ^ ... ^ a[z]
 11. f(x, z) = g(z) ^ g(x - 1)
 12. a[y] ^ f(x, z) > f(x, z) 
 13. a[y] ^ g(z) ^ g(x-1) > g(z) ^ g(x-1)
 14. g[z] ^ (a[y] ^ g(x-1)) > g(z) ^ g(x-1)
 15. g[z][i] = 1, 
 16.   g(x-1)[i] = 1, a[y][i] = 1 good, else continue
 17.   g(x-1)[i] = 0, a[y][i] = 0 continue， else bad
 18.      这时候重要的是a[y][i]第一次为1的位置， 如果g(x-1)[i] = 1 good, else bad
 19. g[z][i] = 0
 20.   g(x-1)[i] = 0, a[y][i] = 1 good, else continue
 21.   g(x-1)[i] = 1, a[y][i] = 0 cointue, else bad
 22.       这里重要的也是 a[y][i]第一次为1的位置， 如果g(x-1)[i] = 0 good, else bad      
 23. 对于z, 从高位到低位, i, 
 24.   找到前面最高位为i的a[y], 
 25.     如果 a[y][i] = 1, 那么 + g(x-1)[i] = 1的计数， 否则加上 g(x-1)[i] = 0 的奇数
 26. 但是这里会出现重复计数的问题吗？应该不会，因为y肯定是不同的
1.  

### solution
First we may rewrite the inequality as 𝑓(𝑥,𝑧)⊕𝑎𝑦>𝑓(𝑥,𝑧)
. So, for each index 𝑦
, we want to find the total number of subarrays that contain 𝑦
 but whose xor
 decreases when we include the contribution of 𝑎𝑦
.

Including 𝑎𝑦
 in some subarray [𝑥,𝑧]
 (where 𝑥≤𝑦≤𝑧
) can make the xor
 lower only when some set bit of 𝑎𝑦
 cancels out an existing set bit in 𝑓(𝑥,𝑦−1)⊕𝑓(𝑦+1,𝑧)
. Consider the most signicant set bit in 𝑎𝑦
. Call this bit 𝑖
. Including 𝑎𝑦
 would decrease the subarray xor
 in subarrays where 𝑓(𝑥,𝑦−1)
 has 𝑖
 set while 𝑓(𝑦+1,𝑧)
 is unset or vice-versa.

Now, for the subarrays where both the prefix subarray ([𝑥…𝑦−1]
) as well as the suffix subarray ([𝑦+1…𝑧]
) where either both have bit 𝑖
 set or both have it unset, by including 𝑎𝑦
, we increment the xor by at least 2𝑖
. Even if by including 𝑎𝑦
, every other smaller bit gets unset (because of cancelling out with 𝑎𝑦
), this sum of these decrements is still less than 2𝑖
 thereby resulting in an overall positive contribution by including 𝑎𝑦
.