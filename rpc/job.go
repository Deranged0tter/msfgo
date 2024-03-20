package rpc

import "github.com/deranged0tter/msfgo/methods"

type job struct {
	rpc *RPC
}

func (j *job) List() (*JobListResp, error) {
	req := &GenericRequest{
		Method: methods.JobList,
		Token:  j.rpc.GetToken(),
	}

	var resp *JobListResp
	err := j.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (j *job) Info(jid int) (*JobInfoResp, error) {
	req := &JobInfoReq{
		Method: methods.JobInfo,
		Token:  j.rpc.GetToken(),
		Jid:    jid,
	}

	var resp *JobInfoResp
	err := j.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (j *job) Stop(jid int) (*GenericResponse, error) {
	req := &JobStopReq{
		Method: methods.JobStop,
		Token:  j.rpc.GetToken(),
		Jid:    jid,
	}

	var resp *GenericResponse
	err := j.rpc.Call(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
