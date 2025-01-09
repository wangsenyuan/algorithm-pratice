Now Vasya is taking an exam in mathematics. In order to get a good mark, Vasya needs to guess the matrix that the teacher has constructed!

Vasya knows that the matrix consists of n rows and m columns. For each row, he knows the xor (bitwise excluding or) of the elements in this row. The sequence a1, a2, ..., an denotes the xor of elements in rows with indices 1, 2, ..., n, respectively. Similarly, for each column, he knows the xor of the elements in this column. The sequence b1, b2, ..., bm denotes the xor of elements in columns with indices 1, 2, ..., m, respectively.

Help Vasya! Find a matrix satisfying the given constraints or tell him that there is no suitable matrix.



### ideas
1. a[1] = x[1][1] ^ ... ^ x[1][m]
2. a[i] = x[i][1] ^ ... ^ x[i][m]
3. b[1] = x[1][1] ^ ... ^ x[n][1]
4. b[j] = x[1][j] ^ ... ^ x[n][j]
5. s = a[1] ^ a[2] ... ^ a[n] = b[1] ^ b[2] ... ^ b[m] (这个是必要条件)
6. 完全没有想法～
7. 不妨设 x[1][1] = a[1], x[2][1] = a[2]... x[i][1] = a[i]...
8. 如果 let w[1] = b[1] ^ a[1] ^ a[2] ^ ... ^ a[n]
9. 如果 w[1] != 0, 那么就让 x[1][1] = a[1] ^ w[1], a[1] = w[1], a[2..n] = 0
10. 这个时候，b[1]满足的， a[1]也是满足的
11. 如果w[2] = b[2] ^ a[1] ^ a[2] ^ ... ^ a[n]
12. 如果w[2] != 0, 那么久让x[2][2] = a[2] ^ w[2], a[2] = w[2], a[1...n] = 0
13. 为啥不用x[1][2]? 好像也是ok的。