package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr := []string{
		"empid", "employee", "address", "hours worked",
		"hourly rate", "manager",
	}
	result := csvHdrCol(hdr)
	fmt.Printf("Result:\n%v\n", result)
	hdr2 := []string{
		"employee", "empid", "hours worked",
		"address", "manager", "hourly rate",
	}
	result = csvHdrCol(hdr2)
	fmt.Printf("Result:\n%v\n", result)
}

func csvHdrCol(header []string) map[int]string {
	csvIdxToCol := make(map[int]string)
	for i, v := range header {
		v = strings.TrimSpace(v)
		switch strings.ToLower(v) {
		case "employee":
			csvIdxToCol[i] = v
		case "hours worked":
			csvIdxToCol[i] = v
		case "hourly rate":
			csvIdxToCol[i] = v
		}
	}
	return csvIdxToCol
}
