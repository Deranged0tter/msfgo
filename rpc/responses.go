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
type ModuleInfoResp struct {
	Name        string     `msgpack:"name"`
	Description string     `msgpack:"description"`
	License     string     `msgpack:"license"`
	FilePath    string     `msgpack:"filepath"`
	Version     string     `msgpack:"version"`
	Rank        string     `msgpack:"rank"`
	References  [][]string `msgpack:"references"`
	Authors     []string   `msgpack:"authors"`
}

// module.compatible_payloads
type ModuleCompatiblePayloadsResp struct {
	Payloads []string `msgpack:"payloads"`
}

// module.compatible_sessions
type ModuleCompatibleSessionsResp struct {
	Sessions []string `msgpack:"sessions"`
}

// module.options
type ModuleOptionsResp map[string]struct {
	Type     string      `msgpack:"type"`
	Required bool        `msgpack:"required"`
	Advanced bool        `msgpack:"advanced"`
	Evasion  bool        `msgpack:"evasion"`
	Desc     string      `msgpack:"desc"`
	Default  interface{} `msgpack:"default"`
	Enums    []string    `msgpack:"enums,omitempty"`
}

// module.execute
type ModuleExecuteResp struct {
	Jid  int    `msgpack:"job_id"`
	UUID string `msgpack:"uuid"`
}

// module.search
type ModuleSearchResp struct {
	Matches []string `msgpack:"matches"`
}

// session.list
type SessionListResp map[int]struct {
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit"`
	Description string `msgpack:"desc"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"session_host"`
	SessionPort int    `msgpack:"session_port"`
	TargetHost  string `msgpack:"target_host"`
	Username    string `msgpack:"Username"`
	Uuid        string `msgpack:"uuid"`
	ExploitUuid string `msgpack:"exploit_uuid"`
	Routes      string `msgpack:"routes"`
	Platform    string `msgpack:"platform"`
}

// session.shell_read
type SessionShellReadResp struct {
	Seq  string `msgpack:"seq"`
	Data string `msgpack:"data"`
}

// session.shell_write
type SessionShellWriteResp struct {
	WriteCount int `msgpack:"write_count"`
}

// session.interactive_read
type MeterpreterReadResp struct {
	Data string `msgpack:"data"`
}
