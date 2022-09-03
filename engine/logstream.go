package engine

type LogStream interface {
	Write(interface{})
}

var (
	logStream LogStream
)
