You are given a rectangular matrix of size 𝑛×𝑚
 consisting of integers from 1
 to 2⋅105
.

In one move, you can:

choose any element of the matrix and change its value to any integer between 1
 and 𝑛⋅𝑚
, inclusive;
take any column and shift it one cell up cyclically (see the example of such cyclic shift below).
A cyclic shift is an operation such that you choose some 𝑗
 (1≤𝑗≤𝑚
) and set 𝑎1,𝑗:=𝑎2,𝑗,𝑎2,𝑗:=𝑎3,𝑗,…,𝑎𝑛,𝑗:=𝑎1,𝑗
 simultaneously.

Example of cyclic shift of the first column
You want to perform the minimum number of moves to make this matrix look like this:


In other words, the goal is to obtain the matrix, where 𝑎1,1=1,𝑎1,2=2,…,𝑎1,𝑚=𝑚,𝑎2,1=𝑚+1,𝑎2,2=𝑚+2,…,𝑎𝑛,𝑚=𝑛⋅𝑚
 (i.e. 𝑎𝑖,𝑗=(𝑖−1)⋅𝑚+𝑗
) with the minimum number of moves performed.

Input
The first line of the input contains two integers 𝑛
 and 𝑚
 (1≤𝑛,𝑚≤2⋅105,𝑛⋅𝑚≤2⋅105
) — the size of the matrix.

The next 𝑛
 lines contain 𝑚
 integers each. The number at the line 𝑖
 and position 𝑗
 is 𝑎𝑖,𝑗
 (1≤𝑎𝑖,𝑗≤2⋅105
).


### ideas
1. 有个观察是，对于任何一列，shift，都应该发生在最后一步
2. 或者说，shift放在最后，不影响结果
3. 可以每一列单独考虑。它需要的结果是固定的（i * m + j) 
4. 如何计算最少的操作数，得到想要的结果呢？
5. got，固定shift次数（也就固定了一个区间，看这个区间内，有多少个数不满足条件，就是要做的replace的次数）