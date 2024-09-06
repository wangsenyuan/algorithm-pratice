You are given an array 𝑎
 of length 𝑛
 and array 𝑏
 of length 𝑚
 both consisting of only integers 0
 and 1
. Consider a matrix 𝑐
 of size 𝑛×𝑚
 formed by following rule: 𝑐𝑖,𝑗=𝑎𝑖⋅𝑏𝑗
 (i.e. 𝑎𝑖
 multiplied by 𝑏𝑗
). It's easy to see that 𝑐
 consists of only zeroes and ones too.

How many subrectangles of size (area) 𝑘
 consisting only of ones are there in 𝑐
?

A subrectangle is an intersection of a consecutive (subsequent) segment of rows and a consecutive (subsequent) segment of columns. I.e. consider four integers 𝑥1,𝑥2,𝑦1,𝑦2
 (1≤𝑥1≤𝑥2≤𝑛
, 1≤𝑦1≤𝑦2≤𝑚
) a subrectangle 𝑐[𝑥1…𝑥2][𝑦1…𝑦2]
 is an intersection of the rows 𝑥1,𝑥1+1,𝑥1+2,…,𝑥2
 and the columns 𝑦1,𝑦1+1,𝑦1+2,…,𝑦2
.

The size (area) of a subrectangle is the total number of cells in it.

### ideas
1. 先要观察特点
2. 任何一段长度为l和r的1的连续段，都可以贡献
3. 1 * 1, 1 * 2, ... 1 * r
4. ....
5. l * 1, l * 2, .... l * r 的计数
6. 那么这个里面共有多少个超过了k呢？共有l * r 个面积 - a * b (<k 的个数)
7. 不止l * r 个， 起点有l * r个
8. 当起点为 (a, b) w * h < k时，这样的起点有多少个
9. 当w = 1时，h = k - 1
10. 当 w = 2时， h = (k - 1) / 2
11. 当 w = 3时， h = (k - 1) / 3
12. .. w = k-1 时, h = 1
13. 1 * (k - 1) * (k - 1) 
14. 但是这里有几个点， 2 * (k - 1) / 2不一定 = k - 1 有可能等于k-2
15. 数学没学好，好烦呐～
16. 如果这个公式能解出来，应该可以在 sqrt(n)时间内完成
17. 这个只和长度有关系，所以按照长度分分组后，总的组数不会超过sqrt(n)
18. 问题就在于，那个公式要咋算呢？
19. (n−h+1)×(m−w+1)