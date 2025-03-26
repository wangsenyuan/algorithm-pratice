AquaMoon gives RiverHamster a sequence of integers 𝑎1,𝑎2,…,𝑎𝑛
, and RiverHamster gives you 𝑞
 queries. Each query is expressed by two integers 𝑙
 and 𝑟
.

For each query independently, you can take any continuous segment of the sequence and subtract an identical non-negative value from all the numbers of this segment. You can do so multiple (possibly, zero) times. However, you may not choose two intersecting segments which are not included in one another. Your goal is to convert to 0
 all numbers whose initial value was within the range [𝑙,𝑟]
. You must do so in the minimum number of operations.

Please note that the queries are independent, the numbers in the array are restored to their initial values between the queries.

Formally, for each query, you are to find the smallest 𝑚
 such that there exists a sequence {(𝑥𝑗,𝑦𝑗,𝑧𝑗)}𝑚𝑗=1
 satisfying the following conditions:

for any 1≤𝑗≤𝑚
, 𝑧𝑗≥0
 and 1≤𝑥𝑗≤𝑦𝑗≤𝑛
 (here [𝑥𝑗,𝑦𝑗]
 correspond to the segment of the sequence);
for any 1≤𝑗<𝑘≤𝑚
, it is true that [𝑥𝑗,𝑦𝑗]⊆[𝑥𝑘,𝑦𝑘]
, or [𝑥𝑘,𝑦𝑘]⊆[𝑥𝑗,𝑦𝑗]
, or [𝑥𝑗,𝑦𝑗]∩[𝑥𝑘,𝑦𝑘]=∅
;
for any 1≤𝑖≤𝑛
, such that 𝑙≤𝑎𝑖≤𝑟
, it is true that
𝑎𝑖=∑1≤𝑗≤𝑚𝑥𝑗≤𝑖≤𝑦𝑗𝑧𝑗.


### ideas
1. 

### solution
First, we consider only the sequence of elements to be manipulated. We claim that it is optimal to operate on the whole sequence so that the minimum elements are all decreased to 0
, and then solve the problem on the segments separated by the $0$s recursively.

A straightforward corollary of the claim is that two equal elements between which there are no smaller elements are handled in a single operation, so the final answer is ℓ
 (the number of elements to be manipulated) minus the number of such adjacent pairs.

Proof: Define 𝑆(𝑙,𝑟)
 to be the answer for the segment [𝑙,𝑟]
. Do induction on the length of the segment. If the first operation does not manipulate the whole segment, the segment will be separated into several independent parts by the first operation, for the non-inclusive operations cannot intersect, and the final answer, which is the sum of 𝑆
 of the parts, will not be less than the initial answer according to the corollary, for some of the equal pairs are separated.

Now the original problem is converted into a data structure task. Consider all adjacent equal pairs (𝑙,𝑟)
, and 𝑚
 is the maximum element between (𝑙,𝑟)
, $m < a_l$. (𝑙,𝑟)
 decrease the answer if and only if 𝑚<𝐿≤𝑎𝑙=𝑎𝑟≤𝑅
, which can be easily calculated with a segment tree and sweeping lines.

