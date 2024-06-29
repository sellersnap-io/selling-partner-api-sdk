// Package merchantFulfillment provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package merchantFulfillment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"github.com/amzapi/selling-partner-api-sdk/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetAdditionalSellerInputs request  with any body
	GetAdditionalSellerInputsWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	GetAdditionalSellerInputs(ctx context.Context, body GetAdditionalSellerInputsJSONRequestBody) (*http.Response, error)

	// GetEligibleShipmentServicesOld request  with any body
	GetEligibleShipmentServicesOldWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	GetEligibleShipmentServicesOld(ctx context.Context, body GetEligibleShipmentServicesOldJSONRequestBody) (*http.Response, error)

	// GetEligibleShipmentServices request  with any body
	GetEligibleShipmentServicesWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	GetEligibleShipmentServices(ctx context.Context, body GetEligibleShipmentServicesJSONRequestBody) (*http.Response, error)

	// GetAdditionalSellerInputsOld request  with any body
	GetAdditionalSellerInputsOldWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	GetAdditionalSellerInputsOld(ctx context.Context, body GetAdditionalSellerInputsOldJSONRequestBody) (*http.Response, error)

	// CreateShipment request  with any body
	CreateShipmentWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	CreateShipment(ctx context.Context, body CreateShipmentJSONRequestBody) (*http.Response, error)

	// CancelShipment request
	CancelShipment(ctx context.Context, shipmentId string) (*http.Response, error)

	// GetShipment request
	GetShipment(ctx context.Context, shipmentId string) (*http.Response, error)

	// CancelShipmentOld request
	CancelShipmentOld(ctx context.Context, shipmentId string) (*http.Response, error)
}

func (c *Client) GetAdditionalSellerInputsWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewGetAdditionalSellerInputsRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetAdditionalSellerInputs(ctx context.Context, body GetAdditionalSellerInputsJSONRequestBody) (*http.Response, error) {
	req, err := NewGetAdditionalSellerInputsRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetEligibleShipmentServicesOldWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewGetEligibleShipmentServicesOldRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetEligibleShipmentServicesOld(ctx context.Context, body GetEligibleShipmentServicesOldJSONRequestBody) (*http.Response, error) {
	req, err := NewGetEligibleShipmentServicesOldRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetEligibleShipmentServicesWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewGetEligibleShipmentServicesRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetEligibleShipmentServices(ctx context.Context, body GetEligibleShipmentServicesJSONRequestBody) (*http.Response, error) {
	req, err := NewGetEligibleShipmentServicesRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetAdditionalSellerInputsOldWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewGetAdditionalSellerInputsOldRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetAdditionalSellerInputsOld(ctx context.Context, body GetAdditionalSellerInputsOldJSONRequestBody) (*http.Response, error) {
	req, err := NewGetAdditionalSellerInputsOldRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) CreateShipmentWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewCreateShipmentRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CreateShipment(ctx context.Context, body CreateShipmentJSONRequestBody) (*http.Response, error) {
	req, err := NewCreateShipmentRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) CancelShipment(ctx context.Context, shipmentId string) (*http.Response, error) {
	req, err := NewCancelShipmentRequest(c.Endpoint, shipmentId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetShipment(ctx context.Context, shipmentId string) (*http.Response, error) {
	req, err := NewGetShipmentRequest(c.Endpoint, shipmentId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CancelShipmentOld(ctx context.Context, shipmentId string) (*http.Response, error) {
	req, err := NewCancelShipmentOldRequest(c.Endpoint, shipmentId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewGetAdditionalSellerInputsRequest calls the generic GetAdditionalSellerInputs builder with application/json body
func NewGetAdditionalSellerInputsRequest(endpoint string, body GetAdditionalSellerInputsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetAdditionalSellerInputsRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewGetAdditionalSellerInputsRequestWithBody generates requests for GetAdditionalSellerInputs with any type of body
func NewGetAdditionalSellerInputsRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/mfn/v0/additionalSellerInputs")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetEligibleShipmentServicesOldRequest calls the generic GetEligibleShipmentServicesOld builder with application/json body
func NewGetEligibleShipmentServicesOldRequest(endpoint string, body GetEligibleShipmentServicesOldJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetEligibleShipmentServicesOldRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewGetEligibleShipmentServicesOldRequestWithBody generates requests for GetEligibleShipmentServicesOld with any type of body
func NewGetEligibleShipmentServicesOldRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/mfn/v0/eligibleServices")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetEligibleShipmentServicesRequest calls the generic GetEligibleShipmentServices builder with application/json body
func NewGetEligibleShipmentServicesRequest(endpoint string, body GetEligibleShipmentServicesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetEligibleShipmentServicesRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewGetEligibleShipmentServicesRequestWithBody generates requests for GetEligibleShipmentServices with any type of body
func NewGetEligibleShipmentServicesRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/mfn/v0/eligibleShippingServices")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetAdditionalSellerInputsOldRequest calls the generic GetAdditionalSellerInputsOld builder with application/json body
func NewGetAdditionalSellerInputsOldRequest(endpoint string, body GetAdditionalSellerInputsOldJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetAdditionalSellerInputsOldRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewGetAdditionalSellerInputsOldRequestWithBody generates requests for GetAdditionalSellerInputsOld with any type of body
func NewGetAdditionalSellerInputsOldRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/mfn/v0/sellerInputs")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewCreateShipmentRequest calls the generic CreateShipment builder with application/json body
func NewCreateShipmentRequest(endpoint string, body CreateShipmentJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateShipmentRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewCreateShipmentRequestWithBody generates requests for CreateShipment with any type of body
func NewCreateShipmentRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/mfn/v0/shipments")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewCancelShipmentRequest generates requests for CancelShipment
func NewCancelShipmentRequest(endpoint string, shipmentId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "shipmentId", shipmentId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/mfn/v0/shipments/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetShipmentRequest generates requests for GetShipment
func NewGetShipmentRequest(endpoint string, shipmentId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "shipmentId", shipmentId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/mfn/v0/shipments/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCancelShipmentOldRequest generates requests for CancelShipmentOld
func NewCancelShipmentOldRequest(endpoint string, shipmentId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "shipmentId", shipmentId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/mfn/v0/shipments/%s/cancel", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetAdditionalSellerInputs request  with any body
	GetAdditionalSellerInputsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetAdditionalSellerInputsResp, error)

	GetAdditionalSellerInputsWithResponse(ctx context.Context, body GetAdditionalSellerInputsJSONRequestBody) (*GetAdditionalSellerInputsResp, error)

	// GetEligibleShipmentServicesOld request  with any body
	GetEligibleShipmentServicesOldWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetEligibleShipmentServicesOldResp, error)

	GetEligibleShipmentServicesOldWithResponse(ctx context.Context, body GetEligibleShipmentServicesOldJSONRequestBody) (*GetEligibleShipmentServicesOldResp, error)

	// GetEligibleShipmentServices request  with any body
	GetEligibleShipmentServicesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetEligibleShipmentServicesResp, error)

	GetEligibleShipmentServicesWithResponse(ctx context.Context, body GetEligibleShipmentServicesJSONRequestBody) (*GetEligibleShipmentServicesResp, error)

	// GetAdditionalSellerInputsOld request  with any body
	GetAdditionalSellerInputsOldWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetAdditionalSellerInputsOldResp, error)

	GetAdditionalSellerInputsOldWithResponse(ctx context.Context, body GetAdditionalSellerInputsOldJSONRequestBody) (*GetAdditionalSellerInputsOldResp, error)

	// CreateShipment request  with any body
	CreateShipmentWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateShipmentResp, error)

	CreateShipmentWithResponse(ctx context.Context, body CreateShipmentJSONRequestBody) (*CreateShipmentResp, error)

	// CancelShipment request
	CancelShipmentWithResponse(ctx context.Context, shipmentId string) (*CancelShipmentResp, error)

	// GetShipment request
	GetShipmentWithResponse(ctx context.Context, shipmentId string) (*GetShipmentResp, error)

	// CancelShipmentOld request
	CancelShipmentOldWithResponse(ctx context.Context, shipmentId string) (*CancelShipmentOldResp, error)
}

type GetAdditionalSellerInputsResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetAdditionalSellerInputsResponse
}

// Status returns HTTPResponse.Status
func (r GetAdditionalSellerInputsResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAdditionalSellerInputsResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetEligibleShipmentServicesOldResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetEligibleShipmentServicesResponse
}

// Status returns HTTPResponse.Status
func (r GetEligibleShipmentServicesOldResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetEligibleShipmentServicesOldResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetEligibleShipmentServicesResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetEligibleShipmentServicesResponse
}

// Status returns HTTPResponse.Status
func (r GetEligibleShipmentServicesResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetEligibleShipmentServicesResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAdditionalSellerInputsOldResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetAdditionalSellerInputsResponse
}

// Status returns HTTPResponse.Status
func (r GetAdditionalSellerInputsOldResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAdditionalSellerInputsOldResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateShipmentResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CreateShipmentResponse
}

// Status returns HTTPResponse.Status
func (r CreateShipmentResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateShipmentResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CancelShipmentResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CancelShipmentResponse
}

// Status returns HTTPResponse.Status
func (r CancelShipmentResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CancelShipmentResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetShipmentResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetShipmentResponse
}

// Status returns HTTPResponse.Status
func (r GetShipmentResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetShipmentResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CancelShipmentOldResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CancelShipmentResponse
}

// Status returns HTTPResponse.Status
func (r CancelShipmentOldResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CancelShipmentOldResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAdditionalSellerInputsWithBodyWithResponse request with arbitrary body returning *GetAdditionalSellerInputsResponse
func (c *ClientWithResponses) GetAdditionalSellerInputsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetAdditionalSellerInputsResp, error) {
	rsp, err := c.GetAdditionalSellerInputsWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseGetAdditionalSellerInputsResp(rsp)
}

func (c *ClientWithResponses) GetAdditionalSellerInputsWithResponse(ctx context.Context, body GetAdditionalSellerInputsJSONRequestBody) (*GetAdditionalSellerInputsResp, error) {
	rsp, err := c.GetAdditionalSellerInputs(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseGetAdditionalSellerInputsResp(rsp)
}

// GetEligibleShipmentServicesOldWithBodyWithResponse request with arbitrary body returning *GetEligibleShipmentServicesOldResponse
func (c *ClientWithResponses) GetEligibleShipmentServicesOldWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetEligibleShipmentServicesOldResp, error) {
	rsp, err := c.GetEligibleShipmentServicesOldWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseGetEligibleShipmentServicesOldResp(rsp)
}

func (c *ClientWithResponses) GetEligibleShipmentServicesOldWithResponse(ctx context.Context, body GetEligibleShipmentServicesOldJSONRequestBody) (*GetEligibleShipmentServicesOldResp, error) {
	rsp, err := c.GetEligibleShipmentServicesOld(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseGetEligibleShipmentServicesOldResp(rsp)
}

// GetEligibleShipmentServicesWithBodyWithResponse request with arbitrary body returning *GetEligibleShipmentServicesResponse
func (c *ClientWithResponses) GetEligibleShipmentServicesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetEligibleShipmentServicesResp, error) {
	rsp, err := c.GetEligibleShipmentServicesWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseGetEligibleShipmentServicesResp(rsp)
}

func (c *ClientWithResponses) GetEligibleShipmentServicesWithResponse(ctx context.Context, body GetEligibleShipmentServicesJSONRequestBody) (*GetEligibleShipmentServicesResp, error) {
	rsp, err := c.GetEligibleShipmentServices(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseGetEligibleShipmentServicesResp(rsp)
}

// GetAdditionalSellerInputsOldWithBodyWithResponse request with arbitrary body returning *GetAdditionalSellerInputsOldResponse
func (c *ClientWithResponses) GetAdditionalSellerInputsOldWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetAdditionalSellerInputsOldResp, error) {
	rsp, err := c.GetAdditionalSellerInputsOldWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseGetAdditionalSellerInputsOldResp(rsp)
}

func (c *ClientWithResponses) GetAdditionalSellerInputsOldWithResponse(ctx context.Context, body GetAdditionalSellerInputsOldJSONRequestBody) (*GetAdditionalSellerInputsOldResp, error) {
	rsp, err := c.GetAdditionalSellerInputsOld(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseGetAdditionalSellerInputsOldResp(rsp)
}

// CreateShipmentWithBodyWithResponse request with arbitrary body returning *CreateShipmentResponse
func (c *ClientWithResponses) CreateShipmentWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateShipmentResp, error) {
	rsp, err := c.CreateShipmentWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateShipmentResp(rsp)
}

func (c *ClientWithResponses) CreateShipmentWithResponse(ctx context.Context, body CreateShipmentJSONRequestBody) (*CreateShipmentResp, error) {
	rsp, err := c.CreateShipment(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateShipmentResp(rsp)
}

// CancelShipmentWithResponse request returning *CancelShipmentResponse
func (c *ClientWithResponses) CancelShipmentWithResponse(ctx context.Context, shipmentId string) (*CancelShipmentResp, error) {
	rsp, err := c.CancelShipment(ctx, shipmentId)
	if err != nil {
		return nil, err
	}
	return ParseCancelShipmentResp(rsp)
}

// GetShipmentWithResponse request returning *GetShipmentResponse
func (c *ClientWithResponses) GetShipmentWithResponse(ctx context.Context, shipmentId string) (*GetShipmentResp, error) {
	rsp, err := c.GetShipment(ctx, shipmentId)
	if err != nil {
		return nil, err
	}
	return ParseGetShipmentResp(rsp)
}

// CancelShipmentOldWithResponse request returning *CancelShipmentOldResponse
func (c *ClientWithResponses) CancelShipmentOldWithResponse(ctx context.Context, shipmentId string) (*CancelShipmentOldResp, error) {
	rsp, err := c.CancelShipmentOld(ctx, shipmentId)
	if err != nil {
		return nil, err
	}
	return ParseCancelShipmentOldResp(rsp)
}

// ParseGetAdditionalSellerInputsResp parses an HTTP response from a GetAdditionalSellerInputsWithResponse call
func ParseGetAdditionalSellerInputsResp(rsp *http.Response) (*GetAdditionalSellerInputsResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetAdditionalSellerInputsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetAdditionalSellerInputsResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetEligibleShipmentServicesOldResp parses an HTTP response from a GetEligibleShipmentServicesOldWithResponse call
func ParseGetEligibleShipmentServicesOldResp(rsp *http.Response) (*GetEligibleShipmentServicesOldResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetEligibleShipmentServicesOldResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetEligibleShipmentServicesResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetEligibleShipmentServicesResp parses an HTTP response from a GetEligibleShipmentServicesWithResponse call
func ParseGetEligibleShipmentServicesResp(rsp *http.Response) (*GetEligibleShipmentServicesResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetEligibleShipmentServicesResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetEligibleShipmentServicesResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetAdditionalSellerInputsOldResp parses an HTTP response from a GetAdditionalSellerInputsOldWithResponse call
func ParseGetAdditionalSellerInputsOldResp(rsp *http.Response) (*GetAdditionalSellerInputsOldResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetAdditionalSellerInputsOldResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetAdditionalSellerInputsResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCreateShipmentResp parses an HTTP response from a CreateShipmentWithResponse call
func ParseCreateShipmentResp(rsp *http.Response) (*CreateShipmentResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CreateShipmentResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CreateShipmentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCancelShipmentResp parses an HTTP response from a CancelShipmentWithResponse call
func ParseCancelShipmentResp(rsp *http.Response) (*CancelShipmentResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CancelShipmentResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CancelShipmentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetShipmentResp parses an HTTP response from a GetShipmentWithResponse call
func ParseGetShipmentResp(rsp *http.Response) (*GetShipmentResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetShipmentResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetShipmentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCancelShipmentOldResp parses an HTTP response from a CancelShipmentOldWithResponse call
func ParseCancelShipmentOldResp(rsp *http.Response) (*CancelShipmentOldResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CancelShipmentOldResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CancelShipmentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}
