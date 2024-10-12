This is an easier version of the next problem. The difference is only in constraints.

You are given a rectangular 𝑛×𝑚
 matrix 𝑎
. In one move you can choose any column and cyclically shift elements in this column. You can perform this operation as many times as you want (possibly zero). You can perform this operation to a column multiple times.

After you are done with cyclical shifts, you compute for every row the maximal value in it. Suppose that for 𝑖
-th row it is equal 𝑟𝑖
. What is the maximal possible value of 𝑟1+𝑟2+…+𝑟𝑛
?

Input
The first line contains an integer 𝑡
 (1≤𝑡≤40
), the number of test cases in the input.

The first line of each test case contains integers 𝑛
 and 𝑚
 (1≤𝑛≤4
, 1≤𝑚≤100
) — the number of rows and the number of columns in the given matrix 𝑎
.

Each of the following 𝑛
 lines contains 𝑚
 integers, the elements of 𝑎
 (1≤𝑎𝑖,𝑗≤105
).


### ideas
1. 只有4行
2. 可以确定第一行里面的最大值是哪个，然后fix这一列（它不能动了）
3. 然后在剩余列中，找到第二大的数
4. m * m * m * m ~ 1e8
5. 能勉强通过吧