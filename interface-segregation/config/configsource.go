package config

type ConfigSource interface {
	Get(key string) string
	GetInt(key string) int
}

type ZookeeperConfigSource struct{}

func (z *ZookeeperConfigSource) Get(key string) string {
	// get string config info
	return "mock_value"
}

func (z *ZookeeperConfigSource) GetInt(key string) int {
	// get int config info
	return 100
}
