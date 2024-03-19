package errors

import "errors"

// auth errors
var (
	ErrAuthLoginFailed         error = errors.New("failed to authenticate")
	ErrAuthLogoutFailed        error = errors.New("failed to logout")
	ErrAuthTokenAddFailed      error = errors.New("failed to add new token")
	ErrAuthTokenGenerateFailed error = errors.New("failed to generate new token")
	ErrAuthTokenRemoveFailed   error = errors.New("failed to remove token")
)

// console errors
var (
	ErrConsoleDestroyFailed       error = errors.New("failed to destroy console")
	ErrConsoleSessionDetachFailed error = errors.New("failed to detach session")
	ErrConsoleSessionKillFailed   error = errors.New("failed to kill session")
)

// core errors
var (
	ErrCoreSetGFailed       error = errors.New("failed to set global variable")
	ErrCoreUnsetGFailed     error = errors.New("failed to unset global variable")
	ErrCoreSaveFailed       error = errors.New("failed to save framework settings")
	ErrCoreThreadKillFailed error = errors.New("failed to kill thread")
)

// db errors

// job errors

// module errors

// plugin errors

// session errors
