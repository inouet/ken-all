package util

import "testing"

func TestNormalizeString(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{input: "ｲﾜﾃｹﾝ", expected: "イワテケン"},
		{input: "１２３４５６７８９０", expected: "1234567890"},
		{input: "６丁目１－２アーバンネット札幌ビル２Ｆ", expected: "6丁目1-2アーバンネット札幌ビル2F"},
		{input: "株式会社　日本経済新聞社　札幌支社", expected: "株式会社 日本経済新聞社 札幌支社"},
	}

	for _, c := range cases {
		actual := NormalizeString(c.input)
		if actual != c.expected {
			t.Errorf("want '%s', got '%s'\n", c.expected, actual)
		}
	}
}

func TestGetPrefCode(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{input: "東京都", expected: "13"},
		{input: "テスト", expected: ""},
	}
	for _, c := range cases {
		actual := GetPrefCode(c.input)
		if actual != c.expected {
			t.Errorf("want '%s', got '%s'\n", c.expected, actual)
		}
	}
}
