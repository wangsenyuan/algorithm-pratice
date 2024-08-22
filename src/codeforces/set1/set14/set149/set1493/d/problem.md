You are given an array 𝑎
 of length 𝑛
. You are asked to process 𝑞
 queries of the following format: given integers 𝑖
 and 𝑥
, multiply 𝑎𝑖
 by 𝑥
.

After processing each query you need to output the greatest common divisor (GCD) of all elements of the array 𝑎
.

Since the answer can be too large, you are asked to output it modulo 109+7
.

Input
The first line contains two integers — 𝑛
 and 𝑞
 (1≤𝑛,𝑞≤2⋅105
).

The second line contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤2⋅105
) — the elements of the array 𝑎
 before the changes.

The next 𝑞
 lines contain queries in the following format: each line contains two integers 𝑖
 and 𝑥
 (1≤𝑖≤𝑛
, 1≤𝑥≤2⋅105
).

Output
Print 𝑞
 lines: after processing each query output the GCD of all elements modulo 109+7
 on a separate line.

