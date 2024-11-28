One of Arkady's friends works at a huge radio telescope. A few decades ago the telescope has sent a signal 𝑠
 towards a faraway galaxy. Recently they've received a response 𝑡
 which they believe to be a response from aliens! The scientists now want to check if the signal 𝑡
 is similar to 𝑠
.

The original signal 𝑠
 was a sequence of zeros and ones (everyone knows that binary code is the universe-wide language). The returned signal 𝑡
, however, does not look as easy as 𝑠
, but the scientists don't give up! They represented 𝑡
 as a sequence of English letters and say that 𝑡
 is similar to 𝑠
 if you can replace all zeros in 𝑠
 with some string 𝑟0
 and all ones in 𝑠
 with some other string 𝑟1
 and obtain 𝑡
. The strings 𝑟0
 and 𝑟1
 must be different and non-empty.

Please help Arkady's friend and find the number of possible replacements for zeros and ones (the number of pairs of strings 𝑟0
 and 𝑟1
) that transform 𝑠
 to 𝑡
.

### ideas
1. 假设 r0, r1, 那么有 cnt[0] * len(r0) + cnt[1] * len(r1) = len(t)
2. 且这个t的前缀[:i] 肯定是r0(假设0开始)
3. 那么就可以计算出来 len(r1)
4. 可以找到，r1理论上也可以找出来
5. 但是问题在于，这样的r0有至少n-1个
6. 如果s[0] = s[n-1] = '0' 
7. 所以要快速的处理
8. 要用hash，这样子，就可以快速的进行检查