You are playing a game where your character should overcome different obstacles. The current problem is to come down from a cliff. The cliff has height ℎ
, and there is a moving platform on each height 𝑥
 from 1
 to ℎ
.

Each platform is either hidden inside the cliff or moved out. At first, there are 𝑛
 moved out platforms on heights 𝑝1,𝑝2,…,𝑝𝑛
. The platform on height ℎ
 is moved out (and the character is initially standing there).

If you character is standing on some moved out platform on height 𝑥
, then he can pull a special lever, which switches the state of two platforms: on height 𝑥
 and 𝑥−1
. In other words, the platform you are currently standing on will hide in the cliff and the platform one unit below will change it state: it will hide if it was moved out or move out if it was hidden. In the second case, you will safely land on it. Note that this is the only way to move from one platform to another.

Your character is quite fragile, so it can safely fall from the height no more than 2
. In other words falling from the platform 𝑥
 to platform 𝑥−2
 is okay, but falling from 𝑥
 to 𝑥−3
 (or lower) is certain death.

Sometimes it's not possible to come down from the cliff, but you can always buy (for donate currency) several magic crystals. Each magic crystal can be used to change the state of any single platform (except platform on height ℎ
, which is unaffected by the crystals). After being used, the crystal disappears.

### ideas
1. dp[i] = 安全到达第i个平台时，需要花费的最少的宝石
2. dp[i] = 如果 p[i-1] - 1 = p[i], 那么必须花费一个宝石，先让 p[i]隐藏起来，才能使用lever
3.     => dp[i] = dp[i-1] + 1
4.     如果 p[i-2] - 2 = p[i], 那么可以直接在 p[i-2]的地方使用lever, 从那里掉到 p[i], 
5.     =》 dp[i] = dp[i-2]
6.    如果 p[i-1] - p[i] > 2, 那么， 从 p[-1], 可以一直安全的到到 p[i] - 1 的位置
7.    此时，必须花费一个额外的宝石，让 p[i]隐藏起来，或者 p[i] - 2 的位置, 让 p[i] - 1显示出来
8.    => dp[i] = dp[i-1] + 1