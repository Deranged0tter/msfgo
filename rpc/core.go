package rpc

import "github.com/deranged0tter/msfgo/methods"

type core struct {
	rpc *RPC
}

func (c *core) Version() (*CoreVersionResp, error) {
	req := &GenericRequest{
		Method: methods.CoreVersion,
		Token:  c.rpc.GetToken(),
	}

	var resp *CoreVersionResp
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *core) Stop() error {
	req := &GenericRequest{
		Method: methods.CoreStop,
		Token:  c.rpc.GetToken(),
	}

	var resp *GenericResponse
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return err
	}

	return nil
}

func (c *core) SetG(v string, val string) (*GenericResponse, error) {
	req := &CoreSetGReq{
		Method: methods.CoreSetG,
		Token:  c.rpc.GetToken(),
		V:      v,
		Val:    val,
	}

	var resp *GenericResponse
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *core) UnsetG(v string) (*GenericResponse, error) {
	req := &CoreUnsetGReq{
		Method: methods.CoreUnsetG,
		Token:  c.rpc.GetToken(),
		V:      v,
	}

	var resp *GenericResponse
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *core) Save() (*GenericResponse, error) {
	req := &GenericRequest{
		Method: methods.CoreSave,
		Token:  c.rpc.GetToken(),
	}

	var resp *GenericResponse
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *core) ReloadModules() (*CoreReloadModulesResp, error) {
	req := &GenericRequest{
		Method: methods.CoreReloadModules,
		Token:  c.rpc.GetToken(),
	}

	var resp *CoreReloadModulesResp
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *core) ModuleStats() (*CoreModuleStatsResp, error) {
	req := &GenericRequest{
		Method: methods.CoreModuleStats,
		Token:  c.rpc.GetToken(),
	}

	var resp *CoreModuleStatsResp
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *core) AddModulePath(path string) (*CoreAddModulePathResp, error) {
	req := &CoreModuleAddReq{
		Method: methods.CoreAddModulePath,
		Token:  c.rpc.GetToken(),
		Path:   path,
	}

	var resp *CoreAddModulePathResp
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *core) TheadList() (*CoreThreadListResp, error) {
	req := &GenericRequest{
		Method: methods.CoreThreadList,
		Token:  c.rpc.GetToken(),
	}

	var resp *CoreThreadListResp
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *core) TheadKill(tid int) (*GenericResponse, error) {
	req := &CoreThreadKillReq{
		Method: methods.CoreThreadKill,
		Token:  c.rpc.GetToken(),
		Tid:    tid,
	}

	var resp *GenericResponse
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
