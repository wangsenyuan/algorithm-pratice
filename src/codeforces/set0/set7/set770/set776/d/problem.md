Moriarty has trapped n people in n distinct rooms in a hotel. Some rooms are locked, others are unlocked. But, there is
a condition that the people in the hotel can only escape when all the doors are unlocked at the same time. There are m
switches. Each switch control doors of some rooms, but each door is controlled by exactly two switches.

You are given the initial configuration of the doors. Toggling any switch, that is, turning it ON when it is OFF, or
turning it OFF when it is ON, toggles the condition of the doors that this switch controls. Say, we toggled switch 1,
which was connected to room 1, 2 and 3 which were respectively locked, unlocked and unlocked. Then, after toggling the
switch, they become unlocked, locked and locked.

You need to tell Sherlock, if there exists a way to unlock all doors at the same time.

### ideas

1. room和switch之间连一条先， room有两条线
2. 如果room是locked，那么它必须要使用一个switch 处理一下
3. 如果两个switch都take off，相当于对于room没有影响（这种room只能是非locked)
4. 也就是说locked的room，要求它的两个switch的颜色要相反（只有一个被触发了）
5. 非locked的room，要求它的两个switch的颜色要相同（同时触发/或不触发）
6. 所以要对switch进行着色，如果存在一个合理的着色方案，就是可行的，否则就是不可行的