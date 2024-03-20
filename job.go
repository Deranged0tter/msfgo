package msfgo

import (
	"github.com/deranged0tter/msfgo/errors"
	"github.com/deranged0tter/msfgo/rpc"
)

type JobManager struct {
	rpc *rpc.RPC
}

// job.list
func (jb *JobManager) List() (map[string]string, error) {
	resp, err := jb.rpc.Job.List()
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

type JobInfo struct {
	Jid       int
	Name      string
	StartTime int
	Datastore map[string]interface{}
}

// job.info
func (jb *JobManager) Info(jid int) (JobInfo, error) {
	jobInfo := JobInfo{}

	resp, err := jb.rpc.Job.Info(jid)
	if err != nil {
		return jobInfo, err
	}

	jobInfo.Jid = resp.Jid
	jobInfo.Name = resp.Name
	jobInfo.StartTime = resp.StartTime
	jobInfo.Datastore = resp.Datastore

	return jobInfo, nil
}

// job.stop
func (jb *JobManager) Stop(jid int) error {
	resp, err := jb.rpc.Job.Stop(jid)
	if err != nil {
		return err
	}

	if resp.Result == rpc.Failure {
		return errors.ErrJobStopFailed
	}

	return nil
}
