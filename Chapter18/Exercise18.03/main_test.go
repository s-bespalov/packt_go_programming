package main

import (
	"bytes"
	"log"
	"testing"
)

func Test_Main(t *testing.T) {
	for range 10000 {
		var s bytes.Buffer
		log.SetOutput(&s)
		log.SetFlags(0)
		main()
		if s.String() != "5050\n" {
			t.Error(s.String())
		}
	}
}
