package test

import (
	"fmt"
	"testing"
)

func TestEncoding(t *testing.T) {
	path := []byte(`/abwadz/wd/`)
	for i, ch := range path {
		path[i] = ch + 1
	}
	fmt.Println(string(path))
	for i, ch := range path {
		path[i] = ch - 1
	}
	fmt.Println(string(path))
}
