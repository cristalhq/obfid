package obfid

import (
	"errors"
	"strings"
)

var ErrNoPrefix = errors.New("")

type Stringer struct {
	gen    *Generator
	set    string
	prefix string
}

func NewStringer(g *Generator, set string) (*Stringer, error) {
	return NewStringerWithPrefix(g, set, "")
}

func NewStringerWithPrefix(gen *Generator, set, prefix string) (*Stringer, error) {
	// TODO: check for duplicates in set
	s := &Stringer{
		gen:    gen,
		set:    set,
		prefix: prefix,
	}
	return s, nil
}

func (s *Stringer) Generate(number uint64) string {
	obf := s.gen.Encode(number)
	_ = obf
	return ""
}

func (s *Stringer) Parse(str string) (number uint64, err error) {
	if s.prefix != "" {
		if !strings.HasPrefix(str, s.prefix) {
			return 0, ErrNoPrefix
		}
		str = str[len(s.prefix):]
	}

	return 0, nil
}
