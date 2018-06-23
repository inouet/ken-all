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
	for _, i := range []int{1, 2, 4, 5, 6} {
		record[i] = util.NormalizeString(record[i])
	}
	return record, err
}
