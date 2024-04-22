package rpc

// Generic Request for most RPC Calls
type GenericRequest struct {
	Method string
	Token  string
}

// auth.login request
type AuthLoginReq struct {
	Method   string
	Username string
	Password string
}

// auth.token_add request
type AuthTokenAddReq struct {
	Method   string
	Token    string
	NewToken string
}

// auth.token_remove request
type AuthTokenRemoveReq struct {
	Method   string
	Token    string
	DelToken string
}

// console.destroy request
type ConsoleDestroyReq struct {
	Method string
	Token  string
	Id     int
}

// console.read request
type ConsoleReadReq struct {
	Method string
	Token  string
	Id     int
}

// console.write request
type ConsoleWriteReq struct {
	Method string
	Token  string
	Id     int
	Data   string
}

// console.tab request
type ConsoleTabReq struct {
	Method string
	Token  string
	Id     int
	Line   string
}

// console.session_detach request
type ConsoleSessionDetachReq struct {
	Method string
	Token  string
	Id     int
}

// console.session_kill request
type ConsoleSessionKillReq struct {
	Method string
	Token  string
	Id     int
}

// core.setg request
type CoreSetGReq struct {
	Method string
	Token  string
	V      string
	Val    string
}

// core.unsetg request
type CoreUnsetGReq struct {
	Method string
	Token  string
	V      string
}

// core.add_module_path request
type CoreModuleAddReq struct {
	Method string
	Token  string
	Path   string
}

// core.thread_kill request
type CoreThreadKillReq struct {
	Method string
	Token  string
	Tid    int
}

// job.info request
type JobInfoReq struct {
	Method string
	Token  string
	Jid    int
}

// job.stop request
type JobStopReq struct {
	Method string
	Token  string
	Jid    int
}

// module.info request
type ModuleInfoReq struct {
	Method     string
	Token      string
	ModuleType string
	ModuleName string
}

// module.compatible_payloads request
type ModuleCompatiblePayloadsReq struct {
	Method     string
	Token      string
	ModuleName string
}

// module.compatible_sessions request
type ModuleCompatibleSessionsReq struct {
	Method     string
	Token      string
	ModuleName string
}

// module.options request
type ModuleOptionsReq struct {
	Method     string
	Token      string
	ModuleType string
	ModuleName string
}

// module.execute request
type ModuleExecuteReq struct {
	Method     string
	Token      string
	ModuleType string
	ModuleName string
	Options    map[string]string
}

// module.search request
type ModuleSearchReq struct {
	Method string
	Token  string
	Match  string
}

// module.check request
type ModuleCheckReq struct {
	Method     string
	Token      string
	ModuleType string
	ModuleName string
	Options    map[string]string
}

// session.stop request
type SessionStopReq struct {
	Method string
	Token  string
	Sid    int
}

// session.shell_read
type SessionShellReadReq struct {
	Method string
	Token  string
	Sid    int
}

// session.shell_write
type SessionShellWriteReq struct {
	Method string
	Token  string
	Sid    int
	Data   string
}

// session.shell_upgrade
type SessionShellUpgradeReq struct {
	Method string
	Token  string
	Sid    int
	LHost  string
	LPort  int
}

// session.interactive_read
type SessionMeterpreterReadReq struct {
	Method string
	Token  string
	Sid    int
}

// session.interactive_write
type SessionMeterpreterWriteReq struct {
	Method string
	Token  string
	Sid    int
	Data   string
}
