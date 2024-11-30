package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

const (
	DEFAULT_EXPIRED_TIME_INTERVAL = 1 * 60 * 1000
)

type AuthToken struct {
	token              string
	createdTime        time.Time
	expireTimeInterval time.Duration
}

// instantiation
func NewAuthToken(token string, created time.Time, expire time.Duration) *AuthToken {
	return &AuthToken{
		token:              token,
		createdTime:        created,
		expireTimeInterval: expire,
	}
}

// generate token
func GenerateToken(url, appID, password string, timestamp int64) string {
	msg := url + appID + password + strconv.Itoa(int(timestamp))
	hash := hmac.New(sha256.New, []byte(password))
	hash.Write([]byte(msg))
	// return signature
	return hex.EncodeToString(hash.Sum(nil))
}

func (a *AuthToken) isExpire() bool {
	return time.Now().After(a.createdTime.Add(a.expireTimeInterval))
}

func (a *AuthToken) Match(token string) bool {
	return a.token == token
}

func (a *AuthToken) GetToken() string {
	return a.token
}
