package msfgo

import (
	"errors"

	"github.com/deranged0tter/msfgo/rpc"
)

var (
	ErrFailedAuth        error = errors.New("failed to authenticate")
	ErrFailedLogout      error = errors.New("failed to logout")
	ErrTokenAddFail      error = errors.New("failed to add new token")
	ErrTokenGenerateFail error = errors.New("failed to generate new token")
	ErrTokenRemoveFail   error = errors.New("failed to remove token")
)

type AuthMangager struct {
	rpc *rpc.RPC
}

// auth.login
func (am *AuthMangager) Login(username string, password string) (string, error) {
	resp, err := am.rpc.Auth.Login(username, password)
	if err != nil {
		return "", err
	}

	if resp.Result == rpc.Failure || resp.Token == "" {
		return "", ErrFailedAuth
	}

	return resp.Token, nil
}

// auth.logout
func (am *AuthMangager) Logout() error {
	resp, err := am.rpc.Auth.Logout()
	if err != nil {
		return err
	}

	if resp.Result == rpc.Failure {
		return ErrFailedLogout
	}

	return nil
}

// auth.token_list
func (am *AuthMangager) TokenList() ([]string, error) {
	resp, err := am.rpc.Auth.TokenList()
	if err != nil {
		return nil, err
	}

	return resp.Tokens, nil
}

// auth.token_add
func (am *AuthMangager) TokenAdd(token string) error {
	resp, err := am.rpc.Auth.TokenAdd(token)
	if err != nil {
		return err
	}

	if resp.Result == rpc.Failure {
		return ErrTokenAddFail
	}

	return nil
}

// auth.token_generate
func (am *AuthMangager) TokenGenerate() (string, error) {
	resp, err := am.rpc.Auth.TokenGenerate()
	if err != nil {
		return "", err
	}

	if resp.Result == rpc.Failure {
		return "", ErrTokenGenerateFail
	}

	return resp.Token, nil
}

// auth.token_remove
func (am *AuthMangager) TokenRemove(token string) error {
	resp, err := am.rpc.Auth.TokenRemove(token)
	if err != nil {
		return err
	}

	if resp.Result == rpc.Failure {
		return ErrTokenRemoveFail
	}

	return nil
}
