There are 𝑛
 people in the programming contest chat. Chat participants are ordered by activity, but each person sees himself at the top of the list.

For example, there are 4
 participants in the chat, and their order is [2,3,1,4]
. Then

1
-st user sees the order [1,2,3,4]
.
2
-nd user sees the order [2,3,1,4]
.
3
-rd user sees the order [3,2,1,4]
.
4
-th user sees the order [4,2,3,1]
.
𝑘
 people posted screenshots in the chat, which show the order of participants shown to this user. The screenshots were taken within a short period of time, and the order of participants has not changed.

Your task is to determine whether there is a certain order that all screenshots correspond to.

### ideas
1. 除去自己，其他的顺序是固定的，那直接检查，不就可以了？
2. 1的位置，有可能在任何地方