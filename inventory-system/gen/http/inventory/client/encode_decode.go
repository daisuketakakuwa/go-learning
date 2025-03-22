// Code generated by goa v3.7.2, DO NOT EDIT.
//
// inventory HTTP client encoders and decoders
//
// Command:
// $ goa gen inventory-system/api/svc/design

package client

import (
	"bytes"
	"context"
	inventory "inventory-system/gen/inventory"
	inventoryviews "inventory-system/gen/inventory/views"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "inventory" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateInventoryPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("inventory", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the inventory
// create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*inventory.CreateProductPayload)
		if !ok {
			return goahttp.ErrInvalidType("inventory", "create", "*inventory.CreateProductPayload", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("inventory", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the
// inventory create endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("inventory", "create", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("inventory", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "inventory" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		productID int
	)
	{
		p, ok := v.(*inventory.UpdateProductPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("inventory", "update", "*inventory.UpdateProductPayload", v)
		}
		productID = p.ProductID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateInventoryPath(productID)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("inventory", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the inventory
// update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*inventory.UpdateProductPayload)
		if !ok {
			return goahttp.ErrInvalidType("inventory", "update", "*inventory.UpdateProductPayload", v)
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("inventory", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the
// inventory update endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("inventory", "update", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("inventory", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildFindRequest instantiates a HTTP request object with method and path set
// to call the "inventory" service "find" endpoint
func (c *Client) BuildFindRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		productID int
	)
	{
		p, ok := v.(*inventory.FindPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("inventory", "find", "*inventory.FindPayload", v)
		}
		productID = p.ProductID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: FindInventoryPath(productID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("inventory", "find", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeFindResponse returns a decoder for responses returned by the inventory
// find endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeFindResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body FindResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("inventory", "find", err)
			}
			p := NewFindproductresultViewOK(&body)
			view := "default"
			vres := &inventoryviews.Findproductresult{Projected: p, View: view}
			if err = inventoryviews.ValidateFindproductresult(vres); err != nil {
				return nil, goahttp.ErrValidationError("inventory", "find", err)
			}
			res := inventory.NewFindproductresult(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("inventory", "find", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "inventory" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		productID int
	)
	{
		p, ok := v.(*inventory.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("inventory", "delete", "*inventory.DeletePayload", v)
		}
		productID = p.ProductID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteInventoryPath(productID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("inventory", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteResponse returns a decoder for responses returned by the
// inventory delete endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("inventory", "delete", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("inventory", "delete", resp.StatusCode, string(body))
		}
	}
}
