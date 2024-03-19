package rpc

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"reflect"

	"github.com/vmihailenco/msgpack/v5"
)

type Result string

var (
	Failure Result = "failure"
	Success Result = "success"
)

var ErrNotAuth error = errors.New("client is not authenticated (no token)")

// RPC object
type RPC struct {
	http    *http.Client
	url     string
	token   string
	Auth    *auth
	Console *console
}

// create RPC object
func NewRPC(http *http.Client, url string) *RPC {
	// create new RPC object
	rpc := &RPC{
		http: http,
		url:  url,
	}

	rpc.Auth = &auth{rpc: rpc}

	return rpc
}

// make a call using rpc object
func (r *RPC) Call(request, response interface{}) error {
	method := rpcMethod(request)
	if method != "auth.login" && method != "health.check" {
		if r.token == "" {
			return ErrNotAuth
		}
	}

	buffer := new(bytes.Buffer)
	encoder := msgpack.NewEncoder(buffer)
	encoder.UseArrayEncodedStructs(true)

	err := encoder.Encode(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", r.url, buffer)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "binary/message-pack")

	resp, err := r.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if response != nil {
		err = msgpack.NewDecoder(resp.Body).Decode(response)
		if err != nil {
			return err
		}
	}

	return nil
}

// get token from RPC client
func (r *RPC) GetToken() string {
	return r.token
}

// set token on RPC client
func (r *RPC) SetToken(token string) {
	r.token = token
}

func rpcMethod(request interface{}) string {
	stype := reflect.ValueOf(request).Elem()
	field := stype.FieldByName("Method")

	if field.IsValid() {
		return field.String()
	}

	return ""
}
