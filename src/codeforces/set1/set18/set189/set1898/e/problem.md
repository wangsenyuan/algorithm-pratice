Sofia has a string 𝑠
of length 𝑛
, consisting only of lowercase English letters. She can perform operations of the following types with this string.

Select an index 1≤𝑖≤|𝑠|
and remove the character 𝑠𝑖
from the string.
Select a pair of indices (𝑙,𝑟)
(1≤𝑙≤𝑟≤|𝑠|
) and sort the substring 𝑠𝑙𝑠𝑙+1…𝑠𝑟
in alphabetical order.
Here, |𝑠|
denotes the current length of 𝑠
. In particular, |𝑠|=𝑛
before the first operation. For example, if 𝑠=𝚜𝚘𝚏𝚒𝚊
, then performing the operation of the first type with 𝑖=4
results in 𝑠
becoming 𝚜𝚘𝚏𝚊
, and performing the operation of the second type with (𝑙,𝑟)=(2,4)
after that results in 𝑠
becoming 𝚜𝚊𝚏𝚘
.
Sofia wants to obtain the string 𝑡
of length 𝑚
after performing zero or more operations on string 𝑠
as described above. Please determine whether it is possible or not.

### thoughts

1. 是不是操作1先执行，对结果是没有影响的呢？直觉上，应该是ok的
2. 感觉操作2，就是可以随便排序啊
3. 那只要看剩余的字符，是不是和t的字符集一致，就OK了？
4. 操作2，不是随便排序，而是只能顺序
5. 如果在t中， t[i] > t[i+1] 那么i和 i+1不能出现在一个区间呢
6. 上面5的结论不大对
7. 如果s[i] 要出现在位置t[j], 那么假设当前位置是s[k]， 那么 s[j] <= s[k...j)
8. 所以，这里尽量要保留，只在需要删除时，进行删除