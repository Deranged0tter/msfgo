package rpc

// Generic Response Most RPC Calls receive
type GenericResponse struct {
	Result Result `msgpack:"result"`
}

// auth.login response
type AuthLoginResp struct {
	Result Result `msgpack:"result"`
	Token  string `msgpack:"token"`
}

// auth.token_list response
type AuthTokenListResp struct {
	Tokens []string `msgpack:"tokens"`
}

// auth.token_generate response
type AuthTokenGenerateResp struct {
	Result Result `msgpack:"result"`
	Token  string `msgpack:"token"`
}

// console.create response
type ConsoleCreateResp struct {
	Id     int    `msgpack:"id"`
	Prompt string `msgpack:"prompt"`
	Busy   bool   `msgpack:"busy"`
}

// console.list response
type ConsoleListResp map[string][]struct {
	Id     int    `msgpack:"id"`
	Prompt string `msgpack:"prompt"`
	Busy   bool   `msgpack:"busy"`
}

// conosle.read response
type ConsoleReadResp struct {
	Data   string `msgpack:"data"`
	Prompt string `msgpack:"prompt"`
	Busy   bool   `msgpack:"busy"`
}

// console.write response
type ConsoleWriteResp struct {
	Wrote uint `msgpack:"wrote"`
}

// console.tab response
type ConsoleTabResp struct {
	Tabs string `msgpack:"tabs"`
}

// core.version response
type CoreVersionResp struct {
	Version string `msgpack:"version"`
	Ruby    string `msgpack:"ruby"`
	Api     string `msgpack:"api"`
}

// core.reload_modules response
type CoreReloadModulesResp struct {
	Exploits  uint `msgpack:"exploits"`
	Auxiliary uint `msgpack:"auxiliary"`
	Post      uint `msgpack:"post"`
	Encoders  uint `msgpack:"encoders"`
	NOPs      uint `msgpack:"nops"`
	Payloads  uint `msgpack:"payloads"`
	Evasions  uint `msgpack:"evasions"`
}

// core.module_stats
type CoreModuleStatsResp struct {
	Exploits  uint `msgpack:"exploits"`
	Auxiliary uint `msgpack:"auxiliary"`
	Post      uint `msgpack:"post"`
	Encoders  uint `msgpack:"encoders"`
	NOPs      uint `msgpack:"nops"`
	Payloads  uint `msgpack:"payloads"`
	Evasions  uint `msgpack:"evasions"`
}

// core.add_module_path
type CoreAddModulePathResp struct {
	Exploits  uint `msgpack:"exploits"`
	Auxiliary uint `msgpack:"auxiliary"`
	Post      uint `msgpack:"post"`
	Encoders  uint `msgpack:"encoders"`
	NOPs      uint `msgpack:"nops"`
	Payloads  uint `msgpack:"payloads"`
	Evasions  uint `msgpack:"evasions"`
}

// core.thread_list
type CoreThreadListResp map[string][]struct {
	Status   string `msgpack:"status"`
	Critical bool   `msgpack:"critical"`
	Name     string `msgpack:"name"`
	Started  string `msgpack:"started"`
}

// job.list
type JobListResp map[string]string

// job.info
type JobInfoResp struct {
	Jid       int                    `msgpack:"jid"`
	Name      string                 `msgpack:"name"`
	StartTime int                    `msgpack:"start_time"`
	Datastore map[string]interface{} `msgpack:"datastore,omitempty"`
}

// module.exploits
type ModuleExploitsResp struct {
	Modules []string `msgpack:"modules"`
}

// module.evasion
type ModuleEvasionResp struct {
	Modules []string `msgpack:"modules"`
}

// module.auxiliary
type ModuleAuxiliaryResp struct {
	Modules []string `msgpack:"modules"`
}

// module.payloads
type ModulePayloadsResp struct {
	Modules []string `msgpack:"modules"`
}

// module.encoders
type ModuleEncodersResp struct {
	Modules []string `msgpack:"modules"`
}

// module.nops
type ModuleNopsResp struct {
	Modules []string `msgpack:"modules"`
}

// module.platforms
type ModulePlatformsResp struct {
	Platforms []string `msgpack:"platforms"`
}

// module.post
type ModulePostResp struct {
	Modules []string `msgpack:"modules"`
}

// module.info
type ModuleInfoRes struct {
	Name        string     `msgpack:"name"`
	Description string     `msgpack:"description"`
	License     string     `msgpack:"license"`
	FilePath    string     `msgpack:"filepath"`
	Version     string     `msgpack:"version"`
	Rank        string     `msgpack:"rank"`
	References  [][]string `msgpack:"references"`
	Authors     []string   `msgpack:"authors"`
}
