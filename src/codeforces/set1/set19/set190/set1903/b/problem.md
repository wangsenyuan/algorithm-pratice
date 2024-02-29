In Cyprus, the weather is pretty hot. Thus, Theofanis saw this as an opportunity to create an ice cream company.

He keeps the ice cream safe from other ice cream producers by locking it inside big storage rooms. However, he forgot
the password. Luckily, the lock has a special feature for forgetful people!

It gives you a table 𝑀
with 𝑛
rows and 𝑛
columns of non-negative integers, and to open the lock, you need to find an array 𝑎
of 𝑛
elements such that:

0≤𝑎𝑖<230
, and
𝑀𝑖,𝑗=𝑎𝑖|𝑎𝑗
for all 𝑖≠𝑗
, where |
denotes the bitwise OR operation.
The lock has a bug, and sometimes it gives tables without any solutions. In that case, the ice cream will remain frozen
for the rest of eternity.

Can you find an array to open the lock?

### thoughts

1. 每一位可以单独考虑
2. 如果 m[i,j][d] = 0, 那么可以确定 a[i][d] = 0, a[j][d] = 0
3. m[i, j][d] = 1, a[i][d] = 1 or a[j][d] = 1
4. 所以，可以先把这一位都为0的行找出来，然后剩下的部分，进行判断，是否能够满足条件