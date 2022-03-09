package apierror

import (
	"fmt"
	"strings"
	jsonerror "github.com/googleapis/gax-go/v2/apierror/internal/proto"
	"google.golang.org/api/googleapi"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

type ErrDetails struct {
	ErrorInfo           *errdetails.ErrorInfo
	BadRequest          *errdetails.BadRequest
	PreconditionFailure *errdetails.PreconditionFailure
	QuotaFailure        *errdetails.QuotaFailure
	RetryInfo           *errdetails.RetryInfo
	ResourceInfo        *errdetails.ResourceInfo
	RequestInfo         *errdetails.RequestInfo
	DebugInfo           *errdetails.DebugInfo
	Help                *errdetails.Help
	LocalizedMessage    *errdetails.LocalizedMessage
	Unknown             []interface {
	}
}

func (e ErrDetails) gologoo__String_ec0d6f91a5100afafb5164bd75c1e4e7() string {
	var d strings.Builder
	if e.ErrorInfo != nil {
		d.WriteString(fmt.Sprintf("error details: name = ErrorInfo reason = %s domain = %s metadata = %s\n", e.ErrorInfo.GetReason(), e.ErrorInfo.GetDomain(), e.ErrorInfo.GetMetadata()))
	}
	if e.BadRequest != nil {
		v := e.BadRequest.GetFieldViolations()
		var f []string
		var desc []string
		for _, x := range v {
			f = append(f, x.GetField())
			desc = append(desc, x.GetDescription())
		}
		d.WriteString(fmt.Sprintf("error details: name = BadRequest field = %s desc = %s\n", strings.Join(f, " "), strings.Join(desc, " ")))
	}
	if e.PreconditionFailure != nil {
		v := e.PreconditionFailure.GetViolations()
		var t []string
		var s []string
		var desc []string
		for _, x := range v {
			t = append(t, x.GetType())
			s = append(s, x.GetSubject())
			desc = append(desc, x.GetDescription())
		}
		d.WriteString(fmt.Sprintf("error details: name = PreconditionFailure type = %s subj = %s desc = %s\n", strings.Join(t, " "), strings.Join(s, " "), strings.Join(desc, " ")))
	}
	if e.QuotaFailure != nil {
		v := e.QuotaFailure.GetViolations()
		var s []string
		var desc []string
		for _, x := range v {
			s = append(s, x.GetSubject())
			desc = append(desc, x.GetDescription())
		}
		d.WriteString(fmt.Sprintf("error details: name = QuotaFailure subj = %s desc = %s\n", strings.Join(s, " "), strings.Join(desc, " ")))
	}
	if e.RequestInfo != nil {
		d.WriteString(fmt.Sprintf("error details: name = RequestInfo id = %s data = %s\n", e.RequestInfo.GetRequestId(), e.RequestInfo.GetServingData()))
	}
	if e.ResourceInfo != nil {
		d.WriteString(fmt.Sprintf("error details: name = ResourceInfo type = %s resourcename = %s owner = %s desc = %s\n", e.ResourceInfo.GetResourceType(), e.ResourceInfo.GetResourceName(), e.ResourceInfo.GetOwner(), e.ResourceInfo.GetDescription()))
	}
	if e.RetryInfo != nil {
		d.WriteString(fmt.Sprintf("error details: retry in %s\n", e.RetryInfo.GetRetryDelay().AsDuration()))
	}
	if e.Unknown != nil {
		var s []string
		for _, x := range e.Unknown {
			s = append(s, fmt.Sprintf("%v", x))
		}
		d.WriteString(fmt.Sprintf("error details: name = Unknown  desc = %s\n", strings.Join(s, " ")))
	}
	if e.DebugInfo != nil {
		d.WriteString(fmt.Sprintf("error details: name = DebugInfo detail = %s stack = %s\n", e.DebugInfo.GetDetail(), strings.Join(e.DebugInfo.GetStackEntries(), " ")))
	}
	if e.Help != nil {
		var desc []string
		var url []string
		for _, x := range e.Help.Links {
			desc = append(desc, x.GetDescription())
			url = append(url, x.GetUrl())
		}
		d.WriteString(fmt.Sprintf("error details: name = Help desc = %s url = %s\n", strings.Join(desc, " "), strings.Join(url, " ")))
	}
	if e.LocalizedMessage != nil {
		d.WriteString(fmt.Sprintf("error details: name = LocalizedMessage locale = %s msg = %s\n", e.LocalizedMessage.GetLocale(), e.LocalizedMessage.GetMessage()))
	}
	return d.String()
}

type APIError struct {
	err     error
	status  *status.Status
	httpErr *googleapi.Error
	details ErrDetails
}

func (a *APIError) gologoo__Details_ec0d6f91a5100afafb5164bd75c1e4e7() ErrDetails {
	return a.details
}
func (a *APIError) gologoo__Unwrap_ec0d6f91a5100afafb5164bd75c1e4e7() error {
	return a.err
}
func (a *APIError) gologoo__Error_ec0d6f91a5100afafb5164bd75c1e4e7() string {
	var msg string
	if a.status != nil {
		msg = a.err.Error()
	} else if a.httpErr != nil {
		msg = fmt.Sprintf("googleapi: Error %d: %s", a.httpErr.Code, a.httpErr.Message)
	}
	return strings.TrimSpace(fmt.Sprintf("%s\n%s", msg, a.details))
}
func (a *APIError) gologoo__GRPCStatus_ec0d6f91a5100afafb5164bd75c1e4e7() *status.Status {
	return a.status
}
func (a *APIError) gologoo__Reason_ec0d6f91a5100afafb5164bd75c1e4e7() string {
	return a.details.ErrorInfo.GetReason()
}
func (a *APIError) gologoo__Domain_ec0d6f91a5100afafb5164bd75c1e4e7() string {
	return a.details.ErrorInfo.GetDomain()
}
func (a *APIError) gologoo__Metadata_ec0d6f91a5100afafb5164bd75c1e4e7() map[string]string {
	return a.details.ErrorInfo.GetMetadata()
}
func gologoo__FromError_ec0d6f91a5100afafb5164bd75c1e4e7(err error) (*APIError, bool) {
	if err == nil {
		return nil, false
	}
	ae := APIError{err: err}
	st, isStatus := status.FromError(err)
	herr, isHTTPErr := err.(*googleapi.Error)
	switch {
	case isStatus:
		ae.status = st
		ae.details = parseDetails(st.Details())
	case isHTTPErr:
		ae.httpErr = herr
		ae.details = parseHTTPDetails(herr)
	default:
		return nil, false
	}
	return &ae, true
}
func gologoo__parseDetails_ec0d6f91a5100afafb5164bd75c1e4e7(details []interface {
}) ErrDetails {
	var ed ErrDetails
	for _, d := range details {
		switch d := d.(type) {
		case *errdetails.ErrorInfo:
			ed.ErrorInfo = d
		case *errdetails.BadRequest:
			ed.BadRequest = d
		case *errdetails.PreconditionFailure:
			ed.PreconditionFailure = d
		case *errdetails.QuotaFailure:
			ed.QuotaFailure = d
		case *errdetails.RetryInfo:
			ed.RetryInfo = d
		case *errdetails.ResourceInfo:
			ed.ResourceInfo = d
		case *errdetails.RequestInfo:
			ed.RequestInfo = d
		case *errdetails.DebugInfo:
			ed.DebugInfo = d
		case *errdetails.Help:
			ed.Help = d
		case *errdetails.LocalizedMessage:
			ed.LocalizedMessage = d
		default:
			ed.Unknown = append(ed.Unknown, d)
		}
	}
	return ed
}
func gologoo__parseHTTPDetails_ec0d6f91a5100afafb5164bd75c1e4e7(gae *googleapi.Error) ErrDetails {
	e := &jsonerror.Error{}
	if err := protojson.Unmarshal([]byte(gae.Body), e); err != nil {
		return ErrDetails{}
	}
	details := []interface {
	}{}
	for _, any := range e.GetError().GetDetails() {
		m, err := any.UnmarshalNew()
		if err != nil {
			continue
		}
		details = append(details, m)
	}
	return parseDetails(details)
}
func (e ErrDetails) String() string {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__String_ec0d6f91a5100afafb5164bd75c1e4e7")
	log.Printf("Input : (none)\n")
	r0 := e.gologoo__String_ec0d6f91a5100afafb5164bd75c1e4e7()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (a *APIError) Details() ErrDetails {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Details_ec0d6f91a5100afafb5164bd75c1e4e7")
	log.Printf("Input : (none)\n")
	r0 := a.gologoo__Details_ec0d6f91a5100afafb5164bd75c1e4e7()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (a *APIError) Unwrap() error {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Unwrap_ec0d6f91a5100afafb5164bd75c1e4e7")
	log.Printf("Input : (none)\n")
	r0 := a.gologoo__Unwrap_ec0d6f91a5100afafb5164bd75c1e4e7()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (a *APIError) Error() string {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Error_ec0d6f91a5100afafb5164bd75c1e4e7")
	log.Printf("Input : (none)\n")
	r0 := a.gologoo__Error_ec0d6f91a5100afafb5164bd75c1e4e7()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (a *APIError) GRPCStatus() *status.Status {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__GRPCStatus_ec0d6f91a5100afafb5164bd75c1e4e7")
	log.Printf("Input : (none)\n")
	r0 := a.gologoo__GRPCStatus_ec0d6f91a5100afafb5164bd75c1e4e7()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (a *APIError) Reason() string {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Reason_ec0d6f91a5100afafb5164bd75c1e4e7")
	log.Printf("Input : (none)\n")
	r0 := a.gologoo__Reason_ec0d6f91a5100afafb5164bd75c1e4e7()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (a *APIError) Domain() string {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Domain_ec0d6f91a5100afafb5164bd75c1e4e7")
	log.Printf("Input : (none)\n")
	r0 := a.gologoo__Domain_ec0d6f91a5100afafb5164bd75c1e4e7()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (a *APIError) Metadata() map[string]string {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Metadata_ec0d6f91a5100afafb5164bd75c1e4e7")
	log.Printf("Input : (none)\n")
	r0 := a.gologoo__Metadata_ec0d6f91a5100afafb5164bd75c1e4e7()
	log.Printf("Output: %v\n", r0)
	return r0
}
func FromError(err error) (*APIError, bool) {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__FromError_ec0d6f91a5100afafb5164bd75c1e4e7")
	log.Printf("Input : %v\n", err)
	r0, r1 := gologoo__FromError_ec0d6f91a5100afafb5164bd75c1e4e7(err)
	log.Printf("Output: %v %v\n", r0, r1)
	return r0, r1
}
func parseDetails(details []interface {
}) ErrDetails {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__parseDetails_ec0d6f91a5100afafb5164bd75c1e4e7")
	log.Printf("Input : %v\n", details)
	r0 := gologoo__parseDetails_ec0d6f91a5100afafb5164bd75c1e4e7(details)
	log.Printf("Output: %v\n", r0)
	return r0
}
func parseHTTPDetails(gae *googleapi.Error) ErrDetails {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__parseHTTPDetails_ec0d6f91a5100afafb5164bd75c1e4e7")
	log.Printf("Input : %v\n", gae)
	r0 := gologoo__parseHTTPDetails_ec0d6f91a5100afafb5164bd75c1e4e7(gae)
	log.Printf("Output: %v\n", r0)
	return r0
}
