package writer

import (
	"encoding/csv"
	"io"
)

type Writer interface {
	Write(Writable) error
	Flush()
}

func NewWriter(writer io.Writer, outputType string) Writer {
	var w Writer

	if outputType == "json" {
		w = JSONWriter{
			w: writer,
		}
	} else if outputType == "csv" || outputType == "tsv" {
		tmp := csv.NewWriter(writer)
		if outputType == "tsv" {
			tmp.Comma = '\t'
		}

		w = CsvWriter{
			w: tmp,
		}
	}
	return w
}
