package address

import (
	"regexp"
	"strings"

	"github.com/inouet/ken-all/util"
)

type Row struct {
	RegionID       string `json:"region_id"`     //  0. 全国地方公共団体コード
	Zip5           string `json:"-"`             //  1. 郵便番号（5桁）
	Zip7           string `json:"zip"`           //  2. 郵便番号（7桁）
	PrefKana       string `json:"pref_kana"`     //  3. 都道府県名(カナ)
	CityKana       string `json:"city_kana"`     //  4. 市区町村名(カナ)
	TownKana       string `json:"town_kana"`     //  5. 町域名(カナ)
	Pref           string `json:"pref"`          //  6. 都道府県名
	City           string `json:"city"`          //  7. 市区町村名
	Town           string `json:"town"`          //  8. 町域名
	IsMultiZip     string `json:"-"`             //  9. 一町域が二以上の郵便番号で表される場合の表示（1: 該当、0: 該当せず）
	HasKoazaBanchi string `json:"-"`             // 10. 小字毎に番地が起番されている町域の表示　（1: 該当、0: 該当せず）
	HasChome       string `json:"-"`             // 11. 丁目を有する町域の場合の表示　（1: 該当、0: 該当せず）
	IsMultiTown    string `json:"-"`             // 12. 一つの郵便番号で二以上の町域を表す場合の表示　（1: 該当、0:該当せず）
	UpdateStatus   string `json:"update_status"` // 13. 更新の表示 (0: 変更なし、1:変更あり、2: 廃止)
	UpdateReason   string `json:"update_reason"`
	// 14. 変更理由(0: 変更なし、1: 市政・区政・町政・分区・政令指定都市施行、2:住居表示の実施、3:区画整理、4:郵便区調整等、5:訂正、6:廃止)
	RawTown  string `json:"-"`
	PrefCode string `json:"pref_code"` // xx. 都道府県コード
}

func NewRow(cols []string) Row {
	row := Row{
		RegionID:       cols[0],
		Zip5:           cols[1],
		Zip7:           cols[2],
		PrefKana:       cols[3],
		CityKana:       cols[4],
		TownKana:       cols[5],
		Pref:           cols[6],
		City:           cols[7],
		Town:           cols[8],
		IsMultiZip:     cols[9],
		HasKoazaBanchi: cols[10],
		HasChome:       cols[11],
		IsMultiTown:    cols[12],
		UpdateStatus:   cols[13],
		UpdateReason:   cols[14],
	}

	row.fixTown()
	row.setPrefCode()

	return row
}

func (row Row) Array() []string {
	cols := []string{
		row.RegionID,
		row.Zip5,
		row.Zip7,
		row.PrefKana,
		row.CityKana,
		row.TownKana,
		row.Pref,
		row.City,
		row.Town,
		row.IsMultiZip,
		row.HasKoazaBanchi,
		row.HasChome,
		row.IsMultiTown,
		row.UpdateStatus,
		row.UpdateReason,
		row.PrefCode,
	}
	return cols
}

func (row *Row) setPrefCode() {
	row.PrefCode = util.GetPrefCode(row.Pref)
}

func (row *Row) fixTown() {
	row.patch()

	// 以下に掲載がない場合
	if row.Town == "以下に掲載がない場合" {
		row.Town = ""
		row.TownKana = ""
	}

	// xx一円 、xxの次に番地がくる場合
	if strings.HasSuffix(row.Town, "町一円") || strings.HasSuffix(row.Town, "村一円") || strings.HasSuffix(row.Town, "の次に番地がくる場合") {
		row.Town = ""
		row.TownKana = ""
	}

	// 「」内消す
	row.Town = deleteKeyBracket(row.Town)
	row.TownKana = deleteKeyBracketKana(row.TownKana)

	// 地割
	row.fixTownChiwari()

	// trim
	row.Town = strings.Trim(row.Town, " ")
	row.TownKana = strings.Trim(row.TownKana, " ")
}

// 元データの不具合を補正する
func (row *Row) patch() {
	// 6511102  」の後に、が足りないので修正する
	if row.Zip7 == "6511102" &&
		row.Town == "山田町下谷上(大上谷、修法ケ原、中一里山「9番地の4、12番地を除く」長尾山、再度公園)" {
		row.Town = "山田町下谷上(大上谷、修法ケ原、中一里山「9番地の4、12番地を除く」、長尾山、再度公園)"
		row.TownKana = "ヤマダチョウシモタニガミ(オオカミダニ、シュウホウガハラ、" +
			"ナカイチリヤマ<9バンチノ4、12バンチヲノゾク>、ナガオヤマ、フタタビコウエン)"
	}
}

// 岩手県の第n地割　もしくは、 n地割 以降は削除する
var (
	regexpChiwari1 = regexp.MustCompile(`(\()?(第)?[0-9]+地割.*`)
	regexpChiwari2 = regexp.MustCompile(`(\()?(ダイ)?[0-9]+チワリ.*`)
)

func (row *Row) fixTownChiwari() {
	if row.Pref == "岩手県" && strings.Contains(row.Town, "地割") {
		row.Town = regexpChiwari1.ReplaceAllString(row.Town, "")
		// カナ
		row.TownKana = regexpChiwari2.ReplaceAllString(row.TownKana, "")
	}
}

func (row *Row) HasBrackets() bool {
	return strings.Contains(row.Town, "(")
}

// 「」内を消す
var regexpKeyBracket = regexp.MustCompile(`「([^「」]+)」`)

func deleteKeyBracket(str string) string {
	str = regexpKeyBracket.ReplaceAllString(str, "")
	return str
}

// カナの <>内を消す
var regexpKeyBracketKana = regexp.MustCompile(`<([^<>]+)>`)

func deleteKeyBracketKana(str string) string {
	str = regexpKeyBracketKana.ReplaceAllString(str, "")
	return str
}

var regexpIsBuilding = regexp.MustCompile(`\((.+)階(.*)\)`)

// IsBuilding ビルかどうか
func IsBuilding(s string) bool {
	return regexpIsBuilding.Match([]byte(s))
}
