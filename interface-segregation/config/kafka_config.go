package config

import "sync"

type KafkaConfig struct {
	mu           sync.RWMutex
	configSource ConfigSource
	address      string
	topic        string
}

func NewKafkaConfig(configSource ConfigSource) *KafkaConfig {
	return &KafkaConfig{configSource: configSource}
}

func (k *KafkaConfig) Update() {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.address = k.configSource.Get("kafka.address")
	k.topic = k.configSource.Get("kafka.topic")
}
