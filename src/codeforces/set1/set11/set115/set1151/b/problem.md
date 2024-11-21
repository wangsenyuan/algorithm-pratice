Student Dima from Kremland has a matrix 𝑎
 of size 𝑛×𝑚
 filled with non-negative integers.

He wants to select exactly one integer from each row of the matrix so that the bitwise exclusive OR of the selected integers is strictly greater than zero. Help him!

Formally, he wants to choose an integers sequence 𝑐1,𝑐2,…,𝑐𝑛
 (1≤𝑐𝑗≤𝑚
) so that the inequality 𝑎1,𝑐1⊕𝑎2,𝑐2⊕…⊕𝑎𝑛,𝑐𝑛>0
 holds, where 𝑎𝑖,𝑗
 is the matrix element from the 𝑖
-th row and the 𝑗
-th column.

Here 𝑥⊕𝑦
 denotes the bitwise XOR operation of integers 𝑥
 and 𝑦
.


### ideas
1. 似乎挺简单的？
2. 从1到10位，看是否有某一位，可以得到奇数的排列就可
3. 