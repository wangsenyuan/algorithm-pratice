XXI Berland Annual Fair is coming really soon! Traditionally fair consists of 𝑛
booths, arranged in a circle. The booths are numbered 1
through 𝑛
clockwise with 𝑛
being adjacent to 1
. The 𝑖
-th booths sells some candies for the price of 𝑎𝑖
burles per item. Each booth has an unlimited supply of candies.

Polycarp has decided to spend at most 𝑇
burles at the fair. However, he has some plan in mind for his path across the booths:

at first, he visits booth number 1
;
if he has enough burles to buy exactly one candy from the current booth, then he buys it immediately;
then he proceeds to the next booth in the clockwise order (regardless of if he bought a candy or not).
Polycarp's money is finite, thus the process will end once he can no longer buy candy at any booth.

Calculate the number of candies Polycarp will buy.

### ideas

1. 如果某个a[i]太大，那么它有可能永远也不会参与进来
2. 进过一轮后，假设消耗了x，那么下一轮消耗的，只会比x（等于或）更小
3. 如果最大的能被处理x轮，那么所有其他的也至少可以被处理x轮