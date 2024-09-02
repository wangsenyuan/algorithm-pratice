Phoenix has decided to become a scientist! He is currently investigating the growth of bacteria.

Initially, on day 1
, there is one bacterium with mass 1
.

Every day, some number of bacteria will split (possibly zero or all). When a bacterium of mass 𝑚
 splits, it becomes two bacteria of mass 𝑚2
 each. For example, a bacterium of mass 3
 can split into two bacteria of mass 1.5
.

Also, every night, the mass of every bacteria will increase by one.

Phoenix is wondering if it is possible for the total mass of all the bacteria to be exactly 𝑛
. If it is possible, he is interested in the way to obtain that mass using the minimum possible number of nights. Help him become the best scientist!


### ideas
1. use binary search
2. 每天都要分裂，分裂的越早越好
3. 假设用了x个nights，那么就有x个day
4. 然后模拟这个过程就可以了