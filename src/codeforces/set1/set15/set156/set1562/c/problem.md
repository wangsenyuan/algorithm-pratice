Somewhere in a parallel Middle-earth, when Saruman caught Frodo, he only found 𝑛
 rings. And the 𝑖
-th ring was either gold or silver. For convenience Saruman wrote down a binary string 𝑠
 of 𝑛
 characters, where the 𝑖
-th character was 0 if the 𝑖
-th ring was gold, and 1 if it was silver.

Saruman has a magic function 𝑓
, which takes a binary string and returns a number obtained by converting the string into a binary number and then converting the binary number into a decimal number. For example, 𝑓(001010)=10,𝑓(111)=7,𝑓(11011101)=221
.

Saruman, however, thinks that the order of the rings plays some important role. He wants to find 2
 pairs of integers (𝑙1,𝑟1),(𝑙2,𝑟2)
, such that:

1≤𝑙1≤𝑛
, 1≤𝑟1≤𝑛
, 𝑟1−𝑙1+1≥⌊𝑛2⌋
1≤𝑙2≤𝑛
, 1≤𝑟2≤𝑛
, 𝑟2−𝑙2+1≥⌊𝑛2⌋
Pairs (𝑙1,𝑟1)
 and (𝑙2,𝑟2)
 are distinct. That is, at least one of 𝑙1≠𝑙2
 and 𝑟1≠𝑟2
 must hold.
Let 𝑡
 be the substring 𝑠[𝑙1:𝑟1]
 of 𝑠
, and 𝑤
 be the substring 𝑠[𝑙2:𝑟2]
 of 𝑠
. Then there exists non-negative integer 𝑘
, such that 𝑓(𝑡)=𝑓(𝑤)⋅𝑘
.
Here substring 𝑠[𝑙:𝑟]
 denotes 𝑠𝑙𝑠𝑙+1…𝑠𝑟−1𝑠𝑟
, and ⌊𝑥⌋
 denotes rounding the number down to the nearest integer.

Help Saruman solve this problem! It is guaranteed that under the constraints of the problem at least one solution exists.

### ideas
1. 如果在后半段，存在一个s[i] = 0, 可以直接选择[l1, r1] = [0...i], [l2, r2] = [0...i-1] 即可 (k = 2)
2. [1100] t = 110, w = 11即可
3. 否则的话，就是后半段都是1，如果全部是1，那么可以选择k=1
4. 那么 l2, r2 = (n/2 + 1, n )
5. 