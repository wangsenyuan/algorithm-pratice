You are given a positive integer number 𝑛
. You really love good numbers so you want to find the smallest good number greater than or equal to 𝑛
.

The positive integer is called good if it can be represented as a sum of distinct powers of 3
 (i.e. no duplicates of powers of 3
 are allowed).

For example:

30
 is a good number: 30=33+31
,
1
 is a good number: 1=30
,
12
 is a good number: 12=32+31
,
but 2
 is not a good number: you can't represent it as a sum of distinct powers of 3
 (2=30+30
),
19
 is not a good number: you can't represent it as a sum of distinct powers of 3
 (for example, the representations 19=32+32+30=32+31+31+31+30
 are invalid),
20
 is also not a good number: you can't represent it as a sum of distinct powers of 3
 (for example, the representation 20=32+32+30+30
 is invalid).
Note, that there exist other representations of 19
 and 20
 as sums of powers of 3
 but none of them consists of distinct powers of 3
.

For the given positive integer 𝑛
 find such smallest 𝑚
 (𝑛≤𝑚
) that 𝑚
 is a good number.

You have to answer 𝑞
 independent queries.

### ideas
1. 先计算出来最小的3倍数 <= n, 然后再加上3