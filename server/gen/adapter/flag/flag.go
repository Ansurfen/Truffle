package flag

import "strings"

type sliceValue []string

func NewSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}

func (s *sliceValue) String() string {
	*s = sliceValue(strings.Split("", ","))
	return ""
}
