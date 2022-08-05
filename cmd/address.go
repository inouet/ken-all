package cmd

import (
	"errors"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/spf13/cobra"

	"github.com/inouet/ken-all/address"
	"github.com/inouet/ken-all/util"
	"github.com/inouet/ken-all/writer"
)

// http://text.baldanders.info/golang/using-and-testing-cobra/

func init() {
	output = os.Stdout
}

func newAddressCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "address [KEN_ALL.CSV]",
		Short: "Convert KEN_ALL.CSV into other format.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			inputFile := args[0]

			outputType, err := cmd.Flags().GetString("type")
			if err != nil {
				return err
			}

			if !isValidOutputType(outputType) {
				return errors.New("type must be json or csv or tsv")
			}

			err = execAddressCmd(output, inputFile, outputType)
			return err
		},
	}

	cmd.Flags().StringP("type", "t", "csv", "output type [json,csv,tsv]")

	return cmd
}

func execAddressCmd(w io.Writer, inputFile, outputType string) error {

	ioReader, err := os.Open(inputFile)

	defer func() {
		err := ioReader.Close()
		if err != nil {
			log.Println("can't close ioReader", err)
		}
	}()

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

			err := wtr.Write(row)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
