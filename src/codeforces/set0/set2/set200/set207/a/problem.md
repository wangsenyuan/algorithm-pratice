The Smart Beaver from ABBYY has once again surprised us! He has developed a new calculating device, which he called the "Beaver's Calculator 1.0". It is very peculiar and it is planned to be used in a variety of scientific problems.

To test it, the Smart Beaver invited n scientists, numbered from 1 to n. The i-th scientist brought ki calculating problems for the device developed by the Smart Beaver from ABBYY. The problems of the i-th scientist are numbered from 1 to ki, and they must be calculated sequentially in the described order, since calculating each problem heavily depends on the results of calculating of the previous ones.

Each problem of each of the n scientists is described by one integer ai, j, where i (1 ≤ i ≤ n) is the number of the scientist, j (1 ≤ j ≤ ki) is the number of the problem, and ai, j is the number of resource units the calculating device needs to solve this problem.

The calculating device that is developed by the Smart Beaver is pretty unusual. It solves problems sequentially, one after another. After some problem is solved and before the next one is considered, the calculating device allocates or frees resources.

The most expensive operation for the calculating device is freeing resources, which works much slower than allocating them. It is therefore desirable that each next problem for the calculating device requires no less resources than the previous one.

You are given the information about the problems the scientists offered for the testing. You need to arrange these problems in such an order that the number of adjacent "bad" pairs of problems in this list is minimum possible. We will call two consecutive problems in this list a "bad pair" if the problem that is performed first requires more resources than the one that goes after it. Do not forget that the problems of the same scientist must be solved in a fixed order.

### ideas
1. sum of (ki) <= 25000000 (25 * 1e6) 所以可以先算出来，再排序
2. 但是感觉用heap，更方便一些