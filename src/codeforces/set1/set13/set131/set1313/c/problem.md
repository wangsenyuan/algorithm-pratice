This is a harder version of the problem. In this version 𝑛≤500000

The outskirts of the capital are being actively built up in Berland. The company "Kernel Panic" manages the construction of a residential complex of skyscrapers in New Berlskva. All skyscrapers are built along the highway. It is known that the company has already bought 𝑛
 plots along the highway and is preparing to build 𝑛
 skyscrapers, one skyscraper per plot.

Architects must consider several requirements when planning a skyscraper. Firstly, since the land on each plot has different properties, each skyscraper has a limit on the largest number of floors it can have. Secondly, according to the design code of the city, it is unacceptable for a skyscraper to simultaneously have higher skyscrapers both to the left and to the right of it.

Formally, let's number the plots from 1
 to 𝑛
. Then if the skyscraper on the 𝑖
-th plot has 𝑎𝑖
 floors, it must hold that 𝑎𝑖
 is at most 𝑚𝑖
 (1≤𝑎𝑖≤𝑚𝑖
). Also there mustn't be integers 𝑗
 and 𝑘
 such that 𝑗<𝑖<𝑘
 and 𝑎𝑗>𝑎𝑖<𝑎𝑘
. Plots 𝑗
 and 𝑘
 are not required to be adjacent to 𝑖
.

The company wants the total number of floors in the built skyscrapers to be as large as possible. Help it to choose the number of floors for each skyscraper in an optimal way, i.e. in such a way that all requirements are fulfilled, and among all such construction plans choose any plan with the maximum possible total number of floors.

### ideas
1. 只能是递增，然后递减，这样一个结构
2. 如果以i为最大值f(i) = f(i-1) + a[i] 如果 a[i] >= a[i-1]
3. 否则的话， f(i) = f(j) + (i - j) * a[i] 如果a[j] <= a[i]
4. j是i左边的，a[j] <= a[i]的值
5. 