BerPhone X is almost ready for release with 𝑛
 applications being preinstalled on the phone. A category of an application characterizes a genre or a theme of this application (like "game", "business", or "education"). The categories are given as integers between 1
 and 𝑛
, inclusive; the 𝑖
-th application has category 𝑐𝑖
.

You can choose 𝑚
 — the number of screens and 𝑠
 — the size of each screen. You need to fit all 𝑛
 icons of the applications (one icon representing one application) meeting the following requirements:

On each screen, all the icons must belong to applications of the same category (but different screens can contain icons of applications of the same category);
Each screen must be either completely filled with icons (the number of icons on the screen is equal to 𝑠
) or almost filled with icons (the number of icons is equal to 𝑠−1
).
Your task is to find the minimal possible number of screens 𝑚
.


### ideas
1. 能二分吗？
2. 假设category x有k个个icons
3. 且它共使用了w个屏幕, 
4. w * (s - 1) <= k
5. w * s >= k
6. 那么这样子可以计算出s的大小 (其实还没有，因为不知道w的数量)
7. 理论上s越大，使用的屏幕越少
8. 所以，这里 s <= min(k) + 1
9. 且，必须满足, 上面的式子
10. 