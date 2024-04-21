You are given two huge binary integer numbers 𝑎
and 𝑏
of lengths 𝑛
and 𝑚
respectively. You will repeat the following process: if 𝑏>0
, then add to the answer the value 𝑎 & 𝑏
and divide 𝑏
by 2
rounding down (i.e. remove the last digit of 𝑏
), and repeat the process again, otherwise stop the process.

The value 𝑎 & 𝑏
means bitwise AND of 𝑎
and 𝑏
. Your task is to calculate the answer modulo 998244353
.

Note that you should add the value 𝑎 & 𝑏
to the answer in decimal notation, not in binary. So your task is to calculate the answer in decimal notation. For
example, if 𝑎=10102 (1010)
and 𝑏=10002 (810)
, then the value 𝑎 & 𝑏
will be equal to 8
, not to 1000
.

### ideas

1. 计算a中每位1的贡献