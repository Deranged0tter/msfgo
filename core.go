package msfgo

import (
	"errors"

	"github.com/deranged0tter/msfgo/rpc"
)

var (
	ErrFailedSetG       error = errors.New("failed to set global variable")
	ErrFailedUnsetG     error = errors.New("failed to unset global variable")
	ErrFailedSave       error = errors.New("failed to save framework settings")
	ErrThreadKillFailed error = errors.New("failed to kill thread")
)

type CoreManager struct {
	rpc *rpc.RPC
}

// core.version
func (cm *CoreManager) Version() (string, error) {
	resp, err := cm.rpc.Core.Version()
	if err != nil {
		return "", err
	}

	return resp.Version, nil
}

// core.version
func (cm *CoreManager) ApiVersion() (string, error) {
	resp, err := cm.rpc.Core.Version()
	if err != nil {
		return "", err
	}

	return resp.Api, nil
}

// core.stop
func (cm *CoreManager) Stop() error {
	return cm.rpc.Core.Stop()
}

// core.setg
func (cm *CoreManager) SetG(v string, val string) error {
	resp, err := cm.rpc.Core.SetG(v, val)
	if err != nil {
		return err
	}

	if resp.Result == rpc.Failure {
		return ErrFailedSetG
	}

	return nil
}

// core.unsetg
func (cm *CoreManager) UnsetG(v string) error {
	resp, err := cm.rpc.Core.UnsetG(v)
	if err != nil {
		return err
	}

	if resp.Result == rpc.Failure {
		return ErrFailedUnsetG
	}

	return nil
}

// core.save
func (cm *CoreManager) Save() error {
	resp, err := cm.rpc.Core.Save()
	if err != nil {
		return err
	}

	if resp.Result == rpc.Failure {
		return ErrFailedSave
	}

	return nil
}

// core.reload_modules
func (cm *CoreManager) ReloadModules() (*rpc.CoreReloadModulesResp, error) {
	resp, err := cm.rpc.Core.ReloadModules()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// core.module_stats
func (cm *CoreManager) ModuleStats() (*rpc.CoreModuleStatsResp, error) {
	resp, err := cm.rpc.Core.ModuleStats()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// core.add_module_path
func (cm *CoreManager) AddModulePath(path string) (*rpc.CoreAddModulePathResp, error) {
	resp, err := cm.rpc.Core.AddModulePath(path)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// core.thread_list
func (cm *CoreManager) ThreadList() (*rpc.CoreThreadListResp, error) {
	resp, err := cm.rpc.Core.TheadList()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// core.thread_kill
func (cm *CoreManager) ThreadKill(tid int) error {
	resp, err := cm.rpc.Core.TheadKill(tid)
	if err != nil {
		return err
	}

	if resp.Result == rpc.Failure {
		return ErrThreadKillFailed
	}

	return nil
}
