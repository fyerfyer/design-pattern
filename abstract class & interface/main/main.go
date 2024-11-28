package main

import "abstract_class/abstract_class"

func main() {
	// fileLogger
	fileLogger := abstract_class.NewFileLogger("FileLogger", true, 3, "log.txt")
	fileLogger.Log(4, "This is a file log message")

	// msgQueueLogger
	msgQueueLogger := abstract_class.NewMessageQueueLogger("MQLogger", true, 2, abstract_class.MsgQueueClient)
	msgQueueLogger.Log(3, "This is a message queue log message")
}
