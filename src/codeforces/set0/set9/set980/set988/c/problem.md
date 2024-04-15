You are given 𝑘
sequences of integers. The length of the 𝑖
-th sequence equals to 𝑛𝑖
.

You have to choose exactly two sequences 𝑖
and 𝑗
(𝑖≠𝑗
) such that you can remove exactly one element in each of them in such a way that the sum of the changed sequence 𝑖
(its length will be equal to 𝑛𝑖−1
) equals to the sum of the changed sequence 𝑗
(its length will be equal to 𝑛𝑗−1
).

Note that it's required to remove exactly one element in each of the two chosen sequences.

Assume that the sum of the empty (of the length equals 0
) sequence is 0
.

Input
The first line contains an integer 𝑘
(2≤𝑘≤2⋅105
) — the number of sequences.

Then 𝑘
pairs of lines follow, each pair containing a sequence.

The first line in the 𝑖
-th pair contains one integer 𝑛𝑖
(1≤𝑛𝑖<2⋅105
) — the length of the 𝑖
-th sequence. The second line of the 𝑖
-th pair contains a sequence of 𝑛𝑖
integers 𝑎𝑖,1,𝑎𝑖,2,…,𝑎𝑖,𝑛𝑖
.

The elements of sequences are integer numbers from −104
to 104
.

The sum of lengths of all given sequences don't exceed 2⋅105
, i.e. 𝑛1+𝑛2+⋯+𝑛𝑘≤2⋅105
.

Output
If it is impossible to choose two sequences such that they satisfy given conditions, print "NO" (without quotes).
Otherwise in the first line print "YES" (without quotes), in the second line — two integers 𝑖
, 𝑥
(1≤𝑖≤𝑘,1≤𝑥≤𝑛𝑖
), in the third line — two integers 𝑗
, 𝑦
(1≤𝑗≤𝑘,1≤𝑦≤𝑛𝑗
). It means that the sum of the elements of the 𝑖
-th sequence without the element with index 𝑥
equals to the sum of the elements of the 𝑗
-th sequence without the element with index 𝑦
.

Two chosen sequences must be distinct, i.e. 𝑖≠𝑗
. You can print them in any order.

If there are multiple possible answers, print any of them.

