package office

import "github.com/inouet/ken-all/util"

type Row struct {
	JisCode    string `json:"jis_code"`    //  0. 大口事業所の所在地のJISコード（5バイト）
	OfficeKana string `json:"kana"`        //  1. 大口事業所名（カナ）（100バイト）
	OfficeName string `json:"name"`        //  2. 大口事業所名（漢字）（160バイト）
	Pref       string `json:"pref"`        //  3. 都道府県名（漢字）（8バイト）
	City       string `json:"city"`        //  4. 市区町村名（漢字）（24バイト）
	Town       string `json:"town"`        //  5. 町域名（漢字）（24バイト）
	Address    string `json:"address"`     //  6. 小字名、丁目、番地等（漢字）（124バイト）
	Zip7       string `json:"zip7"`        //  7. 大口事業所個別番号（7バイト）
	Zip5       string `json:"zip5"`        //  8. 旧郵便番号（5バイト）
	PostOffice string `json:"post_office"` //  9. 取扱局（漢字）（40バイト）
	Type       string `json:"type"`        // 10. 個別番号の種別の表示（1バイト）「0」大口事業所 「1」私書箱
	// 11. 複数番号の有無（1バイト）
	//	「0」複数番号無し
	//	「1」複数番号を設定している場合の個別番号の1
	//	「2」複数番号を設定している場合の個別番号の2
	//	「3」複数番号を設定している場合の個別番号の3
	// 	一つの事業所が同一種別の個別番号を複数持つ場合に複数番号を設定しているものとします。
	// 	従って、一つの事業所で大口事業所、私書箱の個別番号をそれぞれ一つづつ設定している場合は 12）は「0」となります。
	IsMulti string `json:"is_multi"`

	UpdateStatus string `json:"update_status"` // 12. 修正コード（1バイト） 「0」修正なし/「1」新規追加/「5」廃止
	PrefCode     string `json:"pref_code"`     // xx. 都道府県コード
}

func NewRow(cols []string) Row {
	row := Row{
		JisCode:      cols[0],
		OfficeKana:   cols[1],
		OfficeName:   cols[2],
		Pref:         cols[3],
		City:         cols[4],
		Town:         cols[5],
		Address:      cols[6],
		Zip7:         cols[7],
		Zip5:         cols[8],
		PostOffice:   cols[9],
		Type:         cols[10],
		IsMulti:      cols[11],
		UpdateStatus: cols[12],
	}

	row.setPrefCode()

	return row
}

func (row Row) Array() []string {
	cols := []string{
		row.JisCode,
		row.OfficeKana,
		row.OfficeName,
		row.Pref,
		row.City,
		row.Town,
		row.Address,
		row.Zip7,
		row.Zip5,
		row.PostOffice,
		row.Type,
		row.IsMulti,
		row.UpdateStatus,
		row.PrefCode,
	}
	return cols
}

func (row *Row) setPrefCode() {
	row.PrefCode = util.GetPrefCode(row.Pref)
}
