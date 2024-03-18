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
