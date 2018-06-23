package cmd

// https://deeeet.com/writing/2014/12/18/golang-cli-test/
// https://qiita.com/kami_zh/items/ff636f15da87dabebe6c

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var testDataDir string

func init() {
	testDataDir = "../testdata"
}

func TestBuild(t *testing.T) {

	cases := []struct {
		input      string
		output     string
		outputType string
	}{
		{
			input:      "test_001.csv",
			output:     "test_001_out.json",
			outputType: "json",
		},
	}

	for _, c := range cases {
		buffer := &bytes.Buffer{}
		inputFile := filepath.Join(testDataDir, c.input)
		outputFile := filepath.Join(testDataDir, c.output)

		execConvertAddressCmd(buffer, inputFile, c.outputType)

		b, err := ioutil.ReadFile(outputFile)
		if err != nil {
			t.Errorf("File not found %s", outputFile)
		}

		if buffer.String() != string(b) {
			t.Errorf("want '%s', got '%s'\n", string(b), buffer)
		}
	}
}
