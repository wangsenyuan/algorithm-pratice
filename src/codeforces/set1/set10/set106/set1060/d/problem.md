You invited 𝑛
guests to dinner! You plan to arrange one or more circles of chairs. Each chair is going to be either occupied by one
guest, or be empty. You can make any number of circles.

Your guests happen to be a little bit shy, so the 𝑖
-th guest wants to have a least 𝑙𝑖
free chairs to the left of his chair, and at least 𝑟𝑖
free chairs to the right. The "left" and "right" directions are chosen assuming all guests are going to be seated
towards the center of the circle. Note that when a guest is the only one in his circle, the 𝑙𝑖
chairs to his left and 𝑟𝑖
chairs to his right may overlap.

What is smallest total number of chairs you have to use?

### ideas

1. 如果一个桌子上只有一个人，那么这个桌子的大小至少 max(l, r) + 1
2. 但是如果是多个人，是否将它们安排在一个桌子上更合理？或者要安排在一个桌子上，什么情况下更省椅子？
3. 假设按照顺时针安排，最后一个人安排是i, 现在安排是j (j不需要=i+1)
4. 那么如果 r[j] = l[i] 那么它们之间没有浪费， 且如果 l[j] = r[first]，那么就可以完美完成一个圈
5. 