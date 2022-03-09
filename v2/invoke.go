package gax

import (
	"context"
	"strings"
	"time"
	"github.com/googleapis/gax-go/v2/apierror"
	"log"
)

type APICall func(context.Context, CallSettings) error

func gologoo__Invoke_997373c1c44cd44a6802317db0de3b09(ctx context.Context, call APICall, opts ...CallOption) error {
	var settings CallSettings
	for _, opt := range opts {
		opt.Resolve(&settings)
	}
	return invoke(ctx, call, settings, Sleep)
}
func gologoo__Sleep_997373c1c44cd44a6802317db0de3b09(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	select {
	case <-ctx.Done():
		t.Stop()
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

type sleeper func(ctx context.Context, d time.Duration) error

func gologoo__invoke_997373c1c44cd44a6802317db0de3b09(ctx context.Context, call APICall, settings CallSettings, sp sleeper) error {
	var retryer Retryer
	for {
		err := call(ctx, settings)
		if err == nil {
			return nil
		}
		if strings.Contains(err.Error(), "x509: certificate signed by unknown authority") {
			return err
		}
		if apierr, ok := apierror.FromError(err); ok {
			err = apierr
		}
		if settings.Retry == nil {
			return err
		}
		if retryer == nil {
			if r := settings.Retry(); r != nil {
				retryer = r
			} else {
				return err
			}
		}
		if d, ok := retryer.Retry(err); !ok {
			return err
		} else if err = sp(ctx, d); err != nil {
			return err
		}
	}
}
func Invoke(ctx context.Context, call APICall, opts ...CallOption) error {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Invoke_997373c1c44cd44a6802317db0de3b09")
	log.Printf("Input : %v %v %v\n", ctx, call, opts)
	r0 := gologoo__Invoke_997373c1c44cd44a6802317db0de3b09(ctx, call, opts...)
	log.Printf("Output: %v\n", r0)
	return r0
}
func Sleep(ctx context.Context, d time.Duration) error {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Sleep_997373c1c44cd44a6802317db0de3b09")
	log.Printf("Input : %v %v\n", ctx, d)
	r0 := gologoo__Sleep_997373c1c44cd44a6802317db0de3b09(ctx, d)
	log.Printf("Output: %v\n", r0)
	return r0
}
func invoke(ctx context.Context, call APICall, settings CallSettings, sp sleeper) error {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__invoke_997373c1c44cd44a6802317db0de3b09")
	log.Printf("Input : %v %v %v %v\n", ctx, call, settings, sp)
	r0 := gologoo__invoke_997373c1c44cd44a6802317db0de3b09(ctx, call, settings, sp)
	log.Printf("Output: %v\n", r0)
	return r0
}
