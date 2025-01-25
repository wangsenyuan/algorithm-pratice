Vasya’s elder brother Petya loves playing computer games. In one of his favourite computer games Petya reached the final level where a fight with the boss take place.

While playing the game Petya found spell scrolls and now he is about to use them. Let’s describe the way fighting goes on this level:

1) The boss has two parameters: max — the initial amount of health and reg — regeneration rate per second.

2) Every scroll also has two parameters: powi — spell power measured in percents — the maximal amount of health counted off the initial one, which allows to use the scroll (i.e. if the boss has more than powi percent of health the scroll cannot be used); and dmgi the damage per second inflicted upon the boss if the scroll is used. As soon as a scroll is used it disappears and another spell is cast upon the boss that inflicts dmgi of damage per second upon him until the end of the game.

During the battle the actions per second are performed in the following order: first the boss gets the damage from all the spells cast upon him, then he regenerates reg of health (at the same time he can’t have more than max of health), then the player may use another scroll (no more than one per second).

The boss is considered to be defeated if at the end of a second he has nonpositive ( ≤ 0) amount of health.

Help Petya to determine whether he can win with the set of scrolls available to him and if he can, determine the minimal number of seconds he needs to do it.

Input
The first line contains three integers N, max and reg (1 ≤ N, max, reg ≤ 1000) –– the amount of scrolls and the parameters of the boss. The next N lines contain two integers powi and dmgi each — the parameters of the i-th scroll (0 ≤ powi ≤ 100, 1 ≤ dmgi ≤ 2000).

Output
In case Petya can’t complete this level, output in the single line NO.

Otherwise, output on the first line YES. On the second line output the minimal time after which the boss can be defeated and the number of used scrolls. In the next lines for each used scroll output space-separated number of seconds passed from the start of the battle to the moment the scroll was used and the number of the scroll. Scrolls are numbered starting from 1 in the input order. The first scroll is considered to be available to be used after 0 seconds.

Output scrolls in the order they were used. It is not allowed to use scrolls after the boss is defeated.


### ideas
1. 如果没有pow = 100的，scroll，那么 => no
2. 如果在某个时刻，始终无法使的boss的血量下降到下一个能用的scroll => no
3. 按照pow降序，相同pow，按照dam降序,能用就用