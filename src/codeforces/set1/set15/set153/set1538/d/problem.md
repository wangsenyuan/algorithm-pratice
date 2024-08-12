You are given two integers 𝑎
 and 𝑏
. In one turn, you can do one of the following operations:

Take an integer 𝑐
 (𝑐>1
 and 𝑎
 should be divisible by 𝑐
) and replace 𝑎
 with 𝑎𝑐
;
Take an integer 𝑐
 (𝑐>1
 and 𝑏
 should be divisible by 𝑐
) and replace 𝑏
 with 𝑏𝑐
.
Your goal is to make 𝑎
 equal to 𝑏
 using exactly 𝑘
 turns.

For example, the numbers 𝑎=36
 and 𝑏=48
 can be made equal in 4
 moves:

𝑐=6
, divide 𝑏
 by 𝑐
 ⇒
 𝑎=36
, 𝑏=8
;
𝑐=2
, divide 𝑎
 by 𝑐
 ⇒
 𝑎=18
, 𝑏=8
;
𝑐=9
, divide 𝑎
 by 𝑐
 ⇒
 𝑎=2
, 𝑏=8
;
𝑐=4
, divide 𝑏
 by 𝑐
 ⇒
 𝑎=2
, 𝑏=2
.
For the given numbers 𝑎
 and 𝑏
, determine whether it is possible to make them equal using exactly 𝑘
 turns.

 ### ideas
 1. 