package abstract_class

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Level int

var globalLevel Level = 4

var MsgQueueClient = MessageQueueClient{queueName: "public msgqueue client"}

// Logger interface defines the behavior that subclasses must implement
type Logger interface {
	Log(level Level, message string) // Abstract method for logging
	IsLoggable(level Level) bool
}

// BaseLogger provides common logic, simulating an abstract class
type BaseLogger struct {
	name              string
	enabled           bool
	minPermittedLevel Level
}

// initBaseLogger initializes the common logic, only for use by subclasses
func (b *BaseLogger) initBaseLogger(name string, enabled bool, minPermittedLevel Level) {
	b.name = name
	b.enabled = enabled
	b.minPermittedLevel = minPermittedLevel
}

func (b *BaseLogger) IsLoggable(level Level) bool {
	return b.enabled && int(b.minPermittedLevel) <= int(level)
}

// FileLogger subclass: embeds BaseLogger and implements Logger interface
type FileLogger struct {
	BaseLogger
	fileWriter io.Writer
}

// NewFileLogger
func NewFileLogger(name string, enabled bool, minPermittedLevel Level, filepath string) *FileLogger {
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}

	logger := &FileLogger{
		fileWriter: file,
	}

	// go will bind the call to embedded BaseLogger field
	logger.initBaseLogger(name, enabled, minPermittedLevel)
	return logger
}

// Log implements the Logger interface's Log method for FileLogger
func (f *FileLogger) Log(level Level, message string) {
	if !f.IsLoggable(level) {
		return
	}
	logMessage := fmt.Sprintf("[%d] %s\n", level, message)
	_, _ = f.fileWriter.Write([]byte(logMessage))
}

// MessageQueueLogger subclass
type MessageQueueLogger struct {
	BaseLogger
	msgQueueClient MessageQueueClient
}

// MessageQueueClient simulates a message queue middleware client
type MessageQueueClient struct {
	queueName string
}

// NewMessageQueueLogger
func NewMessageQueueLogger(name string, enabled bool, minPermittedLevel Level, msgQueueClient MessageQueueClient) *MessageQueueLogger {
	logger := &MessageQueueLogger{
		msgQueueClient: msgQueueClient,
	}
	logger.initBaseLogger(name, enabled, minPermittedLevel)
	return logger
}

// Log
func (m *MessageQueueLogger) Log(level Level, message string) {
	if !m.IsLoggable(level) {
		return
	}
	logMessage := fmt.Sprintf("[%d] %s", level, message)
	m.msgQueueClient.Send(logMessage)
}

// Send simulates sending a message to the message queue
func (mq *MessageQueueClient) Send(message string) {
	fmt.Printf("Message sent to queue %s: %s\n", mq.queueName, message)
}
