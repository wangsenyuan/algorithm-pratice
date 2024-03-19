On a strip of land of length 𝑛
there are 𝑘
air conditioners: the 𝑖
-th air conditioner is placed in cell 𝑎𝑖
(1≤𝑎𝑖≤𝑛
). Two or more air conditioners cannot be placed in the same cell (i.e. all 𝑎𝑖
are distinct).

Each air conditioner is characterized by one parameter: temperature. The 𝑖
-th air conditioner is set to the temperature 𝑡𝑖
.

Example of strip of length 𝑛=6
, where 𝑘=2
, 𝑎=[2,5]
and 𝑡=[14,16]
.
For each cell 𝑖
(1≤𝑖≤𝑛
) find it's temperature, that can be calculated by the formula
min1≤𝑗≤𝑘(𝑡𝑗+|𝑎𝑗−𝑖|),
where |𝑎𝑗−𝑖|
denotes absolute value of the difference 𝑎𝑗−𝑖
.

In other words, the temperature in cell 𝑖
is equal to the minimum among the temperatures of air conditioners, increased by the distance from it to the cell 𝑖
.

Let's look at an example. Consider that 𝑛=6,𝑘=2
, the first air conditioner is placed in cell 𝑎1=2
and is set to the temperature 𝑡1=14
and the second air conditioner is placed in cell 𝑎2=5
and is set to the temperature 𝑡2=16
. In that case temperatures in cells are:

temperature in cell 1
is: min(14+|2−1|,16+|5−1|)=min(14+1,16+4)=min(15,20)=15
;
temperature in cell 2
is: min(14+|2−2|,16+|5−2|)=min(14+0,16+3)=min(14,19)=14
;
temperature in cell 3
is: min(14+|2−3|,16+|5−3|)=min(14+1,16+2)=min(15,18)=15
;
temperature in cell 4
is: min(14+|2−4|,16+|5−4|)=min(14+2,16+1)=min(16,17)=16
;
temperature in cell 5
is: min(14+|2−5|,16+|5−5|)=min(14+3,16+0)=min(17,16)=16
;
temperature in cell 6
is: min(14+|2−6|,16+|5−6|)=min(14+4,16+1)=min(18,17)=17
.
For each cell from 1
to 𝑛
find the temperature in it.

### ideas

1. res[i] = min(t[j] + abs(a[j] - i))
2. when a[j] > i => min(t[j] + a[j] - i)
3. when a[j] < i => min(t[j] + i - a[j])
4. 所以，这里需要按照a[j]排个序，计算i的时候，找到它后面a[j]的值，和前面a[j]的值
5. a[j] <= 1e9, 但是n比较小，超过n的部分，都找到其中的最小的t[j] + a[j]即可