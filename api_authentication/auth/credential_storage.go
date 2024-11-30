package auth

import "errors"

type CredentialStorage interface {
	getPasswordByAppID(appID string) (string, error)
}

type InMemoryCredentialStorage struct {
	storage map[string]string
}

func NewInMemoryCredentialStorage() *InMemoryCredentialStorage {
	return &InMemoryCredentialStorage{storage: make(map[string]string)}
}

func (cs *InMemoryCredentialStorage) AddCredential(appID, password string) {
	cs.storage[appID] = password
}

func (cs *InMemoryCredentialStorage) getPasswordByAppID(appID string) (string, error) {
	password, exists := cs.storage[appID]
	if !exists {
		return "", errors.New("appID not found")
	}
	return password, nil
}
