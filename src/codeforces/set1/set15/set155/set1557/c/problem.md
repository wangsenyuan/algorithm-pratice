Moamen and Ezzat are playing a game. They create an array 𝑎
 of 𝑛
 non-negative integers where every element is less than 2𝑘
.

Moamen wins if 𝑎1&𝑎2&𝑎3&…&𝑎𝑛≥𝑎1⊕𝑎2⊕𝑎3⊕…⊕𝑎𝑛
.

Here &
 denotes the bitwise AND operation, and ⊕
 denotes the bitwise XOR operation.

Please calculate the number of winning for Moamen arrays 𝑎
.

As the result may be very large, print the value modulo 1000000007
 (109+7
).

Input
The first line contains a single integer 𝑡
 (1≤𝑡≤5
)— the number of test cases.

Each test case consists of one line containing two integers 𝑛
 and 𝑘
 (1≤𝑛≤2⋅105
, 0≤𝑘≤2⋅105
).

Output
For each test case, print a single value — the number of different arrays that Moamen wins with.

Print the result modulo 1000000007
 (109+7
).

### ideas
1. a & b >= a ^ b
2. 从高到低，第i位如果不是全部为1，那么 a[i] & = 0, 如果为1的个数为奇数 a[i] ^ = 1
3. 所以，就是从高到低，在保证高位一致的情况下，看能否在当前位得到 a[i] & > a[i] ^ 的情况
4. a[i]& = 1 (表示全部 = 1)， n是偶数 =》 good 
5. 高位只能全部为1，或者偶数个为0 
6. 如果 n 是奇数, n = 1, 那么就是 pow(2, k)
7. else n = 3, 貌似答案也是 pow(2, k) (所有的数必须相同)
8. 最高位，如果 and(a[i]) > xor(a[i]) => a[i] = 1, 但是因为是奇数，所以xor(a[i]) = 1
9. 但是如果有一个为0， and(a[i]) = 0, xor(a[i]) = 0, 那么就必须选择偶数个1设置为1
10. 也就是说n为奇数时，and = 1（表明所有为都为1, xor = 1)成立，如果有一个为0，那么必须是奇数个0（1， 3， 5.。。。n)
11. 只能保证 =, 不可能出现 and(a[i]) > xor(a[i])的情况
12. 如果n为偶数，and(a[i]) = 1时，xor(a[i]) = 0此时，低位可以随便 pow(1 << i, n) 每个数的低位部分可取0...1<<i - 1
13. else and(a[i]) = 0, 要保证xor(a[i]) = 0, 就必须取（2，4， 6， 8, n 个0） 