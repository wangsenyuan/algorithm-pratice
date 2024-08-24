There are two arrays 𝑎
 and 𝑏
 of length 𝑛
. Initially, an 𝑎𝑛𝑠
 is equal to 0
 and the following operation is defined:

Choose position 𝑖
 (1≤𝑖≤𝑛
);
Add 𝑎𝑖
 to 𝑎𝑛𝑠
;
If 𝑏𝑖≠−1
 then add 𝑎𝑖
 to 𝑎𝑏𝑖
.
What is the maximum 𝑎𝑛𝑠
 you can get by performing the operation on each 𝑖
 (1≤𝑖≤𝑛
) exactly once?

Uncle Bogdan is eager to get the reward, so he is asking your help to find the optimal order of positions to perform the operation on them.

Input
The first line contains the integer 𝑛
 (1≤𝑛≤2⋅105
) — the length of arrays 𝑎
 and 𝑏
.

The second line contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (−106≤𝑎𝑖≤106
).

The third line contains 𝑛
 integers 𝑏1,𝑏2,…,𝑏𝑛
 (1≤𝑏𝑖≤𝑛
 or 𝑏𝑖=−1
).

Additional constraint: it's guaranteed that for any 𝑖
 (1≤𝑖≤𝑛
) the sequence 𝑏𝑖,𝑏𝑏𝑖,𝑏𝑏𝑏𝑖,…
 is not cyclic, in other words it will always end with −1
.

Output
In the first line, print the maximum 𝑎𝑛𝑠
 you can get.

In the second line, print the order of operations: 𝑛
 different integers 𝑝1,𝑝2,…,𝑝𝑛
 (1≤𝑝𝑖≤𝑛
). The 𝑝𝑖
 is the position which should be chosen at the 𝑖
-th step. If there are multiple orders, print any of them.