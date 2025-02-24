Shohag has an integer 𝑛
 and a set 𝑆
 of 𝑚
 unique integers. Help him find the lexicographically largest∗
 integer array 𝑎1,𝑎2,…,𝑎𝑛
 such that 𝑎𝑖∈𝑆
 for each 1≤𝑖≤𝑛
 and 𝑎gcd(𝑖,𝑗)≠gcd(𝑎𝑖,𝑎𝑗)
†
 is satisfied over all pairs 1≤𝑖<𝑗≤𝑛
, or state that no such array exists.

### ideas
1. a[gcd(i, j)] != gcd(a[i], a[j])
2. gcd(i, i + 1) = 1
3. a[1] != gcd(a[i], a[i+1])
4. a[2] != gcd(a[i], a[j]) where gcd(i, j) = 2
5. gcd(i, j) <= i and j 的
6. 如果 a[gcd(i, j)] > a[i] and a[j] 那不就成立了吗？
7. 但是比如 a[3] = x, a[6] = x, a[9] = x, 那么就出现冲突了
8. 2的倍数，比如4,6,8,10...
9. 它们都必须要 < a[2], 
10. 那么 a[i] < 它的所有的因子
11. 如果找不到这样的数，就无解
12. 