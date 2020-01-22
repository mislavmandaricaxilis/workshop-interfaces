package domain

type LoggerV2 interface {
	Log(s string, data string)
}

type Logger interface {
	Log(s string)
}

func NewLogger(loggerV2 LoggerV2) Logger {
	return loggerAdapter{
		loggerV2: loggerV2,
	}
}

type loggerAdapter struct {
	loggerV2 LoggerV2
}

func (l loggerAdapter) Log(s string) {
	l.loggerV2.Log(s, "")
}