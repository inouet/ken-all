package writer

import (
	"encoding/json"
	"io"
)

type JSONWriter struct {
	w io.Writer
}

func (writer JSONWriter) Write(row Writable) error {
	b, err := json.Marshal(row)

	if err != nil {
		return err
	}

	b = append(b, "\n"...)

	_, err = writer.w.Write(b)

	if err != nil {
		return err
	}

	return err
}

func (writer JSONWriter) Flush() {
}
