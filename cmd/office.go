package cmd

import (
	"errors"
	"io"
	"os"

	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/japanese"
	"github.com/spf13/cobra"

	"github.com/inouet/ken-all/office"
	"github.com/inouet/ken-all/writer"
)

func init() {
	output = os.Stdout
}

func newOfficeCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "office [JIGYOSYO.CSV]",
		Short: "Convert JIGYOSYO.CSV into other format.",
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

			err = execOfficeCmd(output, inputFile, outputType)

			return err
		},
	}

	cmd.Flags().StringP("type", "t", "csv", "output type [json,csv,tsv]")

	return cmd
}

func execOfficeCmd(w io.Writer, inputFile, outputType string) error {

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
