package p2157

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, words []string, expect []int) {
	res := groupStrings(words)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"a", "b", "ab", "cde"}
	expect := []int{2, 3}
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"a", "ab", "abc"}
	expect := []int{1, 3}
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{"qamp", "am", "khdrn"}
	expect := []int{3, 1}
	runSample(t, words, expect)
}

func TestSample4(t *testing.T) {
	words := []string{"xo", "t", "uhc", "gf"}
	expect := []int{4, 1}
	runSample(t, words, expect)
}

func TestSample5(t *testing.T) {
	words := []string{"xsgrqznel", "zy", "yk", "qdpoijn", "sgcevyinq", "m", "gfencukmaiohbw", "vf", "xucfho", "ewqrtk", "kvtbcnzwyqusper", "qkbmo", "df", "ic", "jlbfxu", "hcdkxspqu", "bthluwdy", "epzrg", "vfyuxlazsmqkoj", "pmuyiqnljdgr", "dqcgrsuhnw", "eujcordq", "neirphybmdj", "emhbnkvd", "o", "f", "cjfwdbqlna", "hxofjbmplrkvws", "qfjlnkouxea", "tiwcguvqnk", "rnwjtm", "qbpyuhtalwi", "qs", "vawlbkp", "mvhfgbxwpy", "wxqnpybocldmghv", "cz", "hdtk", "h", "snbkyrxvocl", "oqrhbxjwslkczat", "tfbzesgrmxjhpo", "ebflnptqjuymaz", "wuapyhkfidtbj", "dimjq", "xtjnmyfie", "rg", "qstjxw", "kvzjobspxaq", "cfusjbmdxyn", "pnmh", "nqkvm", "vd", "qlmvhaobxscj", "tbsacwrqjvmiez", "b", "terzxvjbsnlahi", "fyksj", "cn", "qi", "qbilgophxz", "lroeqf", "wzorubxqaimydf", "qrxfstophynlb", "vcpjhg", "tcdljo", "bjftgakuicvers", "pjckyvgxwd", "chuqwfjvzdkmny", "ltvwmbo", "dehxajgkruoc", "prhcdkubmfvi", "lsca", "btnhaxyzg", "rbeyzqo", "lhcd", "rjqlkwycous", "sidbtalx", "xsdlqfbemzjg", "g", "j", "er", "xr", "mtebgwoundsz", "desrzflgb", "rdacmnbwktzh", "dhtsipjenl", "h", "cajfihzxvbm", "esjnrf", "zwldiuomcqh", "bjzwnpvh", "tywrh", "olchpugdxwktqs", "xsoydzt", "jdeknm", "avgoyne", "jbchopakq", "xo", "tuenhryzvl", "sxhmblpkud", "xmtezwqsglkuny", "iqwgpxcrm", "ngtjxldckfib", "hyj", "j", "efygxmqzo", "kbhvructy", "n", "nsdtguzr", "dpomkierl", "pz", "fybsnv", "kbs", "amzflkrqpustob", "qjinbtxyhvf", "s", "lowpevqai", "ng", "gbq", "qa", "fakeqlspvm", "agoihubckqpe", "qitbu", "xmud", "vxbemrfnpsu", "bthd", "icvszbrkqowg", "gsy", "qlbtnydakcomj", "yehr", "isdevbtfj", "uybrepaz", "czwxnriekq", "xfh", "oh", "mv", "cbdvrwegtzos", "oljuhzabgn", "cyjukdaxnfrv", "eigopwxyarnjd", "vaioetdhyxjmq", "ozfrah", "lwmcjpsub", "vlxdimprea", "wc", "hqjlfzyga", "vgjnmbu", "uskrlmaovg", "hodurfpcnljbe", "erantguib", "bsyqwpe", "ceoszfb", "jioraw", "isoetbavyjnpwqg", "uydjgm", "bycklfdtwonqjsm", "yvxtr", "dvmcotwkl", "wz", "wftjrymx", "vb", "sr", "uozc", "aobqdn", "pcyzrhna", "nxghatvesplb", "hon", "wxghdmczsfentaj", "nh", "telcwryxgpmbaki", "ejdairchlfubgnz", "xkbteuhszwn", "twkdiaqbcmp", "rkzgquehfv", "gawtp", "jupzlmotahd", "bpcamwtsyzkgjr", "rqhxj", "lzei", "vceyjwoautsfn", "zbqryoctxhw", "zr", "tzfkiuvmrchex", "dfpxny", "urjyisnd", "pkqxlchoifwut", "e", "zorwincxmhvabks", "gheyvwl", "dv", "bmneaupgfwyvk", "rtbgnfmlpwezvdh", "vobwsh", "lm", "ihzre", "lcqx", "hltjbaofnirpkqe", "mnwguxjrok", "vztxlmgr", "pjcaudey", "zu", "dumsh", "vyawij", "r", "kasdqjow", "fygbxcqurhtpmzi", "wqpchufjimrk", "knotgfrdpvzi", "xsckbilvngpr", "wvuxmylkrcp", "cphxrywvqb", "grol", "sgx", "t", "ehsglmtwou", "vhbgzqopxjs", "sfu", "nl", "enbduyahkgr", "fdiwqrkcxepobt", "sorycnhgw", "tfzupnsgyadchj", "ybvtfoljhcrqzxa", "iek", "xoepzblswa", "h", "xpvglfir", "lrkavotxwqbsfp", "yelwtdfirxc", "yqmfwvcuthbjrs", "jtwydkclhpr", "uhjzcw", "sq", "cpdiutomye", "obauxqkptmv", "rqxcbnlvgy", "kx", "pimdjnvqawelfht", "okxchfd", "tzcfpgjnoydrqk", "wfbchxtoaslyge", "bidoznvgre", "qa", "vo", "z", "fixbshrjgwlukc", "kys", "rfnxdaqwemoku", "vlcbnzkjdu", "pxcnkthjuo", "jvneowk", "dbjuyna", "xzms", "onimdjafpywthcu", "gtlq", "b", "timldacy", "rqibohpfn", "sbrmncopl", "teulxjwocnpfmd", "bgesqofc", "rgihaeuvbtqwxsp", "ytfqvkgdpalcwos", "hzswoxudjynti", "ysqfkbcoaevdt", "kgylrwixvhpufj", "sdljhzqrew", "oetgkxcuqszlm", "jwqo", "rpxgdyltscem", "gbps", "bzkjlfhedqow", "oqructwf", "xflmqutnakvs", "ifltcujwgdnq", "qrjoy", "zgbkrnayiqm", "vmdflxhbkrjzcpa", "brdfjz", "ybzl", "kafrzvgi", "ukgievpfwysdxht", "g", "xmd", "gob", "rbugxvj", "cmujo", "wipoldcryjtguv", "lwr", "lpso", "z", "azlqnmuwtgdeypv", "dtrjogxiehav", "swucyq", "wfzpxvr", "rlg", "m", "cqsteouydag", "uedblvogsnzqjw", "lbjhgkdmpafenvr", "uc", "nfpbjqrmakx", "z", "pizhqxcktl", "ptz", "jwtm", "kwcgsmvju", "bu", "kpgnts", "myedtiv", "czhqeksgmiwavpx", "balfes", "qtvlaywzfm", "wrkjfvtyuxe", "nlmbjshuey", "ilwdfkz", "shyrtdbfp", "qonzei", "gmlhtxy", "wqn", "ozubsndlcaq", "kolh", "ec", "dpkwharimlt", "bcslveydjazfm", "u", "eyupasolhm", "zyhgua", "okesmyliz", "kloty", "jvzdhqbntaes", "ajsn", "e", "yxjpqhgelictak", "gphsxyaobrfi", "aqxoel", "ob", "igvn", "mewhgxkrzjsacoq", "cznrta", "hlpxs", "g", "dwujic", "kisobxwedva", "cmsbyjhfzr", "jsi", "nifpuadmglxyj", "zqtiopvydkl", "ykazbpoqdt", "ehra", "rsh", "qmxnjve", "mjdnevub", "widltpchvsqj", "rdleyfiwm", "pnojberkgtzaqsx", "dcfeytu", "oshpnqfjgd", "nqpzlhktwcfreyx", "ea", "iscojwk", "kaqbmpscliuznyo", "ik", "szgbiqldm", "zokbpqeanjigmd", "xwynvidpauqeb", "kcgdn", "bogecpwkrhs", "ebjyzocdmf", "aplxv", "skydjpwcqobrgt", "tyohmeviqzbsa", "hidfuyxwzvrol", "qr", "urqsvpihdezgyj", "xgtapi", "ex", "fnqzdvwlparshc", "aonmk", "mjxsvf", "nvbtshyludcfikq", "mkgfjeypxuwqbaz", "ykh", "zdyevmgw", "uoihtq", "nxluw", "y", "ckretvjihg", "skxqrdzfhnji", "bctqnfakhiwy", "fbcidznyrghwko", "tehngdcxskwif", "uhbkteoz", "lajnwedyirfuxh", "qfskpxmlcvg", "ny", "zdsvolmtcignqey", "borjtufaz", "el", "uegqcapdjkw", "exfatpqdhyrw", "qu", "dwolzpcu", "pxzmj", "gbjosynv", "atg", "hcurgaeisnb", "vnqtw", "fteypzhogku", "mkbsndy", "cusx", "zrhkxawesp", "uhj", "wbekzotrnsfxl", "xi", "tuxh", "atkpfwb", "j", "to", "npwsuzkbr", "afozpnkjb", "lhdrfkmtzjives", "g", "l", "qrnp", "xokqi", "swqrtmofkcy", "pltqrayvhsmnzb", "rnusqme", "pg", "u", "vcuoksf", "potmc", "z", "z", "yikwclu", "fgeqamscvxjy", "sozmctlkbrn", "xudkc", "vfclegih", "kwljfa", "pcaiwbvenhtsoxq", "rnsuvydizc", "crqaudvtiz", "ztn", "vnorpcd", "xqntzjosimpkdue", "feyn", "dwblmj", "sqvepk", "owblcikqyt", "h", "yfkpsivtqjrax", "vfsardnejxqgmbo", "vhqbfi", "ptqvlhfwajrcmdz", "ulksnmjrxf", "nmowlrp", "tidglc", "dzintjvymbxphl", "c", "xkwpdaqsot", "jzlcbghsxvdpry", "uips", "aliykzrgvom", "vbcjezrlktfihd", "srwc", "csieramlthfbw", "mbw", "hbcal", "nrdam", "o", "asxuvch", "vntxmya", "oikcesbvtg", "t", "pkgdwusoaml", "ileohznaykqu", "pan", "hmgxb", "m", "roslevxhndc", "l", "viegpuat", "sdeqjthrcuxyg", "bfukpgneawl", "ogyrpdlekjhb", "objdp", "mqtrezahf", "rcnytu", "etlgchqjxour", "wlfbeaiznjstqy", "vh", "rwebumvxypdigh", "gouvcknwsfb", "lkwzycsbfmpt", "qyxmczsuvnlgfd", "tumgcfqiolyew", "xd", "v", "cxirbmwvfueg", "ubsitqeyrmvoklg", "okfgiuyhqdwlxb", "qao", "vqcpbjnso", "zilesa", "ljficq", "obn", "q", "ukxwa", "uxwvlm", "dpsbmnwxfe", "htjlpmuicn", "fo", "vtzpqclhukgfnj", "kd", "kzfjp", "jymf", "lucxwf", "yh", "dgmpswviyfxubtr", "tvprzh", "lbxpwumzskgoie", "gofriqpkzcjlnhe", "gfnitulkdv", "wnqxszuh", "kpzaqhoi", "efksj", "ajqx", "fvqlx", "dfxitr", "hyowcrme", "hjxzvrycapwke", "jeqasvhlk", "tol", "xf", "ufaicyd", "njmfblorexszhci", "jyxrdb", "yst", "mfpotuqr", "jixkncpztm", "pjqenbmaoudyxcr", "iwncvf", "hornilmegvfcqtb", "u", "xlmpsqar", "rspwijhzkqucme", "kgijfoq", "hvncxaer", "pqh", "qsnmzfkv", "puyov", "vlqo", "bues", "olpdhuvrs", "khjt", "wxsapzudmgcjny", "kzbilygxpfrmsu", "cbuos", "yvshwfrnqckt", "wvyia", "tpnauwblx", "zpavuxkjimdgylh", "m", "apkzrvuhne", "x", "sviz", "byrncofv", "wmtuchz", "vwjrbfxln", "aztuen", "qia", "qoazsrbkpnyvw", "uz", "sqflvmnrcdjo", "fvhtus", "cyzkrdpmafeut", "il", "xockn", "gw", "atxr", "tj", "bqphtlyf", "syvfxmw", "us", "nkrhz", "ujzyvps", "tnibdfskjgprw", "gvo", "fkj", "xvqodkazisbepu", "fu", "yimstcknrqowdvx", "iagenfrwdxvqozm", "t", "ecfbyxdl", "npsqrkdbew", "byfzlter", "qysvfwokgrbex", "ogfdranskxtlbmu", "t", "doqn", "uvfyqtlioperczs", "nu", "ahvdlixubr", "wetjfhk", "cbomgetrz", "zuqle", "qmu", "lmesdnizva", "eubjsyoipghv", "ewxbvzafml", "nxqyoabzvflsmr", "ynicvzd", "xoefsigclabnkp", "rd", "wskmcipz", "esyntl", "jmtprywxulgab", "jobegzqnrlau", "helq", "nzrwjexgtmi", "nqshpzau", "vtharwbiszu", "qopgdrhlmwuejs", "amgwvqiul", "rxmpyflsv", "bzxvryulj", "dgboxlfukcaw", "eb", "inajdykeqxrzolg", "umwxeiar", "xbwfo", "htmrsgxwzdocjk", "clxhenyiz", "mahwipfqcn", "pbsqtlfcmei", "c", "nlvkruxbsi", "stlpmvqeyg", "cbjyxtafsn", "oarzin", "bfiovyjp", "pcwgtoyn", "rydcg", "ghomx", "fishoclwmdxt", "yexuvlzd"}
	expect := []int{510, 135}
	runSample(t, words, expect)
}
