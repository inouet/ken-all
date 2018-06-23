package writer

import (
	"encoding/csv"
)

type Writable interface {
	Array() []string
}

type CsvWriter struct {
	w *csv.Writer
}

func (writer CsvWriter) Write(row Writable) error {
	cols := row.Array()
	return writer.w.Write(cols)
}

func (writer CsvWriter) Flush() {
	writer.w.Flush()
}
