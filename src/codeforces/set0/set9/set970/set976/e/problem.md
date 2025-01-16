Recently Max has got himself into popular CCG "BrainStone". As "BrainStone" is a pretty intellectual game, Max has to solve numerous hard problems during the gameplay. Here is one of them:

Max owns n creatures, i-th of them can be described with two numbers — its health hpi and its damage dmgi. Max also has two types of spells in stock:

Doubles health of the creature (hpi := hpi·2);
Assigns value of health of the creature to its damage (dmgi := hpi).
Spell of first type can be used no more than a times in total, of the second type — no more than b times in total. Spell can be used on a certain creature multiple times. Spells can be used in arbitrary order. It isn't necessary to use all the spells.

Max is really busy preparing for his final exams, so he asks you to determine what is the maximal total damage of all creatures he can achieve if he uses spells in most optimal way.


### ideas
1. if b = 0, 那么就没有必要是要操作1（就是目前所有的damage之和）
2. b > 0, 肯定是先用操作1，（且只能用在一个上面）
3. 然后把那些 hp[i] - dmg[i] 最大的b个给选出来
4. 可以用avl tree