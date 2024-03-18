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
