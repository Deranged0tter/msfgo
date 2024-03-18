package rpc

import "github.com/deranged0tter/msfgo/methods"

type auth struct {
	rpc *RPC
}

// handle client auth
func (a *auth) Login(username string, password string) (*AuthLoginResp, error) {
	req := &AuthLoginReq{
		Method:   methods.AuthLogin,
		Username: username,
		Password: password,
	}

	var resp *AuthLoginResp
	err := a.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *auth) Logout() (*GenericResponse, error) {
	req := &GenericRequest{
		Method: methods.AuthLogout,
		Token:  a.rpc.GetToken(),
	}

	var resp *GenericResponse
	err := a.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *auth) TokenList() (*AuthTokenListResp, error) {
	req := &GenericRequest{
		Method: methods.AuthTokenList,
		Token:  a.rpc.GetToken(),
	}

	var resp *AuthTokenListResp
	err := a.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *auth) TokenAdd(token string) (*GenericResponse, error) {
	req := &AuthTokenAddReq{
		Method:   methods.AuthTokenAdd,
		Token:    a.rpc.GetToken(),
		NewToken: token,
	}

	var resp *GenericResponse
	err := a.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *auth) TokenGenerate() (*AuthTokenGenerateResp, error) {
	req := &GenericRequest{
		Method: methods.AuthTokenGenerate,
		Token:  a.rpc.GetToken(),
	}

	var resp *AuthTokenGenerateResp
	err := a.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, err
}
func (a *auth) TokenRemove(token string) (*GenericResponse, error) {
	req := &AuthTokenRemoveReq{
		Method:   methods.AuthTokenRemove,
		Token:    a.rpc.GetToken(),
		DelToken: token,
	}

	var resp *GenericResponse
	err := a.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
