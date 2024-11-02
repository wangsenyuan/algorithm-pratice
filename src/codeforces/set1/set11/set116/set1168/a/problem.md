Toad Zitz has an array of integers, each integer is between 0
 and 𝑚−1
 inclusive. The integers are 𝑎1,𝑎2,…,𝑎𝑛
.

In one operation Zitz can choose an integer 𝑘
 and 𝑘
 indices 𝑖1,𝑖2,…,𝑖𝑘
 such that 1≤𝑖1<𝑖2<…<𝑖𝑘≤𝑛
. He should then change 𝑎𝑖𝑗
 to ((𝑎𝑖𝑗+1)mod𝑚)
 for each chosen integer 𝑖𝑗
. The integer 𝑚
 is fixed for all operations and indices.

Here 𝑥mod𝑦
 denotes the remainder of the division of 𝑥
 by 𝑦
.

Zitz wants to make his array non-decreasing with the minimum number of such operations. Find this minimum number of operations.

### ideas
1. 有点难呐。
2. 前面的数也可以增加（可以变到0，比自己小的数）
3. 后面的数，也可以增加（比自己大的数）
4. 可以计算最大的那个增加的数（其他的增加，可以在它的增加中同步执行）
5. got了。可以假装执行了操作