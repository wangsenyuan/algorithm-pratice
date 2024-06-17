Polycarp has two favorite integers 𝑥
 and 𝑦
 (they can be equal), and he has found an array 𝑎
 of length 𝑛
.

Polycarp considers a pair of indices ⟨𝑖,𝑗⟩
 (1≤𝑖<𝑗≤𝑛
) beautiful if:

𝑎𝑖+𝑎𝑗
 is divisible by 𝑥
;
𝑎𝑖−𝑎𝑗
 is divisible by 𝑦
.
For example, if 𝑥=5
, 𝑦=2
, 𝑛=6
, 𝑎=
[1,2,7,4,9,6
], then the only beautiful pairs are:

⟨1,5⟩
: 𝑎1+𝑎5=1+9=10
 (10
 is divisible by 5
) and 𝑎1−𝑎5=1−9=−8
 (−8
 is divisible by 2
);
⟨4,6⟩
: 𝑎4+𝑎6=4+6=10
 (10
 is divisible by 5
) and 𝑎4−𝑎6=4−6=−2
 (−2
 is divisible by 2
).
Find the number of beautiful pairs in the array 𝑎
.