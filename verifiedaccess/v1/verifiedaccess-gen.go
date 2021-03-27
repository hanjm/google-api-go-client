// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package verifiedaccess provides access to the Chrome Verified Access API.
//
// For product documentation, see: https://developers.google.com/chrome/verified-access
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/verifiedaccess/v1"
//   ...
//   ctx := context.Background()
//   verifiedaccessService, err := verifiedaccess.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   verifiedaccessService, err := verifiedaccess.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   verifiedaccessService, err := verifiedaccess.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package verifiedaccess // import "google.golang.org/api/verifiedaccess/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "verifiedaccess:v1"
const apiName = "verifiedaccess"
const apiVersion = "v1"
const basePath = "https://verifiedaccess.googleapis.com/"
const mtlsBasePath = "https://verifiedaccess.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Verify your enterprise credentials
	VerifiedaccessScope = "https://www.googleapis.com/auth/verifiedaccess"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/verifiedaccess",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Challenge = NewChallengeService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Challenge *ChallengeService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewChallengeService(s *Service) *ChallengeService {
	rs := &ChallengeService{s: s}
	return rs
}

type ChallengeService struct {
	s *Service
}

// Challenge: Result message for VerifiedAccess.CreateChallenge.
type Challenge struct {
	// AlternativeChallenge: Challenge generated with the old signing key
	// (this will only be present during key rotation)
	AlternativeChallenge *SignedData `json:"alternativeChallenge,omitempty"`

	// Challenge: Generated challenge
	Challenge *SignedData `json:"challenge,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "AlternativeChallenge") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AlternativeChallenge") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Challenge) MarshalJSON() ([]byte, error) {
	type NoMethod Challenge
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated empty messages in your APIs. A typical example is to use
// it as the request or the response type of an API method. For
// instance: service Foo { rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty); } The JSON representation for `Empty` is
// empty JSON object `{}`.
type Empty struct {
}

// SignedData: The wrapper message of any data and its signature.
type SignedData struct {
	// Data: The data to be signed.
	Data string `json:"data,omitempty"`

	// Signature: The signature of the data field.
	Signature string `json:"signature,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Data") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Data") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SignedData) MarshalJSON() ([]byte, error) {
	type NoMethod SignedData
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VerifyChallengeResponseRequest: signed ChallengeResponse
type VerifyChallengeResponseRequest struct {
	// ChallengeResponse: The generated response to the challenge
	ChallengeResponse *SignedData `json:"challengeResponse,omitempty"`

	// ExpectedIdentity: Service can optionally provide identity information
	// about the device or user associated with the key. For an EMK, this
	// value is the enrolled domain. For an EUK, this value is the user's
	// email address. If present, this value will be checked against
	// contents of the response, and verification will fail if there is no
	// match.
	ExpectedIdentity string `json:"expectedIdentity,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ChallengeResponse")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ChallengeResponse") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *VerifyChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	type NoMethod VerifyChallengeResponseRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VerifyChallengeResponseResult: Result message for
// VerifiedAccess.VerifyChallengeResponse.
type VerifyChallengeResponseResult struct {
	// DeviceEnrollmentId: Device enrollment id is returned in this field
	// (for the machine response only).
	DeviceEnrollmentId string `json:"deviceEnrollmentId,omitempty"`

	// DevicePermanentId: Device permanent id is returned in this field (for
	// the machine response only).
	DevicePermanentId string `json:"devicePermanentId,omitempty"`

	// SignedPublicKeyAndChallenge: Certificate Signing Request (in the
	// SPKAC format, base64 encoded) is returned in this field. This field
	// will be set only if device has included CSR in its challenge
	// response. (the option to include CSR is now available for both user
	// and machine responses)
	SignedPublicKeyAndChallenge string `json:"signedPublicKeyAndChallenge,omitempty"`

	// VerificationOutput: For EMCert check, device permanent id is returned
	// here. For EUCert check, signed_public_key_and_challenge [base64
	// encoded] is returned if present, otherwise empty string is returned.
	// This field is deprecated, please use device_permanent_id or
	// signed_public_key_and_challenge fields.
	VerificationOutput string `json:"verificationOutput,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DeviceEnrollmentId")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceEnrollmentId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *VerifyChallengeResponseResult) MarshalJSON() ([]byte, error) {
	type NoMethod VerifyChallengeResponseResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "verifiedaccess.challenge.create":

type ChallengeCreateCall struct {
	s          *Service
	empty      *Empty
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: CreateChallenge API
func (r *ChallengeService) Create(empty *Empty) *ChallengeCreateCall {
	c := &ChallengeCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.empty = empty
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChallengeCreateCall) Fields(s ...googleapi.Field) *ChallengeCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ChallengeCreateCall) Context(ctx context.Context) *ChallengeCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ChallengeCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ChallengeCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210326")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.empty)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/challenge")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "verifiedaccess.challenge.create" call.
// Exactly one of *Challenge or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Challenge.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ChallengeCreateCall) Do(opts ...googleapi.CallOption) (*Challenge, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Challenge{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "CreateChallenge API",
	//   "flatPath": "v1/challenge",
	//   "httpMethod": "POST",
	//   "id": "verifiedaccess.challenge.create",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/challenge",
	//   "request": {
	//     "$ref": "Empty"
	//   },
	//   "response": {
	//     "$ref": "Challenge"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/verifiedaccess"
	//   ]
	// }

}

// method id "verifiedaccess.challenge.verify":

type ChallengeVerifyCall struct {
	s                              *Service
	verifychallengeresponserequest *VerifyChallengeResponseRequest
	urlParams_                     gensupport.URLParams
	ctx_                           context.Context
	header_                        http.Header
}

// Verify: VerifyChallengeResponse API
func (r *ChallengeService) Verify(verifychallengeresponserequest *VerifyChallengeResponseRequest) *ChallengeVerifyCall {
	c := &ChallengeVerifyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.verifychallengeresponserequest = verifychallengeresponserequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChallengeVerifyCall) Fields(s ...googleapi.Field) *ChallengeVerifyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ChallengeVerifyCall) Context(ctx context.Context) *ChallengeVerifyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ChallengeVerifyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ChallengeVerifyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210326")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.verifychallengeresponserequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/challenge:verify")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "verifiedaccess.challenge.verify" call.
// Exactly one of *VerifyChallengeResponseResult or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *VerifyChallengeResponseResult.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ChallengeVerifyCall) Do(opts ...googleapi.CallOption) (*VerifyChallengeResponseResult, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &VerifyChallengeResponseResult{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "VerifyChallengeResponse API",
	//   "flatPath": "v1/challenge:verify",
	//   "httpMethod": "POST",
	//   "id": "verifiedaccess.challenge.verify",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/challenge:verify",
	//   "request": {
	//     "$ref": "VerifyChallengeResponseRequest"
	//   },
	//   "response": {
	//     "$ref": "VerifyChallengeResponseResult"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/verifiedaccess"
	//   ]
	// }

}
