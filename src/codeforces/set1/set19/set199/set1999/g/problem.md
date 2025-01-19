We have a secret ruler that is missing one number 𝑥
 (2≤𝑥≤999
). When you measure an object of length 𝑦
, the ruler reports the following values:

If 𝑦<𝑥
, the ruler (correctly) measures the object as having length 𝑦
.
If 𝑦≥𝑥
, the ruler incorrectly measures the object as having length 𝑦+1
.

The ruler above is missing the number 4
, so it correctly measures the first segment as length 3
 but incorrectly measures the second segment as length 6
 even though it is actually 5
.

You need to find the value of 𝑥
. To do that, you can make queries of the following form:

? 𝑎 𝑏
 — in response, we will measure the side lengths of an 𝑎×𝑏
 rectangle with our ruler and multiply the results, reporting the measured area of the rectangle back to you. For example, if 𝑥=4
 and you query a 3×5
 rectangle, we will measure its side lengths as 3×6
 and report 18
 back to you.
Find the value of 𝑥
. You can ask at most 7
 queries.


 ### ideas
 1. 如果缺少的正好是1
 2. (a, b) => (a + 1) * (b + 1) = a * b + a + b + 1
 3. 如何唯一的确定，肯定是上面的结果呢？
 4. 如果缺少的是2
 5. (a, b), 如果 a = 1, b > 2 => a * (b + 1) = a * b + a
 6. 如果 (1, 1) = 4 => x = 1(不可能是其他的)
 7. 如果 (1, 1) = 1 => x > 1
 8. (1, b) 二分（看看要几次）
 9. (1, b) = b => b < x(但是2分，要10次)
 10. 如果 (512, 512) = 1024 => x > 512 (这样子，还是需要9次)
 11.  所以不大行
 12.  (128, 512) 如果等于 128 * 512 => 512 < x
 13.  如果 128 * 512 = 129 * 512 => 128 >= x
 14.  如果 128 * 512 = 128 * 513 => 128 < x <= 512
 15.  也就是说一次查询，变成了1/3