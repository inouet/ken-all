package address

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/inouet/ken-all/util"
)

// Reader reads records from a CSV-encoded file.
type Reader struct {
	r *csv.Reader
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: csv.NewReader(r),
	}
}

func (reader *Reader) Read() (record []string, err error) {

	inBrackets := false

	idxTownKana := 5
	idxTownName := 8

	townKana := ""
	townName := ""

	for {
		record, err = reader.r.Read()

		if err == io.EOF {
			break
		}

		for _, v := range []int{3, 4, 5, 8} {
			record[v] = util.NormalizeString(record[v])
		}

		// zip5のスペース除去
		record[1] = strings.Trim(record[1], " ")

		if strings.Contains(record[idxTownName], "(") {
			inBrackets = true
		}

		if inBrackets { // カッコ内の場合は結合
			townName = townName + record[idxTownName]
			if townKana != record[idxTownKana] {
				// 6028064 イッチョウメ のように同じものが続く場合は無視する
				townKana = townKana + record[idxTownKana]
			}
		}

		if strings.Contains(record[idxTownName], ")") {
			inBrackets = false
		}

		if !inBrackets { // カッコ内でない場合
			if townKana != "" {
				record[idxTownName] = townName
				record[idxTownKana] = townKana
			}
			return record, err
		}
		continue
	}
	return record, err
}
