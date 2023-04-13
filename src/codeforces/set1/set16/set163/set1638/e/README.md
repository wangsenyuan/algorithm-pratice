You have an array 𝑎1,𝑎2,…,𝑎𝑛
. Each element initially has value 0
and color 1
. You are also given 𝑞
queries to perform:

Color 𝑙
𝑟
𝑐
: Change the color of elements 𝑎𝑙,𝑎𝑙+1,⋯,𝑎𝑟
to 𝑐
(1≤𝑙≤𝑟≤𝑛
, 1≤𝑐≤𝑛
).
Add 𝑐
𝑥
: Add 𝑥
to values of all elements 𝑎𝑖
(1≤𝑖≤𝑛
) of color 𝑐
(1≤𝑐≤𝑛
, −109≤𝑥≤109
).
Query 𝑖
: Print 𝑎𝑖
(1≤𝑖≤𝑛
).
Input
The first line of input contains two integers 𝑛
and 𝑞
(1≤𝑛,𝑞≤106
) — the length of array 𝑎
and the number of queries you have to perform.

Each of the next 𝑞
lines contains the query given in the form described in the problem statement.

Output
Print the answers to the queries of the third type on separate lines.

In the first part, let's consider that for all update operations 𝑙=𝑟
. The idea is not to update each element in an Add operation and instead, keeping an array 𝑙𝑎𝑧𝑦[𝑐𝑜𝑙𝑜𝑟]
which stores for each color the total sum we must add to it (because we didn't do it when we had to).

Lets's discuss each operation:

Update 𝑙
𝑟
𝑐
:
We will use the notation 𝑙=𝑟=𝑖
. In this operation we change the color of element 𝑖
from 𝑐′
to 𝑐
. First, remember that we have the sum 𝑙𝑎𝑧𝑦[𝑐′]
that we haven't added to any of the elements of color 𝑐′
(including 𝑖
), so we better do it now because the color changes: 𝑎[𝑖]:=𝑎[𝑖]+𝑙𝑎𝑧𝑦[𝑐′]
. Now we can change the color to 𝑐
.

But wait, what about 𝑙𝑎𝑧𝑦[𝑐]
? It says that we will need to add some value to element 𝑖
, but this is obviously false, since now it is up to date. We can compensate by subtracting now 𝑙𝑎𝑧𝑦[𝑐]
from element 𝑖
, repairing the mistake we will do later: 𝑎[𝑖]:=𝑎[𝑖]−𝑙𝑎𝑧𝑦[𝑐]
. Finally, don't forget to set 𝑐𝑜𝑙𝑜𝑟[𝑖]:=𝑐
.

Add 𝑐
𝑥
:
This is as simple as it gets: 𝑙𝑎𝑧𝑦[𝑐]:=𝑙𝑎𝑧𝑦[𝑐]+𝑥
.

Query 𝑖
:
The query operation is also not very complicated. We print the value 𝑎[𝑖]
and don't forget about 𝑙𝑎𝑧𝑦[𝑐𝑜𝑙𝑜𝑟[𝑖]]
: 𝑝𝑟𝑖𝑛𝑡(𝑎[𝑖]+𝑙𝑎𝑧𝑦[𝑐𝑜𝑙𝑜𝑟[𝑖]])
.

The time complexity is 𝑂(1)
per query.

Now we get back to the initial problem and remove the restriction 𝑙=𝑟
. Let's keep an array of maximal intervals of elements with the same color. We will name them color intervals. By doing
so, we can keep the 𝑙𝑎𝑧𝑦
value for the a whole color interval. When we change the color of all elements in [𝑙,𝑟]
, there are two kinds of color intervals that interest us in our array:

[𝑙′,𝑟′]⊆[𝑙,𝑟]
:
In this case the whole interval changes its color. First, we add the 𝑙𝑎𝑧𝑦
values to each interval. Then, after changing the color of all these intervals into the same one, we can merge them all.
Now we update the resulting interval similar to how we would update a single element.

𝑙∈[𝑙′,𝑟′]
or 𝑟∈[𝑙′,𝑟′]
(or both):
These are the two (or one) intervals that contain the endpoints 𝑙
and 𝑟
. Here we will first split the color interval into two (or three) smaller ones: outside and inside [𝑙,𝑟]
. Then, we just update the one inside [𝑙,𝑟]
as before.

Notice that in contrast to the solution for 𝑙=𝑟
, here we have to add some value on a range. We can do this using a data structure such as Fenwick tree or segment tree
in 𝑂(𝑙𝑜𝑔2(𝑛))
. Also, for storing the color intervals we can use a set. This allows insertions and deletions, as well as quickly
finding the range of intervals modified in a coloring.

The time complexity is a bit tricky to determine because at first it might seem like it is 𝑂(𝑞⋅𝑛)
, but if we analyze the effect each update has on the long term, it turns out to be much better.

We will further refer to the number of intervals in our array as the potential of the current state. Let's consider that
in our update we found 𝑘
color intervals contained in the update interval, the potential decreases by 𝑘−1
and then it grows by at most 2
(because of the two splits). The number of steps our program performs is proportional to the total changes in potential.

In one operation, the potential can decrease by a lot, but luckily, it can only grow by 2
. Because the potential is always positive, it decreases in total at most as much as it increases. Thus, the total
change in potential is 𝑂(𝑞)
.

Although not described here, there exists another solution to this problem using only a segment tree with lazy
propagation. In this solution, our data structure stops only on monochrome segments. The time complexity is the same.

Time complexity: 𝑂(𝑛+𝑞⋅𝑙𝑜𝑔2(𝑛))
.