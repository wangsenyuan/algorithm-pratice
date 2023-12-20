### description

There is a street with 𝑛
sights, with sight number 𝑖
being 𝑖
miles from the beginning of the street. Sight number 𝑖
has beauty 𝑏𝑖
. You want to start your morning jog 𝑙
miles and end it 𝑟
miles from the beginning of the street. By the time you run, you will see sights you run by (including sights at 𝑙
and 𝑟
miles from the start). You are interested in the 3
most beautiful sights along your jog, but every mile you run, you get more and more tired.

So choose 𝑙
and 𝑟
, such that there are at least 3
sights you run by, and the sum of beauties of the 3
most beautiful sights minus the distance in miles you have to run is maximized. More formally, choose 𝑙
and 𝑟
, such that 𝑏𝑖1+𝑏𝑖2+𝑏𝑖3−(𝑟−𝑙)
is maximum possible, where 𝑖1,𝑖2,𝑖3
are the indices of the three maximum elements in range [𝑙,𝑟]
.

### solution

- r - l + 1 >= 3
- 3个最大值中的两个，必然在l和r上，这是因为，加入不在，通过减少r和增加l，会得到更好的结果
- 在确定l和r的时候，只需要找到中间的最大的i即可
- 但显然，这样的复杂度是 n * n
- bl + bi + br - (r - l) => bl + l + br - r + bi的最大值
- 如果从i的角度看，似乎就是在它前面的最大的 bl + l 和它后面最大的br - r

[source](https://codeforces.com/problemset/problem/1826/D)

