At first, there was a legend related to the name of the problem, but now it's just a formal statement.

You are given 𝑛
 points 𝑎1,𝑎2,…,𝑎𝑛
 on the 𝑂𝑋
 axis. Now you are asked to find such an integer point 𝑥
 on 𝑂𝑋
 axis that 𝑓𝑘(𝑥)
 is minimal possible.

The function 𝑓𝑘(𝑥)
 can be described in the following way:

form a list of distances 𝑑1,𝑑2,…,𝑑𝑛
 where 𝑑𝑖=|𝑎𝑖−𝑥|
 (distance between 𝑎𝑖
 and 𝑥
);
sort list 𝑑
 in non-descending order;
take 𝑑𝑘+1
 as a result.
If there are multiple optimal answers you can print any of them.