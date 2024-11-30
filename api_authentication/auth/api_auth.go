package auth

import (
	"errors"
	"time"
)

type APIAutheneicator interface {
	Auth(URL string) error
	AuthWithRequest(apiRequest APIRequest) error
}

type DefaultAPIAuthenticator struct {
	credentialStorage CredentialStorage
}

func NewDefaultAPIAuthenticator() *DefaultAPIAuthenticator {
	return &DefaultAPIAuthenticator{
		credentialStorage: NewInMemoryCredentialStorage(),
	}
}

func (a *DefaultAPIAuthenticator) Auth(url string) error {
	req, err := NewAPIRequestFromFullURL(url)
	if err != nil {
		return err
	}

	return a.AuthWithRequest(req)
}

func (a *DefaultAPIAuthenticator) AuthWithRequest(req *APIRequest) error {
	appID := req.GetAppID()
	token := req.GetToken()
	timestamp := req.GetTimeStamp()
	baseURL := req.GetBaseUrl()

	authToken := NewAuthToken(token, time.Unix(timestamp, 0), DEFAULT_EXPIRED_TIME_INTERVAL)
	if authToken.isExpire() {
		return errors.New("token expired")
	}

	password, err := a.credentialStorage.getPasswordByAppID(appID)
	if err != nil {
		return err
	}

	// generate token
	serverAuthToken := GenerateToken(baseURL, appID, password, timestamp)
	if !authToken.Match(serverAuthToken) {
		return errors.New("token verification failed")
	}

	return nil
}
