package rpc

import "github.com/deranged0tter/msfgo/methods"

type module struct {
	rpc *RPC
}

func (m *module) Exploits() (*ModuleExploitsResp, error) {
	req := &GenericRequest{
		Method: methods.ModuleExploits,
		Token:  m.rpc.GetToken(),
	}

	var resp *ModuleExploitsResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) Evasion() (*ModuleEvasionResp, error) {
	req := &GenericRequest{
		Method: methods.ModuleEvasion,
		Token:  m.rpc.GetToken(),
	}

	var resp *ModuleEvasionResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) Auxiliary() (*ModuleAuxiliaryResp, error) {
	req := &GenericRequest{
		Method: methods.ModuleAuxiliary,
		Token:  m.rpc.GetToken(),
	}

	var resp *ModuleAuxiliaryResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) Payloads() (*ModulePayloadsResp, error) {
	req := &GenericRequest{
		Method: methods.ModulePayloads,
		Token:  m.rpc.GetToken(),
	}

	var resp *ModulePayloadsResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
