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
