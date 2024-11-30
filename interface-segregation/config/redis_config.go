package config

import (
	"strconv"
	"sync"
)

type RedisConfig struct {
	mu           sync.RWMutex
	configSource ConfigSource
	address      string
	timeout      int
	maxTotal     int
}

func NewRedisConfig(configSource ConfigSource) *RedisConfig {
	return &RedisConfig{configSource: configSource}
}

// implement Update interface to get hot update
func (r *RedisConfig) Update() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.address = r.configSource.Get("redis.address")
	r.timeout = r.configSource.GetInt("redis.timeout")
	r.maxTotal = r.configSource.GetInt("redis.maxTotal")
}

func (r *RedisConfig) OutputInPlainText() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return "Redis Config - Address: " + r.address + ", Timeout: " + strconv.Itoa(r.timeout)
}

func (r *RedisConfig) Output() map[string]string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return map[string]string{
		"address":  r.address,
		"timeout":  strconv.Itoa(r.timeout),
		"maxTotal": strconv.Itoa(r.maxTotal),
	}
}
