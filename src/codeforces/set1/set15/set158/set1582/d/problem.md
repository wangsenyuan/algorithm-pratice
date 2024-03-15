Vupsen and Pupsen were gifted an integer array. Since Vupsen doesn't like the number 0
, he threw away all numbers equal to 0
from the array. As a result, he got an array 𝑎
of length 𝑛
.

Pupsen, on the contrary, likes the number 0
and he got upset when he saw the array without zeroes. To cheer Pupsen up, Vupsen decided to come up with another array
𝑏
of length 𝑛
such that ∑𝑛𝑖=1𝑎𝑖⋅𝑏𝑖=0
. Since Vupsen doesn't like number 0
, the array 𝑏
must not contain numbers equal to 0
. Also, the numbers in that array must not be huge, so the sum of their absolute values cannot exceed 109
. Please help Vupsen to find any such array 𝑏
!

### ideas

1. sum(a[i] * b[i]) = 0
2. 如果是长度是偶数，似乎有个简单的做法
3. (x, y) => (-y, x) 就可以了
4. 现在考虑长度是奇数， 单个数字，因为不能使用0，所以不行，考虑(a, b, c)
5. 不是一般性，考虑 a <= b <= c
6. 如果 a + b = c 那可以简单的 (1, 1, -1)
7. 如果 a + b > c, 比如 2, 2, 3, 考虑(a + b, c) 的公倍数， 假设位l
8. l / (a + b), -(l/c)即可