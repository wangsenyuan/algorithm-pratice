The Squareland national forest is divided into equal 1×1
square plots aligned with north-south and east-west directions. Each plot can be uniquely described by integer Cartesian
coordinates (𝑥,𝑦)
of its south-west corner.

Three friends, Alice, Bob, and Charlie are going to buy three distinct plots of land 𝐴,𝐵,𝐶
in the forest. Initially, all plots in the forest (including the plots 𝐴,𝐵,𝐶
) are covered by trees. The friends want to visit each other, so they want to clean some of the plots from trees. After
cleaning, one should be able to reach any of the plots 𝐴,𝐵,𝐶
from any other one of those by moving through adjacent cleared plots. Two plots are adjacent if they share a side.

### ideas

1. 如果B在A,C的中间， 那么只需要计算A，C的距离
2. 如果B在A，C的外部