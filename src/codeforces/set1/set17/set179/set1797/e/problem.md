Li Hua wants to solve a problem about 𝜑
— Euler's totient function. Please recall that 𝜑(𝑥)=∑𝑖=1𝑥[gcd(𝑖,𝑥)=1]
.†,‡
He has a sequence 𝑎1,𝑎2,⋯,𝑎𝑛
and he wants to perform 𝑚
operations:

"1 𝑙
𝑟
" (1≤𝑙≤𝑟≤𝑛
) — for each 𝑥∈[𝑙,𝑟]
, change 𝑎𝑥
into 𝜑(𝑎𝑥)
.
"2 𝑙
𝑟
" (1≤𝑙≤𝑟≤𝑛
) — find out the minimum changes needed to make sure 𝑎𝑙=𝑎𝑙+1=⋯=𝑎𝑟
. In each change, he chooses one 𝑥∈[𝑙,𝑟]
, change 𝑎𝑥
into 𝜑(𝑎𝑥)
. Each operation of this type is independent, which means the array doesn't actually change.
Suppose you were Li Hua, please solve this problem.

