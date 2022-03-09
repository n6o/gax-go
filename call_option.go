package gax

import (
	v2 "github.com/googleapis/gax-go/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
)

type CallOption v2.CallOption
type Retryer v2.Retryer

func gologoo__WithRetry_379d97ab91819bf53766df2a7569a9e0(fn func() Retryer) CallOption {
	return v2.WithRetry(fn)
}
func gologoo__OnCodes_379d97ab91819bf53766df2a7569a9e0(cc []codes.Code, bo Backoff) Retryer {
	return v2.OnCodes(cc, bo)
}

type Backoff v2.Backoff

func gologoo__WithGRPCOptions_379d97ab91819bf53766df2a7569a9e0(opt ...grpc.CallOption) CallOption {
	return v2.WithGRPCOptions(opt...)
}

type CallSettings v2.CallSettings

func WithRetry(fn func() Retryer) CallOption {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__WithRetry_379d97ab91819bf53766df2a7569a9e0")
	log.Printf("Input : %v\n", fn)
	r0 := gologoo__WithRetry_379d97ab91819bf53766df2a7569a9e0(fn)
	log.Printf("Output: %v\n", r0)
	return r0
}
func OnCodes(cc []codes.Code, bo Backoff) Retryer {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__OnCodes_379d97ab91819bf53766df2a7569a9e0")
	log.Printf("Input : %v %v\n", cc, bo)
	r0 := gologoo__OnCodes_379d97ab91819bf53766df2a7569a9e0(cc, bo)
	log.Printf("Output: %v\n", r0)
	return r0
}
func WithGRPCOptions(opt ...grpc.CallOption) CallOption {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__WithGRPCOptions_379d97ab91819bf53766df2a7569a9e0")
	log.Printf("Input : %v\n", opt)
	r0 := gologoo__WithGRPCOptions_379d97ab91819bf53766df2a7569a9e0(opt...)
	log.Printf("Output: %v\n", r0)
	return r0
}
