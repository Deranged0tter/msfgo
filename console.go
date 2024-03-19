package msfgo

import (
	"github.com/deranged0tter/msfgo/errors"
	"github.com/deranged0tter/msfgo/rpc"
)

type ConsoleManager struct {
	rpc *rpc.RPC
}

type Console struct {
	Id     int
	Prompt string
	Busy   bool
}

type ConsoleRead struct {
	Data   string
	Prompt string
	Busy   bool
}

// console.create
func (cm *ConsoleManager) Create() (Console, error) {
	var console Console

	resp, err := cm.rpc.Console.Create()
	if err != nil {
		return console, nil
	}

	console.Id = resp.Id
	console.Prompt = resp.Prompt
	console.Busy = resp.Busy

	return console, nil
}

// console.destroy
func (cm *ConsoleManager) Destroy(id int) error {
	resp, err := cm.rpc.Console.Destroy(id)
	if err != nil {
		return nil
	}

	if resp.Result == rpc.Failure {
		return errors.ErrConsoleDestroyFailed
	}

	return nil
}

// console.list
func (cm *ConsoleManager) List() (*rpc.ConsoleListResp, error) {
	resp, err := cm.rpc.Console.List()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// console.read
func (cm *ConsoleManager) Read(id int) (ConsoleRead, error) {
	var data ConsoleRead

	resp, err := cm.rpc.Console.Read(id)
	if err != nil {
		return data, err
	}

	data.Data = resp.Data
	data.Prompt = resp.Prompt
	data.Busy = resp.Busy

	return data, nil
}

// console.write
// return number of bytes sent
func (cm *ConsoleManager) Write(id int, data string) (uint, error) {
	resp, err := cm.rpc.Console.Write(id, data)
	if err != nil {
		return 0, err
	}

	return resp.Wrote, nil
}

// console.tab
func (cm *ConsoleManager) Tab(id int, line string) (string, error) {
	resp, err := cm.rpc.Console.Tab(id, line)
	if err != nil {
		return "", err
	}

	return resp.Tabs, nil
}

// console.session_detach
func (cm *ConsoleManager) SessionDetach(id int) error {
	resp, err := cm.rpc.Console.SessionDetach(id)
	if err != nil {
		return nil
	}

	if resp.Result == rpc.Failure {
		return errors.ErrConsoleSessionDetachFailed
	}

	return nil
}

// console.session_kill
func (cm *ConsoleManager) SessionKill(id int) error {
	resp, err := cm.rpc.Console.SessionKill(id)
	if err != nil {
		return nil
	}

	if resp.Result == rpc.Failure {
		return errors.ErrConsoleSessionKillFailed
	}

	return nil
}
