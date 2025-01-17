Find the number, modulo 
998244353, of permutations 
P=(P 
1
​
 ,P 
2
​
 ,⋯,P 
N
​
 ) of 
(1,2,⋯,N) that satisfy all of the following 
M conditions.

The 
i-th condition: The maximum among 
P 
L 
i
​
 
​
 ,P 
L 
i
​
 +1
​
 ,⋯,P 
R 
i
​
 
​
  is not 
P 
X 
i
​
 
​
 . Here, 
L 
i
​
 , 
R 
i
​
 , and 
X 
i
​
  are integers given in the input.


  ### ideas
  1. 考虑n的位置，显然n不能是任何一个condition中的x（假设这样的位置有w个）
  2. 把n放在其中某个位置上后，可能会让一些condition满足（n放在了非x的位置）
  3. 还有一些condition没有被覆盖到
  4. dp[i][j] 表示在前i个位置放置了n,n-1, ...j 时，满足所有覆盖到的条件
  5. 现在考虑放置位置i, n - i + 1
  6. 如果i是某个condition的x，且i也是这个condition的l
  7. 那么就不能使用n-i+1在这里（因为这个是当前最大的数）
  8. n-i+1必须放在剩余空间中，非x的位置（那些未被满足的condition）
  9. dp[i][w] 表示前i个condition被满足的情况下，现在考虑w的位置
  10. 但是前i个conditioin被满足，无法提供更确切的信息，比如哪些位置是还能用的。
  11. 考虑两个条件，如果它们在中间有重叠，且x1,x2在这个重叠区间外
  12. 这种情况，感觉都做不出来～～～
  13. 考虑一个区间的情况
  14. C(n, sz) * P(n - sz) * (P(sz) - P(sz-1)) 选择sz个数，剩余的可以随便排列，中间的sz的数随便排列，但是去掉x是最大值的情况
  15. 单独一个区间，P(sz) - P(sz - 1)  是不是它们乘起来就可以了？但是貌似有重合的
  16. 