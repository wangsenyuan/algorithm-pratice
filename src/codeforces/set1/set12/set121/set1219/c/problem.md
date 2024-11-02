Alice became interested in periods of integer numbers. We say positive 𝑋
 integer number is periodic with length 𝐿
 if there exists positive integer number 𝑃
 with 𝐿
 digits such that 𝑋
 can be written as 𝑃𝑃𝑃𝑃…𝑃
. For example:

𝑋=123123123
 is periodic number with length 𝐿=3
 and 𝐿=9

𝑋=42424242
 is periodic number with length 𝐿=2,𝐿=4
 and 𝐿=8

𝑋=12345
 is periodic number with length 𝐿=5

For given positive period length 𝐿
 and positive integer number 𝐴
, Alice wants to find smallest integer number 𝑋
 strictly greater than 𝐴
 that is periodic with length L.

 ### ideas
 1. X的周期为L，那么X的长度，必然是L的倍数
 2. 如果A的长度为l，给定的L，那么X的长度 = (l + L - 1)/L * L
 3. 如果len(X) > l, 那么可以用100010000 这样子的结构
 4. 如果len(X) = l, 且A = '9999' 那么， len(X) + L and 1000010000这样的结构
 5. 现在A < 9999 （全部是9的情况）
 6. 在保证前面一致的情况下，是否能在后面得到更大的值
 7. 如果在i的后面[i+1...L]中间存在<9 的数，那么在这里肯定可以取值和A相同
 8. 有一种特殊的形式，就是前L和A完全相同，看这个是否可以大于〉A。如果可以，答案就是这个。
 9. 否则，就从L开始找第一个可以变大的数