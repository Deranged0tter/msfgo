package rpc

import "github.com/deranged0tter/msfgo/methods"

type console struct {
	rpc *RPC
}

func (c *console) Create() (*ConsoleCreateResp, error) {
	req := &GenericRequest{
		Method: methods.ConsoleCreate,
		Token:  c.rpc.GetToken(),
	}

	var resp *ConsoleCreateResp
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *console) Destroy(id int) (*GenericResponse, error) {
	req := &ConsoleDestroyReq{
		Method: methods.ConsoleDestroy,
		Token:  c.rpc.GetToken(),
		Id:     id,
	}

	var resp *GenericResponse
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *console) List() (*ConsoleListResp, error) {
	req := &GenericRequest{
		Method: methods.ConsoleCreate,
		Token:  c.rpc.GetToken(),
	}

	var resp *ConsoleListResp
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *console) Read(id int) (*ConsoleReadResp, error) {
	req := &ConsoleReadReq{
		Method: methods.ConsoleDestroy,
		Token:  c.rpc.GetToken(),
		Id:     id,
	}

	var resp *ConsoleReadResp
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *console) Write(id int, data string) (*ConsoleWriteResp, error) {
	data += "\r\n"
	req := &ConsoleWriteReq{
		Method: methods.ConsoleWrite,
		Token:  c.rpc.GetToken(),
		Id:     id,
		Data:   data,
	}

	var resp *ConsoleWriteResp
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *console) Tab(id int, line string) (*ConsoleTabResp, error) {
	req := &ConsoleTabReq{
		Method: methods.ConsoleTab,
		Token:  c.rpc.GetToken(),
		Id:     id,
		Line:   line,
	}

	var resp *ConsoleTabResp
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *console) SessionDetach(id int) (*GenericResponse, error) {
	req := &ConsoleSessionDetachReq{
		Method: methods.ConsoleDestroy,
		Token:  c.rpc.GetToken(),
		Id:     id,
	}

	var resp *GenericResponse
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *console) SessionKill(id int) (*GenericResponse, error) {
	req := &ConsoleSessionKillReq{
		Method: methods.ConsoleDestroy,
		Token:  c.rpc.GetToken(),
		Id:     id,
	}

	var resp *GenericResponse
	err := c.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
