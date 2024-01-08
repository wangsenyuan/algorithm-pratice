Tenzing has 𝑛
balls arranged in a line. The color of the 𝑖
-th ball from the left is 𝑎𝑖
.

Tenzing can do the following operation any number of times:

select 𝑖
and 𝑗
such that 1≤𝑖<𝑗≤|𝑎|
and 𝑎𝑖=𝑎𝑗
,
remove 𝑎𝑖,𝑎𝑖+1,…,𝑎𝑗
from the array (and decrease the indices of all elements to the right of 𝑎𝑗
by 𝑗−𝑖+1
).
Tenzing wants to know the maximum number of balls he can remove.