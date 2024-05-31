Iulia has 𝑛
 glasses arranged in a line. The 𝑖
-th glass has 𝑎𝑖
 units of juice in it. Iulia drinks only from odd-numbered glasses, while her date drinks only from even-numbered glasses.

To impress her date, Iulia wants to find a contiguous subarray of these glasses such that both Iulia and her date will have the same amount of juice in total if only the glasses in this subarray are considered. Please help her to do that.

More formally, find out if there exists two indices 𝑙
, 𝑟
 such that 1≤𝑙≤𝑟≤𝑛
, and 𝑎𝑙+𝑎𝑙+2+𝑎𝑙+4+⋯+𝑎𝑟=𝑎𝑙+1+𝑎𝑙+3+⋯+𝑎𝑟−1
 if 𝑙
 and 𝑟
 have the same parity and 𝑎𝑙+𝑎𝑙+2+𝑎𝑙+4+⋯+𝑎𝑟−1=𝑎𝑙+1+𝑎𝑙+3+⋯+𝑎𝑟
 otherwise.

 ### ideas
 1. sum[i] += a[i] if i is odd
 2. sum[i] -= a[i] if i is even
 3. 假设这段时l...r, 
 4. 如果 l是奇数，r是偶数, sum[l....r] = 0
 5. 所以，只要判断sum[l-1] == sum[r]即可
 6. 如果l是奇数，r也是奇数, 好像也是一样的
 7. 如果l是偶数，也是一样的
 8. 