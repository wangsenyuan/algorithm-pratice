You are given two integers 𝑎
 and 𝑏
. You can perform a sequence of operations: during the first operation you choose one of these numbers and increase it by 1
; during the second operation you choose one of these numbers and increase it by 2
, and so on. You choose the number of these operations yourself.

For example, if 𝑎=1
 and 𝑏=3
, you can perform the following sequence of three operations:

add 1
 to 𝑎
, then 𝑎=2
 and 𝑏=3
;
add 2
 to 𝑏
, then 𝑎=2
 and 𝑏=5
;
add 3
 to 𝑎
, then 𝑎=5
 and 𝑏=5
.
Calculate the minimum number of operations required to make 𝑎
 and 𝑏
 equal.

 ### ideas
 1. assume a < b
 2. 假设sa 是a增加的, sb 是b增加的
 3. sa + sb = s
 4. sa - sb = b - a
 5. 且s = 1 + 2 + ... n (是不是brute force)
 6. 