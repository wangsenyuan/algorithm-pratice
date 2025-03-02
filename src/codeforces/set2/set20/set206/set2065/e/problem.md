Given a binary string∗
 𝑡
, let 𝑥
 represent the number of 𝟶
 in 𝑡
 and 𝑦
 represent the number of 𝟷
 in 𝑡
. Its balance-value is defined as the value of max(𝑥−𝑦,𝑦−𝑥)
.

Skibidus gives you three integers 𝑛
, 𝑚
, and 𝑘
. He asks for your help to construct a binary string 𝑠
 of length 𝑛+𝑚
 with exactly 𝑛
 𝟶
's and 𝑚
 𝟷
's such that the maximum balance-value among all of its substrings†
 is exactly 𝑘
. If it is not possible, output -1.