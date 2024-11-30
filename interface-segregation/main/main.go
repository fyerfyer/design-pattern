package main

import (
	"interface/config"
	"interface/server"
	"interface/updater"
)

func main() {
	configSource := &config.ZookeeperConfigSource{}

	redisConfig := config.NewRedisConfig(configSource)
	kafkaConfig := config.NewKafkaConfig(configSource)
	mysqlConfig := config.NewMysqlConfig(configSource)

	redisUpdater := updater.NewScheduledUpdater(redisConfig, 300)
	go redisUpdater.Run()

	kafkaUpdater := updater.NewScheduledUpdater(kafkaConfig, 60)
	go kafkaUpdater.Run()

	// 启动 HTTP 服务
	httpServer := server.NewSimpleHttpServer("127.0.0.1", 2389)
	httpServer.AddViewer("/config", redisConfig)
	httpServer.AddViewer("/config", mysqlConfig)
	httpServer.Run()
}
