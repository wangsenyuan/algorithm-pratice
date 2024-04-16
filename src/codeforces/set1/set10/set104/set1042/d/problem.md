Petya has an array 𝑎
consisting of 𝑛
integers. He has learned partial sums recently, and now he can calculate the sum of elements on any segment of the array
really fast. The segment is a non-empty sequence of elements standing one next to another in the array.

Now he wonders what is the number of segments in his array with the sum less than 𝑡
. Help Petya to calculate this number.

More formally, you are required to calculate the number of pairs 𝑙,𝑟
(𝑙≤𝑟
) such that 𝑎𝑙+𝑎𝑙+1+⋯+𝑎𝑟−1+𝑎𝑟<𝑡
.

Input
The first line contains two integers 𝑛
and 𝑡
(1≤𝑛≤200000,|𝑡|≤2⋅1014
).

The second line contains a sequence of integers 𝑎1,𝑎2,…,𝑎𝑛
(|𝑎𝑖|≤109
) — the description of Petya's array. Note that there might be negative, zero and positive elements.

Output
Print the number of segments in Petya's array with the sum of elements less than 𝑡
.

