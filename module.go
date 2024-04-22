package msfgo

import "github.com/deranged0tter/msfgo/rpc"

type ModuleManager struct {
	rpc *rpc.RPC
}

// module.exploits
func (mm *ModuleManager) Exploits() ([]string, error) {
	resp, err := mm.rpc.Module.Exploits()
	if err != nil {
		return nil, err
	}

	return resp.Modules, nil
}

// module.evasion
func (mm *ModuleManager) Evasion() ([]string, error) {
	resp, err := mm.rpc.Module.Evasion()
	if err != nil {
		return nil, err
	}

	return resp.Modules, nil
}

// module.auxiliary
func (mm *ModuleManager) Auxiliary() ([]string, error) {
	resp, err := mm.rpc.Module.Auxiliary()
	if err != nil {
		return nil, err
	}

	return resp.Modules, nil
}

// module.payloads
func (mm *ModuleManager) Payloads() ([]string, error) {
	resp, err := mm.rpc.Module.Payloads()
	if err != nil {
		return nil, err
	}

	return resp.Modules, nil
}

// module.encoders
func (mm *ModuleManager) Encoders() ([]string, error) {
	resp, err := mm.rpc.Module.Encoders()
	if err != nil {
		return nil, err
	}

	return resp.Modules, nil
}

// module.nops
func (mm *ModuleManager) Nops() ([]string, error) {
	resp, err := mm.rpc.Module.Nops()
	if err != nil {
		return nil, err
	}

	return resp.Modules, nil
}

// module.platforms
func (mm *ModuleManager) Platforms() ([]string, error) {
	resp, err := mm.rpc.Module.Platforms()
	if err != nil {
		return nil, err
	}

	return resp.Platforms, nil
}

// module.post
func (mm *ModuleManager) Post() ([]string, error) {
	resp, err := mm.rpc.Module.Post()
	if err != nil {
		return nil, err
	}

	return resp.Modules, nil
}

type ModuleInfo struct {
	Name        string
	Description string
	License     string
	FilePath    string
	Version     string
	Rank        string
	References  [][]string
	Authors     []string
}

// module.info
func (mm *ModuleManager) Info(moduleName string, moduleType string) (ModuleInfo, error) {
	var info ModuleInfo

	resp, err := mm.rpc.Module.Info(moduleType, moduleName)
	if err != nil {
		return info, err
	}

	info = ModuleInfo{
		Name:        resp.Name,
		Description: resp.Description,
		License:     resp.License,
		FilePath:    resp.FilePath,
		Version:     resp.Version,
		Rank:        resp.Rank,
		References:  resp.References,
		Authors:     resp.Authors,
	}

	return info, nil
}

// module.compatible_payloads
func (mm *ModuleManager) CompatiblePayloads(moduleName string) ([]string, error) {
	resp, err := mm.rpc.Module.CompatiblePayloads(moduleName)
	if err != nil {
		return nil, err
	}

	return resp.Payloads, nil
}

// module.compatible_sessions
func (mm *ModuleManager) CompatibleSessions(moduleName string) ([]string, error) {
	resp, err := mm.rpc.Module.CompatibleSessions(moduleName)
	if err != nil {
		return nil, err
	}

	return resp.Sessions, nil
}

// module.options
func (mm *ModuleManager) Options(moduleName string, moduleType string) ([]string, error) {
	resp, err := mm.rpc.Module.Options(moduleType, moduleName)
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(*resp))
	for k := range *resp {
		keys = append(keys, k)
	}

	return keys, nil
}

// module.options where option is required
func (mm *ModuleManager) RequiredOptions(moduleName string, moduleType string) ([]string, error) {
	resp, err := mm.rpc.Module.Options(moduleType, moduleName)
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(*resp))
	for k, v := range *resp {
		if v.Required {
			keys = append(keys, k)
		}
	}

	return keys, nil
}

type ModuleExecute struct {
	JID  int
	UUID string
}

// module.execute
func (mm *ModuleManager) Execute(moduleName string, moduleType string, options map[string]string) (ModuleExecute, error) {
	var results ModuleExecute

	resp, err := mm.rpc.Module.Execute(moduleType, moduleName, options)
	if err != nil {
		return results, err
	}

	results = ModuleExecute{
		JID:  resp.Jid,
		UUID: resp.UUID,
	}

	return results, nil
}

// module.search
func (mm *ModuleManager) Search(query string) ([]string, error) {
	resp, err := mm.rpc.Module.Search(query)
	if err != nil {
		return nil, err
	}

	return resp.Matches, nil
}

// module.check
func (mm *ModuleManager) Check(moduleName string, moduleType string, options map[string]string) error {
	return mm.rpc.Module.Check(moduleType, moduleName, options)
}
