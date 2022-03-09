package gax

import (
	v2 "github.com/googleapis/gax-go/v2"
	"log"
)

func gologoo__XGoogHeader_0fcce1757ce52aaf367d27500d3bf074(keyval ...string) string {
	return v2.XGoogHeader(keyval...)
}
func XGoogHeader(keyval ...string) string {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__XGoogHeader_0fcce1757ce52aaf367d27500d3bf074")
	log.Printf("Input : %v\n", keyval)
	r0 := gologoo__XGoogHeader_0fcce1757ce52aaf367d27500d3bf074(keyval...)
	log.Printf("Output: %v\n", r0)
	return r0
}
