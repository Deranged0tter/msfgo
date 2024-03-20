package msfgo

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/deranged0tter/msfgo/rpc"
)

// Object that makes rpc calls
type MsfClient struct {
	Username   string
	Password   string
	ApiVersion string
	Rpc        *rpc.RPC
	Auth       *AuthMangager
	Console    *ConsoleManager
	Core       *CoreManager
	Job        *JobManager
}

// Options for Making Client
type MsfClientOptions struct {
	Timeout         time.Duration
	ProxyUrl        string
	TlsClientConfig *tls.Config
	Token           string
	SSL             bool
	ApiVersion      string
}

// create new msfclient
// returns a MsfClient Object
func NewClient(address string) (*MsfClient, error) {
	// define options for new client
	clientOptions := MsfClientOptions{
		Token:           "",
		SSL:             true,
		TlsClientConfig: &tls.Config{InsecureSkipVerify: true},
		ApiVersion:      "1.0",
	}

	// get url
	url := createUrl(address, clientOptions.SSL, clientOptions.ApiVersion)

	// create new HTTP Client
	http, err := newHttpClient(clientOptions)
	if err != nil {
		return nil, err
	}

	// create new RPC client
	rpc := rpc.NewRPC(http, url)

	if clientOptions.Token != "" {
		rpc.SetToken(clientOptions.Token)
	}

	msfClient := &MsfClient{
		ApiVersion: clientOptions.ApiVersion,
		Rpc:        rpc,
		Auth:       &AuthMangager{rpc: rpc},
		Console:    &ConsoleManager{rpc: rpc},
		Core:       &CoreManager{rpc: rpc},
		Job:        &JobManager{rpc: rpc},
	}

	return msfClient, nil
}

// Login will send the "auth.login" api request. The authentication token expires after
// 5 minutes. The token is automatically renewed when you make a new RPC request.
func (client *MsfClient) Login(username string, password string) error {
	token, err := client.Auth.Login(username, password)
	if err != nil {
		return err
	}

	client.Username = username
	client.Password = password

	client.Rpc.SetToken(token)

	return nil
}

// create http client
func newHttpClient(options MsfClientOptions) (*http.Client, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = options.TlsClientConfig

	if options.ProxyUrl != "" {
		proxyUrl, err := url.Parse(options.ProxyUrl)
		if err != nil {
			return nil, err
		}

		transport.Proxy = http.ProxyURL(proxyUrl)
	}

	return &http.Client{
		Timeout:   options.Timeout,
		Transport: transport,
	}, nil
}

// create a url
func createUrl(addr string, isSsl bool, apiVersion string) string {
	protocol := "http"
	if isSsl {
		protocol = "https"
	}

	return fmt.Sprintf("%s://%s/api/%s", protocol, addr, apiVersion)
}
