In this game, a level consists of 𝑛
 pairs of gates. Each pair contains one left gate and one right gate. Each gate performs one of two operations:

Addition Operation (+ a): Increases the number of people in a lane by a constant amount 𝑎
.
Multiplication Operation (x a): Multiplies the current number of people in a lane by an integer 𝑎
. This means the number of people increases by (𝑎−1)
 times the current count in that lane.
The additional people gained from each operation can be assigned to either lane. However, people already in a lane cannot be moved to the other lane.

Initially, there is one person in each lane. Your task is to determine the maximum total number of people that can be achieved by the end of the level.


### ideas
1. 只用关心乘法