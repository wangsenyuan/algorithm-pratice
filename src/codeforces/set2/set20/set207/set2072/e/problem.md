Akito decided to study a new powerful spell. Since it possesses immeasurable strength, it certainly requires a lot of space and careful preparation. For this, Akito went out into the field. Let's represent the field as a Cartesian coordinate system.

For the spell, Akito needs to place 0≤𝑛≤500
 staffs at distinct integer coordinates in the field such that there will be exactly 𝑘
 pairs (𝑖,𝑗)
 such that 1≤𝑖<𝑗≤𝑛
 and 𝜌(𝑖,𝑗)=𝑑(𝑖,𝑗)
.

Here, for two points with integer coordinates 𝑎=(𝑥𝑎,𝑦𝑎)
 and 𝑏=(𝑥𝑏,𝑦𝑏)
, 𝜌(𝑎,𝑏)=(𝑥𝑎−𝑥𝑏)2+(𝑦𝑎−𝑦𝑏)2‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾√
 and 𝑑(𝑎,𝑏)=|𝑥𝑎−𝑥𝑏|+|𝑦𝑎−𝑦𝑏|
.

Input
The first line of input contains a single number 𝑡
 (1≤𝑡≤1000
) — the number of test cases.

In the only line of each test case, there is a single number 𝑘
 (0≤𝑘≤105
) — the number of pairs of staffs for which the equality 𝜌(𝑖,𝑗)=𝑑(𝑖,𝑗)
 must hold.

Output
For each test case, the first line of output should print the number 𝑛
 (0≤𝑛≤500
) — the number of placed staffs.

In the following 𝑛
 lines, pairs of integers 𝑥𝑖,𝑦𝑖
 (−109≤𝑥𝑖,𝑦𝑖≤109)
 should be printed — the coordinates of the 𝑖
-th staff. The points in which staffs stand must be distinct.