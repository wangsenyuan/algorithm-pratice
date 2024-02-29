You are given 𝑛
strings 𝑠1,𝑠2,…,𝑠𝑛
, consisting of lowercase Latin letters. Let |𝑥|
be the length of string 𝑥
.

Let a collapse 𝐶(𝑎,𝑏)
of two strings 𝑎
and 𝑏
be the following operation:

if 𝑎
is empty, 𝐶(𝑎,𝑏)=𝑏
;
if 𝑏
is empty, 𝐶(𝑎,𝑏)=𝑎
;
if the last letter of 𝑎
is equal to the first letter of 𝑏
, then 𝐶(𝑎,𝑏)=𝐶(𝑎1,|𝑎|−1,𝑏2,|𝑏|)
, where 𝑠𝑙,𝑟
is the substring of 𝑠
from the 𝑙
-th letter to the 𝑟
-th one;
otherwise, 𝐶(𝑎,𝑏)=𝑎+𝑏
, i. e. the concatenation of two strings.
Calculate ∑𝑖=1𝑛∑𝑗=1𝑛|𝐶(𝑠𝑖,𝑠𝑗)|
.