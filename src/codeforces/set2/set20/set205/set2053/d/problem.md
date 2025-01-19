Chris is given two arrays 𝑎
 and 𝑏
, both consisting of 𝑛
 integers.
Iris is interested in the largest possible value of 𝑃=∏𝑖=1𝑛min(𝑎𝑖,𝑏𝑖)
 after an arbitrary rearrangement of 𝑏
. Note that she only wants to know the maximum value of 𝑃
, and no actual rearrangement is performed on 𝑏
.
There will be 𝑞
 modifications. Each modification can be denoted by two integers 𝑜
 and 𝑥
 (𝑜
 is either 1
 or 2
, 1≤𝑥≤𝑛
). If 𝑜=1
, then Iris will increase 𝑎𝑥
 by 1
; otherwise, she will increase 𝑏𝑥
 by 1
.
Iris asks Chris the maximum value of 𝑃
 for 𝑞+1
 times: once before any modification, then after every modification.
Since 𝑃
 might be huge, Chris only needs to calculate it modulo 998244353
.