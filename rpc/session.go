package rpc

import "github.com/deranged0tter/msfgo/methods"

type session struct {
	rpc *RPC
}

func (s *session) List() (*SessionListResp, error) {
	req := &GenericRequest{
		Method: methods.SessionList,
		Token:  s.rpc.GetToken(),
	}

	var resp *SessionListResp
	err := s.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *session) Stop(sid int) (*GenericResponse, error) {
	req := &SessionStopReq{
		Method: methods.SessionStop,
		Token:  s.rpc.GetToken(),
		Sid:    sid,
	}

	var resp *GenericResponse
	err := s.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *session) ShellRead(sid int) (*SessionShellReadResp, error) {
	req := &SessionShellReadReq{
		Method: methods.SessionShellRead,
		Token:  s.rpc.GetToken(),
		Sid:    sid,
	}

	var resp *SessionShellReadResp
	err := s.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *session) ShellWrite(sid int, data string) (*SessionShellWriteResp, error) {
	req := &SessionShellWriteReq{
		Method: methods.SessionShellWrite,
		Token:  s.rpc.GetToken(),
		Sid:    sid,
		Data:   data,
	}

	var resp *SessionShellWriteResp
	err := s.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *session) ShellUpgrade(sid int, lhost string, lport int) (*GenericResponse, error) {
	req := &SessionShellUpgradeReq{
		Method: methods.SessionShellUpgrade,
		Token:  s.rpc.GetToken(),
		Sid:    sid,
		LHost:  lhost,
		LPort:  lport,
	}

	var resp *GenericResponse
	err := s.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *session) MeterpreterRead(sid int) (*MeterpreterReadResp, error) {
	req := &SessionMeterpreterReadReq{
		Method: methods.SessionMeterpreterRead,
		Token:  s.rpc.GetToken(),
		Sid:    sid,
	}

	var resp *MeterpreterReadResp
	err := s.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *session) MeterpreterWrite(sid int, data string) (*GenericResponse, error) {
	req := &SessionMeterpreterWriteReq{
		Method: methods.SessionMeterpreterWrite,
		Token:  s.rpc.GetToken(),
		Sid:    sid,
		Data:   data,
	}

	var resp *GenericResponse
	err := s.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
