package gax

import (
	"bytes"
	"log"
)

func gologoo__XGoogHeader_49d9ee14bcfa0c5421d7129e2e464527(keyval ...string) string {
	if len(keyval) == 0 {
		return ""
	}
	if len(keyval)%2 != 0 {
		panic("gax.Header: odd argument count")
	}
	var buf bytes.Buffer
	for i := 0; i < len(keyval); i += 2 {
		buf.WriteByte(' ')
		buf.WriteString(keyval[i])
		buf.WriteByte('/')
		buf.WriteString(keyval[i+1])
	}
	return buf.String()[1:]
}
func XGoogHeader(keyval ...string) string {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__XGoogHeader_49d9ee14bcfa0c5421d7129e2e464527")
	log.Printf("Input : %v\n", keyval)
	r0 := gologoo__XGoogHeader_49d9ee14bcfa0c5421d7129e2e464527(keyval...)
	log.Printf("Output: %v\n", r0)
	return r0
}
