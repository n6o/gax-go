package gax

import (
	"math/rand"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type CallOption interface {
	Resolve(cs *CallSettings)
}
type Retryer interface {
	Retry(err error) (pause time.Duration, shouldRetry bool)
}
type retryerOption func() Retryer

func (o retryerOption) gologoo__Resolve_08d1d454519b6352dccd24e63e19b547(s *CallSettings) {
	s.Retry = o
}
func gologoo__WithRetry_08d1d454519b6352dccd24e63e19b547(fn func() Retryer) CallOption {
	return retryerOption(fn)
}
func gologoo__OnErrorFunc_08d1d454519b6352dccd24e63e19b547(bo Backoff, shouldRetry func(err error) bool) Retryer {
	return &errorRetryer{shouldRetry: shouldRetry, backoff: bo}
}

type errorRetryer struct {
	backoff     Backoff
	shouldRetry func(err error) bool
}

func (r *errorRetryer) gologoo__Retry_08d1d454519b6352dccd24e63e19b547(err error) (time.Duration, bool) {
	if r.shouldRetry(err) {
		return r.backoff.Pause(), true
	}
	return 0, false
}
func gologoo__OnCodes_08d1d454519b6352dccd24e63e19b547(cc []codes.Code, bo Backoff) Retryer {
	return &boRetryer{backoff: bo, codes: append([]codes.Code(nil), cc...)}
}

type boRetryer struct {
	backoff Backoff
	codes   []codes.Code
}

func (r *boRetryer) gologoo__Retry_08d1d454519b6352dccd24e63e19b547(err error) (time.Duration, bool) {
	st, ok := status.FromError(err)
	if !ok {
		return 0, false
	}
	c := st.Code()
	for _, rc := range r.codes {
		if c == rc {
			return r.backoff.Pause(), true
		}
	}
	return 0, false
}

type Backoff struct {
	Initial    time.Duration
	Max        time.Duration
	Multiplier float64
	cur        time.Duration
}

func (bo *Backoff) gologoo__Pause_08d1d454519b6352dccd24e63e19b547() time.Duration {
	if bo.Initial == 0 {
		bo.Initial = time.Second
	}
	if bo.cur == 0 {
		bo.cur = bo.Initial
	}
	if bo.Max == 0 {
		bo.Max = 30 * time.Second
	}
	if bo.Multiplier < 1 {
		bo.Multiplier = 2
	}
	d := time.Duration(1 + rand.Int63n(int64(bo.cur)))
	bo.cur = time.Duration(float64(bo.cur) * bo.Multiplier)
	if bo.cur > bo.Max {
		bo.cur = bo.Max
	}
	return d
}

type grpcOpt []grpc.CallOption

func (o grpcOpt) gologoo__Resolve_08d1d454519b6352dccd24e63e19b547(s *CallSettings) {
	s.GRPC = o
}
func gologoo__WithGRPCOptions_08d1d454519b6352dccd24e63e19b547(opt ...grpc.CallOption) CallOption {
	return grpcOpt(append([]grpc.CallOption(nil), opt...))
}

type CallSettings struct {
	Retry func() Retryer
	GRPC  []grpc.CallOption
}

func (o retryerOption) Resolve(s *CallSettings) {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Resolve_08d1d454519b6352dccd24e63e19b547")
	log.Printf("Input : %v\n", s)
	o.gologoo__Resolve_08d1d454519b6352dccd24e63e19b547(s)
	log.Printf("🚚 Output: %v\n", "(none)")
	return
}
func WithRetry(fn func() Retryer) CallOption {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__WithRetry_08d1d454519b6352dccd24e63e19b547")
	log.Printf("Input : %v\n", fn)
	r0 := gologoo__WithRetry_08d1d454519b6352dccd24e63e19b547(fn)
	log.Printf("Output: %v\n", r0)
	return r0
}
func OnErrorFunc(bo Backoff, shouldRetry func(err error) bool) Retryer {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__OnErrorFunc_08d1d454519b6352dccd24e63e19b547")
	log.Printf("Input : %v %v\n", bo, shouldRetry)
	r0 := gologoo__OnErrorFunc_08d1d454519b6352dccd24e63e19b547(bo, shouldRetry)
	log.Printf("Output: %v\n", r0)
	return r0
}
func (r *errorRetryer) Retry(err error) (time.Duration, bool) {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Retry_08d1d454519b6352dccd24e63e19b547")
	log.Printf("Input : %v\n", err)
	r0, r1 := r.gologoo__Retry_08d1d454519b6352dccd24e63e19b547(err)
	log.Printf("Output: %v %v\n", r0, r1)
	return r0, r1
}
func OnCodes(cc []codes.Code, bo Backoff) Retryer {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__OnCodes_08d1d454519b6352dccd24e63e19b547")
	log.Printf("Input : %v %v\n", cc, bo)
	r0 := gologoo__OnCodes_08d1d454519b6352dccd24e63e19b547(cc, bo)
	log.Printf("Output: %v\n", r0)
	return r0
}
func (r *boRetryer) Retry(err error) (time.Duration, bool) {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Retry_08d1d454519b6352dccd24e63e19b547")
	log.Printf("Input : %v\n", err)
	r0, r1 := r.gologoo__Retry_08d1d454519b6352dccd24e63e19b547(err)
	log.Printf("Output: %v %v\n", r0, r1)
	return r0, r1
}
func (bo *Backoff) Pause() time.Duration {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Pause_08d1d454519b6352dccd24e63e19b547")
	log.Printf("Input : (none)\n")
	r0 := bo.gologoo__Pause_08d1d454519b6352dccd24e63e19b547()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (o grpcOpt) Resolve(s *CallSettings) {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Resolve_08d1d454519b6352dccd24e63e19b547")
	log.Printf("Input : %v\n", s)
	o.gologoo__Resolve_08d1d454519b6352dccd24e63e19b547(s)
	log.Printf("🚚 Output: %v\n", "(none)")
	return
}
func WithGRPCOptions(opt ...grpc.CallOption) CallOption {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__WithGRPCOptions_08d1d454519b6352dccd24e63e19b547")
	log.Printf("Input : %v\n", opt)
	r0 := gologoo__WithGRPCOptions_08d1d454519b6352dccd24e63e19b547(opt...)
	log.Printf("Output: %v\n", r0)
	return r0
}
