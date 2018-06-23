package address

import (
	"regexp"
	"strings"
)

func NewRows(cols []string) []Row {
	row := NewRow(cols)
	rows := []Row{}

	if row.HasBrackets() {

		//
		// Town
		//

		// 阿三(799の1～867番地) を ( で 分割
		arrTown := strings.SplitN(row.Town, "(", 2)
		town := arrTown[0]

		// カッコ内
		townSub := strings.Replace(arrTown[1], ")", "", -1)
		townSub = strings.Replace(townSub, "(", "", -1)

		// 、で配列に分割
		townSubs := strings.Split(townSub, "、")

		//
		// TownKana
		//

		// ( で分割
		arrTownKana := strings.SplitN(row.TownKana, "(", 2)
		arrTownKana = fixArray(arrTownKana, 2)
		townKana := arrTownKana[0]

		// カナカッコ内
		townKanaSub := strings.Replace(arrTownKana[1], ")", "", -1)
		townKanaSub = strings.Replace(townKanaSub, "(", "", -1)

		townKanaSubs := strings.Split(townKanaSub, "、")
		townKanaSubs = fixArray(townKanaSubs, len(townSubs))

		if !IsBuilding(row.Town) {
			// () 展開前のものを追加
			rowCopy := row
			rowCopy.RawTown = row.Town
			rowCopy.Town = town
			rowCopy.TownKana = townKana
			rows = append(rows, rowCopy)
		}

		// () 内を追加
		for i, sub := range townSubs {
			sub = strings.Trim(sub, " ")

			if isNgString(sub) {
				continue
			}

			if town == sub {
				// 7711231: 富吉（富吉）、 6560514: 賀集（賀集）
				continue
			}

			rowCopy := row

			// Town
			rowCopy.Town = town + sub
			rowCopy.RawTown = row.Town

			// TownKana
			subKana := townKanaSubs[i]
			rowCopy.TownKana = townKana + subKana

			rows = append(rows, rowCopy)
		}

	} else {
		rows = append(rows, row)
	}
	return rows
}

// 指定した長さまで配列を埋める
func fixArray(arr []string, count int) []string {
	size := len(arr)
	if size == count {
		return arr
	}
	for i := size; i < count; i++ {
		arr = append(arr, "")
	}
	return arr
}

var (
	regexpNgStringNumber1 = regexp.MustCompile(`^[0-9の・\-]+$`)
	regexpNgStringNumber2 = regexp.MustCompile(`^第[0-9]+$`)
	regexpNgStringNumber3 = regexp.MustCompile(`^[0-9 ]+丁目[0-9 ]+番?`)
	regexpNgStringNumber4 = regexp.MustCompile(`[0-9]+[\-−〜]+[0-9]+`)
)

func isNgString(s string) bool {
	// 〜 を含んでいたらNG
	if strings.Contains(s, "～") {
		return true
	}

	// ・ を含んでいたらNG
	if strings.Contains(s, "・") {
		return true
	}

	// 数字のみはスキップ
	if regexpNgStringNumber1.Match([]byte(s)) {
		return true
	}

	// 第3
	if regexpNgStringNumber2.Match([]byte(s)) {
		return true
	}

	// 1丁目1番
	if regexpNgStringNumber3.Match([]byte(s)) {
		return true
	}

	// 0482402: 13−4
	if regexpNgStringNumber4.Match([]byte(s)) {
		return true
	}

	if isNgWord(s) {
		return true
	}

	// 特定のサフィックスを持つものはスキップ
	if hasNgSuffix(s) {
		return true
	}
	return false
}

var (
	ngWordList = []string{
		"その他",
		"地階・階層不明",
		"全域",      // 0895865: 厚内（全域）
		"成田国際空港内", // 2820031: 一鍬田（成田国際空港内）
	}
)

func isNgWord(s string) bool {
	for _, ng := range ngWordList {
		if s == ng {
			return true
		}
	}
	return false
}

var (
	ngSuffixList = []string{"以上", "以下", "以降", "以内", "番地", "除く", "丁目", "含む", "その他", "以外", "「その他」"}
)

func hasNgSuffix(s string) bool {
	for _, ng := range ngSuffixList {
		if strings.HasSuffix(s, ng) {
			return true
		}
	}
	return false
}
