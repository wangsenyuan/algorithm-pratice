Alex is participating in the filming of another video of BrMeast, and BrMeast asked Alex to prepare 250 thousand tons of
TNT, but Alex didn't hear him well, so he prepared 𝑛
boxes and arranged them in a row waiting for trucks. The 𝑖
-th box from the left weighs 𝑎𝑖
tons.

All trucks that Alex is going to use hold the same number of boxes, denoted by 𝑘
. Loading happens the following way:

The first 𝑘
boxes goes to the first truck,
The second 𝑘
boxes goes to the second truck,
⋯
The last 𝑘
boxes goes to the 𝑛𝑘
-th truck.
Upon loading is completed, each truck must have exactly 𝑘
boxes. In other words, if at some point it is not possible to load exactly 𝑘
boxes into the truck, then the loading option with that 𝑘
is not possible.

Alex hates justice, so he wants the maximum absolute difference between the total weights of two trucks to be as great
as possible. If there is only one truck, this value is 0
.

Alex has quite a lot of connections, so for every 1≤𝑘≤𝑛
, he can find a company such that each of its trucks can hold exactly 𝑘
boxes. Print the maximum absolute difference between the total weights of any two trucks.

### thoughts

1. k 是n的factor
2. n的factor一共有logn个