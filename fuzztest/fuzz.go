package fuzz

import (
	"github.com/gradienthealth/go-dicom"
)

func Fuzz(data []byte) int {
	p, _ := dicom.NewParserFromBytes(data, nil)
	_, _ = p.Parse(dicom.ParseOptions{})

	return 1
}
