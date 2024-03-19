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
