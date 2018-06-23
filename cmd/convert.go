package cmd

import (
	"errors"
	"io"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/spf13/cobra"

	"github.com/inouet/ken-all/address"
	"github.com/inouet/ken-all/office"
	"github.com/inouet/ken-all/util"
	"github.com/inouet/ken-all/writer"
)

// http://text.baldanders.info/golang/using-and-testing-cobra/

var output io.Writer

func init() {
	output = os.Stdout
}

func newConvertCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "convert",
		Short: "Convert csv file into other format.",
		Long:  "Convert csv file into other format.",
		RunE: func(cmd *cobra.Command, args []string) error {

			println("=== BUILD ===")

			inputFile, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}

			outputFormat, err := cmd.Flags().GetString("output_format")
			if err != nil {
				return err
			}

			inputType, err := cmd.Flags().GetString("input_type")
			if err != nil {
				return err
			}

			if outputFormat != "json" && outputFormat != "csv" && outputFormat != "tsv" {
				return errors.New("type must be json or csv or tsv")
			}

			switch inputType {
			case "office":
				err = execConvertOfficeCmd(output, inputFile, outputFormat)
			case "address":
				err = execConvertAddressCmd(output, inputFile, outputFormat)
			}
			return err
		},
	}

	cmd.Flags().StringP("output_format", "o", "csv", "output format")
	cmd.Flags().StringP("file", "f", "", "input file")
	cmd.Flags().StringP("input_type", "i", "address", "input type")

	return cmd
}

func execConvertAddressCmd(w io.Writer, inputFile, outputType string) error {

	ioReader, err := os.Open(inputFile)

	defer ioReader.Close()

	if err != nil {
		return err
	}

	rdr := address.NewReader(transform.NewReader(ioReader, japanese.ShiftJIS.NewDecoder()))
	wtr := writer.NewWriter(w, outputType)

	defer wtr.Flush()

	uniq := util.NewUniq()

	for {
		cols, err := rdr.Read()

		if err == io.EOF {
			break
		}

		rows := address.NewRows(cols)

		for _, row := range rows {
			// 同じ 郵便番号で同じ住所は出力しない
			key := row.Zip7 + row.Pref + row.City + row.Town
			if !uniq.IsUnique(key) {
				continue
			}

			wtr.Write(row)
		}
	}

	return nil
}

func execConvertOfficeCmd(w io.Writer, inputFile, outputType string) error {

	ioReader, err := os.Open(inputFile)

	defer ioReader.Close()

	if err != nil {
		return err
	}

	rdr := office.NewReader(transform.NewReader(ioReader, japanese.ShiftJIS.NewDecoder()))
	wtr := writer.NewWriter(w, outputType)

	defer wtr.Flush()

	for {
		cols, err := rdr.Read()

		if err == io.EOF {
			break
		}
		row := office.NewRow(cols)
		err = wtr.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
