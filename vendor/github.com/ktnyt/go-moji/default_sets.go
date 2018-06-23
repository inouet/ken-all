package moji

// HE defines the Hankaku Eisuu (i.e. half width english text) Dictionary
var HE = NewRangeDictionary(0x0021, 0x007e)

// ZE defines the Zenkaku Eisuu (i.e. full width english text) Dictionary
var ZE = NewRangeDictionary(0xff01, 0xff5e)

// HG defines the HiraGana Dictionary
var HG = NewRangeDictionary(0x3041, 0x3096)

// KK defines the KataKana Dictionary
var KK = NewRangeDictionary(0x30a1, 0x30f6)

// HK defines the Hankaku Katakana (i.e. half width katakana) Dictionary
var HK = NewDictionary([]string{
	"ｶﾞ", "ｷﾞ", "ｸﾞ", "ｹﾞ", "ｺﾞ",
	"ｻﾞ", "ｼﾞ", "ｽﾞ", "ｾﾞ", "ｿﾞ",
	"ﾀﾞ", "ﾁﾞ", "ﾂﾞ", "ﾃﾞ", "ﾄﾞ",
	"ﾊﾞ", "ﾊﾟ", "ﾋﾞ", "ﾋﾟ", "ﾌﾞ", "ﾌﾟ", "ﾍﾞ", "ﾍﾟ", "ﾎﾞ", "ﾎﾟ",
	"ﾜﾞ", "ｦﾞ", "ｳﾞ",
	"｡", "｢", "｣", "､", "･", "ｰ", "ﾞ", "ﾟ",
	"ｱ", "ｲ", "ｳ", "ｴ", "ｵ",
	"ｶ", "ｷ", "ｸ", "ｹ", "ｺ",
	"ｻ", "ｼ", "ｽ", "ｾ", "ｿ", "ﾀ", "ﾁ", "ﾂ", "ﾃ", "ﾄ",
	"ﾅ", "ﾆ", "ﾇ", "ﾈ", "ﾉ",
	"ﾊ", "ﾋ", "ﾌ", "ﾍ", "ﾎ",
	"ﾏ", "ﾐ", "ﾑ", "ﾒ", "ﾓ",
	"ﾔ", "ﾕ", "ﾖ",
	"ﾗ", "ﾘ", "ﾙ", "ﾚ", "ﾛ",
	"ﾜ", "ｦ", "ﾝ",
	"ｧ", "ｨ", "ｩ", "ｪ", "ｫ",
	"ｬ", "ｭ", "ｮ", "ｯ",
})

// ZK defines the Zenkaku Katakana (i.e. full width katakana) Dictionary
var ZK = NewDictionary([]string{
	"ガ", "ギ", "グ", "ゲ", "ゴ",
	"ザ", "ジ", "ズ", "ゼ", "ゾ",
	"ダ", "ヂ", "ヅ", "デ", "ド",
	"バ", "パ", "ビ", "ピ", "ブ", "プ", "ベ", "ペ", "ボ", "ポ",
	"ヷ", "ヺ", "ヴ",
	"。", "「", "」", "、", "・", "ー", "゛", "゜",
	"ア", "イ", "ウ", "エ", "オ",
	"カ", "キ", "ク", "ケ", "コ",
	"サ", "シ", "ス", "セ", "ソ",
	"タ", "チ", "ツ", "テ", "ト",
	"ナ", "ニ", "ヌ", "ネ", "ノ",
	"ハ", "ヒ", "フ", "ヘ", "ホ",
	"マ", "ミ", "ム", "メ", "モ",
	"ヤ", "ユ", "ヨ",
	"ラ", "リ", "ル", "レ", "ロ",
	"ワ", "ヲ", "ン",
	"ァ", "ィ", "ゥ", "ェ", "ォ",
	"ャ", "ュ", "ョ", "ッ",
})

func stringify(r rune) string {
	return string([]rune{r})
}

// HS defines the Hankaku Space (i.e. half width space) Dictionary
var HS = NewDictionary([]string{" ", stringify(0x00a0)})

// ZS defines the Zenkaku Space (i.e. full width space) Dictionary
var ZS = NewDictionary([]string{stringify(0x3000), stringify(0x3000)})
