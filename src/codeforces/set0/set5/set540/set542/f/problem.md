Polikarp is preparing a quest for his friends. He has already come up with n tasks, for each of which he has estimated its interestingness as an integer q i , and the amount of time t i in minutes needed to complete it.

An interesting feature of his quest is that each participant should get the most suitable task, depending on his preferences. The task is selected on the basis of an interactive questionnaire, consisting of a number of questions, to which the participant must answer "yes" or "no". Depending on the answer to the question, the participant is directed to the next question, or is sent to complete one of the tasks featured in the quest. In other words, the quest is a binary tree, in the nodes of which are questions, and in the leaves of which are the tasks themselves.

It is known that a quest participant spends exactly one minute on any of the questions asked before receiving a task. Polycarp knows that his friends are busy people who cannot participate in the quest for longer than T minutes. Polycarp wants to choose some of the n tasks he has compiled, come up with a corresponding set of questions for them, and form an interactive questionnaire in the form of a binary tree in such a way that, regardless of the quest participant's answers, he will spend no more than T minutes on completing the entire quest (that is, on answering all the questions plus completing the task) . In particular, a quest may not contain any questions, but immediately consist of completing some task. Each task can be used only once (that is, people who give different answers to the questions must receive different tasks).

Polikarp wants the total interestingness of the tasks involved in the quest to be as high as possible. Help him determine the maximum possible total interestingness of the tasks, provided that the quest should be possible to complete in T minutes with any answer to the questions.

### ideas
1. 如果选定了一组task（必须要完成的）那么最大time + 它的路径长度 <= T
2. 所以，每个task的路径的最大距离是可以计算出来的
3. 假设一个节点需要处理的时间为x，它提供的价值为y
4. 那么处理时间越短的那些，那些等待时间越长的，越可以在底部
5. 这里 n < 1000, 所以似乎可以利用这点
6. dp[i][j] 表示把前i个任务，放在深度为j的树中时的最大值？
7. 似乎不大对