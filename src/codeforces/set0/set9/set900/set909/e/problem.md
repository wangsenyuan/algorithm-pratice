You are given a program you want to execute as a set of tasks organized in a dependency graph. The dependency graph is a
directed acyclic graph: each task can depend on results of one or several other tasks, and there are no directed
circular dependencies between tasks. A task can only be executed if all tasks it depends on have already completed.

Some of the tasks in the graph can only be executed on a coprocessor, and the rest can only be executed on the main
processor. In one coprocessor call you can send it a set of tasks which can only be executed on it. For each task of the
set, all tasks on which it depends must be either already completed or be included in the set. The main processor starts
the program execution and gets the results of tasks executed on the coprocessor automatically.

Find the minimal number of coprocessor calls which are necessary to execute the given program.

### ideas

1. bfs
2. 应该尽量的将coprocessor task放在一组内执行
3. 假设都是可以执行的，先执行main任务