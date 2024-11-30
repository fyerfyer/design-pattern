package config

import "sync"

type MysqlConfig struct {
	mu           sync.RWMutex
	configSource ConfigSource
	address      string
	username     string
	password     string
}

func NewMysqlConfig(configSource ConfigSource) *MysqlConfig {
	return &MysqlConfig{configSource: configSource}
}

func (m *MysqlConfig) OutputInPlainText() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return "MySQL Config - Address: " + m.address + ", Username: " + m.username
}

func (m *MysqlConfig) Output() map[string]string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return map[string]string{
		"address":  m.address,
		"username": m.username,
		"password": m.password,
	}
}
