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

func (m *module) Encoders() (*ModuleEncodersResp, error) {
	req := &GenericRequest{
		Method: methods.ModuleEncoders,
		Token:  m.rpc.GetToken(),
	}

	var resp *ModuleEncodersResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) Nops() (*ModuleNopsResp, error) {
	req := &GenericRequest{
		Method: methods.ModuleNops,
		Token:  m.rpc.GetToken(),
	}

	var resp *ModuleNopsResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) Platforms() (*ModulePlatformsResp, error) {
	req := &GenericRequest{
		Method: methods.ModulePlatforms,
		Token:  m.rpc.GetToken(),
	}

	var resp *ModulePlatformsResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) Post() (*ModulePostResp, error) {
	req := &GenericRequest{
		Method: methods.ModuleNops,
		Token:  m.rpc.GetToken(),
	}

	var resp *ModulePostResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) Info(mType string, mName string) (*ModuleInfoResp, error) {
	req := &ModuleInfoReq{
		Method:     methods.ModuleNops,
		Token:      m.rpc.GetToken(),
		ModuleType: mType,
		ModuleName: mName,
	}

	var resp *ModuleInfoResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) CompatiblePayloads(mName string) (*ModuleCompatiblePayloadsResp, error) {
	req := &ModuleCompatiblePayloadsReq{
		Method:     methods.ModuleNops,
		Token:      m.rpc.GetToken(),
		ModuleName: mName,
	}

	var resp *ModuleCompatiblePayloadsResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) CompatibleSessions(mName string) (*ModuleCompatibleSessionsResp, error) {
	req := &ModuleCompatibleSessionsReq{
		Method:     methods.ModuleNops,
		Token:      m.rpc.GetToken(),
		ModuleName: mName,
	}

	var resp *ModuleCompatibleSessionsResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) Options(mType string, mName string) (*ModuleOptionsResp, error) {
	req := &ModuleOptionsReq{
		Method:     methods.ModuleNops,
		Token:      m.rpc.GetToken(),
		ModuleType: mType,
		ModuleName: mName,
	}

	var resp *ModuleOptionsResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) Execute(mType string, mName string, options map[string]string) (*ModuleExecuteResp, error) {
	req := &ModuleExecuteReq{
		Method:     methods.ModuleNops,
		Token:      m.rpc.GetToken(),
		ModuleType: mType,
		ModuleName: mName,
		Options:    options,
	}

	var resp *ModuleExecuteResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) Search(query string) (*ModuleSearchResp, error) {
	req := &ModuleSearchReq{
		Method: methods.ModuleNops,
		Token:  m.rpc.GetToken(),
		Match:  query,
	}

	var resp *ModuleSearchResp
	err := m.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *module) Check(mType string, mName string, options map[string]string) error {
	req := &ModuleCheckReq{
		Method:     methods.ModuleCheck,
		Token:      m.rpc.GetToken(),
		ModuleType: mType,
		ModuleName: mName,
		Options:    options,
	}

	return m.rpc.Call(req, nil)
}
