package gax

import (
	"context"
	"time"
	v2 "github.com/googleapis/gax-go/v2"
	"log"
)

type APICall v2.APICall

func gologoo__Invoke_853194a3f89c4e545d625d7d581d6c3a(ctx context.Context, call APICall, opts ...CallOption) error {
	return v2.Invoke(ctx, call, opts...)
}
func gologoo__Sleep_853194a3f89c4e545d625d7d581d6c3a(ctx context.Context, d time.Duration) error {
	return v2.Sleep(ctx, d)
}
func Invoke(ctx context.Context, call APICall, opts ...CallOption) error {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Invoke_853194a3f89c4e545d625d7d581d6c3a")
	log.Printf("Input : %v %v %v\n", ctx, call, opts)
	r0 := gologoo__Invoke_853194a3f89c4e545d625d7d581d6c3a(ctx, call, opts...)
	log.Printf("Output: %v\n", r0)
	return r0
}
func Sleep(ctx context.Context, d time.Duration) error {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Sleep_853194a3f89c4e545d625d7d581d6c3a")
	log.Printf("Input : %v %v\n", ctx, d)
	r0 := gologoo__Sleep_853194a3f89c4e545d625d7d581d6c3a(ctx, d)
	log.Printf("Output: %v\n", r0)
	return r0
}
