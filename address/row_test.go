package address

import (
	"bytes"
	"testing"
)

func TestDeleteKeyBracketKana(t *testing.T) {

	type Case struct {
		input    string
		expected string
	}

	cases := []Case{
		{
			input:    "シモクボ<174ヲノゾク>",
			expected: "シモクボ",
		},
		{
			input:    "モタイ(1-500<211バンチヲノゾク><フルマチ>、2527-2529<ドトオ>)",
			expected: "モタイ(1-500、2527-2529)",
		},
	}

	for i, row := range cases {
		res := deleteKeyBracketKana(row.input)
		if res != row.expected {
			t.Errorf("#%d: want '%s', got '%s'\n", i, row.expected, res)
		}
	}
}

func TestDeleteKeyBracket(t *testing.T) {
	type Case struct {
		input    string
		expected string
	}

	cases := []Case{
		{
			input:    "葛巻(第40地割「57番地125、176を除く」～第45地割)",
			expected: "葛巻(第40地割～第45地割)",
		},
		{
			input:    "茂田井(1～500「211番地を除く」「古町」、2527～2529「土遠」)",
			expected: "茂田井(1～500、2527～2529)",
		},
	}

	for i, row := range cases {
		res := deleteKeyBracket(row.input)
		if res != row.expected {
			t.Errorf("#%d: want '%s', got '%s'\n", i, row.expected, res)
		}
	}
}

func TestNewRow(t *testing.T) {
	type Case struct {
		name     string
		data     string
		town     string
		townKana string
	}

	cases := []Case{
		// 複数行にまたがる行の連結 ex) 6028062
		{
			name: "6028062",
			data: `26102,"602  ","6028062","ｷｮｳﾄﾌ","ｷｮｳﾄｼｶﾐｷﾞｮｳｸ","ｶﾒﾔﾁｮｳ","京都府","京都市上京区","亀屋町（油小路通上長者町下る、油小路通下長者町上る、油小路通",0,0,0,0,0,0
26102,"602  ","6028062","ｷｮｳﾄﾌ","ｷｮｳﾄｼｶﾐｷﾞｮｳｸ","ｶﾒﾔﾁｮｳ","京都府","京都市上京区","中長者町上る、油小路通中長者町下る、上長者町通油小路西入、上長者町通油小",0,0,0,0,0,0
26102,"602  ","6028062","ｷｮｳﾄﾌ","ｷｮｳﾄｼｶﾐｷﾞｮｳｸ","ｶﾒﾔﾁｮｳ","京都府","京都市上京区","路東入）",0,0,0,0,0,0`,
			town:     "亀屋町(油小路通上長者町下る、油小路通下長者町上る、油小路通中長者町上る、油小路通中長者町下る、上長者町通油小路西入、上長者町通油小路東入)",
			townKana: "カメヤチョウ",
		},

		// 以下に掲載がない場合が消えていること ex) 6000000
		{
			name:     "6000000",
			data:     `26106,"600  ","6000000","ｷｮｳﾄﾌ","ｷｮｳﾄｼｼﾓｷﾞｮｳｸ","ｲｶﾆｹｲｻｲｶﾞﾅｲﾊﾞｱｲ","京都府","京都市下京区","以下に掲載がない場合",0,0,0,0,0,0`,
			town:     "",
			townKana: "",
		},

		// xx一円が消えていること ex) 100-0301
		{
			name:     "1000301",
			data:     `13362,"10003","1000301","ﾄｳｷｮｳﾄ","ﾄｼﾏﾑﾗ","ﾄｼﾏﾑﾗｲﾁｴﾝ","東京都","利島村","利島村一円",0,0,0,0,0,0`,
			town:     "",
			townKana: "",
		},

		// 「一円」は残す ex) 522-0317
		{
			name:     "5220317",
			data:     `25443,"52203","5220317","ｼｶﾞｹﾝ","ｲﾇｶﾐｸﾞﾝﾀｶﾞﾁｮｳ","ｲﾁｴﾝ","滋賀県","犬上郡多賀町","一円",0,0,0,0,0,0`,
			town:     "一円",
			townKana: "イチエン",
		},

		// 地割
		{
			name:     "0287915",
			data:     `03507,"02879","0287915","ｲﾜﾃｹﾝ","ｸﾉﾍｸﾞﾝﾋﾛﾉﾁｮｳ","ﾀﾈｲﾁﾀﾞｲ15ﾁﾜﾘ-ﾀﾞｲ21ﾁﾜﾘ(ｶﾇｶ､ｼｮｳｼﾞｱｲ､ﾐﾄﾞﾘﾁｮｳ､ｵｵｸﾎﾞ､ﾀｶﾄﾘ)","岩手県","九戸郡洋野町","種市第１５地割〜第２１地割（鹿糠、小路合、緑町、大久保、高取）",0,1,0,0,0,0`,
			town:     "種市",
			townKana: "タネイチ",
		},
		{
			name: "0285102",
			data: `03302,"02851","0285102","ｲﾜﾃｹﾝ","ｲﾜﾃｸﾞﾝｸｽﾞﾏｷﾏﾁ","ｸｽﾞﾏｷ(ﾀﾞｲ40ﾁﾜﾘ<57ﾊﾞﾝﾁ125､176ｦﾉｿﾞｸ>-ﾀﾞｲ45","岩手県","岩手郡葛巻町","葛巻（第４０地割「５７番地１２５、１７６を除く」〜第４５",1,1,0,0,0,0
03302,"02851","0285102","ｲﾜﾃｹﾝ","ｲﾜﾃｸﾞﾝｸｽﾞﾏｷﾏﾁ","ﾁﾜﾘ)","岩手県","岩手郡葛巻町","地割）",1,1,0,0,0,0`,
			town:     "葛巻",
			townKana: "クズマキ",
		},
	}

	for i, c := range cases {

		t.Run(c.name, func(t *testing.T) {

			reader := bytes.NewReader([]byte(c.data))
			r := NewReader(reader)
			cols, _ := r.Read()

			row := NewRow(cols)

			if row.Town != c.town {
				t.Errorf("#%d: want '%s', got '%s'\n", i, c.town, row.Town)
			}
			if row.TownKana != c.townKana {
				t.Errorf("#%d: want '%s', got '%s'\n", i, c.townKana, row.TownKana)
			}
		})
	}
}

func TestIsBuilding(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "大崎ThinkParkTower(18階)",
			expected: true,
		},
		{
			input:    "平和通(南)",
			expected: false,
		},
		{
			input:    "中央アエル(地階・階層不明)",
			expected: true,
		},
	}

	for _, c := range cases {
		actual := IsBuilding(c.input)
		if actual != c.expected {
			t.Errorf("#%s: want '%v', got '%v'\n", c.input, c.expected, actual)
		}
	}
}
