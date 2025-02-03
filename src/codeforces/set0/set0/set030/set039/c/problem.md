There are lots of theories concerning the origin of moon craters. Most scientists stick to the meteorite theory, which says that the craters were formed as a result of celestial bodies colliding with the Moon. The other version is that the craters were parts of volcanoes.

An extraterrestrial intelligence research specialist professor Okulov (the namesake of the Okulov, the author of famous textbooks on programming) put forward an alternate hypothesis. Guess what kind of a hypothesis it was –– sure, the one including extraterrestrial mind involvement. Now the professor is looking for proofs of his hypothesis.

Professor has data from the moon robot that moves linearly in one direction along the Moon surface. The moon craters are circular in form with integer-valued radii. The moon robot records only the craters whose centers lay on his path and sends to the Earth the information on the distance from the centers of the craters to the initial point of its path and on the radii of the craters.

According to the theory of professor Okulov two craters made by an extraterrestrial intelligence for the aims yet unknown either are fully enclosed one in the other or do not intersect at all. Internal or external tangency is acceptable. However the experimental data from the moon robot do not confirm this theory! Nevertheless, professor Okulov is hopeful. He perfectly understands that to create any logical theory one has to ignore some data that are wrong due to faulty measuring (or skillful disguise by the extraterrestrial intelligence that will be sooner or later found by professor Okulov!) That’s why Okulov wants to choose among the available crater descriptions the largest set that would satisfy his theory.

### ideas
1. 根据c进行排序
2. 如果c[i]和c[j]相容， dp[i] = dp[j] + 1 (max)
3. 好像不大行，无法处理包含的情况
4. 如果i足够大，感觉它行程了一个独立的区段
5. 它内部可以独立计算
6. dp[i] = 最后一个是i时的最大相容区间 dp[i] = max(dp[j] + ？) 如果 j在i的前面（c[j] +  d[j] <= c[i]- d[i])
7. 增加的部分 = 1 + 只考虑i内部的区间的一个子问题
8. fp[i] = 在i的内部的子问题的解