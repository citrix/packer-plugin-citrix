package common

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hooklift/gowsdl/soap"
)

type ConnectConfig struct {
	ELMServer          string `mapstructure:"elm_server"`
	ELMUsername        string `mapstructure:"elm_username"`
	ELMPassword        string `mapstructure:"elm_password"`
	InsecureConnection bool   `mapstructure:"insecure_connection"`
}

func (c *ConnectConfig) Prepare() []error {
	var errs []error

	if c.ELMUsername == "" {
		errs = append(errs, fmt.Errorf("'elm_username' is required"))
	}
	if c.ELMPassword == "" {
		errs = append(errs, fmt.Errorf("'elm_password' is required"))
	}
	if c.ELMServer == "" {
		errs = append(errs, fmt.Errorf("'elm_server' is required"))
	}

	return errs
}

type StepConnect struct {
	Config *ConnectConfig
}

// Create unidesk soap client and connect to ELM server
func (s *StepConnect) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	UiSayf(ui, "InsecureConnection %t", s.Config.InsecureConnection)
	httpClient := &http.Client{
		Transport: &elmsoap.HeaderCaptureTransport{
			Rt: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: s.Config.InsecureConnection, // 🚨 only use for testing!
				},
			},
		},
	}
	unideskurl, err := elmsoap.BuildServerURL(s.Config.ELMServer)
	if err != nil {
		ui.Errorf("Error building server URL: %v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	UiSayf(ui, "Connecting to ELM server at %s", unideskurl)

	soapClient := soap.NewClient(unideskurl, soap.WithHTTPClient(httpClient))
	client := elmsoap.NewApiSoap(soapClient)

	loginRequest := &elmsoap.Login{
		Command: &elmsoap.LoginCommand{
			UserName:   s.Config.ELMUsername,
			Password:   s.Config.ELMPassword,
			RememberMe: false,
			Culture:    "en-US",
		},
	}
	response, err := client.Login(loginRequest)
	if err != nil {
		ui.Errorf("Error calling Login: %v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}

	// Check for SOAP-level login failure: server returns HTTP 200 but an error in the body,
	// and no Token is issued. An empty Token means authentication was rejected.
	if response.LoginResult == nil || response.LoginResult.Token == "" {
		loginErr := fmt.Errorf("ELM login failed for user '%s': invalid credentials or insufficient privileges", s.Config.ELMUsername)
		ui.Error(loginErr.Error())
		state.Put("error", loginErr)
		return multistep.ActionHalt
	}

	transport := httpClient.Transport.(*elmsoap.HeaderCaptureTransport)
	transport.Unidesk_token = response.LoginResult.Token

	// Use cookie and token already captured by HeaderCaptureTransport during Login.
	// Some ELM server versions do not issue Set-Cookie and rely solely on the
	// Unidesk_token header for session authentication; an empty cookie is not fatal.
	cookie := transport.Cookie
	token := response.LoginResult.Token
	if cookie == "" {
		UiSay(ui, "[INFO] No Set-Cookie in login response; proceeding with token-only authentication")
	}

	helper := &elmsoap.SoapHelper{
		Client:             client,
		Cookie:             cookie,
		Token:              token,
		URL:                unideskurl,
		InsecureSkipVerify: s.Config.InsecureConnection,
	}
	state.Put("soap_helper", helper)
	// Keep individual state vars for builder.go generatedData propagation.
	state.Put("COOKIE", cookie)
	state.Put("UNIDESK_TOKEN", token)
	state.Put("INSECURE_CONNECTION", s.Config.InsecureConnection)
	state.Put("UNIDESK_URL", unideskurl)

	return multistep.ActionContinue
}

func (s *StepConnect) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}
