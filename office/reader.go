package office

import (
	"encoding/csv"
	"io"

	"github.com/inouet/ken-all/util"
)

type Reader struct {
	r *csv.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: csv.NewReader(r),
	}
}

func (reader *Reader) Read() (record []string, err error) {
	record, err = reader.r.Read()
	if err != nil {
		return record, err
	}
	for _, v := range []int{1, 2, 4, 5, 6} {
		record[v] = util.NormalizeString(record[v])
	}
	return record, err
}
