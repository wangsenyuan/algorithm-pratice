In a coordinate space, we want to place three cubes with a side length of 
7 so that the volumes of the regions contained in exactly one, two, three cube(s) are 
V 
1
​
 , 
V 
2
​
 , 
V 
3
​
 , respectively.

For three integers 
a, 
b, 
c, let 
C(a,b,c) denote the cubic region represented by 
(a≤x≤a+7)∧(b≤y≤b+7)∧(c≤z≤c+7).

Determine whether there are nine integers 
a 
1
​
 ,b 
1
​
 ,c 
1
​
 ,a 
2
​
 ,b 
2
​
 ,c 
2
​
 ,a 
3
​
 ,b 
3
​
 ,c 
3
​
  that satisfy all of the following conditions, and find one such tuple if it exists.

∣a 
1
​
 ∣,∣b 
1
​
 ∣,∣c 
1
​
 ∣,∣a 
2
​
 ∣,∣b 
2
​
 ∣,∣c 
2
​
 ∣,∣a 
3
​
 ∣,∣b 
3
​
 ∣,∣c 
3
​
 ∣≤100
Let 
C 
i
​
 =C(a 
i
​
 ,b 
i
​
 ,c 
i
​
 ) (i=1,2,3).
The volume of the region contained in exactly one of 
C 
1
​
 ,C 
2
​
 ,C 
3
​
  is 
V 
1
​
 .
The volume of the region contained in exactly two of 
C 
1
​
 ,C 
2
​
 ,C 
3
​
  is 
V 
2
​
 .
The volume of the region contained in all of 
C 
1
​
 ,C 
2
​
 ,C 
3
​
  is 
V 
3
​
 .


### ideas
1.  v1是，a, b, c包含只在任意一个区域中（不在重叠区域）的点的数量
2.  v2是在任意两个重叠区域中的点的数量
3.  v3时在3个重叠区域中的点的数量
4.  V * 3 = v1 - v2 + v3
5.     = 840 + 84 + 7 = 931
6.     6 * 6 * 7 + 6 * 7 * 7 + 6 * 7 * 7 = 840
7.  v1 + v2 + v3 <= V * 3 = 1029
8.  vx =  V * 3 - v2 - 2 * v3 = 表示的是，在任意一个区域中的点的数量
9.     = 1029 - 84 - 14 = 952 = v1 + v2 + v3 = 931
10.    3 * V = v1 + 2 * v2 + 3 * v3 成立时，始终是结果的
11. 第一个立方体的位置是确定的, 就在原点
12. 考虑v3， = h * w * l 分别表示，任意两个立方体重叠的距离
13.  考虑到v3比较小(h, w, l <= 7) 似乎是可以迭代的
14.  其中 (h, w) 是和第一个立方体重叠的高度和深度
15.  有了这两个信息，可以确定 第二个立方体的位置
16.  同样的，可以把 （h, l）来确定第三个立方体的位置（并用 w, l来验证2和3的位置）