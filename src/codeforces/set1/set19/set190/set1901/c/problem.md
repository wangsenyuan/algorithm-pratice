You are given an integer array 𝑎1,𝑎2,…,𝑎𝑛
(0≤𝑎𝑖≤109
). In one operation, you can choose an integer 𝑥
(0≤𝑥≤1018
) and replace 𝑎𝑖
with ⌊𝑎𝑖+𝑥2⌋
(⌊𝑦⌋
denotes rounding 𝑦
down to the nearest integer) for all 𝑖
from 1
to 𝑛
. Pay attention to the fact that all elements of the array are affected on each operation.

Print the smallest number of operations required to make all elements of the array equal.

If the number of operations is less than or equal to 𝑛
, then print the chosen 𝑥
for each operation. If there are multiple answers, print any of them.

### thoughts

1. (a + b) / 2
2. 如果a是偶数，且b是偶数 (a + b) / 2 = a / 2 + b / 2
3. 如果a是偶数，b是奇数 (a + b) / 2 = a / 2 + b / 2
4. 如果a是奇数, b是奇数 (a + b) / 2 = a / 2 + b / 2 + 1
5. 