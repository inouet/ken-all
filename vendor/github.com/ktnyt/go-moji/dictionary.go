package moji

import (
	"strings"
)

// Dictionary defines an interface for mapping between a string and index
type Dictionary interface {
	Encode([]MaybeIndex) string
	Decode(string) []MaybeIndex
}

type defaultDictionary []string

func (d defaultDictionary) encode(m MaybeIndex) string {
	if m.i < 0 || m.i > len(d) {
		return m.s
	}
	return d[m.i]
}

func (d defaultDictionary) decode(s string) (MaybeIndex, string) {
	for i, p := range d {
		if strings.HasPrefix(s, p) {
			return MaybeIndex{i: i, s: p}, strings.Replace(s, p, "", 1)
		}
	}
	rs := []rune(s)
	head := string(rs[:1])
	tail := string(rs[1:])
	return MaybeIndex{i: -1, s: head}, tail
}

func (d defaultDictionary) Encode(ms []MaybeIndex) string {
	s := make([]byte, 0)
	for _, m := range ms {
		s = append(s, d.encode(m)...)
	}
	return string(s)
}

func (d defaultDictionary) Decode(s string) []MaybeIndex {
	ms := make([]MaybeIndex, 0)
	var m MaybeIndex
	for len(s) != 0 {
		m, s = d.decode(s)
		ms = append(ms, m)
	}
	return ms
}

// NewDictionary creates a dictionary from the given string slice
func NewDictionary(d []string) Dictionary {
	return defaultDictionary(d)
}

// NewRangeDictionary creates a dictionary from the given range
func NewRangeDictionary(s, e rune) Dictionary {
	if s > e {
		panic("NewRangeDictionary: range is invaid")
	}
	d := make([]string, 0)
	for r := s; r != e; r++ {
		d = append(d, string([]rune{r}))
	}
	return NewDictionary(d)
}
