// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostLoginRequestBody struct {
	Apikey string  `json:"apikey"`
	Pin    *string `json:"pin,omitempty"`
}

func (o *PostLoginRequestBody) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

func (o *PostLoginRequestBody) GetPin() *string {
	if o == nil {
		return nil
	}
	return o.Pin
}

type Data struct {
	Token *string `json:"token,omitempty"`
}

func (o *Data) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

// PostLoginResponseBody - response
type PostLoginResponseBody struct {
	Data   *Data   `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}

func (o *PostLoginResponseBody) GetData() *Data {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *PostLoginResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type PostLoginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *PostLoginResponseBody
}

func (o *PostLoginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostLoginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostLoginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostLoginResponse) GetObject() *PostLoginResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
